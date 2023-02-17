package main

import (
	"main/ToDoList/conf"
	"main/ToDoList/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	r.Run(conf.HttpPort)
}
