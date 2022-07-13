package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pankaj-nupay/ginAndGorm/controllers"
)

func Router() {
	router := gin.Default()
	// router.GET("/books", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "hey there, it's working fine",
	// 	})
	// })
	router.GET("/books", controllers.GetAllBooks)
	router.GET("/book/:id", controllers.GetBookById)
	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.DELETE("book/:id", controllers.DeleteBook)

	router.Run(":5000")
}
