package formats

import "encoding/binary"

//OpCode corresponds to a byte representing the operation related to a
//particular format
type OpCode uint8

//ByteOrder is the byte order for handling requests
var ByteOrder = binary.BigEndian
