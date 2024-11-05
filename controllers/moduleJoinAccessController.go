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

/* module metadata table join to access data table */
type ModuleMetadataJoinAccessDataPagination struct {
	*utils.Pagination
	GradeID         int    `json:"grade_id,omitempty" query:"grade_id"`
	UserID          int    `json:"user_id,omitempty" query:"user_id"`
	PositionID      int    `json:"position_id,omitempty" query:"position_id"`
	CompanyID       int    `json:"company_id,omitempty" query:"company_id"`
	LearningJourney string `json:"learning_journey,omitempty" query:"learning_journey"`
	Category        string `json:"category,omitempty" query:"category"`
}

func ModuleMetadataJoinAccessDataPaginate(value interface{}, pagination *ModuleMetadataJoinAccessDataPagination, db *gorm.DB, ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	if pagination.GradeID != 0 || pagination.UserID != 0 || pagination.PositionID != 0 || pagination.CompanyID != 0 {
		if pagination.Category != "" {
			return func(db *gorm.DB) *gorm.DB {
				return db.Select("DISTINCT *").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Table("t_module_metadata").Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_metadata.n_id left join t_module_data on t_module_data.n_module_meta_id = t_module_metadata.n_id").Where("((( ? = ANY(n_array_grade_id) OR ? = ANY(n_array_user_id) OR ? = ANY(n_array_position_id) OR n_company_id = ? ) AND c_learning_journey ILIKE ? ) OR ? = ANY(n_array_user_id) OR c_learning_journey = 'umum' ) AND c_category ILIKE ? AND b_ispublished = ?", pagination.GradeID, pagination.UserID, pagination.PositionID, pagination.CompanyID, pagination.LearningJourney, pagination.UserID, pagination.Category, true)
			}
		}
		return func(db *gorm.DB) *gorm.DB {
			return db.Select("DISTINCT *").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Table("t_module_metadata").Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_metadata.n_id left join t_module_data on t_module_data.n_module_meta_id = t_module_metadata.n_id").Where("(( ? = ANY(n_array_grade_id) OR ? = ANY(n_array_user_id) OR ? = ANY(n_array_position_id) OR n_company_id = ?) AND c_learning_journey ILIKE ? ) OR ? = ANY(n_array_user_id) OR c_learning_journey = 'umum' AND b_ispublished = ?", pagination.GradeID, pagination.UserID, pagination.PositionID, pagination.CompanyID, pagination.LearningJourney, pagination.UserID, true)
		}
	}
	if pagination.Category != "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Select("DISTINCT *").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Table("t_module_metadata").Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_metadata.n_id left join t_module_data on t_module_data.n_module_meta_id = t_module_metadata.n_id").Where("c_category ILIKE ? AND b_ispublished = ?", pagination.Category, true)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Select("DISTINCT *").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Table("t_module_metadata").Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_metadata.n_id left join t_module_data on t_module_data.n_module_meta_id = t_module_metadata.n_id").Where("b_ispublished = ?", true)
	}
}

