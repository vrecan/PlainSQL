package execute

import (
	"fmt"

	client "github.com/vrecan/PlainSQL/pkg/plainsql/client"
	"github.com/vrecan/PlainSQL/pkg/request"
)

//Execute will execute what is sent to it
type Execute struct {
	client.ClientOpts
}

// Execute satisfies the go-flags Commander interface and indicates the command was actually invoked
func (e *Execute) Execute(args []string) error {
	fmt.Println("args: ", args)
	if len(args) <= 0 {
		return fmt.Errorf("no args passed to echo")
	}

	cli := client.NewClient(e.ClientOpts)
	err := cli.Init()
	if err != nil {
		return err
	}
	defer cli.Close()

	resp, err := cli.Execute(request.Request{Query: string(args[0])})
	fmt.Println("ERROR: ", err, " RESP: ", resp)

	return nil
}
