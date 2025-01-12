-- name: CreateFeed :one

insert into feeds(title,url,user_id) values ($1,$2,$3) returning *;

-- name: GetFeeds :many

select * from feeds;