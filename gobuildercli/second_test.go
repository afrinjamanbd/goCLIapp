package main

import (
	"testing"

	"github.com/afrinjamanbd/goCLIapp/gobuildercli/cmd"
)

func TESTTwo(t *testing.T) {

	t.Errorf("Error from test two")
	cmd.Skip(t)
}
