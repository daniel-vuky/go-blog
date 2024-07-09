package admin

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Post struct {
	PostID           int64       `json:"post_id"`
	Name             string      `json:"name"`
	ShortDescription pgtype.Text `json:"short_description"`
	Description      pgtype.Text `json:"description"`
	Content          pgtype.Text `json:"content"`
	UrlKey           pgtype.Text `json:"url_key"`
	Thumbnail        pgtype.Text `json:"thumbnail"`
	AuthorID         int64       `json:"author_id"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}
