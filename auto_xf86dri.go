package xgb

/*
	This file was generated by xf86dri.xml on May 6 2012 5:48:47pm EDT.
	This file is automatically generated. Edit at your peril!
*/

// Xf86driInit must be called before using the XFree86-DRI extension.
func (c *Conn) Xf86driInit() error {
	reply, err := c.QueryExtension(11, "XFree86-DRI").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return newError("No extension named XFree86-DRI could be found on on the server.")
	}

	c.extLock.Lock()
	c.extensions["XFree86-DRI"] = reply.MajorOpcode
	for evNum, fun := range newExtEventFuncs["XFree86-DRI"] {
		newEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	c.extLock.Unlock()

	return nil
}

func init() {
	newExtEventFuncs["XFree86-DRI"] = make(map[int]newEventFun)
}

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

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// 'Xf86driDrmClipRect' struct definition
// Size: 8
type Xf86driDrmClipRect struct {
	X1 int16
	Y1 int16
	X2 int16
	X3 int16
}

// Struct read Xf86driDrmClipRect
func ReadXf86driDrmClipRect(buf []byte, v *Xf86driDrmClipRect) int {
	b := 0

	v.X1 = int16(Get16(buf[b:]))
	b += 2

	v.Y1 = int16(Get16(buf[b:]))
	b += 2

	v.X2 = int16(Get16(buf[b:]))
	b += 2

	v.X3 = int16(Get16(buf[b:]))
	b += 2

	return b
}

// Struct list read Xf86driDrmClipRect
func ReadXf86driDrmClipRectList(buf []byte, dest []Xf86driDrmClipRect) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Xf86driDrmClipRect{}
		b += ReadXf86driDrmClipRect(buf[b:], &dest[i])
	}
	return pad(b)
}

// Struct write Xf86driDrmClipRect
func (v Xf86driDrmClipRect) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	Put16(buf[b:], uint16(v.X1))
	b += 2

	Put16(buf[b:], uint16(v.Y1))
	b += 2

	Put16(buf[b:], uint16(v.X2))
	b += 2

	Put16(buf[b:], uint16(v.X3))
	b += 2

	return buf
}

// Write struct list Xf86driDrmClipRect
func Xf86driDrmClipRectListBytes(buf []byte, list []Xf86driDrmClipRect) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}
	return b
}

// Request Xf86driQueryVersion
// size: 4
type Xf86driQueryVersionCookie struct {
	*cookie
}

func (c *Conn) Xf86driQueryVersion() Xf86driQueryVersionCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xf86driQueryVersionRequest(), cookie)
	return Xf86driQueryVersionCookie{cookie}
}

func (c *Conn) Xf86driQueryVersionUnchecked() Xf86driQueryVersionCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xf86driQueryVersionRequest(), cookie)
	return Xf86driQueryVersionCookie{cookie}
}

// Request reply for Xf86driQueryVersion
// size: 16
type Xf86driQueryVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	DriMajorVersion uint16
	DriMinorVersion uint16
	DriMinorPatch   uint32
}

// Waits and reads reply data from request Xf86driQueryVersion
func (cook Xf86driQueryVersionCookie) Reply() (*Xf86driQueryVersionReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xf86driQueryVersionReply(buf), nil
}

// Read reply into structure from buffer for Xf86driQueryVersion
func xf86driQueryVersionReply(buf []byte) *Xf86driQueryVersionReply {
	v := new(Xf86driQueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.DriMajorVersion = Get16(buf[b:])
	b += 2

	v.DriMinorVersion = Get16(buf[b:])
	b += 2

	v.DriMinorPatch = Get32(buf[b:])
	b += 4

	return v
}

func (cook Xf86driQueryVersionCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driQueryVersion
func (c *Conn) xf86driQueryVersionRequest() []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request Xf86driQueryDirectRenderingCapable
// size: 8
type Xf86driQueryDirectRenderingCapableCookie struct {
	*cookie
}

func (c *Conn) Xf86driQueryDirectRenderingCapable(Screen uint32) Xf86driQueryDirectRenderingCapableCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xf86driQueryDirectRenderingCapableRequest(Screen), cookie)
	return Xf86driQueryDirectRenderingCapableCookie{cookie}
}

func (c *Conn) Xf86driQueryDirectRenderingCapableUnchecked(Screen uint32) Xf86driQueryDirectRenderingCapableCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xf86driQueryDirectRenderingCapableRequest(Screen), cookie)
	return Xf86driQueryDirectRenderingCapableCookie{cookie}
}

// Request reply for Xf86driQueryDirectRenderingCapable
// size: 9
type Xf86driQueryDirectRenderingCapableReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	IsCapable bool
}

