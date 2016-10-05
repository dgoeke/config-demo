package stages

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/mitchellh/multistep"
)

var (
	registeredStages = make(StageMap)
)

type StageConstructor func() Stage
type StageMap map[string]StageConstructor

type Stage interface {
	Name() string
	Validate() error
	Run(state multistep.StateBag) multistep.StepAction
}

func Reify(rawStages []map[string]interface{}) ([]Stage, error) {
	var realStages []Stage

	for i, rawStage := range rawStages {
		result, err := convertStage(rawStage)
		if err != nil {
			return []Stage{}, fmt.Errorf("Error in stage %v definition: %v", i+1, err.Error())
		}

		realStages = append(realStages, *result)
	}

	return realStages, nil
}

func convertStage(rawStage map[string]interface{}) (*Stage, error) {
	rawName, ok := rawStage["name"]
	if !ok {
		return nil, fmt.Errorf("Stage does not have a name: %+v", rawStage)
	}

	name, ok := rawName.(string)
	if !ok {
		return nil, fmt.Errorf("Stage name \"%v\" must be a string", rawName)
	}

	stageFn, ok := registeredStages[name]
	if !ok {
		return nil, fmt.Errorf("No stage named \"%v\" is defined", name)
	}

	newStage := stageFn()
	if err := mapstructure.Decode(rawStage, newStage); err != nil {
		return nil, err
	}

	if err := newStage.Validate(); err != nil {
		return nil, fmt.Errorf("%v: %v", newStage.Name(), err)
	}

	return &newStage, nil
}

func getStage(name string) *StageConstructor {
	stage, ok := registeredStages[strings.ToLower(name)]
	if !ok {
		return nil
	}

	return &stage
}

func MustRegister(name string, createFn StageConstructor) {
	if stage := getStage(name); stage != nil {
		panic("Duplicate stage name registered")
	}

	registeredStages[strings.ToLower(name)] = createFn
}
