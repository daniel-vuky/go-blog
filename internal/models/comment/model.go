package admin

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Comment struct {
	CommentID int64       `json:"comment_id"`
	PostID    int64       `json:"post_id"`
	UserID    int64       `json:"user_id"`
	ParentID  pgtype.Int8 `json:"parent_id"`
	Comment   string      `json:"comment"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
