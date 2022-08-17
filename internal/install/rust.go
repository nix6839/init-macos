package install

import (
	"github.com/nix6839/macos-init/pkg/command"
	"github.com/nix6839/macos-init/pkg/utils"
)

func fetchRustUpScript() string {
	return utils.FetchBody("https://sh.rustup.rs")
}

func Rust() {
	command.New("sh").
		Arg("-c").
		Arg(fetchRustUpScript()).
		Arg("--").
		Arg("--no-modify-path").
		Arg("-y").
		Run()
}
