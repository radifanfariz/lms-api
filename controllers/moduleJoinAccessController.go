package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
	"github.com/radifanfariz/lms-api/utils"
	"gorm.io/gorm"
)

type ModuleMetadataJoinAccssedDataPagination struct {
	*utils.Pagination
	GradeID    int `json:"grade_id,omitempty" query:"grade_id"`
	UserID     int `json:"user_id,omitempty" query:"user_id"`
	PositionID int `json:"position_id,omitempty" query:"position_id"`
}

func ModuleMetadataJoinAccessDataPaginate(value interface{}, pagination *ModuleMetadataJoinAccssedDataPagination, db *gorm.DB, ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	if pagination.GradeID != 0 || pagination.UserID != 0 || pagination.PositionID != 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Table("t_module_metadata").Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_metadata.n_id").Where("? = ANY(n_array_grade_id) OR ? = ANY(n_array_user_id) OR ? = ANY(n_array_position_id)", pagination.GradeID, pagination.UserID, pagination.PositionID)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func ModuleMeatadataJoinAccessDataFindPaging(ctx *gin.Context) {

	limit, _ := strconv.Atoi(ctx.Query("per_page"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	sort := ctx.Query("sort")
	var gradeId int
	if ctx.Query("grade_id") != "" {
		gradeId, _ = strconv.Atoi(ctx.Query("grade_id"))
	}
	var userId int
	if ctx.Query("user_id") != "" {
		userId, _ = strconv.Atoi(ctx.Query("user_id"))
	}
	var positionId int
	if ctx.Query("posiition_id") != "" {
		positionId, _ = strconv.Atoi(ctx.Query("posiition_id"))
	}

	params := ModuleMetadataJoinAccssedDataPagination{
		Pagination: &utils.Pagination{
			Limit: limit,
			Page:  page,
			Sort:  sort,
		},
		GradeID:    gradeId,
		UserID:     userId,
		PositionID: positionId,
	}
	paramsNoPageNoLimit := ModuleMetadataJoinAccssedDataPagination{
		Pagination: &utils.Pagination{
			Limit: -1,
			Page:  0,
			Sort:  sort,
		},
		GradeID:    gradeId,
		UserID:     userId,
		PositionID: positionId,
	}

	var moduleMetadata []models.ModuleMetadata
	res := initializers.DB.Scopes(ModuleMetadataJoinAccessDataPaginate(moduleMetadata, &params, initializers.DB, ctx)).Find(&moduleMetadata)

	/* for total actual total data not total data per page */
	var moduleMetadataNoPageNoLimit []models.ModuleMetadata
	resNoPageNoLimit := initializers.DB.Scopes(ModuleMetadataJoinAccessDataPaginate(moduleMetadata, &paramsNoPageNoLimit, initializers.DB, ctx)).Find(&moduleMetadataNoPageNoLimit)

	totalRows := resNoPageNoLimit.RowsAffected
	params.TotalData = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(params.Limit)))
	params.TotalPages = totalPages

	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error find all.",
		})
		return
	}

	params.Data = moduleMetadata

	ctx.JSON(http.StatusOK, params)
}
