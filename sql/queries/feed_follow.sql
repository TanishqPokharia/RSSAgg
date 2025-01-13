-- name: CreateFeedFollow :one
insert into feed_follows(user_id,feed_id) values ($1,$2) returning *;

-- name: GetFeedFollows :many
select * from feed_follows where user_id = $1;

-- name: DeleteFeedFollows :exec
delete from feed_follows where id = $1 and user_id = $2;