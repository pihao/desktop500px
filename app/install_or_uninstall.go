package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	programFileName = "desktop500px"
	plistFileName   = "com.pihao.desktop500px.launchd.plist"
	programFile     = path.Join(AppDir, programFileName)
	plistDir        = path.Join(os.Getenv("HOME"), "Library/LaunchAgents")
	plistFile       = path.Join(plistDir, plistFileName)
)

func Reinstall() {
	Uninstall()
	Install()
}

func Uninstall() {
	Cmd("launchctl", "unload", plistFile)
	Cmd("rm", plistFile)
	Cmd("rm", "-rf", AppDir)
	fmt.Println("Uninstall Complete.")
}

func Install() {
	fmt.Println("Generate key and plist file...")
	Cmd("mkdir", "-p", AppDir)
	Cmd("mkdir", "-p", PictureDir)
	generateKeyFile()
	generatePlistFile()

	fmt.Println("Config LaunchAgents...")
	wd, _ := os.Getwd()
	Cmd("cp", path.Join(wd, programFileName), AppDir)
	Cmd("chmod", "u+x", programFile)
	Cmd("launchctl", "load", plistFile)

	fmt.Println("Install Complete.")
}

func generateKeyFile() {
	if _, err := os.Stat(KeyFile); os.IsNotExist(err) {
		err = Cmd("cp", "key.json", KeyFile)
		if err != nil {
			fmt.Println("Check your key file: key.json")
			log.Fatal(err)
		}
	}
}

func generatePlistFile() {
	str := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>%v</string>
    <key>EnvironmentVariables</key>
    <dict>
      <key>PATH</key>
      <string>%v</string>
    </dict>
    <key>ProgramArguments</key>
    <array>
        <string>%v</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>StartInterval</key>
    <integer>3600</integer>
    <key>StandardOutPath</key>
    <string>%v/out.log</string>
    <key>StandardErrorPath</key>
    <string>%v/err.log</string>
</dict>
</plist>
`, plistFileName, os.Getenv("PATH"), programFile, AppDir, AppDir)

	err := ioutil.WriteFile(plistFile, []byte(str), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
