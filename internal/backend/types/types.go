package types

type ValueType uint8

const (
	Int32 ValueType = iota
	Int64
	Float32
	Float64

	UintPtr
	List
)
