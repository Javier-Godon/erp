package main

import (
	// sql "database/sql"
	deleteCategoryById "erp-back/catalog/usecases/delete_category_by_id/rest"
	getAllmainCategories "erp-back/catalog/usecases/get_all_main_categories/rest"
	getCategoryById "erp-back/catalog/usecases/get_category_by_id/rest"
	deleteCategoryLink "erp-back/catalog/usecases/link/delete_category_link_by_id/rest"
	getCategoryLinkById "erp-back/catalog/usecases/link/get_category_link_by_id/rest"
	upsertCategoryLink "erp-back/catalog/usecases/link/upsert_category_link/rest"
	updateCategory "erp-back/catalog/usecases/upsert_category/rest"
	"erp-back/framework"

	"github.com/gin-gonic/gin"

	"log"

	migrate "github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	framework.ReadConfig()

	dbSource := framework.AppConfig.PostgresUrl.URI
	serverPort := framework.AppConfig.ServerPort.PORT

	// const (
	// 	// dbDriver = "postgres"
	// 	// dbSource = "postgres://postgres:postgres@blue-env.com:31543/blue_postgres?sslmode=disable&search_path=public"
	// 	// dbSource      = "postgres://postgres:postgres@localhost:5444/erp?sslmode=disable&search_path=public"
	// 	serverAddress = "0.0.0.0:8799"
	// )

	m, err := migrate.New(
		"file://db/migration",
		dbSource)
	// "postgres://postgres:postgres@blue-env.com:31543/blue_postgres?sslmode=disable&search_path=public")
	if err != nil {
		log.Fatal("Failed to initialize migration: ", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("No changes to apply in migration: ", err)
	}

	// conn, err := sql.Open(dbDriver, dbSource)
	// if err != nil {
	// 	log.Fatal("cannot connect to db:", err)
	// }
	// defer func(conn *sql.DB) {
	// 	err := conn.Close()
	// 	if err != nil {
	// 		log.Fatal("cannot close connection to db:", err)
	// 	}
	// }(conn)

	framework.InitDatabase()

	router := gin.Default()
	//createcategory.RouteCreateCategory(router)
	updateCategory.RouteUpsertCategory(router)
	deleteCategoryById.RouteDeleteCategoryById(router)
	getCategoryById.RouteGetCategoryById(router)
	getAllmainCategories.RouteGetAllMainCategories(router)
	deleteCategoryLink.RouteDeleteCategoryLinkById(router)
	upsertCategoryLink.RouteUpsertCategoryLink(router)
	getCategoryLinkById.RouteGetCategoryLinkById(router)

	err = router.Run("0.0.0.0:" + serverPort)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
