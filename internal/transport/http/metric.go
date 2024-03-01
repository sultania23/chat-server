package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (h *Handler) initMetricRoutes(metrics *gin.RouterGroup) {
	metrics.GET("/prometheus", gin.WrapH(promhttp.Handler()))
}
