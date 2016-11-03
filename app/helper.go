package app

import (
	"fmt"
	"log"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

func Trace(arg ...interface{}) {
	if Debug {
		fmt.Println(arg...)
	}
}

func Cmd(name string, arg ...string) (stdout string, err error) {
	c := exec.Command(name, arg...)
	out, err := c.Output()
	if err != nil {
		Trace("ERROR::Cmd:", name, arg)
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func Osascript(script string) (stdout string) {
	arr := strings.Split(script, "\n")
	var arg []string
	for _, l := range arr {
		arg = append(arg, "-e", l)
	}
	out, err := Cmd("osascript", arg...)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func GetDesktopCount() int {
	scpt := `tell application "System Events" to copy count of desktops to stdout`
	c, err := strconv.ParseInt(Osascript(scpt), 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	Trace("Desktop Count:", c)
	return int(c)
}

func ApplyDesktop(picture string, index int) {
	scpt := fmt.Sprintf(`tell application "System Events"
    tell desktop %v
        set picture to "%v"
    end tell
end tell`, index+1, picture)
	Osascript(scpt)
}

// Fix Mac desktop picture cache by dynamic picture name
func GenerateFilePaths(count int) *[]string {
	arr := []string{}
	name := 0
	if Debug {
		name = time.Now().Second()
	} else {
		name = time.Now().Hour()
	}
	for i := 0; i < count; i++ {
		arr = append(arr, path.Join(PictureDir, fmt.Sprintf("%v_%v.jpg", name, i)))
	}
	return &arr
}
