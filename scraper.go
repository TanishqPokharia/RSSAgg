package main

import (
	"RSSAggregator/internal/database"
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"
)

func beginScraping(
	db *database.Queries,
	concurrency int,
	timeGap time.Duration,
) {
	ticker := time.NewTicker(timeGap)
	for ; ; <-ticker.C { // ticker channel receives a tick every timeGap unit
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println(err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1) // increment to wait group to indicate a goroutine has been initiated
			go scrapeFeed(db, &feed, wg)
		}
		wg.Wait() // wait for all goroutines to finish
	}

}

func scrapeFeed(db *database.Queries, feed *database.Feed, wg *sync.WaitGroup) {
	defer wg.Done() // decrement wg counter indicating that a goroutine finished execution
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched id: %s \n %s", feed.ID, err)
		return
	}
	rssFeed, err := urlToFeed(feed.Url.String)
	if err != nil {
		log.Printf("Error fetching feed: %s", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldnt parse date %s\n", err)
			continue
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			Description: description,
			PublishedAt: t,
			Url:         item.Link,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				continue
			}
			log.Printf("failed to create post %v\n", err)
		}

	}
}
