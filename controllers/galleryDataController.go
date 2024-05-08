package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
	"gorm.io/gorm"
)

type GalleryDataBody struct {
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	Src       string    `json:"src"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GalleryDataCreate(ctx *gin.Context) {
	var body GalleryDataBody

	ctx.Bind(&body)

	fmt.Println(&body)

	post := models.GalleryData{
		UserID:    body.UserID,
		Name:      body.Name,
		Src:       body.Src,
		CreatedBy: body.CreatedBy,
		CreatedAt: body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Gallery Data created successfully.", "data": &post})
}

func GalleryDataFindById(ctx *gin.Context) {
	var GalleryData models.GalleryData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.Preload("UserData").First(&GalleryData, uint(id))
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gallery Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": GalleryData})
}
func GalleryDataFindByUserId(ctx *gin.Context) {
	var GalleryData []models.GalleryData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("user_id")) {
		userId, _ := strconv.Atoi(ctx.Param("user_id"))
		findByIdResult = initializers.DB.Where("n_user_id = ?", uint(userId)).Preload("UserData").Find(&GalleryData)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gallery Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": GalleryData})
}
func GalleryDataFindByParams(ctx *gin.Context) {

	var body struct {
		Name string `json:"name"`
		Src  string `json:"src"`
	}

	ctx.Bind(&body)

	fmt.Println(&body)

	var GalleryData []models.GalleryData

	findByIdResult := initializers.DB.Where("c_name = ?", &body.Name).Or("c_src = ?", &body.Src).Preload("UserData").Find(&GalleryData)

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gallery Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": GalleryData})
}
func GalleryDataFindAll(ctx *gin.Context) {
	var GalleryData []models.GalleryData
	result := initializers.DB.Preload("UserData").Find(&GalleryData)

	// fmt.Println(GalleryData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": GalleryData})
}

func GalleryDataUpdate(ctx *gin.Context) {
	var body GalleryDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.GalleryData{
		UserID:    body.UserID,
		Name:      body.Name,
		Src:       body.Src,
		UpdatedBy: body.UpdatedBy,
		UpdatedAt: body.UpdatedAt,
	}

	var current models.GalleryData
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&current, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gallery Data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Gallery Data.",
		})
		return
	}

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResultAfterUpdate = initializers.DB.First(&current, "c_global_id = ?", id)
	}

	if findByIdResultAfterUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gallery Data not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Gallery Data updated successfully.", "data": &current})
}

func GalleryDataUpsert(ctx *gin.Context) {
	var body GalleryDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.GalleryData
	var upsertResult *gorm.DB
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	}

	if findByIdResult.Error != nil { /* create */ /* if url params is id then global_id can be provided in JSON Body Req */
		if govalidator.IsNumeric(ctx.Param("id")) {
			upsert := models.GalleryData{
				UserID:    body.UserID,
				Name:      body.Name,
				Src:       body.Src,
				CreatedBy: body.CreatedBy,
				CreatedAt: body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Gallery Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Gallery Data created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.GalleryData{
			ID:        current.ID,
			UserID:    body.UserID,
			Name:      body.Name,
			Src:       body.Src,
			UpdatedBy: body.UpdatedBy,
			UpdatedAt: body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating Gallery Data.",
			})
			return
		}

		if govalidator.IsNumeric(ctx.Param("id")) {
			id, _ := strconv.Atoi(ctx.Param("id"))
			findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))
		}

		if findByIdResultAfterUpdate.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Gallery Data not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Gallery Data updated successfully.", "data": &current})
	}
}

func GalleryDataDelete(ctx *gin.Context) {
	var current models.GalleryData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gallery Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Gallery Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Gallery Data deleted successfully.", "deletedData": &current})

}
