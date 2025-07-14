-- name: ListArticles :many
SELECT * FROM articles
ORDER BY created_at;

-- name: GetArticle :one
SELECT * FROM articles
WHERE id = $1
LIMIT 1;

-- name: CreateArticle :one
INSERT INTO articles (
  id, title, content, author, published
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateArticle :exec
UPDATE articles
SET modified_at = $1
WHERE id = $2
RETURNING *;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = $1;
