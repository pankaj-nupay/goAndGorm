package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pankaj-nupay/ginAndGorm/models"
)

func GetAllBooks(c *gin.Context) {
	fmt.Println("hello")
	c.JSON(http.StatusOK, gin.H{
		"message": "all books data",
	})
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "the book id is: " + id,
	})
}

func CreateBook(c *gin.Context) {
	bookObj := models.Book{}
	c.Bind(&bookObj)
	// fmt.Fprintf("%v", bookObj)
	finalJson, err := json.MarshalIndent(bookObj, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func UpdateBook(c *gin.Context) {

}

func DeleteBook(c *gin.Context) {

}
