package controller

import (
	"go-pagination/entity"
	"go-pagination/service"
	"go-pagination/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SiteController interface {
	GetTotalRows(ctx *gin.Context)
	GetAllSites(ctx *gin.Context)
	GetAllSitesPaginatedOffset(ctx *gin.Context)
	GetAllSitesPaginatedSeek(ctx *gin.Context)
}

type siteController struct {
	siteService service.SiteService
}

func NewSiteController(ss service.SiteService) SiteController {
	return &siteController{
		siteService: ss,
	}
}

func (c *siteController) GetTotalRows(ctx *gin.Context) {
	result, err := c.siteService.GetTotalRows(ctx.Request.Context())
	if err != nil {
		res := utils.BuildErrorResponse("Failed to get row count", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *siteController) GetAllSites(ctx *gin.Context) {
	result, err := c.siteService.GetAllSites(ctx.Request.Context())
	if err != nil {
		res := utils.BuildErrorResponse("Failed to get rows", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *siteController) GetAllSitesPaginatedOffset(ctx *gin.Context) {

	var paginationRequest entity.Pagination
	// pick
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page <= 0 {
		page = 1
	}

	paginationRequest.Page = page

	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit < 10 {
		limit = 10
	}

	paginationRequest.Limit = limit

	result, err := c.siteService.GetAllSitesPaginatedOffset(ctx.Request.Context(), paginationRequest)
	if err != nil {
		res := utils.BuildErrorResponse("Failed to get rows", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *siteController) GetAllSitesPaginatedSeek(ctx *gin.Context) {

	var paginationRequest entity.Pagination
	// pick
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page <= 0 {
		page = 1
	}

	paginationRequest.Page = page

	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit < 10 {
		limit = 10
	}

	paginationRequest.Limit = limit

	result, err := c.siteService.GetAllSitesPaginatedSeek(ctx.Request.Context(), paginationRequest)
	if err != nil {
		res := utils.BuildErrorResponse("Failed to get rows", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}
