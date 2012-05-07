package xgb

/*
	This file was generated by xvmc.xml on May 6 2012 5:48:48pm EDT.
	This file is automatically generated. Edit at your peril!
*/

// Imports are not necessary for XGB because everything is 
// in one package. They are still listed here for reference.
// import "xv"

// XvmcInit must be called before using the XVideo-MotionCompensation extension.
func (c *Conn) XvmcInit() error {
	reply, err := c.QueryExtension(25, "XVideo-MotionCompensation").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return newError("No extension named XVideo-MotionCompensation could be found on on the server.")
	}

	c.extLock.Lock()
	c.extensions["XVideo-MotionCompensation"] = reply.MajorOpcode
	for evNum, fun := range newExtEventFuncs["XVideo-MotionCompensation"] {
		newEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	c.extLock.Unlock()

	return nil
}

func init() {
	newExtEventFuncs["XVideo-MotionCompensation"] = make(map[int]newEventFun)
}

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Id'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping resource definition of 'Context'

// Skipping resource definition of 'Surface'

// Skipping resource definition of 'Subpicture'

// 'XvmcSurfaceInfo' struct definition
// Size: 24
type XvmcSurfaceInfo struct {
	Id                  Id
	ChromaFormat        uint16
	Pad0                uint16
	MaxWidth            uint16
	MaxHeight           uint16
	SubpictureMaxWidth  uint16
	SubpictureMaxHeight uint16
	McType              uint32
	Flags               uint32
}

// Struct read XvmcSurfaceInfo
func ReadXvmcSurfaceInfo(buf []byte, v *XvmcSurfaceInfo) int {
	b := 0

	v.Id = Id(Get32(buf[b:]))
	b += 4

	v.ChromaFormat = Get16(buf[b:])
	b += 2

	v.Pad0 = Get16(buf[b:])
	b += 2

	v.MaxWidth = Get16(buf[b:])
	b += 2

	v.MaxHeight = Get16(buf[b:])
	b += 2

	v.SubpictureMaxWidth = Get16(buf[b:])
	b += 2

	v.SubpictureMaxHeight = Get16(buf[b:])
	b += 2

	v.McType = Get32(buf[b:])
	b += 4

	v.Flags = Get32(buf[b:])
	b += 4

	return b
}

// Struct list read XvmcSurfaceInfo
func ReadXvmcSurfaceInfoList(buf []byte, dest []XvmcSurfaceInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = XvmcSurfaceInfo{}
		b += ReadXvmcSurfaceInfo(buf[b:], &dest[i])
	}
	return pad(b)
}

// Struct write XvmcSurfaceInfo
func (v XvmcSurfaceInfo) Bytes() []byte {
	buf := make([]byte, 24)
	b := 0

	Put32(buf[b:], uint32(v.Id))
	b += 4

	Put16(buf[b:], v.ChromaFormat)
	b += 2

	Put16(buf[b:], v.Pad0)
	b += 2

	Put16(buf[b:], v.MaxWidth)
	b += 2

	Put16(buf[b:], v.MaxHeight)
	b += 2

	Put16(buf[b:], v.SubpictureMaxWidth)
	b += 2

	Put16(buf[b:], v.SubpictureMaxHeight)
	b += 2

	Put32(buf[b:], v.McType)
	b += 4

	Put32(buf[b:], v.Flags)
	b += 4

	return buf
}

// Write struct list XvmcSurfaceInfo
func XvmcSurfaceInfoListBytes(buf []byte, list []XvmcSurfaceInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}
	return b
}

// Request XvmcQueryVersion
// size: 4
type XvmcQueryVersionCookie struct {
	*cookie
}

func (c *Conn) XvmcQueryVersion() XvmcQueryVersionCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xvmcQueryVersionRequest(), cookie)
	return XvmcQueryVersionCookie{cookie}
}

func (c *Conn) XvmcQueryVersionUnchecked() XvmcQueryVersionCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xvmcQueryVersionRequest(), cookie)
	return XvmcQueryVersionCookie{cookie}
}

// Request reply for XvmcQueryVersion
// size: 16
type XvmcQueryVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Major uint32
	Minor uint32
}

// Waits and reads reply data from request XvmcQueryVersion
func (cook XvmcQueryVersionCookie) Reply() (*XvmcQueryVersionReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xvmcQueryVersionReply(buf), nil
}

