package apicontrollers

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hsynakin/GORM/dbconnect"
	"github.com/hsynakin/GORM/models"
)

var xmlUsers models.Users

func TaxNoResults(c *gin.Context) {
	var TaxNo = c.Params.ByName("id")

	var eIUsers = models.EInvoiceUsers{}
	dbconnect.DB.First(&eIUsers, TaxNo)

	if eIUsers.Identifier != "" {
		c.JSON(http.StatusOK, eIUsers)
	} else {
		c.JSON(http.StatusNotFound, "Kayıt Bulunamadı.Kayıtları güncelleyip tekrar deneyiniz.")

	}
}

func FirstCreate(c *gin.Context) {
	var FirstTime = []models.EInvoiceUsers{}
	var Variable = c.Params.ByName("date")

	convertTime, err := time.Parse("2006-01-02", Variable)
	convertTime = convertTime.Local().Add(time.Hour * -3)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Hatalı bir zaman girilmiştir. Zaman formatı Yıl-Ay-Gün şeklinde olması gerekmektedir. Ör: 2018-11-28")
		return
	}

	dbconnect.DB.Where("FirstCreationTime>= ?", Variable).Find(&FirstTime)
	if len(FirstTime) > 0 {
		c.JSON(http.StatusOK, FirstTime) //return values user
	} else {
		c.JSON(http.StatusNotFound, "kayıt bulunamadı")
	}

}
func LastFirstCreation(c *gin.Context) {
	var user = models.EInvoiceUsers{}
	dbconnect.DB.Order("first_creation_time desc").First(&user)
	year := user.FirstCreationTime.Format("2006-01-02")

	if user.Identifier != "" {
		c.JSON(http.StatusOK, year)
	} else {
		c.JSON(http.StatusNotFound, "Kayıt bulunamadı.")
	}
}
func POSTEInvoiceUsers(c *gin.Context) {
	log.Println("burdayımmm")
	var user = models.EInvoiceUsers{}

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Hatalı gönderim yaptınız.")
		return
	}

	t2 := models.EInvoiceUsers{}
	dbconnect.DB.First(&t2, user.Identifier)
	if t2.Identifier == "" {
		dbconnect.DB.Create(&user)
		c.JSON(http.StatusCreated, "E-Fatura kullanıcısı başarıyla eklenmiştir.")
	} else {
		c.JSON(http.StatusBadRequest, "E-Fatura kullanıcısı zaten var.")
	}
}

func PostXml(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, "Dosya Yüklenemedi")
		return
	}

	filePath := "./upload"
	exist, _ := FileOrDirectoryExists(filePath)

	if exist == false {
		log.Println("FilePath: ", filePath)
		os.Mkdir(filePath, 0700)
	}

	filePath += "/" + file.Filename

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, filePath)

	InputFunction := updateUsersArray(true)

	if InputFunction {
		c.JSON(http.StatusOK, "Xml başarılı bir şekilde güncellenmiştir.")
	} else {
		c.JSON(http.StatusNotFound, "Dosya bulunamadı.")
	}
}

func updateUsersArray(last bool) bool {
	xmlFile, err := os.Open("./upload/Users.xml")
	if err != nil {
		return false
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &xmlUsers)
	var xmlNewUsers = []models.EInvoiceUsers{}

	xmlNewUsers = xmlUsers.Users

	//var addUser = models.EInvoiceUsers{}
	for _, t := range xmlNewUsers {
		t2 := models.EInvoiceUsers{}
		dbconnect.DB.First(&t2, t.Identifier)
		if dbconnect.DB.NewRecord(&t2) {
			dbconnect.DB.Create(&t)
		}
	}
	return true
}
func FileOrDirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
