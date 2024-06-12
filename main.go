package main

import (
	_ "database/sql"
	createcategory "erp-back/catalog/usecases/create_category/rest"
	deletecategorybyid "erp-back/catalog/usecases/delete_category_by_id/rest"
	getallmaincategories "erp-back/catalog/usecases/get_all_main_categories/rest"
	getcategorybyid "erp-back/catalog/usecases/get_category_by_id/rest"
	updatecategory "erp-back/catalog/usecases/upsert_category/rest"
	"erp-back/framework"
	"github.com/gin-gonic/gin"
	//"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {

	const (
		dbDriver = "postgres"
		//dbSource      = "postgres://postgres:postgres@localhost:31544/postgres?sslmode=disable&search_path=public"
		dbSource      = "postgres://postgres:postgres@localhost:5444/erp?sslmode=disable&search_path=public"
		serverAddress = "0.0.0.0:8799"
	)

	//m, err := migrate.New(
	//	"file://db/migration",
	//	//"postgres://postgres:postgres@localhost:31544/postgres?sslmode=disable&search_path=public")
	//	"postgres://postgres:postgres@localhost:5444/erp?sslmode=disable&search_path=public")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err := m.Up(); err != nil {
	//	log.Fatal(err)
	//}

	//conn, err := sql.Open(dbDriver, dbSource)
	//if err != nil {
	//	log.Fatal("cannot connect to db:", err)
	//}
	//defer func(conn *sql.DB) {
	//	err := conn.Close()
	//	if err != nil {
	//		log.Fatal("cannot close connection to db:", err)
	//	}
	//}(conn)
	framework.InitDatabase()

	router := gin.Default()
	createcategory.RouteCreateCategory(router)
	updatecategory.RouteUpsertCategory(router)
	deletecategorybyid.RouteDeleteCategoryById(router)
	getcategorybyid.RouteGetCategoryById(router)
	getallmaincategories.RouteGetAllMainCategories(router)
	err := router.Run(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
