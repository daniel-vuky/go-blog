package admin

import (
	"context"
	adminModel "github.com/daniel-vuky/go-blog/internal/models/admin"
)

type Reader interface {
	GetAdmin(ctx context.Context, email string) (adminModel.Admin, error)
	GetListAdmin(ctx context.Context, arg *adminModel.GetListAdminParams) ([]adminModel.Admin, error)
	IsAdminActive(ctx context.Context, email string) (bool, error)
}

type Writer interface {
	CreateAdmin(ctx context.Context, arg *adminModel.CreateAdminParams) (adminModel.Admin, error)
	DeleteAdmin(ctx context.Context, email string) (adminModel.Admin, error)
	UpdateAdmin(ctx context.Context, arg *adminModel.UpdateAdminParams) (adminModel.Admin, error)
}

type UseCase interface {
	Reader
	Writer
}
