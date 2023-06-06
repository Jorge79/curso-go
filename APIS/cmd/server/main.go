package main

import "github.com/Jorge79/estudos-go/APIS/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
