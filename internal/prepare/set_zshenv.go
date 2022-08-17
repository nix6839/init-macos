package prepare

import (
	"os"
	"regexp"

	"github.com/nix6839/macos-init/pkg/utils"
)

func fetchZshenv() string {
	return utils.FetchBody("https://raw.githubusercontent.com/nix6839/dotfiles-macos/main/dot_zshenv")
}

var exportRegex = regexp.MustCompile("(?m)^export (?P<key>.+)=\"(?P<value>.+)\"$")

func SetZshenv() {
	for _, s := range exportRegex.FindAllStringSubmatch(fetchZshenv(), -1) {
		key, value := s[1], s[2]
		os.Setenv(key, os.ExpandEnv(value))
	}
}
