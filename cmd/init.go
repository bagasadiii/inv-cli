package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Init() {
	ok := UserApproval()
	if !ok {
		fmt.Println("You need to create json file to store data")
		os.Exit(0)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filepath := filepath.Join(homeDir, ".inventory.json")
	_, err = os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err){
			file, err := os.Create(filepath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			fmt.Println("Create file successfull")
		} else {
			log.Fatal("Unknown error.")
		}
	} else {
		fmt.Println("Already exist")
	}
}