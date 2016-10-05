package main

import (
	"fmt"

	"github.com/dgoeke/config-demo/pkg/config"
	"github.com/kr/pretty"
	"github.com/mitchellh/multistep"
)

func main() {
	cfg, err := config.Parse("test.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println("Config is:\n", pretty.Sprint(cfg))

	bag := &multistep.BasicStateBag{}
	for i, stage := range cfg.RealStages() {
		fmt.Printf("%d. Running stage \"%v\":\n", i, stage.Name())
		stage.Run(bag)
	}
}
