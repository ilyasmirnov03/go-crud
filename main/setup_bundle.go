package main

import (
	"log"

	"github.com/evanw/esbuild/pkg/api"
)

func setup_bundler() {
	js_result := api.Build(api.BuildOptions{
		EntryPoints:       []string{"public/js/main.js"},
		Bundle:            true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Outfile:           "_dist/build.js",
		Engines: []api.Engine{
			{Name: api.EngineChrome, Version: "58"},
			{Name: api.EngineFirefox, Version: "57"},
			{Name: api.EngineSafari, Version: "11"},
			{Name: api.EngineEdge, Version: "16"},
		},
		Write: true,
	})
	css_result := api.Build(api.BuildOptions{
		EntryPoints:       []string{"public/css/global.css"},
		Bundle:            true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Outfile:           "_dist/build.css",
		Engines: []api.Engine{
			{Name: api.EngineChrome, Version: "58"},
			{Name: api.EngineFirefox, Version: "57"},
			{Name: api.EngineSafari, Version: "11"},
			{Name: api.EngineEdge, Version: "16"},
		},
		Write: true,
	})
	if len(js_result.Errors) > 0 && len(css_result.Errors) > 0 {
		log.Fatal("Assets build error", js_result.Errors, css_result.Errors)
	}
}
