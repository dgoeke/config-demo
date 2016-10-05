package main

import (
	"fmt"

	"github.com/dgoeke/config-demo/pkg/config"
	"github.com/dgoeke/config-demo/pkg/stages"
	"github.com/kr/pretty"
	"github.com/mitchellh/multistep"
)

func main() {
	cfg, err := config.Parse("test.yaml")
	if err != nil {
		panic(err)
	}

	realStages, err := stages.Reify(cfg.Stages)
	if err != nil {
		panic(err)
	}

	fmt.Println("Config is:\n", pretty.Sprint(cfg))
	fmt.Println("Real stages are:\n", pretty.Sprint(realStages))

	bag := &multistep.BasicStateBag{}
	for i, stage := range realStages {
		fmt.Printf("%d. Running stage \"%v\":\n", i+1, stage.Name())
		stage.Run(bag)
	}
}
