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
	s.router.POST("/", s.createAdmin)
}

// LoadAdminRoutes
// Load all admin routes
func LoadAdminRoutes(s *Server) {
	adminGroup := s.router.Group("/admin")
	{
		adminGroup.GET("/:email", s.getAdmin)
		adminGroup.GET("/", s.getListAdmin)
		adminGroup.POST("/", s.createAdmin)
		adminGroup.PUT("/", s.updateAdmin)
		adminGroup.DELETE("/:email", s.deleteAdmin)
	}
}
