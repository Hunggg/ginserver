package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)
type Env struct {
	Host string
}

func (env *Env) GetConfig() {
	viper.SetConfigFile("local.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	env.Host = viper.GetString("Host")
}

func GetEmployee(g *gin.Context) {
	g.JSON(200, gin.H{
		"message" : "Hung dep trai",
	})
}
