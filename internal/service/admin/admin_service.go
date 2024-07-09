package admin

import (
	"context"
	model "github.com/daniel-vuky/go-blog/internal/models/admin"
	"github.com/daniel-vuky/go-blog/internal/repository/admin"
)

// Service
// Wraps the Repository struct from the repository package.
type Service struct {
	AdminRepo admin.Repository
}

// NewService
// Returns a new instance of Service.
func NewService(repo admin.Repository) *Service {
	return &Service{AdminRepo: repo}
}

// CreateAdmin
// Creates a new admin.
// @param c context.Context
// @param arg *model.CreateAdminParams
// @return model.Admin
func (s *Service) CreateAdmin(c context.Context, arg *model.CreateAdminParams) (model.Admin, error) {
	return s.AdminRepo.Create(c, arg)
}

// DeleteAdmin
// Deletes an admin.
// @param c context.Context
// @param email string
// @return model.Admin
func (s *Service) DeleteAdmin(c context.Context, email string) (model.Admin, error) {
	return s.AdminRepo.Delete(c, email)
}

// GetAdmin
// Returns an admin.
// @param c context.Context
// @param email string
// @return model.Admin
func (s *Service) GetAdmin(c context.Context, email string) (model.Admin, error) {
	return s.AdminRepo.Get(c, email)
}

// GetListAdmin
// Returns a list of admins.
// @param c context.Context
// @param arg *model.GetListAdminParams
// @return []model.Admin
func (s *Service) GetListAdmin(c context.Context, arg *model.GetListAdminParams) ([]model.Admin, error) {
	return s.AdminRepo.GetList(c, arg)
}

// UpdateAdmin
// Updates an admin.
// @param c context.Context
// @param arg *model.UpdateAdminParams
// @return model.Admin
func (s *Service) UpdateAdmin(c context.Context, arg *model.UpdateAdminParams) (model.Admin, error) {
	return s.AdminRepo.Update(c, arg)
}

// IsAdminActive
// Checks if an admin is active.
// @param c context.Context
// @param email string
// @return bool
func (s *Service) IsAdminActive(c context.Context, email string) (bool, error) {
	adminUser, err := s.AdminRepo.Get(c, email)
	if err != nil {
		return false, err
	}
	return adminUser.Active.Bool, nil
}
