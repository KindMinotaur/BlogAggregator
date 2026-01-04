-- name: UnfollowFeed :one
DELETE FROM feed_follows
WHERE feed_follows.user_id = $1 AND feed_follows.feed_id = $2
RETURNING *;
