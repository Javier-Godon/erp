package rest

import (
	"erp-back/catalog/usecases/get_all_main_categories"
	"erp-back/catalog/usecases/get_all_main_categories/mediator"
	"erp-back/shared_kernell/catalog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteGetAllMainCategories(route *gin.Engine) (routes gin.IRoutes) {
	GetAllMainCategoriesRoute := route.GET("/categories/main", func(context *gin.Context) {

		GetAllMainCategoriesResult := mediator.Send(buildGetAllMainCategoriesToCommand())
		context.JSON(http.StatusOK, mapResultToResponse(GetAllMainCategoriesResult))

	})

	return GetAllMainCategoriesRoute
}

func mapResultToResponse(categoryResults []get_all_main_categories.GetAllMainCategoriesResult) (categoryResponses []catalog.CategoryResponse) {
	for _, catalogCategory := range categoryResults {
		categoryResponses = append(categoryResponses, fromGetAllMainCategoriesResultToResponse(catalogCategory))
	}
	return categoryResponses
}

func fromGetAllMainCategoriesResultToResponse(result get_all_main_categories.GetAllMainCategoriesResult) catalog.CategoryResponse {
	return catalog.CategoryResponse{
		Id:                  result.Id,
		CategoryName:        result.Name,
		CategoryDescription: result.Description,
	}
}

func buildGetAllMainCategoriesToCommand() get_all_main_categories.GetAllMainCategoriesCommand {
	return get_all_main_categories.GetAllMainCategoriesCommand{}
}

//https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files
//https://www.youtube.com/watch?v=BkAoT2XZM24
