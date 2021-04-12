package main

import (
	"fmt"
	"github.com/fcorrionero/europcar/domain"
)

func main() {
	v, err := domain.NewVehicle("12345678901234567", "m3a", "MBMR")
	fmt.Println(err)
	fmt.Println(v)
}
