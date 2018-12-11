package dbconnect

import (
	"fmt"
	"log"

	"github.com/hsynakin/GORM/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	appDBName     = "gorm"
	appDBHost     = "localhost"
	appDBUserName = "postgres"
	appDBPassword = "gorm"
	DB            *gorm.DB
)

func Dbase() {
	cnnString := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s sslmode=disable",
		appDBHost,
		appDBUserName,
		appDBPassword,
		appDBName,
	)

	var err error
	DB, err = gorm.Open("postgres", cnnString)
	if err != nil {
		log.Println("DB Error", err)
	}
	DB.CreateTable(models.EInvoiceUsers{})
	log.Println("DB Connected")
}
