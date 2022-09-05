package parser

import (
	"errors"
	"fmt"

	"github.com/acekingke/Decartes/perm"
)

type StepType struct {
	Name string
	Cmds string
}
type PermType struct {
	StepNames []string
	Pre       string
	Post      string
}

var StepMap map[string]*StepType = make(map[string]*StepType)

func NewStep(name, cmds string) error {
	if _, ok := StepMap[name]; ok {
		return errors.New("Step " + name + " already exists")
	}

	StepMap[name] = &StepType{
		Name: name,
		Cmds: cmds,
	}
	return nil
}

func (p *PermType) Execute() error {
	// Check the Steps are all exist
	for _, name := range p.StepNames {
		if _, ok := StepMap[name]; !ok {
			return errors.New("Step " + name + " not found")
		}
	}
	// Add pre and post op
	perm.Perm(p.StepNames, func(a []string) {
		if len(p.Pre) != 0 {
			PushContex()
			ParserInit()
			_ = Parser(p.Pre + "\n")
			PopContex()
		}
		for _, name := range a {
			cmd := StepMap[name].Cmds
			cmdstr := fmt.Sprintf("%s\n", cmd[1:len(cmd)-1])
			PushContex()
			ParserInit()
			_ = Parser(cmdstr)
			PopContex()
		}
		if len(p.Post) != 0 {
			PushContex()
			ParserInit()
			_ = Parser(p.Post + "\n")
			PopContex()
		}
	})
	return nil
}
