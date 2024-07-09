package admin

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type UrlRewriteEntity string

const (
	UrlRewriteEntity1 UrlRewriteEntity = "1"
	UrlRewriteEntity2 UrlRewriteEntity = "2"
)

func (e *UrlRewriteEntity) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UrlRewriteEntity(s)
	case string:
		*e = UrlRewriteEntity(s)
	default:
		return fmt.Errorf("unsupported scan type for UrlRewriteEntity: %T", src)
	}
	return nil
}

type NullUrlRewriteEntity struct {
	UrlRewriteEntity UrlRewriteEntity `json:"url_rewrite_entity"`
	Valid            bool             `json:"valid"` // Valid is true if UrlRewriteEntity is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUrlRewriteEntity) Scan(value interface{}) error {
	if value == nil {
		ns.UrlRewriteEntity, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UrlRewriteEntity.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUrlRewriteEntity) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UrlRewriteEntity), nil
}

type UrlRewrite struct {
	UrlRewriteID int64                `json:"url_rewrite_id"`
	EntityType   NullUrlRewriteEntity `json:"entity_type"`
	EntityID     pgtype.Int8          `json:"entity_id"`
	UrlKey       pgtype.Text          `json:"url_key"`
	CreatedAt    time.Time            `json:"created_at"`
}
