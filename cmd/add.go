package cmd

import (
	"flag"
	"fmt"
	"library/models"
	"log"
	"os"
)

func AddItem(inv *models.Inventory, args[]string){

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addObject := addCmd.String("Object", "", "New object")
	addPrice := addCmd.Int("Price", 0, "Price of the object")

	addCmd.Parse(args)

	if len(*addObject) == 0 {
		fmt.Println("Error: Object required")

		os.Exit(1)
	}

	inv.Add(*addObject, *addPrice)
	err := inv.Store(CreateJson())
	if err != nil {
		log.Fatal(err)
	}

	inv.Print()
}
