package command

type Command interface {
	Execute()
}

type Invoker struct {
	commandList []Command
}

func (this *Invoker) AddCommand(cmd Command) {
	this.commandList = append(this.commandList, cmd)
}

func (this *Invoker) ExecCommand() {
	for i := 0; i < len(this.commandList); i++ {
		this.commandList[i].Execute()
	}
}
