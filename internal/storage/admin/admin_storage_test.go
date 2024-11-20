package admin

import (
	"context"
	model "github.com/daniel-vuky/go-blog/internal/models/admin"
	"github.com/daniel-vuky/go-blog/pkg/config"
	goRandom "github.com/daniel-vuky/go-random"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
	"time"
)

var repository *Repository

// TestMain
// Initializes the repository and closes the connection pool after all tests have run.
func TestMain(m *testing.M) {
	loadedConfig, err := config.LoadConfig("../../../")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	connPool, err := loadedConfig.ConnectToPgxPool()
	if err != nil {
		log.Fatalf("failed to create connection pool: %v", err)
	}
	repository = NewAdminRepository(connPool)
	code := m.Run()
	repository.connPool.Close()
	os.Exit(code)
}

// compareAdmin
// Compare two admins to check if they are equal.
func compareAdmin(t *testing.T, fromAdmin *model.Admin, targetAdmin *model.Admin) {
	require.Equal(t, fromAdmin.RoleID, targetAdmin.RoleID)
	require.Equal(t, fromAdmin.Email, targetAdmin.Email)
	require.Equal(t, fromAdmin.Firstname, targetAdmin.Firstname)
	require.Equal(t, fromAdmin.Lastname.String, targetAdmin.Lastname.String)
	require.Equal(t, fromAdmin.Active.Bool, targetAdmin.Active.Bool)
	require.WithinDuration(t, fromAdmin.LockExpires.Time, targetAdmin.LockExpires.Time, 5*time.Second)
}

// createRandomAdmin
// Creates a random admin for testing.
func createRandomAdmin(t *testing.T) model.Admin {
	arg := &model.CreateAdminParams{
		RoleID:         1,
		Email:          goRandom.RandomEmail(),
		HashedPassword: goRandom.RandomString(10),
		Firstname:      goRandom.RandomString(10),
		Lastname: pgtype.Text{
			String: goRandom.RandomString(10),
			Valid:  true,
		},
		Active: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
		LockExpires: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}
	createdAdmin, err := repository.Create(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdAdmin)
	compareAdmin(t, &model.Admin{
		RoleID:      arg.RoleID,
		Email:       arg.Email,
		Firstname:   arg.Firstname,
		Lastname:    arg.Lastname,
		Active:      arg.Active,
		LockExpires: arg.LockExpires,
	}, &createdAdmin)

	return createdAdmin
}

// TestRepository_Create_Success
// Tests the Create method with success result
func TestRepository_Create_Success(t *testing.T) {
	createRandomAdmin(t)
}

// TestRepository_Create_Duplicate
// Tests the Create method with a duplicate email.
func TestRepository_Create_Duplicate(t *testing.T) {
	newAdmin := createRandomAdmin(t)
	arg := &model.CreateAdminParams{
		RoleID:         1,
		Email:          newAdmin.Email,
		HashedPassword: goRandom.RandomString(10),
		Firstname:      goRandom.RandomString(10),
		Lastname: pgtype.Text{
			String: goRandom.RandomString(10),
			Valid:  true,
		},
		Active: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
		LockExpires: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}
	createdAdmin, err := repository.Create(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, createdAdmin)
}

// TestRepository_Create_InvalidParam
// Tests the Create method with invalid parameters.
func TestRepository_Create_InvalidParam(t *testing.T) {
	arg := &model.CreateAdminParams{
		RoleID:         0,
		Email:          goRandom.RandomString(10),
		HashedPassword: goRandom.RandomString(10),
		Firstname:      goRandom.RandomString(10),
		Lastname: pgtype.Text{
			String: goRandom.RandomString(10),
			Valid:  true,
		},
		Active: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
		LockExpires: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}
	createdAdmin, err := repository.Create(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, createdAdmin)
}

// TestRepository_Delete_Success
// Tests the Delete method.
func TestRepository_Delete_Success(t *testing.T) {
	randomAdmin := createRandomAdmin(t)
	deletedAdmin, err := repository.Delete(context.Background(), randomAdmin.Email)
	require.NoError(t, err)
	require.NotEmpty(t, deletedAdmin)
	compareAdmin(t, &randomAdmin, &deletedAdmin)
	deletedAdmin, err = repository.Get(context.Background(), randomAdmin.Email)
	require.Error(t, err)
	require.Empty(t, deletedAdmin)
}

// TestRepository_Delete_NonExistedAdmin
// Tests the Delete method with a non-existed admin.
func TestRepository_Delete_NonExistedAdmin(t *testing.T) {
	deletedAdmin, err := repository.Delete(context.Background(), goRandom.RandomString(20))
	require.Error(t, err)
	require.Empty(t, deletedAdmin)
}

// TestRepository_Get_Success
// Tests the Get method.
func TestRepository_Get_Success(t *testing.T) {
	randomAdmin := createRandomAdmin(t)
	fetchedAdmin, err := repository.Get(context.Background(), randomAdmin.Email)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedAdmin)
	compareAdmin(t, &randomAdmin, &fetchedAdmin)
}

// TestRepository_Get_NonExistedAdmin
// Tests the Get method with a non-existed admin.
func TestRepository_Get_NonExistedAdmin(t *testing.T) {
	randomAdmin := createRandomAdmin(t)
	deletedAdmin, err := repository.Delete(context.Background(), randomAdmin.Email)
	require.NoError(t, err)
	require.NotEmpty(t, deletedAdmin)
	fetchedAdmin, err := repository.Get(context.Background(), deletedAdmin.Email)
	require.Error(t, err)
	require.Empty(t, fetchedAdmin)
}

// TestRepository_GetList_Success
// Tests the GetList method.
func TestRepository_GetList_Success(t *testing.T) {
	var listCreatedAdmin []model.Admin
	for i := 0; i < 5; i++ {
		listCreatedAdmin = append(listCreatedAdmin, createRandomAdmin(t))
	}
	arg := &model.GetListAdminParams{
		PageSize:       5,
		CurrentPage:    1,
		OrderBy:        "admin_id",
		OrderDirection: "desc",
		Filter: &model.GetListAdminFilterParams{
			Email: pgtype.Text{String: listCreatedAdmin[0].Email, Valid: true},
		},
	}
	admins, _, err := repository.GetList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, admins)
	//require.Equal(t, len(listCreatedAdmin), len(admins))
}

// TestRepository_GetList_Empty
// Tests the GetList method with an empty list.
func TestRepository_GetList_InvalidParam(t *testing.T) {
	var listCreatedAdmin []model.Admin
	for i := 0; i < 5; i++ {
		listCreatedAdmin = append(listCreatedAdmin, createRandomAdmin(t))
	}
	arg := &model.GetListAdminParams{
		PageSize:       0,
		CurrentPage:    0,
		OrderBy:        "admin_id",
		OrderDirection: "desc",
		Filter:         &model.GetListAdminFilterParams{},
	}
	admins, _, err := repository.GetList(context.Background(), arg)
	require.NoError(t, err)
	require.Empty(t, admins)
}
