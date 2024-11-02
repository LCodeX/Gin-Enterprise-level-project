package main

import (
	"fmt"
	"yky-gin/config"
	"yky-gin/db"
	"yky-gin/router"
	"yky-gin/utils/logger"
	"yky-gin/validator"
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
	validator.InitValidator()
	serverMode := configMap["env"].(string)
	r := router.Router(serverMode)
	serverConfig := configMap["server"].(map[interface{}]interface{})
	port := serverConfig["port"].(int)
	r.Run(fmt.Sprintf(":%d", port))
}
