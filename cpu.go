package main

var CPU struct {
	memory struct {
		hi, lo []byte
	}
	hidden struct {
		hi, lo, pc, epc, cause, badvaddr int32
	}
}
