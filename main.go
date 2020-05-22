package main

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	// Get the HOST AND PORT from .env file.
	hostport := os.Getenv("FTP_HOSTPORT")

	c, err := ftp.Dial(hostport, ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(os.Getenv("FTP_USER"), os.Getenv("FTP_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}

	data := bytes.NewBufferString("Hello World")
	err = c.Stor("./assets/img/avatars/test-file.txt", data)
	if err != nil {
		panic(err)
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
