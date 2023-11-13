//go:build mage_docs

package main

import (
	"github.com/spf13/cobra/doc"

	"github.com/khulnasoft/tunnel/pkg/commands"
	"github.com/khulnasoft/tunnel/pkg/flag"
	"github.com/khulnasoft/tunnel/pkg/log"
)

// Generate CLI references
func main() {
	// Set a dummy path for the documents
	flag.CacheDirFlag.Default = "/path/to/cache"
	flag.ModuleDirFlag.Default = "$HOME/.tunnel/modules"

	cmd := commands.NewApp()
	cmd.DisableAutoGenTag = true
	if err := doc.GenMarkdownTree(cmd, "./docs/docs/references/configuration/cli"); err != nil {
		log.Fatal(err)
	}
}
