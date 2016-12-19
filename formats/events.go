package formats

//Event types
const (
	KeyPress OpCode = iota + 2
	KeyRelease
	ButtonPress
	ButtonRelease
	MotionNotify
	EnterNotify
	LeaveNotify
	FocusIn
	FocusOut
	KeymapNotify
	Expose
	GraphicsExposure
	NoExposure
	VisibilityNotify
	CreateNotify
	DestroyNotify
	UnmapNotify
	MapNotify
	MapRequest
	ReparentNotify
	ConfigureNotify
	ConfigureRequest
	GravityNotify
	ResizeRequest
	CirculateNotify
	CirculateRequest
	PropertyNotify
	SelectionClear
	SelectionRequest
	SelectionNotify
	ColormapNotify
	ClientMessage
	MappingNotify
)

//Event is the format for events (X11proto p. 142-149)
type Event interface {
	Type() string
	Marshal() ([]byte, error)
	UnMarshal([]byte) error
}
