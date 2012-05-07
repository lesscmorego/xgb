package xgb

/*
	This file was generated by composite.xml on May 6 2012 5:48:46pm EDT.
	This file is automatically generated. Edit at your peril!
*/

// Imports are not necessary for XGB because everything is 
// in one package. They are still listed here for reference.
// import "xproto"
// import "xfixes"

// CompositeInit must be called before using the Composite extension.
func (c *Conn) CompositeInit() error {
	reply, err := c.QueryExtension(9, "Composite").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return newError("No extension named Composite could be found on on the server.")
	}

	c.extLock.Lock()
	c.extensions["Composite"] = reply.MajorOpcode
	for evNum, fun := range newExtEventFuncs["Composite"] {
		newEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	c.extLock.Unlock()

	return nil
}

func init() {
	newExtEventFuncs["Composite"] = make(map[int]newEventFun)
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

const (
	CompositeRedirectAutomatic = 0
	CompositeRedirectManual    = 1
)

// Request CompositeQueryVersion
// size: 12
type CompositeQueryVersionCookie struct {
	*cookie
}

func (c *Conn) CompositeQueryVersion(ClientMajorVersion uint32, ClientMinorVersion uint32) CompositeQueryVersionCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.compositeQueryVersionRequest(ClientMajorVersion, ClientMinorVersion), cookie)
	return CompositeQueryVersionCookie{cookie}
}

func (c *Conn) CompositeQueryVersionUnchecked(ClientMajorVersion uint32, ClientMinorVersion uint32) CompositeQueryVersionCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.compositeQueryVersionRequest(ClientMajorVersion, ClientMinorVersion), cookie)
	return CompositeQueryVersionCookie{cookie}
}

// Request reply for CompositeQueryVersion
// size: 32
type CompositeQueryVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	MajorVersion uint32
	MinorVersion uint32
	// padding: 16 bytes
}

// Waits and reads reply data from request CompositeQueryVersion
func (cook CompositeQueryVersionCookie) Reply() (*CompositeQueryVersionReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return compositeQueryVersionReply(buf), nil
}

// Read reply into structure from buffer for CompositeQueryVersion
func compositeQueryVersionReply(buf []byte) *CompositeQueryVersionReply {
	v := new(CompositeQueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.MajorVersion = Get32(buf[b:])
	b += 4

	v.MinorVersion = Get32(buf[b:])
	b += 4

	b += 16 // padding

	return v
}

func (cook CompositeQueryVersionCookie) Check() error {
	return cook.check()
}

// Write request to wire for CompositeQueryVersion
func (c *Conn) compositeQueryVersionRequest(ClientMajorVersion uint32, ClientMinorVersion uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["COMPOSITE"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], ClientMajorVersion)
	b += 4

	Put32(buf[b:], ClientMinorVersion)
	b += 4

	return buf
}

// Request CompositeRedirectWindow
// size: 12
type CompositeRedirectWindowCookie struct {
	*cookie
}

// Write request to wire for CompositeRedirectWindow
func (c *Conn) CompositeRedirectWindow(Window Id, Update byte) CompositeRedirectWindowCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.compositeRedirectWindowRequest(Window, Update), cookie)
	return CompositeRedirectWindowCookie{cookie}
}

func (c *Conn) CompositeRedirectWindowChecked(Window Id, Update byte) CompositeRedirectWindowCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.compositeRedirectWindowRequest(Window, Update), cookie)
	return CompositeRedirectWindowCookie{cookie}
}

func (cook CompositeRedirectWindowCookie) Check() error {
	return cook.check()
}

// Write request to wire for CompositeRedirectWindow
func (c *Conn) compositeRedirectWindowRequest(Window Id, Update byte) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["COMPOSITE"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Window))
	b += 4

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}

// Request CompositeRedirectSubwindows
// size: 12
type CompositeRedirectSubwindowsCookie struct {
	*cookie
}

