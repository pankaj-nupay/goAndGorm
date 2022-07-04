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
	// router.GET("/book/:id", getBookById)
	// router.POST("/book", createBook)
	// router.PUT("/book/:id", updateBook)
	// router.DELETE("book/:id", deleteBook)

	router.Run(":5000")
}
