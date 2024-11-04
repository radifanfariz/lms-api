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

type CertificateMasterDataBody struct {
	Name      string    `json:"name"`
	Src       string    `json:"src"`
	IsActive  *bool     `json:"is_active"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CertificateMasterDataCreate(ctx *gin.Context) {
	var body CertificateMasterDataBody

	ctx.Bind(&body)

	post := models.CertificateMasterData{
		Name:      body.Name,
		Src:       body.Src,
		IsActive:  body.IsActive,
		CreatedBy: body.CreatedBy,
		CreatedAt: time.Now(),
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Certificate master data created successfully.", "data": &post})
}

func CertificateMasterDataFindById(ctx *gin.Context) {
	var certificaMasterData models.CertificateMasterData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.Order("n_id desc").Find(&certificaMasterData, uint(id))
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Certificate master data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": certificaMasterData})
}
func CertificateMasterDataFindByIsActive(ctx *gin.Context) {
	var certificaMasterData models.CertificateMasterData

	findByIsActiveResult := initializers.DB.Order("n_id desc").Find(&certificaMasterData, "b_isactive = ?", true)

	if findByIsActiveResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Certificate master data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": certificaMasterData})
}

func CertificateMasterDataFindAll(ctx *gin.Context) {
	var certificaMasterData []models.CertificateMasterData
	result := initializers.DB.Order("n_id desc").Find(&certificaMasterData)

	// fmt.Println(certificaMasterData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": certificaMasterData})
}

func CertificateMasterDataUpdate(ctx *gin.Context) {
	var body CertificateMasterDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.CertificateMasterData{
		Name:      body.Name,
		Src:       body.Src,
		IsActive:  body.IsActive,
		UpdatedBy: body.CreatedBy,
		UpdatedAt: time.Now(),
	}

	var current models.CertificateMasterData
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Certificate master data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Certificate master data.",
		})
		return
	}

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResultAfterUpdate = initializers.DB.First(&current, uint(id))
	}

	if findByIdResultAfterUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Certificate master data not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Certificate master data updated successfully.", "data": &current})
}

func CertificateMasterDataUpsert(ctx *gin.Context) {
	var body CertificateMasterDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.CertificateMasterData
	var upsertResult *gorm.DB
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	}

	if findByIdResult.Error != nil { /* create */ /* if url params is id then global_id can be provided in JSON Body Req */
		if govalidator.IsNumeric(ctx.Param("id")) {
			upsert := models.CertificateMasterData{
				Name:      body.Name,
				Src:       body.Src,
				IsActive:  body.IsActive,
				CreatedBy: body.CreatedBy,
				CreatedAt: time.Now(),
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
		upsert := models.CertificateMasterData{
			ID:        current.ID,
			Name:      body.Name,
			Src:       body.Src,
			IsActive:  body.IsActive,
			CreatedBy: current.CreatedBy,
			CreatedAt: current.CreatedAt,
			UpdatedBy: body.UpdatedBy,
			UpdatedAt: time.Now(),
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

func CertificateMasterDataDelete(ctx *gin.Context) {
	var current models.CertificateMasterData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Certificate master data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Certificate master data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Certificate master data deleted successfully.", "deletedData": &current})

}
