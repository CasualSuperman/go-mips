package main

import (
	"github.com/CasualSuperman/go-mips/cpu"
)

func main() {
	c := cpu.New(syscalls)
	c.SetRegister("v0", 1)
	c.SetRegister("a0", 1337)
	c.Syscall()
	c.SetRegister("v0", 10)
	c.Syscall()
}
