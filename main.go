package main

import (
	"flag"
	"fmt"
	"github.com/pihao/desktop500px/app"
	"github.com/pihao/desktop500px/px500"
)

func main() {
	i := flag.Bool("i", false, "enable/disable install mode.")
	u := flag.Bool("u", false, "enable/disable uninstall mode.")
	r := flag.Bool("r", false, "enable/disable reinstall mode.")
	d := flag.Bool("d", false, "enable/disable debug mode.")
	v := flag.Bool("v", false, "show version.")
	flag.Parse()

	app.Debug = *d

	if *i {
		app.Install()
	} else if *u {
		app.Uninstall()
	} else if *r {
		app.Reinstall()
	} else if *v {
		fmt.Println(app.VERSION)
	} else {
		px500.Run()
	}
}
