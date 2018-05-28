package main

import (
	"fmt"
	"io/ioutil"
)

func GetVersion() {
	var fileVersion string
	var dbVersion string

	data, err := ioutil.ReadFile("src/versions/current")
	if err != nil {
		fmt.Println("Could not read version from panel")
		fmt.Println(err)
		fileVersion = "error"
	} else {
		fileVersion = string(data)
	}

	db, err := OpenDatabase()

	if err != nil {
		fmt.Println("Could not read version from database")
		fmt.Println(err)
		dbVersion = "error"
	} else {
		dbVersion, err = db.GetVersion()
		if err != nil {
			fmt.Println("Could not read version from database")
			fmt.Println(err)
			dbVersion = "error"
		}
	}

	fmt.Println("PufferPanel installer detected following versions: ")
	fmt.Printf("Local files: %s\n", fileVersion)
	fmt.Printf("Database: %s\n", dbVersion)
}