func ModuleMetadataJoinAccessDataFindPaging(ctx *gin.Context) {

	limit, _ := strconv.Atoi(ctx.Query("per_page"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	// sort := ctx.Query("sort")
	sort := "n_seq asc"
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
	var learningJourney string
	if ctx.Query("learning_journey") != "" {
		learningJourney = ctx.Query("learning_journey")
	}
	var category string
	if ctx.Query("category") != "" {
		category = ctx.Query("category")
	}

	params := ModuleMetadataJoinAccessDataPagination{
		Pagination: &utils.Pagination{
			Limit: limit,
			Page:  page,
			Sort:  sort,
		},
		GradeID:         gradeId,
		UserID:          userId,
		PositionID:      positionId,
		CompanyID:       companyId,
		LearningJourney: learningJourney,
		Category:        category,
	}
	paramsNoPageNoLimit := ModuleMetadataJoinAccessDataPagination{
		Pagination: &utils.Pagination{
			Limit: -1,
			Page:  0,
			Sort:  sort,
		},
		GradeID:         gradeId,
		UserID:          userId,
		PositionID:      positionId,
		CompanyID:       companyId,
		LearningJourney: learningJourney,
		Category:        category,
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

/*---------------------------------------------- */

/* module data table join to access data table */
type ModuleDataJoinAccessDataPagination struct {
	*utils.Pagination
	GradeID         int    `json:"grade_id,omitempty" query:"grade_id"`
	UserID          int    `json:"user_id,omitempty" query:"user_id"`
	PositionID      int    `json:"position_id,omitempty" query:"position_id"`
	CompanyID       int    `json:"company_id,omitempty" query:"company_id"`
	LearningJourney string `json:"learning_journey,omitempty" query:"learning_journey"`
	Category        string `json:"category,omitempty" query:"category"`
}

func ModuleDataJoinAccessDataPaginate(value interface{}, pagination *ModuleDataJoinAccessDataPagination, db *gorm.DB, ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	if pagination.GradeID != 0 || pagination.UserID != 0 || pagination.PositionID != 0 || pagination.CompanyID != 0 || pagination.LearningJourney != "" {
		/* considered dangerous */
		// return func(db *gorm.DB) *gorm.DB {
		// 	return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Table("t_module_metadata").Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_data.n_module_meta_id left join t_module_metadata on t_module_metadata.n_id = t_module_data.n_module_meta_id").Where(pagination.GetFilterColumn()+" ILIKE ? AND "+"( ? = ANY(n_array_grade_id) OR ? = ANY(n_array_user_id) OR ? = ANY(n_array_position_id) OR ? = n_company_id = ?" + "OR c_category ILIKE ?", "%"+pagination.GetFilter()+"%", pagination.GradeID, pagination.UserID, pagination.PositionID, pagination.CompanyID, pagination.Category)
		// }
		/*--------------------------*/
		if pagination.Category != "" {
			return func(db *gorm.DB) *gorm.DB {
				return db.Select("DISTINCT *").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Table("t_module_data").Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_data.n_module_meta_id left join t_module_metadata on t_module_metadata.n_id = t_module_data.n_module_meta_id").Where("((( ? = ANY(n_array_grade_id) OR ? = ANY(n_array_user_id) OR ? = ANY(n_array_position_id) OR n_company_id = ? ) AND c_learning_journey ILIKE ? ) OR ? = ANY(n_array_user_id) OR c_learning_journey = 'umum' ) AND c_category ILIKE ? AND b_ispublished = ?", pagination.GradeID, pagination.UserID, pagination.PositionID, pagination.CompanyID, pagination.LearningJourney, pagination.UserID, pagination.Category, true)
			}
		}
		return func(db *gorm.DB) *gorm.DB {
			return db.Select("DISTINCT *").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Table("t_module_data").Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_data.n_module_meta_id left join t_module_metadata on t_module_metadata.n_id = t_module_data.n_module_meta_id").Where("(( ? = ANY(n_array_grade_id) OR ? = ANY(n_array_user_id) OR ? = ANY(n_array_position_id) OR n_company_id = ?) AND c_learning_journey ILIKE ? ) OR ? = ANY(n_array_user_id) OR c_learning_journey = 'umum' AND b_ispublished = ?", pagination.GradeID, pagination.UserID, pagination.PositionID, pagination.CompanyID, pagination.LearningJourney, pagination.UserID, true)
		}
	}
	if pagination.Category != "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Select("DISTINCT *").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Table("t_module_data").Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_data.n_module_meta_id left join t_module_metadata on t_module_metadata.n_id = t_module_data.n_module_meta_id").Where("c_category ILIKE ? AND b_ispublished = ?", pagination.Category, true)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Select("DISTINCT *").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Table("t_module_data").Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_data.n_module_meta_id left join t_module_metadata on t_module_metadata.n_id = t_module_data.n_module_meta_id").Where("b_ispublished = ?", true)
	}
}

func ModuleDataJoinAccessDataFindPaging(ctx *gin.Context) {

	limit, _ := strconv.Atoi(ctx.Query("per_page"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	// sort := ctx.Query("sort")
	sort := "n_seq asc"
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
	var learningJourney string
	if ctx.Query("learning_journey") != "" {
		learningJourney = ctx.Query("learning_journey")
	}
	var category string
	if ctx.Query("category") != "" {
		category = ctx.Query("category")
	}

	params := ModuleDataJoinAccessDataPagination{
		Pagination: &utils.Pagination{
			Limit: limit,
			Page:  page,
			Sort:  sort,
		},
		GradeID:         gradeId,
		UserID:          userId,
		PositionID:      positionId,
		CompanyID:       companyId,
		LearningJourney: learningJourney,
		Category:        category,
	}
	paramsNoPageNoLimit := ModuleDataJoinAccessDataPagination{
		Pagination: &utils.Pagination{
			Limit: -1,
			Page:  0,
			Sort:  sort,
		},
		GradeID:         gradeId,
		UserID:          userId,
		PositionID:      positionId,
		CompanyID:       companyId,
		LearningJourney: learningJourney,
		Category:        category,
	}

	var moduleData []models.ModuleData
	res := initializers.DB.Scopes(ModuleDataJoinAccessDataPaginate(moduleData, &params, initializers.DB, ctx)).
		Preload("Metadata").
		Preload("UserData").
		Preload("PreTestMetadata").
		Preload("MateriMetadata").
		Preload("PostTestMetadata").
		Preload("PreTestData").
		Preload("MateriData").
		Preload("PostTestData").
		Find(&moduleData)

	/* for total actual total data not total data per page */
	var moduleDataNoLimit []models.ModuleData
	resNoPageNoLimit := initializers.DB.Scopes(ModuleDataJoinAccessDataPaginate(moduleData, &paramsNoPageNoLimit, initializers.DB, ctx)).
		Preload("Metadata").
		Preload("UserData").
		Preload("PreTestMetadata").
		Preload("MateriMetadata").
		Preload("PostTestMetadata").
		Preload("PreTestData").
		Preload("MateriData").
		Preload("PostTestData").
		Find(&moduleDataNoLimit)

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

	params.Data = moduleData

	ctx.JSON(http.StatusOK, params)
}

/*---------------------------------------------- */

func ModuleMetadataJoinAccessDataFindByIdWithParams(ctx *gin.Context) {
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

	learningJourney := ctx.Query("learning_journey")
	globalId := ctx.Param("id")

	var moduleMetadata models.ModuleMetadata
	if gradeId != 0 || userId != 0 || positionId != 0 || companyId != 0 || globalId != "" || learningJourney != "" {
		findByIdWithParamsModuleJoinAccess := initializers.DB.Table("t_module_metadata").
			Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_metadata.n_id").
			Where("t_module_data.c_global_id = ?", globalId).
			Where("(( ? = ANY(n_array_grade_id) OR ? = ANY(n_array_user_id) OR ? = ANY(n_array_position_id) OR n_company_id = ?) AND c_learning_journey ILIKE ? ) OR ? = ANY(n_array_user_id) OR c_learning_journey = 'umum' AND b_ispublished = ?", gradeId, userId, positionId, companyId, learningJourney, userId, true).
			Find(&moduleMetadata)

		if findByIdWithParamsModuleJoinAccess.RowsAffected <= 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Data not found !",
			})
			return
		}

		if findByIdWithParamsModuleJoinAccess.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error find ByIdWithParams.",
			})
			return
		}

	} else {
		findByIdWithParamsModuleJoinAccess := initializers.DB.Table("t_module_metadata").Where("t_access_data.c_global_id = ? AND b_ispublished = ?", globalId, true).
			Find(&moduleMetadata)

		if findByIdWithParamsModuleJoinAccess.RowsAffected <= 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Data not found !",
			})
			return
		}

		if findByIdWithParamsModuleJoinAccess.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error find ByIdWithParams.",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": moduleMetadata})
}
func ModuleDataJoinAccessDataFindByIdWithParams(ctx *gin.Context) {
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

	learningJourney := ctx.Query("learning_journey")
	globalId := ctx.Param("id")

	var moduleData models.ModuleData
	if gradeId != 0 || userId != 0 || positionId != 0 || companyId != 0 || learningJourney != "" {

		findByIdWithParamsModuleJoinAccess := initializers.DB.
			Table("t_module_data").
			Joins("left join t_access_data on t_access_data.n_module_meta_id = t_module_data.n_module_meta_id left join t_module_metadata on t_module_metadata.n_id = t_module_data.n_module_meta_id").
			Where("t_module_data.c_global_id = ?", globalId).
			Where("(( ? = ANY(n_array_grade_id) OR ? = ANY(n_array_user_id) OR ? = ANY(n_array_position_id) OR n_company_id = ?) AND c_learning_journey ILIKE ? ) OR ? = ANY(n_array_user_id) OR c_learning_journey = 'umum' AND b_ispublished = ?", gradeId, userId, positionId, companyId, learningJourney, userId, true).
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
			Find(&moduleData)

		if findByIdWithParamsModuleJoinAccess.RowsAffected <= 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Data not found !",
			})
			return
		}

		if findByIdWithParamsModuleJoinAccess.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error find ByIdWithParams.",
			})
			return
		}

	} else {
		findByIdWithParamsModuleJoinAccess := initializers.DB.
			Table("t_module_data").Where("t_module_data.c_global_id = ? AND b_ispublished = ?", globalId, true).
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
			Find(&moduleData)

		if findByIdWithParamsModuleJoinAccess.RowsAffected <= 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Data not found !",
			})
			return
		}

		if findByIdWithParamsModuleJoinAccess.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error find ByIdWithParams.",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": moduleData})
}
