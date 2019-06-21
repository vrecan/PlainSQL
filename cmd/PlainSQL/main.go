package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	E "github.com/vrecan/PlainSQL/pkg/plainsql/echo"
	EX "github.com/vrecan/PlainSQL/pkg/plainsql/execute"
)

//Opts are all the flags that caan be passed to the PlainSQL command line client.
type Opts struct {
	Verbose []bool     `short:"v" long:"verbose" description:"Show verbose debug information"`
	Echo    E.Echo     `command:"echo"                description:"sent echo request"`
	Execute EX.Execute `command:"execute" description:"execute a query"`
}

func main() {
	parser := flags.NewParser(&Opts{}, flags.HelpFlag)
	_, err := parser.Parse() // Also executes the sub-command
	if flags.WroteHelp(err) {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}
	if err != nil {
		log.Fatalf("%s", err)
	}

}
