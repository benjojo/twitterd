package main

import (
	"os/user"
	"time"
)

func CheckIfUserIsRoot() {
	CurrentUsr, _ := user.Current()
	if CurrentUsr.Uid == "0" {
		for i := 0; i < 10; i++ {
			for hashes := 0; hashes < i; hashes++ {
				Logger.Print("#")
			}
			Logger.Println("!!WARNING!! YOU ARE RUNNING THIS AS ROOT, THIS ISNT A GOOD IDEA")
		}
		Logger.Println("pausing for 5 seconds because of this.")
		time.Sleep(time.Second * 5)
	}
}
