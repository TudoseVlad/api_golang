package main

import (
	"fmt"
	"golang_api/src/api"
	"golang_api/src/data"
)

func main() {
	fmt.Println("Rest API")
	data.InitData()
	api.InitApi()
}
