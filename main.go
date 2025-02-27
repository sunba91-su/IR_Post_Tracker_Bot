package main

import (
	"irptb/controller"
	"irptb/handler"
)

func main() {
	handler.LoadEnv()
	controller.StartService()

}
