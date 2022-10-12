package main

import (
	"github.com/dimasokta14/gin-rest/connections"
	"github.com/dimasokta14/gin-rest/services/product"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".local.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	handler := connections.Init(dbUrl)

	product.RegisterRoutes(router, handler)

	router.Run(port)
}
