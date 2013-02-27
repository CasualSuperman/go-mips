package main

type registers struct {
	names    map[string]*int32
	integer  [32]int32
	floating [32]float32
	extra    struct {
		hi, lo, pc int32
	}
}

var Registers registers

type instruction func(*registers)

func init() {
	Registers.names = make(map[string]*int32)
	var names = [32]string{
		"zero", "at", "v0", "v1", "a0", "a1", "a2", "a3", "t0", "t1", "t2",
		"t3", "t4", "t5", "t6", "t7", "s0", "s1", "s2", "s3", "s4", "s5", "s6",
		"s7", "t8", "t9", "k0", "k1", "gp", "sp", "fp", "ra",
	}
	for i := 0; i < len(names); i++ {
		Registers.names[names[i]] = &Registers.integer[i]
	}
	Registers.names["hi"] = &Registers.extra.hi
	Registers.names["lo"] = &Registers.extra.lo
	Registers.names["pc"] = &Registers.extra.pc
}
