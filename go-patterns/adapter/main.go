package adapter

import "fmt"

type IAdapter interface {
	Exceute()
}
type OldStruct struct {
}

func (o OldStruct) Run() {
	fmt.Println("Run Old")
}

type Adapter struct {
	oldStruct *OldStruct
}

func (a Adapter) Execute() {
	a.oldStruct.Run()
}

func NewAdapter() *Adapter {
	return &Adapter{new(OldStruct)}
}
