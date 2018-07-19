package hello

type Input struct {
	Name string
}

type Output struct {
	Msg string
}

type EventListener interface {
	Callback(output *Output)
}

func Hello(input *Input, listener EventListener) {
	go func(input *Input, listener EventListener) {
		output := &Output{}
		output.Msg = "callback, name:" + input.Name
		listener.Callback(output)
	}(input, listener)
}
