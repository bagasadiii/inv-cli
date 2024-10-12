package main

import (
	"flag"
	"library/cmd"
	"library/models"
	"os"
)

func main() {
	inv := &models.Inventory{}
	flag.Parse()

	switch flag.Arg(0) {
	case "init":
		cmd.Init()
	case "add":
		cmd.RemindInit(inv)
		cmd.AddItem(inv, os.Args[2:])
	case "delete":
		cmd.RemindInit(inv)
		cmd.DeleteObject(inv, os.Args[2:])
	}
}