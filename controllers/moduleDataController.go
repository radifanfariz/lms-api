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

type ModuleDataBody struct {
	ID             int       `json:"id"`
	GlobalID       string    `json:"global_id"`
	ModuleMetaID   int       `json:"module_meta_id"`
	PreTestMetaID  int       `json:"pretest_meta_id"`
	PreTestID      int       `json:"pretest_id"`
	MateriMetaID   int       `json:"materi_meta_id"`
	MateriID       int       `json:"materi_id"`
	PostTestMetaID int       `json:"posttest_meta_id"`
	PostTestID     int       `json:"posttest_id"`
	UserID         int       `json:"user_id"`
	GradeID        int       `json:"grade_id"`
	IsPublished    *bool     `json:"is_published"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ModuleDataCreate(ctx *gin.Context) {
	var body ModuleDataBody

	ctx.Bind(&body)

	post := models.ModuleData{
		ID:             body.ID,
		GlobalID:       body.GlobalID,
		ModuleMetaID:   body.ModuleMetaID,
		PreTestMetaID:  body.PreTestMetaID,
		PreTestID:      body.PreTestID,
		MateriMetaID:   body.MateriMetaID,
		MateriID:       body.MateriID,
		PostTestMetaID: body.PostTestMetaID,
		PostTestID:     body.PostTestID,
		UserID:         body.UserID,
		GradeID:        body.GradeID,
		IsPublished:    body.IsPublished,
		CreatedBy:      body.CreatedBy,
		CreatedAt:      body.CreatedAt,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Module Data created successfully.", "data": &post})
}

func ModuleDataFindById(ctx *gin.Context) {
	var moduleData models.ModuleData

	var findByIdResult *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.
			Preload("Metadata").
			Preload("UserData").
			Preload("PreTestMetadata").
			Preload("MateriMetadata").
			Preload("PostTestMetadata").
			Preload("PreTestData").
			Preload("MateriData").
			Preload("PostTestData").
			Where("b_ispublished = ?", true).
			Find(&moduleData, uint(id))
	} else {
		id := ctx.Param("id")
		findByIdResult = initializers.DB.
			Preload("Metadata").
			Preload("UserData").
			Preload("PreTestMetadata").
			Preload("MateriMetadata").
			Preload("PostTestMetadata").
			Preload("PreTestData").
			Preload("MateriData").
			Preload("PostTestData").
			Where("b_ispublished = ?", true).
			Find(&moduleData, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": moduleData})
}
func ModuleDataFindPaging(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("per_page"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	sort := ctx.Query("sort")
	filter := ctx.Query("filter")
	filterColumn := ctx.Query("filter_column")
	params := utils.Pagination{
		Limit:        limit,
		Page:         page,
		Sort:         sort,
		FilterColumn: filterColumn,
		Filter:       filter,
	}

	var moduleData []models.ModuleData
	res := initializers.DB.Model(moduleData).Scopes(utils.Paginate(moduleData, &params, initializers.DB)).
		Preload("Metadata").
		Preload("UserData").
		Preload("PreTestMetadata").
		Preload("MateriMetadata").
		Preload("PostTestMetadata").
		Preload("PreTestData").
		Preload("MateriData").
		Preload("PostTestData").
		Where("b_ispublished = ?", true).
		Find(&moduleData)

	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	/* this to get total all data (total all rows) and total pages in pagination */
	if params.Filter != "" && params.FilterColumn != "" {
		var moduleData []models.ModuleData
		totalRows := initializers.DB.Where(params.FilterColumn+" ILIKE ?", "%"+params.Filter+"%").Find(&moduleData).RowsAffected
		params.TotalData = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(params.Limit)))
		if params.Limit < 0 {
			params.TotalPages = 1
		} else {
			params.TotalPages = totalPages
		}
	} else {
		var moduleData []models.ModuleData
		totalRows := initializers.DB.Find(&moduleData).RowsAffected
		params.TotalData = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(params.Limit)))
		if params.Limit < 0 {
			params.TotalPages = 1
		} else {
			params.TotalPages = totalPages
		}
	}
	/*------------------------------------------------------------------------------*/

	params.Data = moduleData

	ctx.JSONP(http.StatusOK, params)
}
func ModuleDataFindAll(ctx *gin.Context) {
	var moduleData []models.ModuleData
	result := initializers.DB.
		Preload("Metadata").
		Preload("UserData").
		Preload("PreTestMetadata").
		Preload("MateriMetadata").
		Preload("PostTestMetadata").
		Preload("PreTestData").
		Preload("MateriData").
		Preload("PostTestData").
		Where("b_ispublished = ?", true).
		Find(&moduleData)

	// fmt.Println(ModuleData)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": moduleData})
}

func ModuleDataUpdate(ctx *gin.Context) {
	var body ModuleDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.ModuleData{
		ID:             body.ID,
		GlobalID:       body.GlobalID,
		ModuleMetaID:   body.ModuleMetaID,
		PreTestMetaID:  body.PreTestMetaID,
		PreTestID:      body.PreTestID,
		MateriMetaID:   body.MateriMetaID,
		MateriID:       body.MateriID,
		PostTestMetaID: body.PostTestMetaID,
		PostTestID:     body.PostTestID,
		UserID:         body.UserID,
		GradeID:        body.GradeID,
		IsPublished:    body.IsPublished,
		UpdatedBy:      body.UpdatedBy,
		UpdatedAt:      body.UpdatedAt,
	}

	var current models.ModuleData
	var findByIdResult *gorm.DB
	var findByIdResultAfterUpdate *gorm.DB

	if govalidator.IsNumeric(ctx.Param("id")) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		findByIdResult = initializers.DB.First(&current, uint(id))
	} else {
		id := ctx.Param("id")
		isPublished := true
		findByIdResult = initializers.DB.Where(models.ModuleData{IsPublished: &isPublished}).First(&current, "c_global_id = ?", id)
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Data not found.",
		})
		return
	}

	updateResult := initializers.DB.Model(&current).Omit("ID").Updates(&updates)

	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating Module Data.",
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
			"message": "Module Data not found.(Something went wrong !)",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Module Data updated successfully.", "data": &current})
}

func ModuleDataUpsert(ctx *gin.Context) {
	var body ModuleDataBody

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var current models.ModuleData
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
			upsert := models.ModuleData{
				GlobalID:       body.GlobalID,
				ModuleMetaID:   body.ModuleMetaID,
				PreTestMetaID:  body.PreTestMetaID,
				PreTestID:      body.PreTestID,
				MateriMetaID:   body.MateriMetaID,
				MateriID:       body.MateriID,
				PostTestMetaID: body.PostTestMetaID,
				PostTestID:     body.PostTestID,
				UserID:         body.UserID,
				GradeID:        body.GradeID,
				IsPublished:    body.IsPublished,
				CreatedBy:      body.CreatedBy,
				CreatedAt:      body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Module Metadata.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Module Data created successfully.", "data": &upsert})
		} else { /* create */ /* if url params is global_id then global_id automatic get from url params, so dont need to provide in JSON Body req */
			id := ctx.Param("id")
			upsert := models.ModuleData{
				GlobalID:       id,
				ModuleMetaID:   body.ModuleMetaID,
				PreTestMetaID:  body.PreTestMetaID,
				PreTestID:      body.PreTestID,
				MateriMetaID:   body.MateriMetaID,
				MateriID:       body.MateriID,
				PostTestMetaID: body.PostTestMetaID,
				PostTestID:     body.PostTestID,
				UserID:         body.UserID,
				GradeID:        body.GradeID,
				IsPublished:    body.IsPublished,
				CreatedBy:      body.CreatedBy,
				CreatedAt:      body.CreatedAt,
			}
			upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)
			if upsertResult.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error updating Module Metadata.",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Module Data created successfully.", "data": &upsert})
		}
	} else { /* update */ /* update in upsert cannot update global_id, so dont need to provide global_id in JSON Body req */
		upsert := models.ModuleData{
			ID:             current.ID,
			GlobalID:       current.GlobalID,
			ModuleMetaID:   body.ModuleMetaID,
			PreTestMetaID:  body.PreTestMetaID,
			PreTestID:      body.PreTestID,
			MateriMetaID:   body.MateriMetaID,
			MateriID:       body.MateriID,
			PostTestMetaID: body.PostTestMetaID,
			PostTestID:     body.PostTestID,
			UserID:         body.UserID,
			GradeID:        body.GradeID,
			IsPublished:    body.IsPublished,
			UpdatedBy:      body.UpdatedBy,
			UpdatedAt:      body.UpdatedAt,
		}
		upsertResult = initializers.DB.Model(&current).Omit("ID").Save(&upsert)

		if upsertResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating Module Data.",
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
				"message": "Module Data not found.(Something went wrong !)",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Module Data updated successfully.", "data": &current})
	}
}

func ModuleDataDelete(ctx *gin.Context) {
	var current models.ModuleData

	id, _ := strconv.Atoi(ctx.Param("id"))
	findByIdResult := initializers.DB.First(&current, uint(id))

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Module Data not found.",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&current, uint(id))

	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting Module Data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Module Data deleted successfully.", "deletedData": &current})

}
