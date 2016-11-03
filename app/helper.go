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

func ApplyDesktop(paths ...string) {
	if len(paths) < 2 {
		log.Fatal("ERROR::ApplyDesktop(): paths length less than 2.")
	}
	scpt := fmt.Sprintf(`set arr to {"%v", "%v"}
tell application "System Events"
    set desktopCount to count of desktops
    repeat with i from 1 to desktopCount
        if i <= count of arr then
            set img to item i of arr
        else
            set img to item 1 of arr
        end if
        tell desktop i
            set picture to img
        end tell
    end repeat
end tell`, paths[0], paths[1])
	Osascript(scpt)
}

// Fix Mac desktop picture cache by dynamic picture name
func ImageFiles(count int) *[]string {
	arr := []string{}
	name := time.Now().Second()
	for i := 0; i < count; i++ {
		arr = append(arr, path.Join(PictureDir, fmt.Sprintf("%v_%v.jpg", name, i)))
	}
	return &arr
}
