package main

import (
	"errors"
	"fmt"

	"github.com/dgoeke/config-demo/pkg/stages"
	"github.com/mitchellh/multistep"
)

type SecondStage struct {
	SecondString string
	SecondInt    int
}

func (fs *SecondStage) Name() string { return "SecondStage" }

func (fs *SecondStage) Run(multistep.StateBag) multistep.StepAction {
	fmt.Println("I'm a second stage!")
	return multistep.ActionContinue
}

func (fs *SecondStage) Validate() error {
	if fs.SecondInt != 12 {
		return errors.New("SecondInt must be 12!")
	}

	return nil
}

func createSecond() stages.Stage {
	return &SecondStage{}
}

func init() {
	stages.MustRegister("second", createSecond)
}
