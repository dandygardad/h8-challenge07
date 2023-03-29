package routes

import (
	"challenge07/controller"
	"challenge07/service"
	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine, app service.ServiceInterface) {
	handler := controller.NewController(app)
	api := router.Group("/books")
	{
		api.POST("", handler.CreateBook)       // Buat buku
		api.GET("", handler.GetAllBooks)       // Ambil semua buku
		api.GET("/:id", handler.GetBook)       // Ambil buku berdasarkan id
		api.PUT("/:id", handler.UpdateBook)    // Update buku
		api.DELETE("/:id", handler.DeleteBook) // Hapus buku
	}
}
