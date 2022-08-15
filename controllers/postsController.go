package controllers

import (
	"blogapi/auth"
	"blogapi/database"
	"blogapi/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CreatePosts func creates a record in DB for each new POSTS
func CreatePosts(context *gin.Context) {
	var posts models.Posts
	if err := context.ShouldBindJSON(&posts); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
	}

	if err := database.DB.Create(&posts).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Payload",
		})
		context.Abort()
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Blog successfully created",
	})
}

//Get all posts at a time
func AllPosts(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	limits := 5
	offset := (page - 1) * limits
	var total int64
	var getPosts []models.Posts
	database.DB.Preload("User").Offset(offset).Limit(limits).Find(&getPosts)
	database.DB.Model(&models.Posts{}).Count(&total)

	context.JSON(http.StatusOK, gin.H{
		"data": getPosts,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"last_page": float64(int(total) / limits),
		},
	})

}

//Get a specific POST details using the Id
func DetailPosts(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var getPosts models.Posts
	database.DB.Where("id=?", id).Preload("User").First(&getPosts)

	context.JSON(http.StatusOK, gin.H{
		"data": getPosts,
	})
}

//Update a POST using the Id
func UpdatePosts(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	posts := models.Posts{
		Id: uint(id),
	}

	if err := context.ShouldBindJSON(&posts); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
	}

	database.DB.Model(&posts).Updates(posts)
	context.JSON(http.StatusOK, posts)
}

//Get the POST details form jwt
func UniquePosts(context *gin.Context) {
	cookie, _ := context.Cookie("jwt")
	id := auth.ValidateToken(cookie)
	var posts []models.Posts
	database.DB.Model(&posts).Where("user_id=?", id).Preload("User").Find(&posts)
	context.JSON(http.StatusOK, posts)

}

//Deletion of a specific POST from DB using the Id
func DeletePosts(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	posts := models.Posts{
		Id: uint(id),
	}
	deletequery := database.DB.Delete(&posts)
	if errors.Is(deletequery.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Record Not Found"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})

}
