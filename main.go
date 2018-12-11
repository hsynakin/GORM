package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hsynakin/GORM/apiroots"
	"github.com/hsynakin/GORM/dbconnect"

	_ "github.com/lib/pq"
)

func main() {

	app := gin.Default()

	dbconnect.Dbase()

	api := app.Group("/api")

	apiroots.Einvoiceservices(api)

	app.Run(fmt.Sprintf(":%d", 1456))
}
