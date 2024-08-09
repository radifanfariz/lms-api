package controllers

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
	"gorm.io/gorm"
)

// considered deprecated
func DuplicateLearningModuleAllData(ctx *gin.Context) {
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
		ctx.JSON(http.StatusOK, gin.H{"data": moduleData})
		return
	}

	// Note the use of tx as the database handle once you are within a transaction
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	globalId := uuid.New()

	moduleData.Metadata.GlobalID = globalId.String()
	if err := tx.Model(models.ModuleMetadata{}).
		Omit("ID").
		Create(&moduleData.Metadata).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	moduleData.PreTestMetadata.GlobalID = globalId.String()
	if err := tx.Model(models.PreTestMetadata{}).
		Omit("ID").
		Create(&moduleData.PreTestMetadata).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	moduleData.PreTestData.GlobalID = globalId.String()
	if err := tx.Model(models.PreTestData{}).
		Omit("ID").
		Create(&moduleData.PreTestData).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	moduleData.MateriMetadata.GlobalID = globalId.String()
	if err := tx.Omit("ID").
		Create(&moduleData.MateriMetadata).
		Model(models.MateriData{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	for _, element := range moduleData.MateriData {
		element.GlobalID = globalId.String()
		if err := tx.Model(models.MateriData{}).
			Omit("ID").
			Create(&element).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
			})
			return
		}
	}

	moduleData.PostTestMetadata.GlobalID = globalId.String()
	if err := tx.Model(models.PostTestMetadata{}).
		Omit("ID").
		Create(&moduleData.PostTestMetadata).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	moduleData.PostTestData.GlobalID = globalId.String()
	if err := tx.Model(models.PostTestData{}).
		Omit("ID").
		Create(&moduleData.PostTestData).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{"message": "Learning Module duplicated successfully."})
}

//

func DuplicateLearningModuleAllDataV2(ctx *gin.Context) {
	var moduleData models.ModuleData

	var findByIdResult *gorm.DB

	var body struct {
		ModuleMetadataBody ModuleMetadataBody `json:"module_metadata_body"`
		AccessDataBody     AccessDataBody     `json:"access_data_body"`
	}

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
			Preload("AccessData").
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
			Preload("AccessData").
			Where("b_ispublished = ?", true).
			Find(&moduleData, "c_global_id = ?", id)
	}

	if findByIdResult.RowsAffected <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Module Data not found !",
		})
		return
	}

	if findByIdResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong !.",
		})
		return
	}

	// Note the use of tx as the database handle once you are within a transaction
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	globalId := uuid.New()

	var moduleMetadataDuplicated = moduleData.Metadata
	moduleMetadataDuplicated.GlobalID = globalId.String()
	if err := tx.Model(models.ModuleMetadata{}).
		Omit("ID").
		Create(&moduleMetadataDuplicated).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	var pretestMetadataDuplicated = moduleData.PreTestMetadata
	pretestMetadataDuplicated.GlobalID = globalId.String()
	if err := tx.Model(models.PreTestMetadata{}).
		Omit("ID").
		Create(&pretestMetadataDuplicated).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	var pretestDataDuplicated = moduleData.PreTestData
	pretestDataDuplicated.GlobalID = globalId.String()
	pretestDataDuplicated.PreTestMetaID = pretestMetadataDuplicated.ID
	if err := tx.Model(models.PreTestData{}).
		Omit("ID").
		Create(&pretestDataDuplicated).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	var materiMetadataDuplicated = moduleData.MateriMetadata
	materiMetadataDuplicated.GlobalID = globalId.String()
	if err := tx.Omit("ID").
		Create(&materiMetadataDuplicated).
		Model(models.MateriData{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	var materiDataDuplicated = moduleData.MateriData
	for _, element := range materiDataDuplicated {
		element.GlobalID = globalId.String()
		element.MateriMetaID = materiMetadataDuplicated.ID
		if err := tx.Model(models.MateriData{}).
			Omit("ID").
			Create(&element).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
			})
			return
		}
	}

	var posttestMetadataDuplicated = moduleData.PostTestMetadata
	posttestMetadataDuplicated.GlobalID = globalId.String()
	if err := tx.Model(models.PostTestMetadata{}).
		Omit("ID").
		Create(&posttestMetadataDuplicated).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	var posttestDataDuplicated = moduleData.PostTestData
	posttestDataDuplicated.GlobalID = globalId.String()
	posttestDataDuplicated.PostTestMetaID = posttestMetadataDuplicated.ID
	if err := tx.Model(models.PostTestData{}).
		Omit("ID").
		Create(&posttestDataDuplicated).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	//create access data

	accessDataDuplicated := moduleData.AccessData
	if accessDataDuplicated == nil || len(accessDataDuplicated) <= 0 {
		accessDataDuplicated = []models.AccessData{
			{
				GlobalID:        globalId.String(),
				ModuleMetaID:    moduleMetadataDuplicated.ID,
				ArrayGradeID:    body.AccessDataBody.ArrayGradeID,
				ArrayPositionID: body.AccessDataBody.ArrayPositionID,
				ArrayUserID:     body.AccessDataBody.ArrayUserID,
				CompanyID:       body.AccessDataBody.CompanyID,
				CreatedBy:       body.AccessDataBody.CreatedBy,
				CreatedAt:       body.AccessDataBody.CreatedAt,
				UpdatedBy:       body.AccessDataBody.UpdatedBy,
				UpdatedAt:       body.AccessDataBody.UpdatedAt,
			},
		}
	} else {
		var accessDataDuplicatedTemp []models.AccessData
		for _, element := range accessDataDuplicated {
			accessDataDuplicatedTemp = append(accessDataDuplicatedTemp, models.AccessData{
				GlobalID:        globalId.String(),
				ModuleMetaID:    moduleMetadataDuplicated.ID,
				ArrayGradeID:    element.ArrayGradeID,
				ArrayPositionID: element.ArrayPositionID,
				ArrayUserID:     element.ArrayUserID,
				CompanyID:       element.CompanyID,
				CreatedBy:       element.CreatedBy,
				CreatedAt:       element.CreatedAt,
				UpdatedBy:       element.UpdatedBy,
				UpdatedAt:       element.UpdatedAt,
			})
		}
		accessDataDuplicated = accessDataDuplicatedTemp
	}

	if err := tx.Model(models.AccessData{}).
		Omit("ID").
		Create(&accessDataDuplicated).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
		})
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{"message": "Learning Module duplicated successfully."})
}

