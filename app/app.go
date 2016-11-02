package app

import (
	"os"
	"path"
)

const (
	VERSION = `desktop500px 0.1`
)

var Debug = false

var AppDir = path.Join(os.Getenv("HOME"), ".desktop500px")
var KeyFile = path.Join(AppDir, "key.json")
