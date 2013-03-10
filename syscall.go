package main

func (c *CPU) Syscall() {
	switch (c.registers.integer[2]) { // $v0
	case 0:
		fmt.Printf("%d", c.registers.integer[4]) // $a0
	}
}