// Read reply into structure from buffer for XvmcQueryVersion
func xvmcQueryVersionReply(buf []byte) *XvmcQueryVersionReply {
	v := new(XvmcQueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.Major = Get32(buf[b:])
	b += 4

	v.Minor = Get32(buf[b:])
	b += 4

	return v
}

func (cook XvmcQueryVersionCookie) Check() error {
	return cook.check()
}

// Write request to wire for XvmcQueryVersion
func (c *Conn) xvmcQueryVersionRequest() []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request XvmcListSurfaceTypes
// size: 8
type XvmcListSurfaceTypesCookie struct {
	*cookie
}

func (c *Conn) XvmcListSurfaceTypes(PortId Id) XvmcListSurfaceTypesCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xvmcListSurfaceTypesRequest(PortId), cookie)
	return XvmcListSurfaceTypesCookie{cookie}
}

func (c *Conn) XvmcListSurfaceTypesUnchecked(PortId Id) XvmcListSurfaceTypesCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xvmcListSurfaceTypesRequest(PortId), cookie)
	return XvmcListSurfaceTypesCookie{cookie}
}

// Request reply for XvmcListSurfaceTypes
// size: (32 + pad((int(Num) * 24)))
type XvmcListSurfaceTypesReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Num uint32
	// padding: 20 bytes
	Surfaces []XvmcSurfaceInfo // size: pad((int(Num) * 24))
}

// Waits and reads reply data from request XvmcListSurfaceTypes
func (cook XvmcListSurfaceTypesCookie) Reply() (*XvmcListSurfaceTypesReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xvmcListSurfaceTypesReply(buf), nil
}

// Read reply into structure from buffer for XvmcListSurfaceTypes
func xvmcListSurfaceTypesReply(buf []byte) *XvmcListSurfaceTypesReply {
	v := new(XvmcListSurfaceTypesReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.Num = Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Surfaces = make([]XvmcSurfaceInfo, v.Num)
	b += ReadXvmcSurfaceInfoList(buf[b:], v.Surfaces)

	return v
}

func (cook XvmcListSurfaceTypesCookie) Check() error {
	return cook.check()
}

// Write request to wire for XvmcListSurfaceTypes
func (c *Conn) xvmcListSurfaceTypesRequest(PortId Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(PortId))
	b += 4

	return buf
}

// Request XvmcCreateContext
// size: 24
type XvmcCreateContextCookie struct {
	*cookie
}

func (c *Conn) XvmcCreateContext(ContextId Id, PortId Id, SurfaceId Id, Width uint16, Height uint16, Flags uint32) XvmcCreateContextCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xvmcCreateContextRequest(ContextId, PortId, SurfaceId, Width, Height, Flags), cookie)
	return XvmcCreateContextCookie{cookie}
}

func (c *Conn) XvmcCreateContextUnchecked(ContextId Id, PortId Id, SurfaceId Id, Width uint16, Height uint16, Flags uint32) XvmcCreateContextCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xvmcCreateContextRequest(ContextId, PortId, SurfaceId, Width, Height, Flags), cookie)
	return XvmcCreateContextCookie{cookie}
}

// Request reply for XvmcCreateContext
// size: (36 + pad((int(Length) * 4)))
type XvmcCreateContextReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	WidthActual  uint16
	HeightActual uint16
	FlagsReturn  uint32
	// padding: 20 bytes
	PrivData []uint32 // size: pad((int(Length) * 4))
}

// Waits and reads reply data from request XvmcCreateContext
func (cook XvmcCreateContextCookie) Reply() (*XvmcCreateContextReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xvmcCreateContextReply(buf), nil
}

// Read reply into structure from buffer for XvmcCreateContext
func xvmcCreateContextReply(buf []byte) *XvmcCreateContextReply {
	v := new(XvmcCreateContextReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.WidthActual = Get16(buf[b:])
	b += 2

	v.HeightActual = Get16(buf[b:])
	b += 2

	v.FlagsReturn = Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.PrivData = make([]uint32, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.PrivData[i] = Get32(buf[b:])
		b += 4
	}
	b = pad(b)

	return v
}

func (cook XvmcCreateContextCookie) Check() error {
	return cook.check()
}

// Write request to wire for XvmcCreateContext
func (c *Conn) xvmcCreateContextRequest(ContextId Id, PortId Id, SurfaceId Id, Width uint16, Height uint16, Flags uint32) []byte {
	size := 24
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(ContextId))
	b += 4

	Put32(buf[b:], uint32(PortId))
	b += 4

	Put32(buf[b:], uint32(SurfaceId))
	b += 4

	Put16(buf[b:], Width)
	b += 2

	Put16(buf[b:], Height)
	b += 2

	Put32(buf[b:], Flags)
	b += 4

	return buf
}

// Request XvmcDestroyContext
// size: 8
type XvmcDestroyContextCookie struct {
	*cookie
}