// Waits and reads reply data from request Xf86driQueryDirectRenderingCapable
func (cook Xf86driQueryDirectRenderingCapableCookie) Reply() (*Xf86driQueryDirectRenderingCapableReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xf86driQueryDirectRenderingCapableReply(buf), nil
}

// Read reply into structure from buffer for Xf86driQueryDirectRenderingCapable
func xf86driQueryDirectRenderingCapableReply(buf []byte) *Xf86driQueryDirectRenderingCapableReply {
	v := new(Xf86driQueryDirectRenderingCapableReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	if buf[b] == 1 {
		v.IsCapable = true
	} else {
		v.IsCapable = false
	}
	b += 1

	return v
}

func (cook Xf86driQueryDirectRenderingCapableCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driQueryDirectRenderingCapable
func (c *Conn) xf86driQueryDirectRenderingCapableRequest(Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Screen)
	b += 4

	return buf
}

// Request Xf86driOpenConnection
// size: 8
type Xf86driOpenConnectionCookie struct {
	*cookie
}

func (c *Conn) Xf86driOpenConnection(Screen uint32) Xf86driOpenConnectionCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xf86driOpenConnectionRequest(Screen), cookie)
	return Xf86driOpenConnectionCookie{cookie}
}

func (c *Conn) Xf86driOpenConnectionUnchecked(Screen uint32) Xf86driOpenConnectionCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xf86driOpenConnectionRequest(Screen), cookie)
	return Xf86driOpenConnectionCookie{cookie}
}

// Request reply for Xf86driOpenConnection
// size: (32 + pad((int(BusIdLen) * 1)))
type Xf86driOpenConnectionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	SareaHandleLow  uint32
	SareaHandleHigh uint32
	BusIdLen        uint32
	// padding: 12 bytes
	BusId string // size: pad((int(BusIdLen) * 1))
}

// Waits and reads reply data from request Xf86driOpenConnection
func (cook Xf86driOpenConnectionCookie) Reply() (*Xf86driOpenConnectionReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xf86driOpenConnectionReply(buf), nil
}

// Read reply into structure from buffer for Xf86driOpenConnection
func xf86driOpenConnectionReply(buf []byte) *Xf86driOpenConnectionReply {
	v := new(Xf86driOpenConnectionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.SareaHandleLow = Get32(buf[b:])
	b += 4

	v.SareaHandleHigh = Get32(buf[b:])
	b += 4

	v.BusIdLen = Get32(buf[b:])
	b += 4

	b += 12 // padding

	{
		byteString := make([]byte, v.BusIdLen)
		copy(byteString[:v.BusIdLen], buf[b:])
		v.BusId = string(byteString)
		b += pad(int(v.BusIdLen))
	}

	return v
}

func (cook Xf86driOpenConnectionCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driOpenConnection
func (c *Conn) xf86driOpenConnectionRequest(Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Screen)
	b += 4

	return buf
}

// Request Xf86driCloseConnection
// size: 8
type Xf86driCloseConnectionCookie struct {
	*cookie
}

// Write request to wire for Xf86driCloseConnection
func (c *Conn) Xf86driCloseConnection(Screen uint32) Xf86driCloseConnectionCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.xf86driCloseConnectionRequest(Screen), cookie)
	return Xf86driCloseConnectionCookie{cookie}
}

func (c *Conn) Xf86driCloseConnectionChecked(Screen uint32) Xf86driCloseConnectionCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.xf86driCloseConnectionRequest(Screen), cookie)
	return Xf86driCloseConnectionCookie{cookie}
}

func (cook Xf86driCloseConnectionCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driCloseConnection
func (c *Conn) xf86driCloseConnectionRequest(Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Screen)
	b += 4

	return buf
}

// Request Xf86driGetClientDriverName
// size: 8
type Xf86driGetClientDriverNameCookie struct {
	*cookie
}

func (c *Conn) Xf86driGetClientDriverName(Screen uint32) Xf86driGetClientDriverNameCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xf86driGetClientDriverNameRequest(Screen), cookie)
	return Xf86driGetClientDriverNameCookie{cookie}
}

func (c *Conn) Xf86driGetClientDriverNameUnchecked(Screen uint32) Xf86driGetClientDriverNameCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xf86driGetClientDriverNameRequest(Screen), cookie)
	return Xf86driGetClientDriverNameCookie{cookie}
}

// Request reply for Xf86driGetClientDriverName
// size: (32 + pad((int(ClientDriverNameLen) * 1)))
type Xf86driGetClientDriverNameReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	ClientDriverMajorVersion uint32
	ClientDriverMinorVersion uint32
	ClientDriverPatchVersion uint32
	ClientDriverNameLen      uint32
	// padding: 8 bytes
	ClientDriverName string // size: pad((int(ClientDriverNameLen) * 1))
}

