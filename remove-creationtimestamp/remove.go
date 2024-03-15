package main

import (
	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/command"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func newProcessor() framework.ResourceListProcessor {
	return framework.SimpleProcessor{
		Filter: kio.FilterFunc(func(nodes []*yaml.RNode) ([]*yaml.RNode, error) {
			for _, node := range nodes {
				var err error
				node, err = node.Pipe(
					yaml.Tee(
						yaml.Lookup("metadata"),
						yaml.Clear("creationTimestamp"),
					),
					yaml.Tee(
						yaml.Lookup("spec", "template", "metadata"),
						yaml.Clear("creationTimestamp"),
					),
				)
				if err != nil {
					return nil, err
				}
			}
			return nodes, nil
		}),
	}
}

func NewCommand() *cobra.Command {
	return command.Build(newProcessor(), command.StandaloneEnabled, false)
}
