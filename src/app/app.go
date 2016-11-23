package app

import (
	"os"
	"path"
)

const (
	VERSION = `desktop500px v0.3.1`
)

var (
	Debug = false

	AppDir     = path.Join(os.Getenv("HOME"), "Applications/desktop500px")
	PictureDir = path.Join(AppDir, "img")
	KeyFile    = path.Join(AppDir, "key.json")
)
