package utils

import (
	"github.com/selefra/selefra/config"
	"github.com/selefra/selefra/pkg/registry"
	"gopkg.in/yaml.v3"
)

func SetProviders(DefaultConfigTemplate string, provider registry.ProviderBinary, config *config.SelefraConfig) error {
	if config.Providers.Kind != yaml.MappingNode {
		config.Providers.Kind = yaml.MappingNode
		config.Providers.Tag = "!!map"
		config.Providers.HeadComment = "provider configurations"
		config.Providers.Value = ""
		config.Providers.Content = make([]*yaml.Node, 0)
	}
	var node yaml.Node
	yaml.Unmarshal([]byte(DefaultConfigTemplate), &node)
	var provNode yaml.Node
	provNode.Content = append([]*yaml.Node{
		{
			Kind:  yaml.ScalarNode,
			Value: provider.Name,
		},
		{
			Kind:    yaml.MappingNode,
			Content: node.Content[0].Content,
		},
	})

	config.Providers.Content = append(config.Providers.Content, provNode.Content...)

	return nil
}