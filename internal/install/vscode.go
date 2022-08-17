package install

import (
	"strings"
	"sync"

	"github.com/nix6839/macos-init/pkg/command"
	"github.com/nix6839/macos-init/pkg/utils"
)

func extensions() []string {
	fetchedExtensions := utils.FetchBody("https://raw.githubusercontent.com/nix6839/dotfiles-macos/main/resources/code-extensions.list")
	return strings.Split(strings.TrimSpace(fetchedExtensions), "\n")
}

func VSCode() {
	extensions := extensions()

	var wg sync.WaitGroup
	wg.Add(len(extensions))
	for _, extension := range extensions {
		go func(extension string) {
			defer wg.Done()
			command.New("code").
				Arg("--install-extension").
				Arg(extension).
				Run()
		}(extension)
	}
	wg.Wait()
}
