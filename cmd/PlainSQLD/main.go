package main

import (
	"flag"

	SYS "syscall"

	"log"

	sqld "github.com/vrecan/PlainSQL/pkg/plainsqld"
	DEATH "github.com/vrecan/death"
)

var port = flag.Int("port", 2001, "port you want to use for PlainSQLD")

func main() {

	death := DEATH.NewDeath(SYS.SIGINT, SYS.SIGTERM)
	flag.Parse()
	srv := sqld.NewPlainSQLD(*port)
	srv.Start()
	death.WaitForDeath(srv)
	log.Println("PlainSQLD stopped")
}
