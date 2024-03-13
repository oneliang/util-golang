package common

type UnsignedIntType = interface {
	uint | uint8 | uint16 | uint32 | uint64
}
type SignedIntType = interface {
	int | int8 | int16 | int32 | int64
}

type IntType = interface {
	UnsignedIntType | SignedIntType
}

type FloatType = interface {
	float32 | float64
}
type NumberType = interface {
	IntType | FloatType
}
type SimpleType = interface {
	NumberType | string
}

type SimpleTypeAndStruct = interface {
	SimpleType | struct{}
}
