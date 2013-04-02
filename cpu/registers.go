package cpu

var nameToIndex = map[string]int {
	"zero": 0,
	"at": 1,
	"v0": 2,
	"v1": 3,
	"a0": 4,
	"a1": 5,
	"a2": 6,
	"a3": 7,
	"t0": 8,
	"t1": 9,
	"t2": 10,
	"t3": 11,
	"t4": 12,
	"t5": 13,
	"t6": 14,
	"t7": 15,
	"s0": 16,
	"s1": 17,
	"s2": 18,
	"s3": 19,
	"s4": 20,
	"s5": 21,
	"s6": 22,
	"s7": 23,
	"t8": 24,
	"t9": 25,
	"k0": 26,
	"k1": 27,
	"gp": 28,
	"sp": 29,
	"fp": 30,
	"ra": 31,
}

type registers struct {
	integer  [32]int32
	floating [32]float32
}

type instruction func(*registers)

func (r *registers) SetRegister(reg string, val int32) {
	if i, ok := nameToIndex[reg]; ok {
		if i != 0 {
			r.integer[nameToIndex[reg]] = val
		}
	} else {
		panic("Invalid register: $" + reg)
	}
}

func (r registers) Register(reg string) int32 {
	if i, ok := nameToIndex[reg]; ok {
		return r.integer[i]
	}
	panic("Invalid register: $" + reg)
}

