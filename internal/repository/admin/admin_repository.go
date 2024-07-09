package admin

import (
	"context"
	adminModel "github.com/daniel-vuky/go-blog/internal/models/admin"
)

type Reader interface {
	Get(ctx context.Context, email string) (adminModel.Admin, error)
	GetList(ctx context.Context, arg *adminModel.GetListAdminParams) ([]adminModel.Admin, error)
}

type Writer interface {
	Create(ctx context.Context, arg *adminModel.CreateAdminParams) (adminModel.Admin, error)
	Delete(ctx context.Context, email string) (adminModel.Admin, error)
	Update(ctx context.Context, arg *adminModel.UpdateAdminParams) (adminModel.Admin, error)
}

type Repository interface {
	Reader
	Writer
}
