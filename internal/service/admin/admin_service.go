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

// convertAdminToModel
// Converts an admin from the repository to the model.
// @param adminUser model.Admin
// @return model.Admin
func convertAdminToModel(adminUser *model.Admin) model.Admin {
	return model.Admin{
		AdminID:           adminUser.AdminID,
		RoleID:            adminUser.RoleID,
		Email:             adminUser.Email,
		Firstname:         adminUser.Firstname,
		Lastname:          adminUser.Lastname,
		Active:            adminUser.Active,
		LockExpires:       adminUser.LockExpires,
		PasswordChangedAt: adminUser.PasswordChangedAt,
		CreatedAt:         adminUser.CreatedAt,
	}
}

// CreateAdmin
// Creates a new admin.
// @param c context.Context
// @param arg *model.CreateAdminParams
// @return model.Admin
func (s *Service) CreateAdmin(c context.Context, arg *model.CreateAdminParams) (model.Admin, error) {
	createdAdmin, err := s.AdminRepo.Create(c, arg)
	if err != nil {
		return createdAdmin, err
	}

	return convertAdminToModel(&createdAdmin), nil
}

// DeleteAdmin
// Deletes an admin.
// @param c context.Context
// @param email string
// @return model.Admin
func (s *Service) DeleteAdmin(c context.Context, email string) (model.Admin, error) {
	deletedAdmin, err := s.AdminRepo.Delete(c, email)
	if err != nil {
		return deletedAdmin, err
	}

	return convertAdminToModel(&deletedAdmin), nil
}

// GetAdmin
// Returns an admin.
// @param c context.Context
// @param email string
// @return model.Admin
func (s *Service) GetAdmin(c context.Context, email string) (model.Admin, error) {
	existedAdmin, err := s.AdminRepo.Get(c, email)
	if err != nil {
		return existedAdmin, err
	}

	return convertAdminToModel(&existedAdmin), err
}

type ListAdminResponse struct {
	Totals int64         `json:"totals"`
	Admins []model.Admin `json:"admins"`
}

// GetListAdmin
// Returns a list of admins.
// @param c context.Context
// @param arg *model.GetListAdminParams
// @return ListAdminResponse
func (s *Service) GetListAdmin(c context.Context, arg *model.GetListAdminParams) (ListAdminResponse, error) {
	var rsp ListAdminResponse
	if arg.OrderBy == "" {
		arg.OrderBy = "admin_id"
	}
	if arg.OrderDirection == "" {
		arg.OrderDirection = "desc"
	}
	listAdmin, totalAdmin, err := s.AdminRepo.GetList(c, arg)
	if err != nil {
		return rsp, err
	}
	rsp = ListAdminResponse{
		Totals: totalAdmin,
		Admins: listAdmin,
	}

	return rsp, nil
}

// UpdateAdmin
// Updates an admin.
// @param c context.Context
// @param arg *model.UpdateAdminParams
// @return model.Admin
func (s *Service) UpdateAdmin(c context.Context, arg *model.UpdateAdminParams) (model.Admin, error) {
	updatedAdmin, err := s.AdminRepo.Update(c, arg)
	if err != nil {
		return updatedAdmin, err
	}

	return convertAdminToModel(&updatedAdmin), nil
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
