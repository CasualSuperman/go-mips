package cpu

type CPU struct {
	registers
	memory struct {
		hi, lo []byte
	}
	hidden struct {
		hi, lo, pc, epc, cause, badvaddr int32
	}
}

func New() (c CPU) {
	c.registers = newRegisters()
	return
}
