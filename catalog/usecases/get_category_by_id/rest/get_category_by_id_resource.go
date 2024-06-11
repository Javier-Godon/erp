package rest

import (
	"erp-back/catalog/usecases/get_category_by_id"
	"erp-back/catalog/usecases/get_category_by_id/mediator"
	"erp-back/shared_kernell/catalog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteGetCategoryById(route *gin.Engine) (routes gin.IRoutes) {
	GetCategoryByIdRoute := route.GET("/category/:id", func(context *gin.Context) {
		id := context.Param("id")

		GetCategoryByIdResult := mediator.Send(buildGetCategoryByIdToCommand(id))
		context.JSON(http.StatusOK, fromGetCategoryByIdResultToResponse(GetCategoryByIdResult))

	})

	return GetCategoryByIdRoute
}

func fromGetCategoryByIdResultToResponse(result get_category_by_id.GetCategoryByIdResult) catalog.CategoryResponse {
	return catalog.CategoryResponse{
		Id:                  result.Id,
		CategoryName:        result.Name,
		CategoryDescription: result.Description,
	}
}

func buildGetCategoryByIdToCommand(id string) get_category_by_id.GetCategoryByIdCommand {
	return get_category_by_id.GetCategoryByIdCommand{
		Id: id,
	}
}

//https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files
//https://www.youtube.com/watch?v=BkAoT2XZM24
