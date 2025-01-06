package rest

import (
	"erp-back/catalog/usecases/link/upsert_category_link"
	"erp-back/catalog/usecases/link/upsert_category_link/mediator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteUpsertCategoryLink(route *gin.Engine) (routes gin.IRoutes) {
	UpsertCategoryLinkRoute := route.POST("/category-link", func(context *gin.Context) {
		var request UpsertCategoryLinkRequest
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		UpsertCategoryLinkResult := mediator.Send(fromUpsertCategoryLinkRequestToCommand(request))
		context.JSON(http.StatusOK, fromUpsertCategoryLinkResultToResponse(UpsertCategoryLinkResult))

	})

	return UpsertCategoryLinkRoute
}

func fromUpsertCategoryLinkResultToResponse(result upsert_category_link.UpsertCategoryLinkResult) UpsertCategoryLinkResponse {
	return UpsertCategoryLinkResponse{
		Id: result.CategoryLinkId,
	}
}

func fromUpsertCategoryLinkRequestToCommand(request UpsertCategoryLinkRequest) upsert_category_link.UpsertCategoryLinkCommand {
	return upsert_category_link.UpsertCategoryLinkCommand{
		Id:               request.Id,
		MainCategoryId:   request.MainCategoryId,
		LinkedCategoryId: request.LinkedCategoryId,
	}
}

//https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files
//https://www.youtube.com/watch?v=BkAoT2XZM24