// Waits and reads reply data from request Xf86driGetClientDriverName
func (cook Xf86driGetClientDriverNameCookie) Reply() (*Xf86driGetClientDriverNameReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xf86driGetClientDriverNameReply(buf), nil
}

// Read reply into structure from buffer for Xf86driGetClientDriverName
func xf86driGetClientDriverNameReply(buf []byte) *Xf86driGetClientDriverNameReply {
	v := new(Xf86driGetClientDriverNameReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.ClientDriverMajorVersion = Get32(buf[b:])
	b += 4

	v.ClientDriverMinorVersion = Get32(buf[b:])
	b += 4

	v.ClientDriverPatchVersion = Get32(buf[b:])
	b += 4

	v.ClientDriverNameLen = Get32(buf[b:])
	b += 4

	b += 8 // padding

	{
		byteString := make([]byte, v.ClientDriverNameLen)
		copy(byteString[:v.ClientDriverNameLen], buf[b:])
		v.ClientDriverName = string(byteString)
		b += pad(int(v.ClientDriverNameLen))
	}

	return v
}

func (cook Xf86driGetClientDriverNameCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driGetClientDriverName
func (c *Conn) xf86driGetClientDriverNameRequest(Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Screen)
	b += 4

	return buf
}

// Request Xf86driCreateContext
// size: 16
type Xf86driCreateContextCookie struct {
	*cookie
}

func (c *Conn) Xf86driCreateContext(Screen uint32, Visual uint32, Context uint32) Xf86driCreateContextCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xf86driCreateContextRequest(Screen, Visual, Context), cookie)
	return Xf86driCreateContextCookie{cookie}
}

func (c *Conn) Xf86driCreateContextUnchecked(Screen uint32, Visual uint32, Context uint32) Xf86driCreateContextCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xf86driCreateContextRequest(Screen, Visual, Context), cookie)
	return Xf86driCreateContextCookie{cookie}
}

// Request reply for Xf86driCreateContext
// size: 12
type Xf86driCreateContextReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	HwContext uint32
}

// Waits and reads reply data from request Xf86driCreateContext
func (cook Xf86driCreateContextCookie) Reply() (*Xf86driCreateContextReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xf86driCreateContextReply(buf), nil
}

// Read reply into structure from buffer for Xf86driCreateContext
func xf86driCreateContextReply(buf []byte) *Xf86driCreateContextReply {
	v := new(Xf86driCreateContextReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.HwContext = Get32(buf[b:])
	b += 4

	return v
}

func (cook Xf86driCreateContextCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driCreateContext
func (c *Conn) xf86driCreateContextRequest(Screen uint32, Visual uint32, Context uint32) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Screen)
	b += 4

	Put32(buf[b:], Visual)
	b += 4

	Put32(buf[b:], Context)
	b += 4

	return buf
}

// Request Xf86driDestroyContext
// size: 12
type Xf86driDestroyContextCookie struct {
	*cookie
}

// Write request to wire for Xf86driDestroyContext
func (c *Conn) Xf86driDestroyContext(Screen uint32, Context uint32) Xf86driDestroyContextCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.xf86driDestroyContextRequest(Screen, Context), cookie)
	return Xf86driDestroyContextCookie{cookie}
}

func (c *Conn) Xf86driDestroyContextChecked(Screen uint32, Context uint32) Xf86driDestroyContextCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.xf86driDestroyContextRequest(Screen, Context), cookie)
	return Xf86driDestroyContextCookie{cookie}
}

func (cook Xf86driDestroyContextCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driDestroyContext
func (c *Conn) xf86driDestroyContextRequest(Screen uint32, Context uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Screen)
	b += 4

	Put32(buf[b:], Context)
	b += 4

	return buf
}

// Request Xf86driCreateDrawable
// size: 12
type Xf86driCreateDrawableCookie struct {
	*cookie
}

func (c *Conn) Xf86driCreateDrawable(Screen uint32, Drawable uint32) Xf86driCreateDrawableCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xf86driCreateDrawableRequest(Screen, Drawable), cookie)
	return Xf86driCreateDrawableCookie{cookie}
}

func (c *Conn) Xf86driCreateDrawableUnchecked(Screen uint32, Drawable uint32) Xf86driCreateDrawableCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xf86driCreateDrawableRequest(Screen, Drawable), cookie)
	return Xf86driCreateDrawableCookie{cookie}
}

