package main

import (
	"parser/parsing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("configerr", err.Error())
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"temperature": parsing.GetTemperature(viper.GetString("url")),
		})
	})

	err := router.Run(viper.GetString("port"))
	if err != nil {
		logrus.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("/home/r_rmarsu/parser/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