// Write request to wire for CompositeRedirectSubwindows
func (c *Conn) CompositeRedirectSubwindows(Window Id, Update byte) CompositeRedirectSubwindowsCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.compositeRedirectSubwindowsRequest(Window, Update), cookie)
	return CompositeRedirectSubwindowsCookie{cookie}
}

func (c *Conn) CompositeRedirectSubwindowsChecked(Window Id, Update byte) CompositeRedirectSubwindowsCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.compositeRedirectSubwindowsRequest(Window, Update), cookie)
	return CompositeRedirectSubwindowsCookie{cookie}
}

func (cook CompositeRedirectSubwindowsCookie) Check() error {
	return cook.check()
}

// Write request to wire for CompositeRedirectSubwindows
func (c *Conn) compositeRedirectSubwindowsRequest(Window Id, Update byte) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["COMPOSITE"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Window))
	b += 4

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}

// Request CompositeUnredirectWindow
// size: 12
type CompositeUnredirectWindowCookie struct {
	*cookie
}

// Write request to wire for CompositeUnredirectWindow
func (c *Conn) CompositeUnredirectWindow(Window Id, Update byte) CompositeUnredirectWindowCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.compositeUnredirectWindowRequest(Window, Update), cookie)
	return CompositeUnredirectWindowCookie{cookie}
}

func (c *Conn) CompositeUnredirectWindowChecked(Window Id, Update byte) CompositeUnredirectWindowCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.compositeUnredirectWindowRequest(Window, Update), cookie)
	return CompositeUnredirectWindowCookie{cookie}
}

func (cook CompositeUnredirectWindowCookie) Check() error {
	return cook.check()
}

// Write request to wire for CompositeUnredirectWindow
func (c *Conn) compositeUnredirectWindowRequest(Window Id, Update byte) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["COMPOSITE"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Window))
	b += 4

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}

// Request CompositeUnredirectSubwindows
// size: 12
type CompositeUnredirectSubwindowsCookie struct {
	*cookie
}

// Write request to wire for CompositeUnredirectSubwindows
func (c *Conn) CompositeUnredirectSubwindows(Window Id, Update byte) CompositeUnredirectSubwindowsCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.compositeUnredirectSubwindowsRequest(Window, Update), cookie)
	return CompositeUnredirectSubwindowsCookie{cookie}
}

func (c *Conn) CompositeUnredirectSubwindowsChecked(Window Id, Update byte) CompositeUnredirectSubwindowsCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.compositeUnredirectSubwindowsRequest(Window, Update), cookie)
	return CompositeUnredirectSubwindowsCookie{cookie}
}

func (cook CompositeUnredirectSubwindowsCookie) Check() error {
	return cook.check()
}

// Write request to wire for CompositeUnredirectSubwindows
func (c *Conn) compositeUnredirectSubwindowsRequest(Window Id, Update byte) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["COMPOSITE"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Window))
	b += 4

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}

// Request CompositeCreateRegionFromBorderClip
// size: 12
type CompositeCreateRegionFromBorderClipCookie struct {
	*cookie
}

// Write request to wire for CompositeCreateRegionFromBorderClip
func (c *Conn) CompositeCreateRegionFromBorderClip(Region Id, Window Id) CompositeCreateRegionFromBorderClipCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.compositeCreateRegionFromBorderClipRequest(Region, Window), cookie)
	return CompositeCreateRegionFromBorderClipCookie{cookie}
}

func (c *Conn) CompositeCreateRegionFromBorderClipChecked(Region Id, Window Id) CompositeCreateRegionFromBorderClipCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.compositeCreateRegionFromBorderClipRequest(Region, Window), cookie)
	return CompositeCreateRegionFromBorderClipCookie{cookie}
}

func (cook CompositeCreateRegionFromBorderClipCookie) Check() error {
	return cook.check()
}

