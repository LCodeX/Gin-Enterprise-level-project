package main

import (
	"fmt"
	"yky-gin/config"
	"yky-gin/db"
	"yky-gin/router"
	"yky-gin/utils/logger"
	"yky-gin/validator"

	"github.com/gin-gonic/gin/binding"
	playgroundValidator "github.com/go-playground/validator/v10"
)

func main() {
	configMap, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("Load config file error: %v", err))
	}
	dbConfig := configMap["database"].(map[interface{}]interface{})
	db.InitDB(dbConfig)
	redisConfig := configMap["redis"].(map[interface{}]interface{})
	db.InitRedis(redisConfig)
	logger.InitLogger()
	if v, ok := binding.Validator.Engine().(*playgroundValidator.Validate); ok {
		v.RegisterValidation("starts_with_letter", validator.StartsWithLetter)
		v.RegisterValidation("zh_phone_number", validator.IsPhoneNumber)
	}
	r := router.Router()
	serverConfig := configMap["server"].(map[interface{}]interface{})
	port := serverConfig["port"].(int)
	r.Run(fmt.Sprintf(":%d", port))
}
