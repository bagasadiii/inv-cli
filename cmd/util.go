package cmd

import (
	"bufio"
	"fmt"
	"library/models"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateJson() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not create file: ",err)
	}
	return filepath.Join(homeDir, ".inventory.json")
}

func UserApproval()bool{
	message := "Need to create json file (Y/N)?"
	
	r := bufio.NewReader(os.Stdin)
	var s string
	fmt.Print(message)
	s, _ = r.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	for {
		if s == "y" || s == "yes" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}
	}
}

func RemindInit(inv *models.Inventory) {
	_, err := os.Stat(CreateJson())
	if err != nil {
		fmt.Println("Please run init to create .inventory.json")
		os.Exit(1)
	} else {
		if err := inv.Load(CreateJson()); err != nil {
			log.Fatal(err)
		}
	}
}