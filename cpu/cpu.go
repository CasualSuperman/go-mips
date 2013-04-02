package cpu

type CPU struct {
	registers
	memory struct {
		hi, lo []byte
	}
	hidden struct {
		hi, lo, pc, epc, cause, badvaddr int32
	}
	os OperatingSystem
}

type Syscall func (c *CPU)
type OperatingSystem map[int32]Syscall

func New(os OperatingSystem) (c CPU) {
	c.os = os
	return
}

func (c *CPU) Syscall() {
	syscallId := c.Register("v0")
	if sys, ok := c.os[syscallId]; ok {
		sys(c)
	} else {
		panic("Syscall not implemented.")
	}
}
