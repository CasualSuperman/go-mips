package cpu

import (
	"fmt"
	"os"
)

type syscall func (c *CPU)

var syscalls []syscall

func init() {
	syscalls = []syscall{unImplemented, syscall1, unImplemented, unImplemented,
		unImplemented, unImplemented, unImplemented, unImplemented,
		unImplemented, unImplemented, syscall10, unImplemented,
		unImplemented, unImplemented, unImplemented, unImplemented,
		unImplemented, syscall17}
}

func (c *CPU) Syscall() error {
	syscallId := int(c.registers.integer[2]) // $v0
	if syscallId < 1 || syscallId > 17 {
		return fmt.Errorf("Syscall %d invalid.", syscallId)
	} else {
		syscalls[syscallId](c)
	}
	return nil
}

func unImplemented(c *CPU) {
	fmt.Fprintln(os.Stderr, "Not implemented")
}

func syscall1(c *CPU) {
	fmt.Printf("%d", c.registers.integer[4]) // $a0
}

func syscall10(c *CPU) {
	os.Exit(0)
}

func syscall17(c *CPU) {
	os.Exit(int(c.registers.integer[4]))
}
