package rest

import (
	"erp-back/catalog/usecases/upsert_category"
	"erp-back/catalog/usecases/upsert_category/mediator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteUpsertCategory(route *gin.Engine) (routes gin.IRoutes) {
	upsertCategoryRoute := route.POST("/category", func(context *gin.Context) {
		var request UpsertCategoryRequest
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		upsertCategoryResult := mediator.Send(fromUpsertCategoryRequestToCommand(request))
		context.JSON(http.StatusOK, fromUpsertCategoryResultToResponse(upsertCategoryResult))

	})

	return upsertCategoryRoute
}

func fromUpsertCategoryResultToResponse(result upsert_category.UpsertCategoryResult) UpsertCategoryResponse {
	return UpsertCategoryResponse{
		Id: result.CategoryId,
	}
}

func fromUpsertCategoryRequestToCommand(request UpsertCategoryRequest) upsert_category.UpsertCategoryCommand {
	return upsert_category.UpsertCategoryCommand{Id: request.Id}
}

//https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files
//https://www.youtube.com/watch?v=BkAoT2XZM24
