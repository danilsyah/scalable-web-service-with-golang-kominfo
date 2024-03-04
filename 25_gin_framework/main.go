package main

import "25_gin_framework/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
