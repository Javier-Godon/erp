package rest

import (
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

	})
	return upsertCategoryRoute
}

//https://stackoverflow.com/questions/42967235/golang-gin-gonic-split-routes-into-multiple-files
//https://www.youtube.com/watch?v=BkAoT2XZM24
