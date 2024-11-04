package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
	"gorm.io/gorm"
)

type CertificateUserDataBody struct {
	VerifiedID string    `json:"verified_id"`
	UserID     int       `json:"user_id"`
	GlobalID   string    `json:"global_id"`
	ModuleID   int       `json:"module_id"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedBy  string    `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func CertificateUserDataCreate(ctx *gin.Context) {
	var body CertificateUserDataBody

	ctx.Bind(&body)

	post := models.CertificateUserData{
		VerifiedID: body.VerifiedID,
		UserID:     body.UserID,
		GlobalID:   body.GlobalID,
		ModuleID:   body.ModuleID,
		CreatedBy:  body.CreatedBy,
		CreatedAt:  time.Now(),
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Certificate user data created successfully.", "data": &post})
}

func CertificateUserDataFindById(ctx *gin.Context) {
	var certificateUserData []models.CertificateUserData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.Order("n_id desc").Preload("UserData").Preload("ModuleData.Metadata").Preload("ModuleData.UserData").Find(&certificateUserData, uint(id))
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Certificate user data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": certificateUserData})
}
func CertificateUserDataFindByUserId(ctx *gin.Context) {
	var certificateUserData []models.CertificateUserData

	var findByIdResult *gorm.DB

	userId := ctx.Param("user_id")
	findByIdResult = initializers.DB.Preload("UserData").Preload("ModuleData.Metadata").Preload("ModuleData.UserData").Find(&certificateUserData, "n_user_id = ?", userId)

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Certificate user data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": certificateUserData})
}
func CertificateUserDataFindByIdAndUserId(ctx *gin.Context) {
	var certificateUserData models.CertificateUserData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		userId, _ := strconv.Atoi(ctx.Param("user_id"))
		findByIdResult = initializers.DB.Where("n_user_id = ?", userId).Order("n_id desc").Preload("UserData").Preload("ModuleData.Metadata").Preload("ModuleData.UserData").Find(&certificateUserData, uint(id))
	} else {
		id := ctx.Param("id")
		userId, _ := strconv.Atoi(ctx.Param("user_id"))
		findByIdResult = initializers.DB.Where("n_user_id = ?", userId).Preload("UserData").Preload("ModuleData.Metadata").Preload("ModuleData.UserData").Find(&certificateUserData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Certificate user data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": certificateUserData})
}
func CertificateUserDataFindAll(ctx *gin.Context) {
	var certificateUserData []models.CertificateUserData
	result := initializers.DB.Order("n_id desc").Preload("UserData").Preload("ModuleData.Metadata").Preload("ModuleData.UserData").Find(&certificateUserData)

	// fmt.Println(certificateUserData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": certificateUserData})
}

func CertificateUserDataUpdate(ctx *gin.Context) {
	var body CertificateUserDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.CertificateUserData{
		VerifiedID: body.VerifiedID,
		UserID:     body.UserID,
		GlobalID:   body.GlobalID,
		ModuleID:   body.ModuleID,
		UpdatedBy:  body.CreatedBy,
		UpdatedAt:  time.Now(),
	}

	var current models.CertificateUserData
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Certificate user data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Certificate user data.",
		})
		return
	}

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))
	}

	if findByIdResultAfterUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Certificate user data not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Certificate user data updated successfully.", "data": &current})
}

func CertificateUserDataUpsert(ctx *gin.Context) {
	var body CertificateUserDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.CertificateUserData
	var upsertResult *gorm.DB
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	}

	if findByIdResult.Error != nil { /* create */ /* if url params is id then global_id can be provided in JSON Body Req */
		if govalidator.IsNumeric(ctx.Param("id")) {
			upsert := models.CertificateUserData{
				VerifiedID: body.VerifiedID,
				UserID:     body.UserID,
				GlobalID:   body.GlobalID,
				ModuleID:   body.ModuleID,
				CreatedBy:  body.CreatedBy,
				CreatedAt:  time.Now(),
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Certificate Data.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Certificate Data created successfully.", "data": &upsert})
		}
	} else {
		upsert := models.CertificateUserData{
			ID:         current.ID,
			VerifiedID: body.VerifiedID,
			UserID:     body.UserID,
			GlobalID:   body.GlobalID,
			ModuleID:   body.ModuleID,
			CreatedBy:  current.CreatedBy,
			CreatedAt:  current.CreatedAt,
			UpdatedBy:  body.UpdatedBy,
			UpdatedAt:  time.Now(),
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating Certificate Data.",
			})
			return
		}

		if govalidator.IsNumeric(ctx.Param("id")) {
			id, _ := strconv.Atoi(ctx.Param("id"))
			findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))
		}

		if findByIdResultAfterUpdate.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Certificate Data not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Certificate Data updated successfully.", "data": &current})
	}
}

func CertificateUserDataDelete(ctx *gin.Context) {
	var current models.CertificateUserData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Certificate user data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Certificate user data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Certificate user data deleted successfully.", "deletedData": &current})

}
