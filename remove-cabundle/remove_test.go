package main

import (
	"testing"

	"sigs.k8s.io/kustomize/kyaml/fn/framework/frameworktestutil"
)

func TestRemoveCreationTimestamp(t *testing.T) {
	checker := frameworktestutil.CommandResultsChecker{
		Command: NewCommand,
	}
	checker.Assert(t)
}
