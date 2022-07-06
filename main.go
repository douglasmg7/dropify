package main

import (
	"io"
	"log"
	"os"
	"path"
	"time"
)

var dropify string
var dropify_host string
var dropify_client_id string
var dropify_client_secret string

var err error
var logPath string

// Production mode.
var production bool

// Brazil time location.
var brLocation *time.Location

func init() {
	// Host.
	dropify_host = os.Getenv("DROPIFY_HOST")
	if dropify_host == "" {
		panic("DROPIFY_HOST env not defined.")
	}
	if err != nil {
		panic(err)
	}

	// App id.
	dropify_client_id = os.Getenv("DROPIFY_CLIENT_ID")
	if dropify_client_id == "" {
		panic("DROPIFY_CLIENT_ID env not defined.")
	}
	// Secret key.
	dropify_client_secret = os.Getenv("DROPIFY_CLIENT_SECRET")
	if dropify_client_secret == "" {
		panic("DROPIFY_CLIENT_SECRET env not defined.")
	}

	// Check if production mode.
	if os.Getenv("RUN_MODE") == "production" {
		production = true
	}

	// Brazil location.
	brLocation, err = time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}

	// Path for log.
	zunkaPathdata := os.Getenv("ZUNKAPATH")
	if zunkaPathdata == "" {
		panic("ZUNKAPATH not defined.")
	}
	logPath := path.Join(zunkaPathdata, "log", "dropify")
	// Create path.
	os.MkdirAll(logPath, os.ModePerm)

	// Log file.
	logFile, err := os.OpenFile(path.Join(logPath, "dropify.log"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	// Log configuration.
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	// log.SetFlags(log.LstdFlags)
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[dropify] ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lmsgprefix)

}

func main() {
	// Log start.
	runMode := "development"
	if production {
		runMode = "production"
	}
	log.Printf("Running in %v mode (version %s)\n", runMode, version)
	getToken()

	// fmt.Println("dropify_host: ", dropify_host)
	// fmt.Println("dropify_client_id: ", dropify_client_id)
	// fmt.Println("dropify_client_secret: ", dropify_client_secret)
}
