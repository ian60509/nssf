package sbi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/nssf/internal/logger"
	"github.com/free5gc/nssf/internal/sbi/processor"
	"github.com/free5gc/nssf/internal/util"
	"github.com/free5gc/openapi/models"
)

func (s *Server) getNsSelectionRoutes() []Route {
	return []Route{
		{
			"Health Check",
			strings.ToUpper("Get"),
			"/",
			func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"status": "Service Available"})
			},
		},

		{
			"NSSelectionGet",
			strings.ToUpper("Get"),
			"/network-slice-information",
			s.NetworkSliceInformationGet,
		},
	}
}

func (s *Server) NetworkSliceInformationGet(c *gin.Context) {
	logger.NsselLog.Infof("Handle NSSelectionGet")

	var query processor.NetworkSliceInformationGetQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		logger.NsselLog.Errorf("BindQuery failed: %+v", err)
		problemDetail := &models.ProblemDetails{
			Title:         "Malformed Request",
			Status:        http.StatusBadRequest,
			Detail:        err.Error(),
			Instance:      "",
			InvalidParams: util.BindErrorInvalidParamsMessages(err),
		}
		util.GinProblemJson(c, problemDetail)
		return
	}

	// query := c.Request.URL.Query()
	s.Processor().NSSelectionSliceInformationGet(c, query)
}
