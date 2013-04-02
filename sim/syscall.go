package main

import (
	"fmt"
	"os"
	"github.com/CasualSuperman/go-mips/cpu"
)

var syscalls = map[int32]cpu.Syscall{
	1: printInt,
	10: exit,
	17: exitWithCode,
}

func unImplemented(c *cpu.CPU) {
	fmt.Fprintln(os.Stderr, "Not implemented")
}

func printInt(c *cpu.CPU) {
	fmt.Printf("%d", c.Register("a0"))
}

func exit(c *cpu.CPU) {
	os.Exit(0)
}

func exitWithCode(c *cpu.CPU) {
	os.Exit(int(c.Register("a0")))
}
