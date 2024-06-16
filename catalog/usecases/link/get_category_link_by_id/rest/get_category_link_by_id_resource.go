package rest

import (
	"erp-back/catalog/usecases/link/get_category_link_by_id"
	"erp-back/catalog/usecases/link/get_category_link_by_id/mediator"
	"erp-back/shared_kernell/catalog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteGetCategoryLinkById(route *gin.Engine) (routes gin.IRoutes) {
	GetCategoryLinkByIdRoute := route.GET("/category-link/:id", func(context *gin.Context) {
		id := context.Param("id")

		GetCategoryLinkByIdResult := mediator.Send(buildGetCategoryLinkByIdToQuery(id))
		context.JSON(http.StatusOK, fromGetCategoryLinkByIdResultToResponse(GetCategoryLinkByIdResult))

	})

	return GetCategoryLinkByIdRoute
}

func fromGetCategoryLinkByIdResultToResponse(result get_category_link_by_id.GetCategoryLinkByIdResult) catalog.CategoryLinkResponse {
	return catalog.CategoryLinkResponse{
		Id:               result.Id,
		MainCategoryId:   result.MainCategoryId,
		LinkedCategoryId: result.LinkedCategoryId,
	}
}

func buildGetCategoryLinkByIdToQuery(id string) get_category_link_by_id.GetCategoryLinkByIdQuery {
	return get_category_link_by_id.GetCategoryLinkByIdQuery{
		Id: id,
	}
}

//https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files
//https://www.youtube.com/watch?v=BkAoT2XZM24
