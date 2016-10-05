package main

import (
	"errors"
	"fmt"

	"github.com/dgoeke/config-demo/pkg/stages"
	"github.com/mitchellh/multistep"
)

type FirstStage struct {
	FirstString string
	FirstInt    int
}

func (fs *FirstStage) Name() string { return "FirstStage" }

func (fs *FirstStage) Run(multistep.StateBag) multistep.StepAction {
	fmt.Println("I am a first stage!")
	return multistep.ActionContinue
}

func (fs *FirstStage) Validate() error {
	if fs.FirstString == "" {
		return errors.New("firstString cannot be empty!")
	}

	return nil
}

func createFirst() stages.Stage {
	return &FirstStage{}
}

func init() {
	stages.MustRegister("first", createFirst)
}
