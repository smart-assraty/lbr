package main

import "log"

func checkError(errName string, e error) {
	if e != nil {
		log.Fatal("Error: "+errName+": ", e)
	}
}
