package cmd

import (
	"flag"
	"library/models"
	"log"
)

func DeleteObject(inv *models.Inventory, args []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.Int("id", 0, "ID to be deleted")

	deleteCmd.Parse(args)
	err := inv.Delete(*deleteID)
	if err != nil {
		log.Fatal(err)
	}

	inv.Print()
}