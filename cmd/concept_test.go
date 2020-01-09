package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func TestConcept(t *testing.T) {
	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p ../_fixtures/examples/api",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "concept",
		Cmd:    "concept",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}