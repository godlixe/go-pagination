package routes

import (
	"go-pagination/controller"

	"github.com/gin-gonic/gin"
)

func SiteRoutes(router *gin.Engine, siteController controller.SiteController) {
	siteRoutes := router.Group("/site")
	{
		siteRoutes.GET("/count", siteController.GetTotalRows)
		siteRoutes.GET("", siteController.GetAllSites)
		siteRoutes.GET("/paginated-offset", siteController.GetAllSitesPaginatedOffset)
		siteRoutes.GET("/paginated-seek", siteController.GetAllSitesPaginatedSeek)

	}
}
