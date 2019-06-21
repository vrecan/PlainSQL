package main

import (
	"os"
	"strconv"

	SYS "syscall"

	"log"

	"github.com/joho/godotenv"
	sqld "github.com/vrecan/PlainSQL/pkg/plainsqld"
	DEATH "github.com/vrecan/death"
)

//Opts are all the flags that caan be passed to the PlainSQL command line client.
type Opts struct {
	Verbose    []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	ServerOpts sqld.Opts
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	death := DEATH.NewDeath(SYS.SIGINT, SYS.SIGTERM)
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalln("Invalid port given")
	}
	certFile := os.Getenv("CERT-FILE")
	log.Println("CERT-FILE:", certFile)
	certKey := os.Getenv("CERT-KEY")
	log.Println("CERT-KEY: ", certKey)
	srv := sqld.NewPlainSQLD(sqld.Opts{
		Port:       int32(port),
		DisableTLS: false,
		CertFile:   certFile,
		CertKey:    certKey,
	})
	srv.Start()
	death.WaitForDeath(srv)
	log.Println("PlainSQLD stopped")
}
