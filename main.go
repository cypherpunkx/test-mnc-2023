package main

import (
	"gorm-practice/config"
	"gorm-practice/delivery"
)

func init() {
	config.InitiliazeConfig()
	config.ConnectDB()
	config.SyncDB()
}

func main() {
	delivery.Server().Run()
}