// Write request to wire for CompositeCreateRegionFromBorderClip
func (c *Conn) compositeCreateRegionFromBorderClipRequest(Region Id, Window Id) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["COMPOSITE"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Region))
	b += 4

	Put32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// Request CompositeNameWindowPixmap
// size: 12
type CompositeNameWindowPixmapCookie struct {
	*cookie
}

// Write request to wire for CompositeNameWindowPixmap
func (c *Conn) CompositeNameWindowPixmap(Window Id, Pixmap Id) CompositeNameWindowPixmapCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.compositeNameWindowPixmapRequest(Window, Pixmap), cookie)
	return CompositeNameWindowPixmapCookie{cookie}
}

func (c *Conn) CompositeNameWindowPixmapChecked(Window Id, Pixmap Id) CompositeNameWindowPixmapCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.compositeNameWindowPixmapRequest(Window, Pixmap), cookie)
	return CompositeNameWindowPixmapCookie{cookie}
}

func (cook CompositeNameWindowPixmapCookie) Check() error {
	return cook.check()
}

// Write request to wire for CompositeNameWindowPixmap
func (c *Conn) compositeNameWindowPixmapRequest(Window Id, Pixmap Id) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["COMPOSITE"]
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Window))
	b += 4

	Put32(buf[b:], uint32(Pixmap))
	b += 4

	return buf
}

// Request CompositeGetOverlayWindow
// size: 8
type CompositeGetOverlayWindowCookie struct {
	*cookie
}

func (c *Conn) CompositeGetOverlayWindow(Window Id) CompositeGetOverlayWindowCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.compositeGetOverlayWindowRequest(Window), cookie)
	return CompositeGetOverlayWindowCookie{cookie}
}

func (c *Conn) CompositeGetOverlayWindowUnchecked(Window Id) CompositeGetOverlayWindowCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.compositeGetOverlayWindowRequest(Window), cookie)
	return CompositeGetOverlayWindowCookie{cookie}
}

// Request reply for CompositeGetOverlayWindow
// size: 32
type CompositeGetOverlayWindowReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	OverlayWin Id
	// padding: 20 bytes
}

// Waits and reads reply data from request CompositeGetOverlayWindow
func (cook CompositeGetOverlayWindowCookie) Reply() (*CompositeGetOverlayWindowReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return compositeGetOverlayWindowReply(buf), nil
}

// Read reply into structure from buffer for CompositeGetOverlayWindow
func compositeGetOverlayWindowReply(buf []byte) *CompositeGetOverlayWindowReply {
	v := new(CompositeGetOverlayWindowReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.OverlayWin = Id(Get32(buf[b:]))
	b += 4

	b += 20 // padding

	return v
}

func (cook CompositeGetOverlayWindowCookie) Check() error {
	return cook.check()
}

// Write request to wire for CompositeGetOverlayWindow
func (c *Conn) compositeGetOverlayWindowRequest(Window Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["COMPOSITE"]
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// Request CompositeReleaseOverlayWindow
// size: 8
type CompositeReleaseOverlayWindowCookie struct {
	*cookie
}

// Write request to wire for CompositeReleaseOverlayWindow
func (c *Conn) CompositeReleaseOverlayWindow(Window Id) CompositeReleaseOverlayWindowCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.compositeReleaseOverlayWindowRequest(Window), cookie)
	return CompositeReleaseOverlayWindowCookie{cookie}
}

func (c *Conn) CompositeReleaseOverlayWindowChecked(Window Id) CompositeReleaseOverlayWindowCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.compositeReleaseOverlayWindowRequest(Window), cookie)
	return CompositeReleaseOverlayWindowCookie{cookie}
}

func (cook CompositeReleaseOverlayWindowCookie) Check() error {
	return cook.check()
}

// Write request to wire for CompositeReleaseOverlayWindow
func (c *Conn) compositeReleaseOverlayWindowRequest(Window Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["COMPOSITE"]
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Window))
	b += 4

	return buf
}
