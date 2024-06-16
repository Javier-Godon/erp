package rest

import (
	"erp-back/catalog/usecases/delete_category_by_id"
	"erp-back/catalog/usecases/delete_category_by_id/mediator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteDeleteCategoryById(route *gin.Engine) (routes gin.IRoutes) {
	DeleteCategoryByIdRoute := route.DELETE("/category/:id", func(context *gin.Context) {
		id := context.Param("id")

		DeleteCategoryByIdResult := mediator.Send(buildDeleteCategoryByIdToCommand(id))
		context.JSON(http.StatusOK, fromDeleteCategoryByIdResultToResponse(DeleteCategoryByIdResult))

	})

	return DeleteCategoryByIdRoute
}

func fromDeleteCategoryByIdResultToResponse(result delete_category_by_id.DeleteCategoryByIdResult) DeleteCategoryByIdResponse {
	return DeleteCategoryByIdResponse{
		Id: result.CategoryId,
	}
}

func buildDeleteCategoryByIdToCommand(id string) delete_category_by_id.DeleteCategoryByIdQuery {
	return delete_category_by_id.DeleteCategoryByIdQuery{
		Id: id,
	}
}

//https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files
//https://www.youtube.com/watch?v=BkAoT2XZM24
