package main

import (
	"flag"
	"fmt"

	"github.com/Tony-Moon/goloader/portal"
	"github.com/Tony-Moon/goloader/reciever"
	"github.com/Tony-Moon/goloader/sender"
)

func main() {
	runMode := flag.String("m", "none", "-m or runMode, tells the program to run in sender, portal or reciever mode.")
	flag.Parse()

	if *runMode == "s" {
		sender.GoSend()
	} else if *runMode == "p" {
		portal.GoPortal()
	} else if *runMode == "r" {
		reciever.GoRecieve()
	} else {
		fmt.Println("Unrecognized command.")
	}

}
