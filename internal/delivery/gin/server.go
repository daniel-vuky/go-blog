package gin

import (
	"context"
	"fmt"
	adminHandler "github.com/daniel-vuky/go-blog/internal/delivery/gin/handler/admin"
	adminService "github.com/daniel-vuky/go-blog/internal/service/admin"
	adminStorage "github.com/daniel-vuky/go-blog/internal/storage/admin"
	"github.com/daniel-vuky/go-blog/pkg/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
)

// service
// Struct to hold all application services
type handlers struct {
	adminHandler *adminHandler.Handler
}

// Server
// Struct to hold all server configuration
type Server struct {
	config  *config.Config
	router  *gin.Engine
	handler *handlers
}

// NewServer
// Create new server instance
// @return *Server, error
func NewServer() (*Server, error) {
	loadedConfig, err := config.LoadConfig("./")
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	connPool, err := loadedConfig.ConnectToPgxPool()
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}
	listHandlers := &handlers{
		adminHandler: adminHandler.NewHandler(
			adminService.NewService(
				adminStorage.NewAdminRepository(connPool),
			),
		),
	}
	newServer := &Server{
		config:  loadedConfig,
		router:  gin.Default(),
		handler: listHandlers,
	}
	newServer.loadRoutes()

	return newServer, nil
}

// Start
// Starting the server with graceful shutdown
// @param ctx context.Context
// @param waitGroup *errgroup.Group
// @return error
func (s *Server) Start(ctx context.Context, waitGroup *errgroup.Group) error {
	server := &http.Server{
		Addr:    s.config.GetServerAddress(),
		Handler: s.router,
	}
	waitGroup.Go(func() error {
		err := server.ListenAndServe()
		return err
	})
	waitGroup.Go(func() error {
		<-ctx.Done()
		err := server.Shutdown(ctx)
		return err
	})
	return nil
}
