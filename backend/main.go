package main

import (
	"fmt"
	"log"

	"os"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/converter"
)

func main() {
	logger := log.New(os.Stdout, "TwoTired", log.Ldate|log.Lshortfile)
	data, err := converter.Load("./_data")
	if err != nil {
		logger.Fatal(err.Error())
	}
	for _, activity := range data {
		fmt.Println(activity.Hash())
	}
}
