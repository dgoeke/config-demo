package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/mitchellh/multistep"
	"gopkg.in/yaml.v2"
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

type Config struct {
	Value1 string
	Value2 string
	Stages []map[string]interface{}

	realStages []Stage
}

func (c *Config) RealStages() []Stage {
	return c.realStages
}

func (c *Config) reifyStages() error {
	for i, rawStage := range c.Stages {
		result, err := convertStage(rawStage)
		if err != nil {
			return fmt.Errorf("Error in stage %v definition: %v", i+1, err.Error())
		}

		c.realStages = append(c.realStages, *result)
	}

	return nil
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

func MustRegisterStage(name string, createFn StageConstructor) {
	if stage := getStage(name); stage != nil {
		panic("Duplicate stage name registered")
	}

	registeredStages[strings.ToLower(name)] = createFn
}

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []byte{}, err
	}

	return ioutil.ReadAll(file)
}

func Parse(filename string) (*Config, error) {
	bytes, err := readFile(filename)
	if err != nil {
		return nil, err
	}

	result := &Config{}
	yaml.Unmarshal(bytes, result)

	if err := result.reifyStages(); err != nil {
		return nil, err
	}

	return result, nil
}
