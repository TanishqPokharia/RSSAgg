-- name: CreateFeed :one

insert into feeds(title,url,user_id) values ($1,$2,$3) returning *;

-- name: GetFeeds :many

select * from feeds;

-- name: GetNextFeedsToFetch :many
select * from feeds order by last_fetched_at asc nulls first limit $1;

-- name: MarkFeedAsFetched :one
update feeds set last_fetched_at = now() where id = $1 returning *;