// Request reply for Xf86driCreateDrawable
// size: 12
type Xf86driCreateDrawableReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	HwDrawableHandle uint32
}

// Waits and reads reply data from request Xf86driCreateDrawable
func (cook Xf86driCreateDrawableCookie) Reply() (*Xf86driCreateDrawableReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xf86driCreateDrawableReply(buf), nil
}

// Read reply into structure from buffer for Xf86driCreateDrawable
func xf86driCreateDrawableReply(buf []byte) *Xf86driCreateDrawableReply {
	v := new(Xf86driCreateDrawableReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.HwDrawableHandle = Get32(buf[b:])
	b += 4

	return v
}

func (cook Xf86driCreateDrawableCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driCreateDrawable
func (c *Conn) xf86driCreateDrawableRequest(Screen uint32, Drawable uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Screen)
	b += 4

	Put32(buf[b:], Drawable)
	b += 4

	return buf
}

// Request Xf86driDestroyDrawable
// size: 12
type Xf86driDestroyDrawableCookie struct {
	*cookie
}

// Write request to wire for Xf86driDestroyDrawable
func (c *Conn) Xf86driDestroyDrawable(Screen uint32, Drawable uint32) Xf86driDestroyDrawableCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.xf86driDestroyDrawableRequest(Screen, Drawable), cookie)
	return Xf86driDestroyDrawableCookie{cookie}
}

func (c *Conn) Xf86driDestroyDrawableChecked(Screen uint32, Drawable uint32) Xf86driDestroyDrawableCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.xf86driDestroyDrawableRequest(Screen, Drawable), cookie)
	return Xf86driDestroyDrawableCookie{cookie}
}

func (cook Xf86driDestroyDrawableCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driDestroyDrawable
func (c *Conn) xf86driDestroyDrawableRequest(Screen uint32, Drawable uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Screen)
	b += 4

	Put32(buf[b:], Drawable)
	b += 4

	return buf
}

// Request Xf86driGetDrawableInfo
// size: 12
type Xf86driGetDrawableInfoCookie struct {
	*cookie
}

func (c *Conn) Xf86driGetDrawableInfo(Screen uint32, Drawable uint32) Xf86driGetDrawableInfoCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xf86driGetDrawableInfoRequest(Screen, Drawable), cookie)
	return Xf86driGetDrawableInfoCookie{cookie}
}

func (c *Conn) Xf86driGetDrawableInfoUnchecked(Screen uint32, Drawable uint32) Xf86driGetDrawableInfoCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xf86driGetDrawableInfoRequest(Screen, Drawable), cookie)
	return Xf86driGetDrawableInfoCookie{cookie}
}

// Request reply for Xf86driGetDrawableInfo
// size: ((36 + pad((int(NumClipRects) * 8))) + pad((int(NumBackClipRects) * 8)))
type Xf86driGetDrawableInfoReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	DrawableTableIndex uint32
	DrawableTableStamp uint32
	DrawableOriginX    int16
	DrawableOriginY    int16
	DrawableSizeW      int16
	DrawableSizeH      int16
	NumClipRects       uint32
	BackX              int16
	BackY              int16
	NumBackClipRects   uint32
	ClipRects          []Xf86driDrmClipRect // size: pad((int(NumClipRects) * 8))
	BackClipRects      []Xf86driDrmClipRect // size: pad((int(NumBackClipRects) * 8))
}

// Waits and reads reply data from request Xf86driGetDrawableInfo
func (cook Xf86driGetDrawableInfoCookie) Reply() (*Xf86driGetDrawableInfoReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xf86driGetDrawableInfoReply(buf), nil
}

// Read reply into structure from buffer for Xf86driGetDrawableInfo
func xf86driGetDrawableInfoReply(buf []byte) *Xf86driGetDrawableInfoReply {
	v := new(Xf86driGetDrawableInfoReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.DrawableTableIndex = Get32(buf[b:])
	b += 4

	v.DrawableTableStamp = Get32(buf[b:])
	b += 4

	v.DrawableOriginX = int16(Get16(buf[b:]))
	b += 2

	v.DrawableOriginY = int16(Get16(buf[b:]))
	b += 2

	v.DrawableSizeW = int16(Get16(buf[b:]))
	b += 2

	v.DrawableSizeH = int16(Get16(buf[b:]))
	b += 2

	v.NumClipRects = Get32(buf[b:])
	b += 4

	v.BackX = int16(Get16(buf[b:]))
	b += 2

	v.BackY = int16(Get16(buf[b:]))
	b += 2

	v.NumBackClipRects = Get32(buf[b:])
	b += 4

	v.ClipRects = make([]Xf86driDrmClipRect, v.NumClipRects)
	b += ReadXf86driDrmClipRectList(buf[b:], v.ClipRects)

	v.BackClipRects = make([]Xf86driDrmClipRect, v.NumBackClipRects)
	b += ReadXf86driDrmClipRectList(buf[b:], v.BackClipRects)

	return v
}

func (cook Xf86driGetDrawableInfoCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driGetDrawableInfo
func (c *Conn) xf86driGetDrawableInfoRequest(Screen uint32, Drawable uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 9 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Screen)
	b += 4

	Put32(buf[b:], Drawable)
	b += 4

	return buf
}

// Request Xf86driGetDeviceInfo
// size: 8
type Xf86driGetDeviceInfoCookie struct {
	*cookie
}

func (c *Conn) Xf86driGetDeviceInfo(Screen uint32) Xf86driGetDeviceInfoCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xf86driGetDeviceInfoRequest(Screen), cookie)
	return Xf86driGetDeviceInfoCookie{cookie}
}

