package router

import (
	"go-rest-example/controller"
	"go-rest-example/dto"
	"go-rest-example/repository"

	"github.com/gin-gonic/gin"
)

// @Description create study group
// @Tags Studies
// @Accept  json
// @Produce  json
// @Param study body dto.Study true "Add study"
// @Success 200
// @Router /studies [post]
func Create(context *gin.Context, repository repository.Repository) {
	var data dto.Study
	context.ShouldBindJSON(&data)
	controller.Create(&data, repository)
}
