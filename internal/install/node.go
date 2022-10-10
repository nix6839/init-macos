package install

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nix6839/macos-init/pkg/command"
	"github.com/nix6839/macos-init/pkg/utils"
)

func fetchNvmScript() string {
	return utils.FetchBody("https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh")
}

var nodePackages = []string{
	"create-next-app", "create-vite", "create-svelte", "create-turbo",
	"@nestjs/cli", "typescript", "ts-node", "nodemon", "localtunnel",
	"npm-check-updates", "sort-package-json",
}

func Node() {
	command.New("bash").
		Arg("-c").
		Arg(fetchNvmScript()).
		Run()

	nvmScriptPath := filepath.Join(os.Getenv("NVM_DIR"), "nvm.sh")

	command.New("zsh").
		Arg("-c").
		Arg(fmt.Sprintf(". '%s'; nvm install --lts --latest-npm", nvmScriptPath)).
		Run()

	command.New("npm").Arg("install").Arg("-g").Arg("corepack").Run()

	command.New("corepack").Arg("enable").Run()
	command.New("corepack").
		Arg("prepare").
		Arg("pnpm@7.13.3").
		Arg("--activate").
		Run()

	command.New("pnpm").Arg("install").Arg("-g").Args(nodePackages)
}
