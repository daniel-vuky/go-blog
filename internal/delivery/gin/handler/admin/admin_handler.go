package admin

import (
	"errors"
	"fmt"
	model "github.com/daniel-vuky/go-blog/internal/models/admin"
	"github.com/daniel-vuky/go-blog/internal/service/admin"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"time"
)

type Handler struct {
	service *admin.Service
}

// NewHandler create a new handler
func NewHandler(s *admin.Service) *Handler {
	return &Handler{
		service: s,
	}
}

// GetAdmin Get admin by email
// @Param email
// @Success 200 {object} model.Admin
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /admin/{email} [get]
func (s *Handler) GetAdmin(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}
	loadedAdmin, err := s.service.GetAdmin(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loadedAdmin)
}

// getListAdminParams
type getListAdminParams struct {
	Email          string `json:"email" form:"email" binding:"omitempty,max=255"`
	Active         bool   `json:"active" form:"active" binding:"omitempty"`
	Firstname      string `json:"firstname" form:"firstname" binding:"omitempty,max=32"`
	Lastname       string `json:"lastname" form:"lastname" binding:"omitempty,max=32"`
	OrderBy        string `json:"order_by" form:"order_by" binding:"omitempty"`
	OrderDirection string `json:"order_direction" form:"order_direction" binding:"omitempty,oneof=asc desc"`
	PageSize       int32  `json:"page_size" form:"page_size" binding:"required,gt=0"`
	CurrentPage    int32  `json:"current_page" form:"current_page" binding:"required,gt=0"`
}

// GetListAdmin Get list of admins
// @Param getListAdminParams
// @Success 200 {object} []model.Admin
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /admin [get]
func (s *Handler) GetListAdmin(ctx *gin.Context) {
	var arg getListAdminParams
	if err := ctx.ShouldBindQuery(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print(arg)
	admins, err := s.service.GetListAdmin(ctx, &model.GetListAdminParams{
		Email:          pgtype.Text{String: arg.Email, Valid: arg.Email != ""},
		Active:         pgtype.Bool{Bool: arg.Active, Valid: true},
		Firstname:      pgtype.Text{String: arg.Firstname, Valid: arg.Firstname != ""},
		Lastname:       pgtype.Text{String: arg.Lastname, Valid: arg.Lastname != ""},
		OrderBy:        arg.OrderBy,
		OrderDirection: arg.OrderDirection,
		PageSize:       arg.PageSize,
		CurrentPage:    arg.CurrentPage,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, admins)
}

// createAdminParams
type createAdminParams struct {
	RoleID      int64     `json:"role_id" binding:"required,gt=0"`
	Email       string    `json:"email" binding:"required,email,max=255"`
	Password    string    `json:"password" binding:"required"`
	Firstname   string    `json:"firstname" binding:"required,max=32"`
	Lastname    string    `json:"lastname" binding:"max=32"`
	Active      bool      `json:"active"`
	LockExpires time.Time `json:"lock_expires"`
}

// CreateAdmin Create a new admin
// @Param createAdminParams
// @Success 200 {object} model.Admin
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /admin [post]
func (s *Handler) CreateAdmin(ctx *gin.Context) {
	var arg createAdminParams
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdAdmin, err := s.service.CreateAdmin(ctx, &model.CreateAdminParams{
		RoleID:         arg.RoleID,
		Email:          arg.Email,
		HashedPassword: arg.Password,
		Firstname:      arg.Firstname,
		Lastname: pgtype.Text{
			String: arg.Lastname,
			Valid:  arg.Lastname != "",
		},
		Active: pgtype.Bool{
			Bool:  arg.Active,
			Valid: true,
		},
		LockExpires: pgtype.Timestamptz{
			Time:  arg.LockExpires,
			Valid: arg.LockExpires != time.Time{},
		},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createdAdmin)
}

type updateAdminParams struct {
	Email       string    `json:"email" binding:"required,email,max=255"`
	RoleID      int64     `json:"role_id" binding:"required,gt=0"`
	Password    string    `json:"password"`
	Firstname   string    `json:"firstname" binding:"max=32"`
	Lastname    string    `json:"lastname" binding:"max=32"`
	Active      bool      `json:"active"`
	LockExpires time.Time `json:"lock_expires"`
}

// UpdateAdmin Update admin params
// @Param updateAdminParams
// @Success 200 {object} model.Admin
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /admin [put]
func (s *Handler) UpdateAdmin(ctx *gin.Context) {
	var arg updateAdminParams
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedAdmin, err := s.service.UpdateAdmin(ctx, &model.UpdateAdminParams{
		Email: arg.Email,
		RoleID: pgtype.Int8{
			Int64: arg.RoleID,
			Valid: true,
		},
		HashedPassword: pgtype.Text{
			String: arg.Password,
			Valid:  arg.Password != "",
		},
		Firstname: pgtype.Text{
			String: arg.Firstname,
			Valid:  arg.Firstname != "",
		},
		Lastname: pgtype.Text{
			String: arg.Lastname,
			Valid:  arg.Lastname != "",
		},
		Active: pgtype.Bool{
			Bool:  arg.Active,
			Valid: true,
		},
		LockExpires: pgtype.Timestamptz{
			Time:  arg.LockExpires,
			Valid: arg.LockExpires != time.Time{},
		},
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedAdmin)
}

// DeleteAdmin Delete an admin
// @Param email
// @Success 200 {object} model.Admin
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /admin/{email} [delete]
func (s *Handler) DeleteAdmin(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}
	deletedAdmin, err := s.service.DeleteAdmin(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, deletedAdmin)
}
