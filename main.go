package main

import (
	setRouter "basicApi/router"
)

func main() {
	router := setRouter.SetupRouter()
	router.Run("[::]:8000")
}