func (c *Conn) Xf86driGetDeviceInfoUnchecked(Screen uint32) Xf86driGetDeviceInfoCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xf86driGetDeviceInfoRequest(Screen), cookie)
	return Xf86driGetDeviceInfoCookie{cookie}
}

// Request reply for Xf86driGetDeviceInfo
// size: (32 + pad((int(DevicePrivateSize) * 4)))
type Xf86driGetDeviceInfoReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	FramebufferHandleLow    uint32
	FramebufferHandleHigh   uint32
	FramebufferOriginOffset uint32
	FramebufferSize         uint32
	FramebufferStride       uint32
	DevicePrivateSize       uint32
	DevicePrivate           []uint32 // size: pad((int(DevicePrivateSize) * 4))
}

// Waits and reads reply data from request Xf86driGetDeviceInfo
func (cook Xf86driGetDeviceInfoCookie) Reply() (*Xf86driGetDeviceInfoReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xf86driGetDeviceInfoReply(buf), nil
}

// Read reply into structure from buffer for Xf86driGetDeviceInfo
func xf86driGetDeviceInfoReply(buf []byte) *Xf86driGetDeviceInfoReply {
	v := new(Xf86driGetDeviceInfoReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.FramebufferHandleLow = Get32(buf[b:])
	b += 4

	v.FramebufferHandleHigh = Get32(buf[b:])
	b += 4

	v.FramebufferOriginOffset = Get32(buf[b:])
	b += 4

	v.FramebufferSize = Get32(buf[b:])
	b += 4

	v.FramebufferStride = Get32(buf[b:])
	b += 4

	v.DevicePrivateSize = Get32(buf[b:])
	b += 4

	v.DevicePrivate = make([]uint32, v.DevicePrivateSize)
	for i := 0; i < int(v.DevicePrivateSize); i++ {
		v.DevicePrivate[i] = Get32(buf[b:])
		b += 4
	}
	b = pad(b)

	return v
}

func (cook Xf86driGetDeviceInfoCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driGetDeviceInfo
func (c *Conn) xf86driGetDeviceInfoRequest(Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 10 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Screen)
	b += 4

	return buf
}

// Request Xf86driAuthConnection
// size: 12
type Xf86driAuthConnectionCookie struct {
	*cookie
}

func (c *Conn) Xf86driAuthConnection(Screen uint32, Magic uint32) Xf86driAuthConnectionCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xf86driAuthConnectionRequest(Screen, Magic), cookie)
	return Xf86driAuthConnectionCookie{cookie}
}

func (c *Conn) Xf86driAuthConnectionUnchecked(Screen uint32, Magic uint32) Xf86driAuthConnectionCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xf86driAuthConnectionRequest(Screen, Magic), cookie)
	return Xf86driAuthConnectionCookie{cookie}
}

// Request reply for Xf86driAuthConnection
// size: 12
type Xf86driAuthConnectionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Authenticated uint32
}

// Waits and reads reply data from request Xf86driAuthConnection
func (cook Xf86driAuthConnectionCookie) Reply() (*Xf86driAuthConnectionReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xf86driAuthConnectionReply(buf), nil
}

// Read reply into structure from buffer for Xf86driAuthConnection
func xf86driAuthConnectionReply(buf []byte) *Xf86driAuthConnectionReply {
	v := new(Xf86driAuthConnectionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.Authenticated = Get32(buf[b:])
	b += 4

	return v
}

func (cook Xf86driAuthConnectionCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xf86driAuthConnection
func (c *Conn) xf86driAuthConnectionRequest(Screen uint32, Magic uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 11 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Screen)
	b += 4

	Put32(buf[b:], Magic)
	b += 4

	return buf
}
