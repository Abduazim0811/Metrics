package router

import (
	"expvar"
	"library/internal/http/api/metrics"
	"library/internal/storage"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	handler := storage.Handler()
	r := gin.Default()

	r.Use(metrics.Metrics())

	r.POST("/books", handler.CreateBook)
	r.GET("/books/:id", handler.GetByIDBook)
	r.GET("/books", handler.GetAllBooks)
	r.PUT("/books/:id", handler.UpdateBook)
	r.DELETE("/books/:id", handler.DeleteBook)
	r.GET("/debug/vars", gin.WrapH(expvar.Handler()))
	r.GET("/", handler.Check)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8888")
}
