package handlers

import (
	"net/http"

	"github.com/CharlesChou03/_git/amt.git/models"
	modelsReq "github.com/CharlesChou03/_git/amt.git/models/request"
	modelsRes "github.com/CharlesChou03/_git/amt.git/models/response"
	"github.com/CharlesChou03/_git/amt.git/services"
	"github.com/gin-gonic/gin"
)

// GetTutorHandler get tutor information
// @Summary get tutor information
// @Description get tutor information
// @Tags Tutor Information
// @Accept json
// @Produce json
// @Param tutor path string true "tutor"
// @Success 200 {object} modelsRes.GetTutorRes "ok"
// @Failure 204 "no content"
// @Failure 400 "bad request"
// @Router /api/tutor/{tutor} [get]
func GetTutorHandler(c *gin.Context) {
	req := modelsReq.GetTutorReq{}
	res := modelsRes.GetTutorRes{}
	c.BindUri(&req)
	statusCode, resp, err := services.GetTutor(&req, &res)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, resp)
	case 204:
		c.JSON(http.StatusNoContent, err)
	case 400:
		c.JSON(http.StatusBadRequest, err)
	default:
		c.JSON(http.StatusInternalServerError, models.InternalServerError)
	}
}