// Write request to wire for XvmcDestroyContext
func (c *Conn) XvmcDestroyContext(ContextId Id) XvmcDestroyContextCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.xvmcDestroyContextRequest(ContextId), cookie)
	return XvmcDestroyContextCookie{cookie}
}

func (c *Conn) XvmcDestroyContextChecked(ContextId Id) XvmcDestroyContextCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.xvmcDestroyContextRequest(ContextId), cookie)
	return XvmcDestroyContextCookie{cookie}
}

func (cook XvmcDestroyContextCookie) Check() error {
	return cook.check()
}

// Write request to wire for XvmcDestroyContext
func (c *Conn) xvmcDestroyContextRequest(ContextId Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(ContextId))
	b += 4

	return buf
}

// Request XvmcCreateSurface
// size: 12
type XvmcCreateSurfaceCookie struct {
	*cookie
}

func (c *Conn) XvmcCreateSurface(SurfaceId Id, ContextId Id) XvmcCreateSurfaceCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xvmcCreateSurfaceRequest(SurfaceId, ContextId), cookie)
	return XvmcCreateSurfaceCookie{cookie}
}

func (c *Conn) XvmcCreateSurfaceUnchecked(SurfaceId Id, ContextId Id) XvmcCreateSurfaceCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xvmcCreateSurfaceRequest(SurfaceId, ContextId), cookie)
	return XvmcCreateSurfaceCookie{cookie}
}

// Request reply for XvmcCreateSurface
// size: (32 + pad((int(Length) * 4)))
type XvmcCreateSurfaceReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	// padding: 24 bytes
	PrivData []uint32 // size: pad((int(Length) * 4))
}

// Waits and reads reply data from request XvmcCreateSurface
func (cook XvmcCreateSurfaceCookie) Reply() (*XvmcCreateSurfaceReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xvmcCreateSurfaceReply(buf), nil
}

// Read reply into structure from buffer for XvmcCreateSurface
func xvmcCreateSurfaceReply(buf []byte) *XvmcCreateSurfaceReply {
	v := new(XvmcCreateSurfaceReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	v.PrivData = make([]uint32, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.PrivData[i] = Get32(buf[b:])
		b += 4
	}
	b = pad(b)

	return v
}

func (cook XvmcCreateSurfaceCookie) Check() error {
	return cook.check()
}

// Write request to wire for XvmcCreateSurface
func (c *Conn) xvmcCreateSurfaceRequest(SurfaceId Id, ContextId Id) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(SurfaceId))
	b += 4

	Put32(buf[b:], uint32(ContextId))
	b += 4

	return buf
}

// Request XvmcDestroySurface
// size: 8
type XvmcDestroySurfaceCookie struct {
	*cookie
}

// Write request to wire for XvmcDestroySurface
func (c *Conn) XvmcDestroySurface(SurfaceId Id) XvmcDestroySurfaceCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.xvmcDestroySurfaceRequest(SurfaceId), cookie)
	return XvmcDestroySurfaceCookie{cookie}
}

func (c *Conn) XvmcDestroySurfaceChecked(SurfaceId Id) XvmcDestroySurfaceCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.xvmcDestroySurfaceRequest(SurfaceId), cookie)
	return XvmcDestroySurfaceCookie{cookie}
}

func (cook XvmcDestroySurfaceCookie) Check() error {
	return cook.check()
}

// Write request to wire for XvmcDestroySurface
func (c *Conn) xvmcDestroySurfaceRequest(SurfaceId Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(SurfaceId))
	b += 4

	return buf
}

// Request XvmcCreateSubpicture
// size: 20
type XvmcCreateSubpictureCookie struct {
	*cookie
}

func (c *Conn) XvmcCreateSubpicture(SubpictureId Id, Context Id, XvimageId uint32, Width uint16, Height uint16) XvmcCreateSubpictureCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xvmcCreateSubpictureRequest(SubpictureId, Context, XvimageId, Width, Height), cookie)
	return XvmcCreateSubpictureCookie{cookie}
}

func (c *Conn) XvmcCreateSubpictureUnchecked(SubpictureId Id, Context Id, XvimageId uint32, Width uint16, Height uint16) XvmcCreateSubpictureCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xvmcCreateSubpictureRequest(SubpictureId, Context, XvimageId, Width, Height), cookie)
	return XvmcCreateSubpictureCookie{cookie}
}

// Request reply for XvmcCreateSubpicture
// size: (32 + pad((int(Length) * 4)))
type XvmcCreateSubpictureReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	WidthActual       uint16
	HeightActual      uint16
	NumPaletteEntries uint16
	EntryBytes        uint16
	ComponentOrder    []byte // size: 4
	// padding: 12 bytes
	PrivData []uint32 // size: pad((int(Length) * 4))
}

