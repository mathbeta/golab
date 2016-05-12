package main

import (
	"fmt"
)

type Singleton struct {
	name string
}

func (s *Singleton) GetName() string {
	return s.name
}

var s Singleton = Singleton{name: "Jack"}

func GetInstance() Singleton {
	return s
}

func main() {
	var s Singleton = GetInstance()
	fmt.Println(s.GetName())
}
