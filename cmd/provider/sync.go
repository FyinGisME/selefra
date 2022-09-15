package provider

import (
	"context"
	"github.com/selefra/selefra/cmd/fetch"
	"github.com/selefra/selefra/config"
	"github.com/selefra/selefra/pkg/registry"
	"github.com/selefra/selefra/pkg/utils"
	"github.com/selefra/selefra/ui"
)

func Sync() error {
	ctx := context.Background()
	var cof = &config.SelefraConfig{}
	err := cof.GetConfig()
	if err != nil {
		return err
	}
	namespace, _, err := utils.Home()
	if err != nil {
		return err
	}
	provider := registry.NewProviderRegistry(namespace)
	for _, p := range cof.Selefra.Providers {
		prov := registry.Provider{
			Name:    p.Name,
			Version: p.Version,
			Source:  "",
		}
		pp, err := provider.Download(ctx, prov, true)
		if err != nil {
			ui.PrintErrorF("%s %s Download failed：%s", p.Name, p.Version, err.Error())
			continue
		} else {
			ui.PrintSuccessF("%s %s Download succeeded", p.Name, p.Version)
		}
		p.Path = pp.Filepath
		p.Version = pp.Version
		ui.PrintSuccessF("%s %s Synchronization ...", p.Name, p.Version)
		err = fetch.Fetch(ctx, cof, p)
		if err != nil {
			ui.PrintErrorF("%s %s Synchronization failed：%s", p.Name, p.Version, err.Error())
			continue
		} else {
			ui.PrintSuccessF("%s %s Synchronization succeeded", p.Name, p.Version)
		}
	}
	return nil
}