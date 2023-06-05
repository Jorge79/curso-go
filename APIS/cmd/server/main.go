package main

import "github.com/Jorge79/estudos-go/APIS/configs"

func main() {
	config := configs.NewConfig()
	println(config.DB_DRIVER)
}
