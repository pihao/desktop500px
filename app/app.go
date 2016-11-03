package app

import (
	"os"
	"path"
)

const (
	VERSION = `desktop500px 0.2`
)

var (
	Debug = false

	AppDir     = path.Join(os.Getenv("HOME"), ".desktop500px")
	PictureDir = path.Join(AppDir, "img")
	KeyFile    = path.Join(AppDir, "key.json")
)
