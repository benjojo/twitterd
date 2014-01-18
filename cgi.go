package main

import (
	"log"
	"os"
)

func CheckForCGIDir() {
	f, e := os.Stat("./cgi")
	if e == nil {
		if !f.IsDir() {
			log.Println(`So you have made a cgi file. not a directory.\n
			 What. Removing your sillyness and doing it the right way`)
			e = os.Remove("./cgi")
			if e != nil {
				log.Fatal("Cannot remove (silly) the cgi file. What have you done!? (Permission probs)")
			}
			e := os.Mkdir("./cgi", 600)
			if e != nil {
				log.Fatalf("Cannot create the cgi dir. I kinda need to stop now. Reason %s", e.Error())
			}
		}
	} else {
		e := os.Mkdir("./cgi", 600)
		if e != nil {
			log.Fatalf("Cannot create the cgi dir. I kinda need to stop now. Reason %s", e.Error())
		}
	}
}
