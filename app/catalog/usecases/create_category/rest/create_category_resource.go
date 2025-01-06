package rest

import (
	"erp-back/catalog/usecases/create_category"
	"erp-back/catalog/usecases/create_category/mediator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteCreateCategory(route *gin.Engine) (routes gin.IRoutes) {
	CreateCategoryRoute := route.POST("/category", func(context *gin.Context) {
		var request CreateCategoryRequest
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		CreateCategoryResult := mediator.Send(fromCreateCategoryRequestToCommand(request))
		context.JSON(http.StatusOK, fromCreateCategoryResultToResponse(CreateCategoryResult))

	})

	return CreateCategoryRoute
}

func fromCreateCategoryResultToResponse(result create_category.CreateCategoryResult) CreateCategoryResponse {
	return CreateCategoryResponse{
		Id: result.CategoryId,
	}
}

func fromCreateCategoryRequestToCommand(request CreateCategoryRequest) create_category.CreateCategoryCommand {
	return create_category.CreateCategoryCommand{
		Id:                  request.Id,
		CategoryName:        request.CategoryName,
		CategoryDescription: request.CategoryDescription,
	}
}

//https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files
//https://www.youtube.com/watch?v=BkAoT2XZM24