// considered deprecated
// func DuplicateLearningModuleAllDataV2(ctx *gin.Context) {
// 	var moduleData models.ModuleData

// 	var findByIdResult *gorm.DB

// 	var globalId string
// 	if ctx.Query("new_global_id") != "" {
// 		globalId = ctx.Query("new_global_id")
// 	}

// 	if govalidator.IsNumeric(ctx.Param("id")) {
// 		id, _ := strconv.Atoi(ctx.Param("id"))
// 		findByIdResult = initializers.DB.
// 			Preload("Metadata").
// 			Preload("UserData").
// 			Preload("PreTestMetadata").
// 			Preload("MateriMetadata").
// 			Preload("PostTestMetadata").
// 			Preload("PreTestData").
// 			Preload("MateriData").
// 			Preload("PostTestData").
// 			Where("b_ispublished = ?", true).
// 			Find(&moduleData, uint(id))
// 	} else {
// 		id := ctx.Param("id")
// 		findByIdResult = initializers.DB.
// 			Preload("Metadata").
// 			Preload("UserData").
// 			Preload("PreTestMetadata").
// 			Preload("MateriMetadata").
// 			Preload("PostTestMetadata").
// 			Preload("PreTestData").
// 			Preload("MateriData").
// 			Preload("PostTestData").
// 			Where("b_ispublished = ?", true).
// 			Find(&moduleData, "c_global_id = ?", id)
// 	}

// 	if findByIdResult.Error != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Module Data not found.",
// 		})
// 		ctx.JSON(http.StatusOK, gin.H{"data": moduleData})
// 		return
// 	}

// 	// Note the use of tx as the database handle once you are within a transaction
// 	tx := initializers.DB.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if err := tx.Error; err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
// 		})
// 		return
// 	}

// 	moduleData.Metadata.GlobalID = globalId
// 	if err := tx.Model(models.ModuleMetadata{}).
// 		Omit("ID").
// 		Create(&moduleData.Metadata).Error; err != nil {
// 		tx.Rollback()
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
// 		})
// 		return
// 	}

// 	moduleData.PreTestMetadata.GlobalID = globalId
// 	if err := tx.Model(models.PreTestMetadata{}).
// 		Omit("ID").
// 		Create(&moduleData.PreTestMetadata).Error; err != nil {
// 		tx.Rollback()
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
// 		})
// 		return
// 	}

// 	moduleData.PreTestData.GlobalID = globalId
// 	if err := tx.Model(models.PreTestData{}).
// 		Omit("ID").
// 		Create(&moduleData.PreTestData).Error; err != nil {
// 		tx.Rollback()
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
// 		})
// 		return
// 	}

// 	moduleData.MateriMetadata.GlobalID = globalId
// 	if err := tx.Omit("ID").
// 		Create(&moduleData.MateriMetadata).
// 		Model(models.MateriData{}).Error; err != nil {
// 		tx.Rollback()
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
// 		})
// 		return
// 	}

// 	for _, element := range moduleData.MateriData {
// 		element.GlobalID = globalId
// 		if err := tx.Model(models.MateriData{}).
// 			Omit("ID").
// 			Create(&element).Error; err != nil {
// 			tx.Rollback()
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
// 			})
// 			return
// 		}
// 	}

// 	moduleData.PostTestMetadata.GlobalID = globalId
// 	if err := tx.Model(models.PostTestMetadata{}).
// 		Omit("ID").
// 		Create(&moduleData.PostTestMetadata).Error; err != nil {
// 		tx.Rollback()
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
// 		})
// 		return
// 	}

// 	moduleData.PostTestData.GlobalID = globalId
// 	if err := tx.Model(models.PostTestData{}).
// 		Omit("ID").
// 		Create(&moduleData.PostTestData).Error; err != nil {
// 		tx.Rollback()
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error creating post. (It is likely that you will have to complete all parts of the module!)",
// 		})
// 		return
// 	}

// 	tx.Commit()

// 	ctx.JSON(http.StatusOK, gin.H{"message": "Learning Module duplicated successfully."})
// }
