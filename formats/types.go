package formats

//Value follows the type spec for the X11 protocol for 'value'
type Value struct {
	Mask  uint32
	Value uint32
}

type ListOfValues struct {
	internalList []Value
}

func AddValue(v value)
