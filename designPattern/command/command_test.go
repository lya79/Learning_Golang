package command

import (
	"fmt"
	"testing"
)

type HelloWorld struct {
	info string
}

func (this *HelloWorld) Execute() {
	fmt.Println(this.info)
}

func Test_cmd(t *testing.T) {
	invoker := &Invoker{}

	fmt.Println("add cmd 1")
	cmd := &HelloWorld{}
	cmd.info = "cmd 1"
	invoker.AddCommand(cmd)

	fmt.Println("add cmd 2")
	cmd2 := &HelloWorld{}
	cmd2.info = "cmd 2"
	invoker.AddCommand(cmd2)

	fmt.Println("ExecCommand")
	invoker.ExecCommand()
}
