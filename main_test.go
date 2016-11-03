package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	i, u, r, d, v := getFlag()
	if *i != false {
		t.Error(`"-i" default value should be false`)
	}
	if *u != false {
		t.Error(`"-u" default value should be false`)
	}
	if *r != false {
		t.Error(`"-r" default value should be false`)
	}
	if *d != false {
		t.Error(`"-d" default value should be false`)
	}
	if *v != false {
		t.Error(`"-v" default value should be false`)
	}
}
