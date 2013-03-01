package main

var CPU struct {
	memory struct {
		hi, lo []byte
	}
	hidden struct {
		hi, lo, pc, epc, cause, badvaddr int32
	}
}

type memory interface {
	GetByte(int32) byte
	GetShort(int32) int16
	GetWord(int32) int32
	GetDouble(int32) float64
}

type sparse_memory struct {
	length int32
	start list_item
}

type list_item struct {
	position int32
	value int32
	next *list_item
}
