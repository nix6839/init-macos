package install

import (
	"github.com/nix6839/macos-init/pkg/command"
	"github.com/nix6839/macos-init/pkg/utils"
)

var pythonPackages = []string{
	"autopep8", "pylint",
}

func fetchPoetryScript() string {
	return utils.FetchBody("https://install.python-poetry.org")
}

func Python() {
	command.New("pip3").
		Arg("install").
		Arg("--upgrade").
		Arg("pip")

	command.New("pip3").
		Arg("install").
		Args(pythonPackages)

	command.New("python3").
		Arg("-c").
		Arg(fetchPoetryScript()).
		Run()
}
