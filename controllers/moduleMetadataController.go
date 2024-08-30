package controllers

import (
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
	"github.com/radifanfariz/lms-api/utils"
	"gorm.io/gorm"
)

type ModuleMetadataBody struct {
	ID              int       `json:"id"`
	GlobalID        string    `json:"global_id"`
	Name            string    `json:"name" binding:"required"`
	Description     string    `json:"description" binding:"required"`
	Src             string    `json:"src"`
	LearningJourney string    `json:"learning_journey" binding:"required"`
	Category        string    `json:"category" binding:"required"`
	MaxMonth        int       `json:"max_month" binding:"required"`
	Seq             int       `json:"seq"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedBy       string    `json:"updated_by"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func ModuleMetadataCreate(ctx *gin.Context) {
	var body ModuleMetadataBody

	ctx.Bind(&body)

	post := models.ModuleMetadata{
		GlobalID:        body.GlobalID,
		Name:            body.Name,
		Description:     body.Description,
		Src:             body.Src,
		LearningJourney: body.LearningJourney,
		Category:        body.Category,
		MaxMonth:        body.MaxMonth,
		Seq:             body.Seq,
		CreatedBy:       body.CreatedBy,
		CreatedAt:       body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Module Metadata created successfully.", "data": &post})
}

func ModuleMetadataCreateAndAccessCreate(ctx *gin.Context) {
	var body struct {
		ModuleMetadataBody ModuleMetadataBody `json:"module_metadata_body"`
		AccessDataBody     AccessDataBody     `json:"access_data_body"`
	}

	ctx.Bind(&body)

	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	post := models.ModuleMetadata{
		GlobalID:        body.ModuleMetadataBody.GlobalID,
		Name:            body.ModuleMetadataBody.Name,
		Description:     body.ModuleMetadataBody.Description,
		Src:             body.ModuleMetadataBody.Src,
		LearningJourney: body.ModuleMetadataBody.LearningJourney,
		Category:        body.ModuleMetadataBody.Category,
		MaxMonth:        body.ModuleMetadataBody.MaxMonth,
		Seq:             body.ModuleMetadataBody.Seq,
		CreatedBy:       body.ModuleMetadataBody.CreatedBy,
		CreatedAt:       body.ModuleMetadataBody.CreatedAt,
	}
	result := tx.Model(models.ModuleMetadata{}).Create(&post)

	if result.Error != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating Module Metadata.",
		})
		return
	}

	var moduleMetadata models.ModuleMetadata

	findByIdResult := tx.Model(models.ModuleMetadata{}).First(&moduleMetadata, "c_global_id = ?", body.ModuleMetadataBody.GlobalID)

	if findByIdResult.Error != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Metadata not found.",
		})
		return
	}

	postAccess := models.AccessData{
		GlobalID:        moduleMetadata.GlobalID,
		ModuleMetaID:    moduleMetadata.ID,
		ArrayGradeID:    body.AccessDataBody.ArrayGradeID,
		ArrayUserID:     body.AccessDataBody.ArrayUserID,
		ArrayPositionID: body.AccessDataBody.ArrayPositionID,
		ArrayCompanyID:  body.AccessDataBody.ArrayCompanyID,
		CompanyID:       body.AccessDataBody.CompanyID,
		CreatedBy:       body.ModuleMetadataBody.CreatedBy,
		CreatedAt:       body.ModuleMetadataBody.CreatedAt,
	}
	resultAccess := tx.Model(models.AccessData{}).Create(&postAccess)

	if resultAccess.Error != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating Access data.",
		})
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{"message": "Module Metadata and Access Data was created successfully."})
}

func ModuleMetadataFindById(ctx *gin.Context) {
	var moduleMetadata models.ModuleMetadata

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&moduleMetadata, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&moduleMetadata, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Metadata not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": moduleMetadata})
}

