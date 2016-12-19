package formats

//Request Op Codes
const (
	CreateWindow OpCode = iota + 1
	ChangeWindowAttributes
	GetWindowAttributes
	DestroyWindow
	DestroySubwindows
	ChangeSaveSet
	ReparentWindow
	MapWindow
	MapSubwindows
	UnmapWindow
	UnmapSubwindows
	ConfigureWindow
	CirculateWindow
	GetGeometry
	QueryTree
	InternAtom
	GetAtomName
	ChangeProperty
	DeleteProperty
	GetProperty
	ListProperties
	SetSelectionOwner
	GetSelectionOwner
	ConvertSelection
	SendEvent
	GrabPointer
	UngrabPointer
	GrabButton
	UngrabButton
	ChangeActivePointerGrab
	GrabKeyboard
	UngrabKeyboard
	GrabKey
	UngrabKey
	AllowEvents
	GrabServer
	UngrabServer
	QueryPointer
	GetMotionEvents
	TranslateCoordinates
	WarpPointer
	SetInputFocus
	GetInputFocus
	QueryKeymap
	OpenFont
	CloseFont
	QueryFont
	QueryTextExtents
	ListFonts
	ListFontsWithInfo
	SetFontPath
	GetFontPath
	CreatePixmap
	FreePixmap
	CreateGC
	ChangeGC
	CopyGC
	SetDashes
	SetClipRectangles
	FreeGC
	ClearArea
	CopyArea
	CopyPlane
	PolyPoint
	PolyLine
	PolySegment
	PolyRectangle
	PolyArc
	FillPoly
	PolyFillRectangle
	PolyFillArc
	PutImage
	GetImage
	PolyText8
	PolyText16
	ImageText8
	ImageText16
	CreateColormap
	CopyColormapAndFree
	InstallColormap
	UninstallColormap
	ListInstalledColormaps
	AllocColor
	AllocNamedColor
	AllocColorCells
	AllocColorPlanes
	FreeColors
	StoreColors
	StoreNamedColor
	QueryColors
	LookupColor
	CreateCursor
	CreateGlyphCursor
	FreeCursor
	RecolorCursor
	QueryBestSize
	QueryExtension
	ListExtension
	ChangeKeyboardMapping
	GetKeyboardMapping
	ChangeKeyboardControl
	GetKeyboardControl
	Bell
	ChangePointerControl
	GetPointerControl
	SetScreenSaver
	GetScreenSaver
	ChangeHosts
	ListHosts
	SetAccessControl
	SetCloseDownMode
	KillClient
	RotateProperties
	ForceScreenSaver
	SetPointerMapping
	GetPointerMapping
	SetModifierMapping
	GetModifierMapping
	NoOperation
)

//Request is the format for requests (X11proto p. 113-142)
type Request interface {
	Marshal() ([]byte, error)
	UnMarshal([]byte) error
}

type CreateWindowRequest struct {
	Depth       uint8
	WindowID    uint32
	Parent      uint32
	X           int16
	Y           int16
	Width       uint16
	Height      uint16
	BorderWidth uint16
	Class       uint16
	Visual      uint32
	valueMask   uint32
	valueList   []Value
}

//Marshal transforms the data in the request into a slice of bytes
func (c *CreateWindowRequest) Marshal() (data []byte, err error) {
	ByteOrder.PutUint32(data, c.WindowID)
	ByteOrder.PutUint32(data, c.Parent)
	ByteOrder.PutUint16(data, uint16(c.X))
	ByteOrder.PutUint16(data, uint16(c.Y))
	ByteOrder.PutUint16(data, c.Width)
	ByteOrder.PutUint16(data, c.Height)
	ByteOrder.PutUint16(data, c.BorderWidth)
	ByteOrder.PutUint16(data, c.Class)
	ByteOrder.PutUint32(data, c.Visual)
	ByteOrder.PutUint32(data, c.valueMask)
	for _, value := range c.valueList {
		ByteOrder.PutUint32(data, value.Value)
	}
	return data, err
}

//Marshal transforms the data in the request into a slice of bytes
func (c *CreateWindowRequest) Unmarshal(data []byte) (err error) {
	ByteOrder.PutUint32(data, c.WindowID)
	ByteOrder.PutUint32(data, c.Parent)
	ByteOrder.PutUint16(data, uint16(c.X))
	ByteOrder.PutUint16(data, uint16(c.Y))
	ByteOrder.PutUint16(data, c.Width)
	ByteOrder.PutUint16(data, c.Height)
	ByteOrder.PutUint16(data, c.BorderWidth)
	ByteOrder.PutUint16(data, c.Class)
	ByteOrder.PutUint32(data, c.Visual)
	ByteOrder.PutUint32(data, c.valueMask)
	for _, value := range c.valueList {
		ByteOrder.PutUint32(data, value.Value)
	}
	return data, err
}
