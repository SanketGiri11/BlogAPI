package controllers

import (
	"blogapi/database"
	"blogapi/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddComments func creates a record in DB for each new Comment
func AddComments(context *gin.Context) {
	var comments models.Comments
	if err := context.ShouldBindJSON(&comments); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
	}

	if err := database.DB.Create(&comments).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Payload",
		})
		context.Abort()
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Comments successfully added",
	})
}

//Get all comments at a time
func AllComments(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	limits := 5
	offset := (page - 1) * limits
	var total int64
	var getComments []models.Comments
	database.DB.Preload("Posts").Offset(offset).Limit(limits).Find(&getComments)
	database.DB.Model(&models.Posts{}).Count(&total)

	context.JSON(http.StatusOK, gin.H{
		"data": getComments,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"last_page": float64(int(total) / limits),
		},
	})

}

//Get a specific comment details using the Id
func DetailComments(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var getComments models.Comments
	database.DB.Where("id=?", id).Preload("Posts").First(&getComments)

	context.JSON(http.StatusOK, gin.H{
		"data": getComments,
	})
}

//Update a comment using Id
func UpdateComments(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	comments := models.Comments{
		Id: uint(id),
	}

	if err := context.ShouldBindJSON(&comments); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
	}

	database.DB.Model(&comments).Updates(comments)
	context.JSON(http.StatusOK, comments)
}

//Deletion of comment using the Id
func DeleteCommnets(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	comments := models.Comments{
		Id: uint(id),
	}
	deletequery := database.DB.Delete(&comments)
	if errors.Is(deletequery.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Record Not Found"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Comments deleted successfully"})

}