func ModuleMetadataFindPaging(ctx *gin.Context) {

	limit, _ := strconv.Atoi(ctx.Query("per_page"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	sort := ctx.Query("sort")
	learningJourney := ctx.Query("learning_journey")
	params := utils.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

	var moduleMetadata []models.ModuleMetadata
	res := initializers.DB.Scopes(utils.Paginate(moduleMetadata, &params, initializers.DB)).Where("c_learning_journey ILIKE ?", "%"+learningJourney+"%").Find(&moduleMetadata)

	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	/* this to get total all data (total all rows) and total pages in pagination */
	if learningJourney != "" {
		var userData []models.ModuleMetadata
		totalRows := initializers.DB.Where("c_learning_journey"+" ILIKE ?", "%"+learningJourney+"%").Find(&userData).RowsAffected
		params.TotalData = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(params.Limit)))
		if params.Limit < 0 {
			params.TotalPages = 1
		} else {
			params.TotalPages = totalPages
		}
	} else {
		var userData []models.ModuleMetadata
		totalRows := initializers.DB.Find(&userData).RowsAffected
		params.TotalData = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(params.Limit)))
		if params.Limit < 0 {
			params.TotalPages = 1
		} else {
			params.TotalPages = totalPages
		}
	}
	/*------------------------------------------------------------------------------*/

	params.Data = moduleMetadata

	ctx.JSON(http.StatusOK, params)
}

func ModuleMetadataFindAll(ctx *gin.Context) {
	var moduleMetadata []models.ModuleMetadata
	result := initializers.DB.Find(&moduleMetadata)

	// fmt.Println(moduleMetadata)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": moduleMetadata})
}

func ModuleMetadataUpdate(ctx *gin.Context) {
	var body ModuleMetadataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.ModuleMetadata{
		GlobalID:        body.GlobalID,
		Name:            body.Name,
		Description:     body.Description,
		Src:             body.Src,
		LearningJourney: body.LearningJourney,
		Category:        body.Category,
		MaxMonth:        body.MaxMonth,
		Seq:             body.Seq,
		UpdatedBy:       body.UpdatedBy,
		UpdatedAt:       body.UpdatedAt,
	}

	var current models.ModuleMetadata
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
			"message": "Module Metadata not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Module Metadata.",
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
			"message": "Module Metadata not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Module Metadata updated successfully.", "data": &current})
}

func ModuleMetadataUpsert(ctx *gin.Context) {
	var body ModuleMetadataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.ModuleMetadata
	var upsertResult *gorm.DB
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.First(&current, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil { /* create */ /* if url params is id then global_id can be provided in JSON Body Req */
		if govalidator.IsNumeric(ctx.Param("id")) {
			upsert := models.ModuleMetadata{
				GlobalID:        body.GlobalID,
				Name:            body.Name,
				Description:     body.Description,
				Src:             body.Src,
				LearningJourney: body.LearningJourney,
				Category:        body.Category,
				MaxMonth:        body.MaxMonth,
				Seq:             body.Seq,
				CreatedBy:       body.CreatedBy,
				CreatedAt:       body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Module Metadata.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Module Metadata created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is global_id then global_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.ModuleMetadata{
				GlobalID:        id,
				Name:            body.Name,
				Description:     body.Description,
				Src:             body.Src,
				LearningJourney: body.LearningJourney,
				Category:        body.Category,
				MaxMonth:        body.MaxMonth,
				Seq:             body.Seq,
				CreatedBy:       body.CreatedBy,
				CreatedAt:       body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Module Metadata.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Module Metadata created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.ModuleMetadata{
			ID:              current.ID,
			GlobalID:        current.GlobalID,
			Name:            body.Name,
			Description:     body.Description,
			Src:             body.Src,
			LearningJourney: body.LearningJourney,
			Category:        body.Category,
			MaxMonth:        body.MaxMonth,
			Seq:             body.Seq,
			CreatedBy:       current.CreatedBy,
			CreatedAt:       current.CreatedAt,
			UpdatedBy:       body.UpdatedBy,
			UpdatedAt:       body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating Module Metadata.",
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
				"message": "Module Metadata not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Module Metadata updated successfully.", "data": &current})
	}
}

func ModuleMetadataDelete(ctx *gin.Context) {
	var current models.ModuleMetadata

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Metadata not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Module Metadata.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Module Metadata deleted successfully.", "deletedData": &current})

}
