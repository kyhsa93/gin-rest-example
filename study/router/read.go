package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/go-rest-example/study/controller"
	"github.com/kyhsa93/go-rest-example/study/repository"
)

// @Tags Studies
// @Accept  json
// @Produce  json
// @Param id path string true "Study ID"
// @Success 200 {object} model.Study
// @Router /studies/{id} [get]
func ReadItem(context *gin.Context, repository repository.Repository) {
	context.JSON(200, controller.ReadItem(context.Param("id"), repository))
}

// @Tags Studies
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Studies
// @Router /studies [get]
func ReadList(context *gin.Context, repository repository.Repository) {
	context.JSON(200, controller.ReadList(repository))
}