package rest

import (
	"erp-back/catalog/usecases/update_category"
	"erp-back/catalog/usecases/update_category/mediator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteUpdateCategory(route *gin.Engine) (routes gin.IRoutes) {
	UpdateCategoryRoute := route.PUT("/category", func(context *gin.Context) {
		var request UpdateCategoryRequest
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		UpdateCategoryResult := mediator.Send(fromUpdateCategoryRequestToCommand(request))
		context.JSON(http.StatusOK, fromUpdateCategoryResultToResponse(UpdateCategoryResult))

	})

	return UpdateCategoryRoute
}

func fromUpdateCategoryResultToResponse(result update_category.UpdateCategoryResult) UpdateCategoryResponse {
	return UpdateCategoryResponse{
		Id: result.CategoryId,
	}
}

func fromUpdateCategoryRequestToCommand(request UpdateCategoryRequest) update_category.UpdateCategoryCommand {
	return update_category.UpdateCategoryCommand{
		Id:                  request.Id,
		CategoryName:        request.CategoryName,
		CategoryDescription: request.CategoryDescription,
	}
}

//https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files
//https://www.youtube.com/watch?v=BkAoT2XZM24