// Waits and reads reply data from request XvmcCreateSubpicture
func (cook XvmcCreateSubpictureCookie) Reply() (*XvmcCreateSubpictureReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xvmcCreateSubpictureReply(buf), nil
}

// Read reply into structure from buffer for XvmcCreateSubpicture
func xvmcCreateSubpictureReply(buf []byte) *XvmcCreateSubpictureReply {
	v := new(XvmcCreateSubpictureReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.WidthActual = Get16(buf[b:])
	b += 2

	v.HeightActual = Get16(buf[b:])
	b += 2

	v.NumPaletteEntries = Get16(buf[b:])
	b += 2

	v.EntryBytes = Get16(buf[b:])
	b += 2

	v.ComponentOrder = make([]byte, 4)
	copy(v.ComponentOrder[:4], buf[b:])
	b += pad(int(4))

	b += 12 // padding

	v.PrivData = make([]uint32, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.PrivData[i] = Get32(buf[b:])
		b += 4
	}
	b = pad(b)

	return v
}

func (cook XvmcCreateSubpictureCookie) Check() error {
	return cook.check()
}

// Write request to wire for XvmcCreateSubpicture
func (c *Conn) xvmcCreateSubpictureRequest(SubpictureId Id, Context Id, XvimageId uint32, Width uint16, Height uint16) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(SubpictureId))
	b += 4

	Put32(buf[b:], uint32(Context))
	b += 4

	Put32(buf[b:], XvimageId)
	b += 4

	Put16(buf[b:], Width)
	b += 2

	Put16(buf[b:], Height)
	b += 2

	return buf
}

// Request XvmcDestroySubpicture
// size: 8
type XvmcDestroySubpictureCookie struct {
	*cookie
}

// Write request to wire for XvmcDestroySubpicture
func (c *Conn) XvmcDestroySubpicture(SubpictureId Id) XvmcDestroySubpictureCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.xvmcDestroySubpictureRequest(SubpictureId), cookie)
	return XvmcDestroySubpictureCookie{cookie}
}

func (c *Conn) XvmcDestroySubpictureChecked(SubpictureId Id) XvmcDestroySubpictureCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.xvmcDestroySubpictureRequest(SubpictureId), cookie)
	return XvmcDestroySubpictureCookie{cookie}
}

func (cook XvmcDestroySubpictureCookie) Check() error {
	return cook.check()
}

// Write request to wire for XvmcDestroySubpicture
func (c *Conn) xvmcDestroySubpictureRequest(SubpictureId Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(SubpictureId))
	b += 4

	return buf
}

// Request XvmcListSubpictureTypes
// size: 12
type XvmcListSubpictureTypesCookie struct {
	*cookie
}

func (c *Conn) XvmcListSubpictureTypes(PortId Id, SurfaceId Id) XvmcListSubpictureTypesCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xvmcListSubpictureTypesRequest(PortId, SurfaceId), cookie)
	return XvmcListSubpictureTypesCookie{cookie}
}

func (c *Conn) XvmcListSubpictureTypesUnchecked(PortId Id, SurfaceId Id) XvmcListSubpictureTypesCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xvmcListSubpictureTypesRequest(PortId, SurfaceId), cookie)
	return XvmcListSubpictureTypesCookie{cookie}
}

// Request reply for XvmcListSubpictureTypes
// size: (32 + XvImageFormatInfoListSize(Types))
type XvmcListSubpictureTypesReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Num uint32
	// padding: 20 bytes
	Types []XvImageFormatInfo // size: XvImageFormatInfoListSize(Types)
}

// Waits and reads reply data from request XvmcListSubpictureTypes
func (cook XvmcListSubpictureTypesCookie) Reply() (*XvmcListSubpictureTypesReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xvmcListSubpictureTypesReply(buf), nil
}

// Read reply into structure from buffer for XvmcListSubpictureTypes
func xvmcListSubpictureTypesReply(buf []byte) *XvmcListSubpictureTypesReply {
	v := new(XvmcListSubpictureTypesReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.Num = Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Types = make([]XvImageFormatInfo, v.Num)
	b += ReadXvImageFormatInfoList(buf[b:], v.Types)

	return v
}

func (cook XvmcListSubpictureTypesCookie) Check() error {
	return cook.check()
}

// Write request to wire for XvmcListSubpictureTypes
func (c *Conn) xvmcListSubpictureTypesRequest(PortId Id, SurfaceId Id) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(PortId))
	b += 4

	Put32(buf[b:], uint32(SurfaceId))
	b += 4

	return buf
}
