package main

import (
	"flag"
	"fmt"

	"github.com/pihao/desktop500px/src/app"
	"github.com/pihao/desktop500px/src/px500"
)

func main() {
	i, u, r, d, v := getFlag()
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

func getFlag() (i, u, r, d, v *bool) {
	i = flag.Bool("i", false, "install.")
	u = flag.Bool("u", false, "uninstall.")
	r = flag.Bool("r", false, "reinstall.")
	d = flag.Bool("d", false, "debug mode.")
	v = flag.Bool("v", false, "show version.")
	flag.Parse()
	return i, u, r, d, v
}
