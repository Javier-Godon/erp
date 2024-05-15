package mediator

import (
	"erp-back/catalog/usecases/upsert_category"
	_ "erp-back/framework"
	"reflect"
	"sync"
)

type key[TRequest any, TResult any] struct {
}

func init() {
	registeredHandlers := sync.Map{}
	k := key[upsert_category.UpsertCategoryCommand, upsert_category.UpsertCategoryResult]{}
	registeredHandlers.LoadOrStore(reflect.TypeOf(k), upsert_category.UpsertCategoryHandlerImpl{})
}
