package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/hsynakin/GORM/apicontrollers"
)

func Einvoiceservices(api *gin.RouterGroup) {
	api.GET("/getUserFromTaxRegistrationNo/:id", apicontrollers.TaxNoResults)
	api.GET("/getUserFirstCreationTime/:date", apicontrollers.FirstCreate)
	api.GET("/getLastFirstCreationTime/:date", apicontrollers.LastFirstCreation)

	api.POST("/PostXmLFile", apicontrollers.PostXml)
	api.POST("/einvoiceusers", apicontrollers.POSTEInvoiceUsers)
}
