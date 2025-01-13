-- name: CreatePost :one
insert into posts(title,description,published_at,url,feed_id) values ($1,$2,$3,$4,$5) returning *;


-- name: GetPostsForUser :many
select P.* from feed_follows FF join posts P using(feed_id) where FF.user_id = $1 order by P.published_at desc;