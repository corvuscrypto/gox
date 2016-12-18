package gox

type Request struct {
	OpCode uint8
	Length uint16
	Data   uint8
}
