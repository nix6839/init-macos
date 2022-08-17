package install

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func installZplug() {
	var ZPLUG_HOME = os.Getenv("ZPLUG_HOME")

	_, err := git.PlainClone(ZPLUG_HOME, false, &git.CloneOptions{
		URL:      "https://github.com/zplug/zplug.git",
		Progress: os.Stdout,
	})
	switch err {
	case nil:
		return

	case git.ErrRepositoryAlreadyExists:
		fmt.Fprintln(os.Stderr, "'"+ZPLUG_HOME+"'"+" path already exists")

	default:
		panic(err)
	}
}

func Zsh() {
	installZplug()
}
