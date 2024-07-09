package gin

import "github.com/gin-gonic/gin"

func (s *Server) getAdmin(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "admin",
	})
}

func (s *Server) getListAdmin(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "admin",
	})
}

func (s *Server) createAdmin(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "admin",
	})
}

func (s *Server) updateAdmin(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "admin",
	})
}

func (s *Server) deleteAdmin(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "admin",
	})
}
