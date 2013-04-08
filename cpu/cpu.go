package cpu

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
	memoryOffset = c.hidden.pc >> 2
	var command int32
	for i := 0; i < 4; i++ {
		command = command & c.memory.program[memoryOffset]
		command = command << 8
	}
	pcChanged := c.Execute(command)
	if !pcChanged {
		c.hidden.pc += 4
	}
}

func (c *CPU) Execute(command int32) bool {

}
