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

//AddTags func creates a record in DB for each new tags creation
func AddTags(context *gin.Context) {
	var tags models.Tags
	if err := context.ShouldBindJSON(&tags); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
	}

	if err := database.DB.Create(&tags).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Payload",
		})
		context.Abort()
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Tags successfully added",
	})
}

//Get all tags from DB
func AllTags(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	limits := 5
	offset := (page - 1) * limits
	var total int64
	var getTags []models.Tags
	database.DB.Preload("Posts").Offset(offset).Limit(limits).Find(&getTags)
	database.DB.Model(&models.Tags{}).Count(&total)

	context.JSON(http.StatusOK, gin.H{
		"data": getTags,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"last_page": float64(int(total) / limits),
		},
	})

}

//Gets a specific Tag info according to the Id
func DetailTags(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var getTags models.Tags
	database.DB.Where("id=?", id).Preload("Posts").First(&getTags)

	context.JSON(http.StatusOK, gin.H{
		"data": getTags,
	})
}

//Update a existing tag using the Id
func UpdateTags(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	tags := models.Tags{
		Id: uint(id),
	}

	if err := context.ShouldBindJSON(&tags); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
	}

	database.DB.Model(&tags).Updates(tags)
	context.JSON(http.StatusOK, tags)
}

//Deletion of tags using the id
func DeleteTags(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	tags := models.Tags{
		Id: uint(id),
	}
	deletequery := database.DB.Delete(&tags)
	if errors.Is(deletequery.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Record Not Found"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Tags deleted successfully"})

}
