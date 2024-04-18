package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/titi0001/GO-GIN-RESTAPI-OPENAI/controllers"
)

// ConfigurarPDFRoutes configura as rotas relacionadas ao PDF.
func ConfigurarPDFRoutes(router *gin.Engine) {
    pdf := router.Group("/pdf")
    {
        pdf.POST("/readpdf", controllers.ReadPDF)
    }
}
