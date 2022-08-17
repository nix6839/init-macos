package command

import (
	"os"
	"os/exec"
)

type command struct {
	cmd *exec.Cmd
}

func New(program string) *command {
	cmd := new(command)
	cmd.cmd = exec.Command(program)
	return cmd
}

func (self *command) Arg(arg string) *command {
	self.cmd.Args = append(self.cmd.Args, arg)
	return self
}

func (self *command) Args(args []string) *command {
	self.cmd.Args = append(self.cmd.Args, args...)
	return self
}

func (self *command) Env(key string, value string) *command {
	self.cmd.Env = append(self.cmd.Environ(), key+"="+value)
	return self
}

func (self command) Run() {
	self.cmd.Stdin = os.Stdin
	self.cmd.Stdout = os.Stdout
	self.cmd.Stderr = os.Stderr

	if err := self.cmd.Start(); err != nil {
		panic(err)
	}

	self.cmd.Wait()
}
