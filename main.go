package main

import (
	"github.com/gin-gonic/gin"
	"github.com/titi0001/GO-GIN-RESTAPI-OPENAI/routes"
)

func main() {

    r := gin.Default()
    routes.ConfigurarPDFRoutes(r)

    r.StaticFile("/", "./views/index.html")

    r.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{
            "message": "Rota não encontrada",
        })
    })

    r.Run(":9000")
}
