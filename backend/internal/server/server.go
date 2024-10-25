package server

import (
	"database/sql"

	"github.com/arizdn234/Marlin-Express/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	// Root endpoint
	r.GET("/", handler.HomeHandler)

	// Tambahkan endpoint lain di sini
	// misalnya r.GET("/packages", handler.GetPackages(db))
}
