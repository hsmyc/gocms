package configs

import (
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func ESBuilder() {
	result := api.Build(api.BuildOptions{
		EntryPoints:      []string{"./frontend/scripts/script.ts", "./frontend/styles/global.css"},
		Bundle:           true,
		MinifyWhitespace: true,
		Outdir:           "./frontend/dist",
		Write:            true,
		GlobalName:       "App",
		Format:           api.FormatESModule,
	})

	if len(result.Errors) > 0 {
		os.Stderr.WriteString("Build failed with errors:\n")
		for _, err := range result.Errors {
			os.Stderr.WriteString(err.Text + "\n")
		}
		os.Exit(1)
	}
}
