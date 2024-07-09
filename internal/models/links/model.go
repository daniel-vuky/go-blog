package admin

import (
	"time"
)

type PostLink struct {
	LinkID     int64     `json:"link_id"`
	CategoryID int64     `json:"category_id"`
	PostID     int64     `json:"post_id"`
	CreatedAt  time.Time `json:"created_at"`
}
