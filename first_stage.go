package main

import (
	"errors"
	"fmt"

	"github.com/dgoeke/config-demo/pkg/config"
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

func createFirst() config.Stage {
	return &FirstStage{}
}

func init() {
	config.MustRegisterStage("first", createFirst)
}
