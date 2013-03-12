package main

import (
	"go-mips/cpu"
)

var c cpu.CPU

func main() {
	c.SetRegister(2, 1)
	c.SetRegister(4, 1337)
	c.Syscall()
	c.SetRegister(2, 10)
	c.Syscall()
}
