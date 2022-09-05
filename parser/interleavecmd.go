package parser

import (
	"errors"
	"fmt"

	interleave "github.com/acekingke/Decartes/interleaveSeq"
)

type InterleaveType struct {
	Left  []string
	Right []string
	Pre   string
	Post  string
}

// Add the execute and test.
func (inter *InterleaveType) Execute() error {
	// Check the Steps are all exist
	for _, name := range inter.Left {
		if _, ok := StepMap[name]; !ok {
			return errors.New("Step " + name + " not found")
		}
	}
	for _, name := range inter.Right {
		if _, ok := StepMap[name]; !ok {
			return errors.New("Step " + name + " not found")
		}
	}
	interleave.Interleave(inter.Left, inter.Right, func(s []string) {
		if len(inter.Pre) != 0 {
			PushContex()
			ParserInit()
			_ = Parser(inter.Pre + "\n")
			PopContex()
		}
		for _, name := range s {
			cmd := StepMap[name].Cmds
			cmdstr := fmt.Sprintf("%s\n", cmd[1:len(cmd)-1])
			PushContex()
			ParserInit()
			_ = Parser(cmdstr)
			PopContex()
		}
		if len(inter.Post) != 0 {
			PushContex()
			ParserInit()
			_ = Parser(inter.Post + "\n")
			PopContex()
		}
	})
	return nil
}
