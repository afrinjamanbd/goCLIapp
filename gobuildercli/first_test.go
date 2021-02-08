package main

import (
	"testing"

	"github.com/afrinjamanbd/goCLIapp/gobuildercli/cmd"
)

func TESTOne(t *testing.T) {
	t.Errorf("Error from test one")
	cmd.Skip(t)
}
