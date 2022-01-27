-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: ListUsersById :many
SELECT * FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (username, bio, avatar, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET bio = $2, avatar = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;


-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id;

-- name: ListPostsById :many
SELECT * FROM posts
WHERE id= $1;

-- name: ListPostsByUserId :many
SELECT * FROM posts
WHERE user_id= $1;

-- name: CreatePost :one
INSERT INTO posts (url, caption, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdatePost :one
UPDATE posts
SET url = $2, caption = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeletePost :one
DELETE FROM posts
WHERE id = $1
RETURNING *;

-- name: ListCommentsById :many
SELECT * FROM comments
WHERE id= $1;

-- name: ListCommentsByPostId :many
SELECT * FROM comments
WHERE post_id= $1;

-- name: GetComment :one
SELECT * FROM comments
WHERE id = $1;

-- name: ListComments :many
SELECT * FROM comments
ORDER BY id;

-- name: CreateComment :one
INSERT INTO comments (contents, user_id, post_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateComment :one
UPDATE comments
SET contents = $2, updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteComment :one
DELETE FROM comments
WHERE id = $1
RETURNING *;

-- name: ListPostLikesById :many
SELECT * FROM post_likes
WHERE id= $1;

-- name: ListPostLikesByPostId :many
SELECT * FROM post_likes
WHERE post_id= $1;

-- name: ListPostLikes :many
SELECT * FROM post_likes
ORDER BY id;

-- name: GetPostLike :one
SELECT * FROM post_likes
WHERE id = $1;

-- name: CreatePostLike :one
INSERT INTO post_likes (user_id, post_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeletePostLike :one
DELETE FROM post_likes
WHERE id = $1
RETURNING *;

-- name: ListCommentLikes :many
SELECT * FROM comment_likes
ORDER BY id;

-- name: ListCommentLikesById :many
SELECT * FROM comment_likes
WHERE id = $1;

-- name: ListCommentLikesByCommentId :many
SELECT * FROM comment_likes
WHERE comment_id = $1;

-- name: GetCommentLike :one
SELECT * FROM comment_likes
WHERE id = $1;

-- name: CreateCommentLike :one
INSERT INTO comment_likes (user_id, comment_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteCommentLike :one
DELETE FROM comment_likes
WHERE id = $1
RETURNING *;

-- name: GetHashtagByTitle :one
SELECT * FROM hashtags
WHERE title = $1;

-- name: GetHashtagById :one
SELECT * FROM hashtags
WHERE id = $1;

-- name: ListHashtags :many
SELECT * FROM hashtags
ORDER BY id;

-- name: ListHashtagsById :many
SELECT * FROM hashtags
WHERE id = $1;

-- name: CreateHashtag :one
INSERT INTO hashtags (title, created_at)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateHashtag :one
UPDATE hashtags
SET title = $2
WHERE id = $1
RETURNING *;

-- name: DeleteHashtag :one
DELETE FROM hashtags
WHERE id = $1
RETURNING *;

-- name: GetHashtagPostById :one
SELECT * FROM hashtag_post
WHERE id = $1;

-- name: ListHashtagsPost :many
SELECT * FROM hashtag_post
ORDER BY id;

-- name: ListHashtagsPostById :many
SELECT * FROM hashtag_post
WHERE id = $1;

-- name: CreateHashtagPost :one
INSERT INTO hashtag_post (hashtag_id, post_id)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteHashtagPost :one
DELETE FROM hashtag_post
WHERE id = $1
RETURNING *;

-- name: GetFollower :one
SELECT * FROM followers
WHERE id = $1;

-- name: ListFollowers :many
SELECT * FROM followers
ORDER BY id;

-- name: ListFollowersById :many
SELECT * FROM followers
WHERE id = $1;

-- name: ListFollowersByLeaderId :many
SELECT * FROM followers
WHERE leader_id = $1;

-- name: CreateFollower :one
INSERT INTO followers (leader_id, follower_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteFollower :one
DELETE FROM followers
WHERE id = $1
RETURNING *;