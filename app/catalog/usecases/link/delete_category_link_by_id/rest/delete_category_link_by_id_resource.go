package rest

import (
	"erp-back/catalog/usecases/link/delete_category_link_by_id"
	"erp-back/catalog/usecases/link/delete_category_link_by_id/mediator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteDeleteCategoryLinkById(route *gin.Engine) (routes gin.IRoutes) {
	DeleteCategoryLinkByIdRoute := route.DELETE("/category-link/:id", func(context *gin.Context) {
		id := context.Param("id")

		DeleteCategoryLinkByIdResult := mediator.Send(buildDeleteCategoryLInkByIdToCommand(id))
		context.JSON(http.StatusOK, fromDeleteCategoryLinkByIdResultToResponse(DeleteCategoryLinkByIdResult))

	})

	return DeleteCategoryLinkByIdRoute
}

func fromDeleteCategoryLinkByIdResultToResponse(result delete_category_link_by_id.DeleteCategoryLinkByIdResult) DeleteCategoryLinkByIdResponse {
	return DeleteCategoryLinkByIdResponse{
		Id: result.CategoryLinkId,
	}
}

func buildDeleteCategoryLInkByIdToCommand(id string) delete_category_link_by_id.DeleteCategoryLinkByIdCommand {
	return delete_category_link_by_id.DeleteCategoryLinkByIdCommand{
		Id: id,
	}
}

//https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files
//https://www.youtube.com/watch?v=BkAoT2XZM24
