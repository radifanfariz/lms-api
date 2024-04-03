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
	CompanyID  int `json:"company_id,omitempty" query:"company_id"`
}

func ModuleMetadataJoinAccessDataPaginate(value interface{}, pagination *ModuleMetadataJoinAccssedDataPagination, db *gorm.DB, ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	if pagination.GradeID != 0 || pagination.UserID != 0 || pagination.PositionID != 0 || pagination.CompanyID != 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Table("t_module_metadata").Joins("inner join t_access_data on t_access_data.n_module_meta_id = t_module_metadata.n_id").Where(pagination.GetFilterColumn()+" ILIKE ? AND "+"( ? = ANY(n_array_grade_id) OR ? = ANY(n_array_user_id) OR ? = ANY(n_array_position_id) OR ? = ANY(n_array_company_id))", "%"+pagination.GetFilter()+"%", pagination.GradeID, pagination.UserID, pagination.PositionID, pagination.CompanyID)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Where(pagination.GetFilterColumn()+" ILIKE ? ", "%"+pagination.GetFilter()+"%")
	}
}

func ModuleMetadataJoinAccessDataFindPaging(ctx *gin.Context) {

	limit, _ := strconv.Atoi(ctx.Query("per_page"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	sort := ctx.Query("sort")
	var filterColumn string = "c_learning_journey"
	if ctx.Query("filter_columns") != "" {
		filterColumn = ctx.Query("filter_columns")
	}
	var filter string
	if ctx.Query("filter") != "" {
		filter = ctx.Query("filter")
	}
	var gradeId int
	if ctx.Query("grade_id") != "" {
		gradeId, _ = strconv.Atoi(ctx.Query("grade_id"))
	}
	var userId int
	if ctx.Query("user_id") != "" {
		userId, _ = strconv.Atoi(ctx.Query("user_id"))
	}
	var positionId int
	if ctx.Query("position_id") != "" {
		positionId, _ = strconv.Atoi(ctx.Query("position_id"))
	}
	var companyId int
	if ctx.Query("company_id") != "" {
		companyId, _ = strconv.Atoi(ctx.Query("company_id"))
	}

	params := ModuleMetadataJoinAccssedDataPagination{
		Pagination: &utils.Pagination{
			Limit:        limit,
			Page:         page,
			Sort:         sort,
			FilterColumn: filterColumn,
			Filter:       filter,
		},
		GradeID:    gradeId,
		UserID:     userId,
		PositionID: positionId,
		CompanyID:  companyId,
	}
	paramsNoPageNoLimit := ModuleMetadataJoinAccssedDataPagination{
		Pagination: &utils.Pagination{
			Limit:        -1,
			Page:         0,
			Sort:         sort,
			FilterColumn: filterColumn,
			Filter:       filter,
		},
		GradeID:    gradeId,
		UserID:     userId,
		PositionID: positionId,
		CompanyID:  companyId,
	}

	var moduleMetadata []models.ModuleMetadata
	res := initializers.DB.Scopes(ModuleMetadataJoinAccessDataPaginate(moduleMetadata, &params, initializers.DB, ctx)).Find(&moduleMetadata)

	/* for total actual total data not total data per page */
	var moduleMetadataNoLimit []models.ModuleMetadata
	resNoPageNoLimit := initializers.DB.Scopes(ModuleMetadataJoinAccessDataPaginate(moduleMetadata, &paramsNoPageNoLimit, initializers.DB, ctx)).Find(&moduleMetadataNoLimit)

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

func ModuleMetadataJoinAccessDataFindFirst(ctx *gin.Context) {
	var gradeId int
	if ctx.Query("grade_id") != "" {
		gradeId, _ = strconv.Atoi(ctx.Query("grade_id"))
	}
	var userId int
	if ctx.Query("user_id") != "" {
		userId, _ = strconv.Atoi(ctx.Query("user_id"))
	}
	var positionId int
	if ctx.Query("position_id") != "" {
		positionId, _ = strconv.Atoi(ctx.Query("position_id"))
	}
	var companyId int
	if ctx.Query("company_id") != "" {
		companyId, _ = strconv.Atoi(ctx.Query("company_id"))
	}
	globalId := ctx.Param("id")

	var moduleMetadata models.ModuleMetadata
	if gradeId != 0 || userId != 0 || positionId != 0 || companyId != 0 || globalId != "" {

		findFirstModuleJoinAccess := initializers.DB.Table("t_module_metadata").Joins("inner join t_access_data on t_access_data.n_module_meta_id = t_module_metadata.n_id").Where("t_access_data.c_global_id = ? AND "+"( ? = ANY(n_array_grade_id) OR ? = ANY(n_array_user_id) OR ? = ANY(n_array_position_id) OR ? = ANY(n_array_company_id))", globalId, gradeId, userId, positionId, companyId).First(&moduleMetadata)
		if findFirstModuleJoinAccess.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error find first.",
			})
			return
		}

	}

	ctx.JSON(http.StatusOK, gin.H{"data": moduleMetadata})
}
func ModuleDataJoinAccessDataFindFirst(ctx *gin.Context) {
	var gradeId int
	if ctx.Query("grade_id") != "" {
		gradeId, _ = strconv.Atoi(ctx.Query("grade_id"))
	}
	var userId int
	if ctx.Query("user_id") != "" {
		userId, _ = strconv.Atoi(ctx.Query("user_id"))
	}
	var positionId int
	if ctx.Query("position_id") != "" {
		positionId, _ = strconv.Atoi(ctx.Query("position_id"))
	}
	var companyId int
	if ctx.Query("company_id") != "" {
		companyId, _ = strconv.Atoi(ctx.Query("company_id"))
	}
	globalId := ctx.Param("id")

	var moduleData models.ModuleData
	if gradeId != 0 || userId != 0 || positionId != 0 || companyId != 0 || globalId != "" {

		findFirstModuleJoinAccess := initializers.DB.
			Preload("Metadata").
			Preload("UserData").
			Preload("PreTestMetadata").
			Preload("MateriMetadata").
			Preload("PostTestMetadata").
			Preload("PreTestData").
			Preload("PreTestData.Metadata").
			Preload("MateriData").
			Preload("MateriData.Metadata").
			Preload("PostTestData").
			Preload("PostTestData.Metadata").
			Table("t_module_data").Joins("inner join t_access_data on t_access_data.n_module_meta_id = t_module_data.n_module_meta_id").Where("t_access_data.c_global_id = ? AND "+"( ? = ANY(n_array_grade_id) OR ? = ANY(n_array_user_id) OR ? = ANY(n_array_position_id) OR ? = ANY(n_array_company_id))", globalId, gradeId, userId, positionId, companyId).First(&moduleData)

		if findFirstModuleJoinAccess.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error find first.",
			})
			return
		}

	}

	ctx.JSON(http.StatusOK, gin.H{"data": moduleData})
}
