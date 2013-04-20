package cpu

import "unsafe"

type CPU struct {
	registers
	memory struct {
		hi, lo []byte
		program []byte
	}
	hidden struct {
		hi, lo, pc, epc, cause, badvaddr int32
	}
	os OperatingSystem
}

type Syscall func (c *CPU)
type OperatingSystem map[int32]Syscall

func New(os OperatingSystem) *CPU {
	c := CPU{os: os}
	return &c
}

func (c *CPU) Syscall() {
	syscallId := c.Register("v0")
	if sys, ok := c.os[syscallId]; ok {
		sys(c)
	} else {
		panic("Syscall not implemented.")
	}
}

func (c *CPU) Tick() {
	memoryOffset := c.hidden.pc >> 2
	command := *((*int32)(unsafe.Pointer(&c.memory.program[memoryOffset])))
	pcChanged := c.Execute(command)
	if !pcChanged {
		c.hidden.pc += 4
	}
}

func (c *CPU) Execute(command int32) bool {
	jumped := false
	opcode := command >> (32 - 6)
	// Binary: 0001XX
	// Binary: 00001X
	// Binary: 00100X
	if opcode & (074) == (004) {
		// This is a branch instruction
	} else if opcode & (076) == 2 || opcode & (076) == 8 {
		// This is a jump instruction
		imm := command & (0x03FFFFFF)
		jumped == true
	} else {

	}
	return jumped
}
