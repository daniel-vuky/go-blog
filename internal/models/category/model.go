package admin

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	CategoryID       int64       `json:"category_id"`
	ParentID         int64       `json:"parent_id"`
	Name             string      `json:"name"`
	UrlKey           pgtype.Text `json:"url_key"`
	ShortDescription pgtype.Text `json:"short_description"`
	Description      pgtype.Text `json:"description"`
	CreatedAt        time.Time   `json:"created_at"`
}
