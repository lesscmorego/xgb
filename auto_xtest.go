package xgb

/*
	This file was generated by xtest.xml on May 6 2012 5:48:48pm EDT.
	This file is automatically generated. Edit at your peril!
*/

// Imports are not necessary for XGB because everything is 
// in one package. They are still listed here for reference.
// import "xproto"

// XtestInit must be called before using the XTEST extension.
func (c *Conn) XtestInit() error {
	reply, err := c.QueryExtension(5, "XTEST").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return newError("No extension named XTEST could be found on on the server.")
	}

	c.extLock.Lock()
	c.extensions["XTEST"] = reply.MajorOpcode
	for evNum, fun := range newExtEventFuncs["XTEST"] {
		newEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	c.extLock.Unlock()

	return nil
}

func init() {
	newExtEventFuncs["XTEST"] = make(map[int]newEventFun)
}

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Id'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

const (
	XtestCursorNone    = 0
	XtestCursorCurrent = 1
)

// Request XtestGetVersion
// size: 8
type XtestGetVersionCookie struct {
	*cookie
}

func (c *Conn) XtestGetVersion(MajorVersion byte, MinorVersion uint16) XtestGetVersionCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xtestGetVersionRequest(MajorVersion, MinorVersion), cookie)
	return XtestGetVersionCookie{cookie}
}

func (c *Conn) XtestGetVersionUnchecked(MajorVersion byte, MinorVersion uint16) XtestGetVersionCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xtestGetVersionRequest(MajorVersion, MinorVersion), cookie)
	return XtestGetVersionCookie{cookie}
}

// Request reply for XtestGetVersion
// size: 10
type XtestGetVersionReply struct {
	Sequence     uint16
	Length       uint32
	MajorVersion byte
	MinorVersion uint16
}

// Waits and reads reply data from request XtestGetVersion
func (cook XtestGetVersionCookie) Reply() (*XtestGetVersionReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xtestGetVersionReply(buf), nil
}

// Read reply into structure from buffer for XtestGetVersion
func xtestGetVersionReply(buf []byte) *XtestGetVersionReply {
	v := new(XtestGetVersionReply)
	b := 1 // skip reply determinant

	v.MajorVersion = buf[b]
	b += 1

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.MinorVersion = Get16(buf[b:])
	b += 2

	return v
}

func (cook XtestGetVersionCookie) Check() error {
	return cook.check()
}

// Write request to wire for XtestGetVersion
func (c *Conn) xtestGetVersionRequest(MajorVersion byte, MinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XTEST"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = MajorVersion
	b += 1

	b += 1 // padding

	Put16(buf[b:], MinorVersion)
	b += 2

	return buf
}

// Request XtestCompareCursor
// size: 12
type XtestCompareCursorCookie struct {
	*cookie
}

func (c *Conn) XtestCompareCursor(Window Id, Cursor Id) XtestCompareCursorCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xtestCompareCursorRequest(Window, Cursor), cookie)
	return XtestCompareCursorCookie{cookie}
}

func (c *Conn) XtestCompareCursorUnchecked(Window Id, Cursor Id) XtestCompareCursorCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xtestCompareCursorRequest(Window, Cursor), cookie)
	return XtestCompareCursorCookie{cookie}
}

// Request reply for XtestCompareCursor
// size: 8
type XtestCompareCursorReply struct {
	Sequence uint16
	Length   uint32
	Same     bool
}

// Waits and reads reply data from request XtestCompareCursor
func (cook XtestCompareCursorCookie) Reply() (*XtestCompareCursorReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xtestCompareCursorReply(buf), nil
}

// Read reply into structure from buffer for XtestCompareCursor
func xtestCompareCursorReply(buf []byte) *XtestCompareCursorReply {
	v := new(XtestCompareCursorReply)
	b := 1 // skip reply determinant

	if buf[b] == 1 {
		v.Same = true
	} else {
		v.Same = false
	}
	b += 1

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	return v
}

func (cook XtestCompareCursorCookie) Check() error {
	return cook.check()
}

// Write request to wire for XtestCompareCursor
func (c *Conn) xtestCompareCursorRequest(Window Id, Cursor Id) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XTEST"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Window))
	b += 4

	Put32(buf[b:], uint32(Cursor))
	b += 4

	return buf
}

// Request XtestFakeInput
// size: 36
type XtestFakeInputCookie struct {
	*cookie
}

// Write request to wire for XtestFakeInput
func (c *Conn) XtestFakeInput(Type byte, Detail byte, Time uint32, Root Id, RootX int16, RootY int16, Deviceid byte) XtestFakeInputCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.xtestFakeInputRequest(Type, Detail, Time, Root, RootX, RootY, Deviceid), cookie)
	return XtestFakeInputCookie{cookie}
}

func (c *Conn) XtestFakeInputChecked(Type byte, Detail byte, Time uint32, Root Id, RootX int16, RootY int16, Deviceid byte) XtestFakeInputCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.xtestFakeInputRequest(Type, Detail, Time, Root, RootX, RootY, Deviceid), cookie)
	return XtestFakeInputCookie{cookie}
}

func (cook XtestFakeInputCookie) Check() error {
	return cook.check()
}

// Write request to wire for XtestFakeInput
func (c *Conn) xtestFakeInputRequest(Type byte, Detail byte, Time uint32, Root Id, RootX int16, RootY int16, Deviceid byte) []byte {
	size := 36
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XTEST"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Type
	b += 1

	buf[b] = Detail
	b += 1

	b += 2 // padding

	Put32(buf[b:], Time)
	b += 4

	Put32(buf[b:], uint32(Root))
	b += 4

	b += 8 // padding

	Put16(buf[b:], uint16(RootX))
	b += 2

	Put16(buf[b:], uint16(RootY))
	b += 2

	b += 7 // padding

	buf[b] = Deviceid
	b += 1

	return buf
}

// Request XtestGrabControl
// size: 8
type XtestGrabControlCookie struct {
	*cookie
}

// Write request to wire for XtestGrabControl
func (c *Conn) XtestGrabControl(Impervious bool) XtestGrabControlCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.xtestGrabControlRequest(Impervious), cookie)
	return XtestGrabControlCookie{cookie}
}

func (c *Conn) XtestGrabControlChecked(Impervious bool) XtestGrabControlCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.xtestGrabControlRequest(Impervious), cookie)
	return XtestGrabControlCookie{cookie}
}

func (cook XtestGrabControlCookie) Check() error {
	return cook.check()
}

// Write request to wire for XtestGrabControl
func (c *Conn) xtestGrabControlRequest(Impervious bool) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XTEST"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	if Impervious {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}
