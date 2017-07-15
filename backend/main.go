package main

import (
	"fmt"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/converter"
)

func main() {
	data, err := converter.Load("./_data")
	if err != nil {
		panic(err)
	}
	for _, activity := range data {
		fmt.Println(activity.Hash())
	}
}
