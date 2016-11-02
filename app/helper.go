package app

import (
	"fmt"
	"log"
	"os/exec"
	"path"
	"strings"
	"time"
)

func Trace(arg ...interface{}) {
	if Debug {
		fmt.Println(arg...)
	}
}

func Cmd(name string, arg ...string) error {
	c := exec.Command(name, arg...)
	err := c.Run()
	if err != nil {
		Trace("ERROR::Cmd:", name, arg)
		return err
	}
	return nil
}

func Osascript(script string) {
	arr := strings.Split(script, "\n")
	var arg []string
	for _, l := range arr {
		arg = append(arg, "-e", l)
	}
	err := Cmd("osascript", arg...)
	if err != nil {
		log.Fatal(err)
	}
	Trace("Complete.")
}

func ApplyDesktop(path string) {
	scpt := fmt.Sprintf(`tell application "System Events"
    set desktopCount to count of desktops
    repeat with i from 1 to desktopCount
        tell desktop i
            set picture to "%v"
        end tell
    end repeat
end tell`, path)
	Osascript(scpt)
}

// Fix Mac desktop picture cache by dynamic picture name
func imageFile() string {
	name := fmt.Sprintf("%v.jpg", time.Now().Hour())
	return path.Join(AppDir, name)
}
