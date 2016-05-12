package main

import (
	"fmt"
	"cigol.org/minise/utils"
)

func main() {
	config := utils.GetConfig("config.ini")
	for k, v := range config {
		fmt.Println(k, "=", v)
	}
}
