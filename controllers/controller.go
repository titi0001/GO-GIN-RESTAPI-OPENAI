package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func ReadPDF(c *gin.Context) {
	file, _, err := c.Request.FormFile("pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao obter o arquivo PDF"})
		return
	}
	defer file.Close()

}

func AskQuestion(c *gin.Context) {
	var requestData struct {
		PDFText  string `json:"pdf_text"`
		Question string `json:"question"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
