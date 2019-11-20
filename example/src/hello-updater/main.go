package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/golubaca/go-selfupdate/selfupdate"
)

var version = "0.1.4"

var updater = &selfupdate.Updater{
	CurrentVersion: version,                  // Manually update the const, or set it using `go build -ldflags="-X main.VERSION=<newver>" -o hello-updater src/hello-updater/main.go`
	ApiURL:         "http://localhost:8080/", // The server hosting `$CmdName/$GOOS-$ARCH.json` which contains the checksum for the binary
	BinURL:         "http://localhost:8080/", // The server hosting the zip file containing the binary application which is a fallback for the patch method
	DiffURL:        "http://localhost:8080/", // The server hosting the binary patch diff for incremental updates
	Dir:            "update/",                // The directory created by the app when run which stores the cktime file
	CmdName:        "hello-updater",          // The app name which is appended to the ApiURL to look for an update
	ForceCheck:     true,                     // For this example, always check for an update unless the version is "dev"
}

func CheckUpdate() {
	for {
		updated, _ := updater.BackgroundRun()

		if updated {
			log.Println("Azurirali smo aplikaciju, gasimo sebe i palimo novi proces")
			args := []string{"-graceful"}
			cmd := exec.Command(os.Args[0], args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			// put socket FD at the first entry
			// cmd.ExtraFiles = []*os.File{f}
			cmd.Start()
			os.Exit(0)
		}
		// fmt.Println(updated, err)
		time.Sleep(time.Second * 10)
	}
}

func init() {
	log.Printf("App Init Done %v", updater.CurrentVersion)
}

func main() {
	log.Printf("Hello world I am currently version %v", updater.CurrentVersion)

	// log.Printf("Next run, I should be %v", updater.Info.Version)
	go CheckUpdate()
	ch := make(chan int, 1)
	<-ch
}
