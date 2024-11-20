package gin

// loadRoutes
// Load all routes of application
func (s *Server) loadRoutes() {
	LoadDefaultAdminRoutes(s)
	LoadAdminRoutes(s)
}

// LoadDefaultAdminRoutes
// Provide the way to create default super admin
func LoadDefaultAdminRoutes(s *Server) {
	s.router.POST("/default_admin", s.handler.adminHandler.CreateAdmin)
}

// LoadAdminRoutes
// Load all admin routes
func LoadAdminRoutes(s *Server) {
	adminGroup := s.router.Group("/admin")
	{
		adminGroup.GET("/:email", s.handler.adminHandler.GetAdmin)
		adminGroup.GET("/", s.handler.adminHandler.GetListAdmin)
		adminGroup.POST("/", s.handler.adminHandler.CreateAdmin)
		adminGroup.PUT("/", s.handler.adminHandler.UpdateAdmin)
		adminGroup.DELETE("/:email", s.handler.adminHandler.DeleteAdmin)
	}
}
