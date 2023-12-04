package migrate

import (
	"github.com/radifanfariz/lms-api/initializers"
	"github.com/radifanfariz/lms-api/models"
)

func inti() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.ModuleMetadata{}, &models.ModuleData{}, &models.PreTestMetadata{}, &models.PreTestData{})
}
