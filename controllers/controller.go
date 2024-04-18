package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/titi0001/GO-GIN-RESTAPI-OPENAI/models"
)

func ReadPDF(c *gin.Context) {

	file, _, err := c.Request.FormFile("pdf")
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao obter o arquivo PDF",
			
		})
		return
	}
	defer file.Close()

	pdfBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao ler o conteudo do arquivo PDF",
		})
		return
	}

	texto := string(pdfBytes)

	pdf := models.PDF{
		Texto: texto,
	}

	c.JSON(http.StatusOK, pdf)
}
