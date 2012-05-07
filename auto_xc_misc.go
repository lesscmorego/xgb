package xgb

/*
	This file was generated by xc_misc.xml on May 6 2012 5:48:47pm EDT.
	This file is automatically generated. Edit at your peril!
*/

// Xc_miscInit must be called before using the XC-MISC extension.
func (c *Conn) Xc_miscInit() error {
	reply, err := c.QueryExtension(7, "XC-MISC").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return newError("No extension named XC-MISC could be found on on the server.")
	}

	c.extLock.Lock()
	c.extensions["XC-MISC"] = reply.MajorOpcode
	for evNum, fun := range newExtEventFuncs["XC-MISC"] {
		newEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	c.extLock.Unlock()

	return nil
}

func init() {
	newExtEventFuncs["XC-MISC"] = make(map[int]newEventFun)
}

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

// Skipping definition for base type 'Void'

// Request Xc_miscGetVersion
// size: 8
type Xc_miscGetVersionCookie struct {
	*cookie
}

func (c *Conn) Xc_miscGetVersion(ClientMajorVersion uint16, ClientMinorVersion uint16) Xc_miscGetVersionCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xc_miscGetVersionRequest(ClientMajorVersion, ClientMinorVersion), cookie)
	return Xc_miscGetVersionCookie{cookie}
}

func (c *Conn) Xc_miscGetVersionUnchecked(ClientMajorVersion uint16, ClientMinorVersion uint16) Xc_miscGetVersionCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xc_miscGetVersionRequest(ClientMajorVersion, ClientMinorVersion), cookie)
	return Xc_miscGetVersionCookie{cookie}
}

// Request reply for Xc_miscGetVersion
// size: 12
type Xc_miscGetVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	ServerMajorVersion uint16
	ServerMinorVersion uint16
}

// Waits and reads reply data from request Xc_miscGetVersion
func (cook Xc_miscGetVersionCookie) Reply() (*Xc_miscGetVersionReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xc_miscGetVersionReply(buf), nil
}

// Read reply into structure from buffer for Xc_miscGetVersion
func xc_miscGetVersionReply(buf []byte) *Xc_miscGetVersionReply {
	v := new(Xc_miscGetVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.ServerMajorVersion = Get16(buf[b:])
	b += 2

	v.ServerMinorVersion = Get16(buf[b:])
	b += 2

	return v
}

func (cook Xc_miscGetVersionCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xc_miscGetVersion
func (c *Conn) xc_miscGetVersionRequest(ClientMajorVersion uint16, ClientMinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XC-MISC"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put16(buf[b:], ClientMajorVersion)
	b += 2

	Put16(buf[b:], ClientMinorVersion)
	b += 2

	return buf
}

// Request Xc_miscGetXIDRange
// size: 4
type Xc_miscGetXIDRangeCookie struct {
	*cookie
}

func (c *Conn) Xc_miscGetXIDRange() Xc_miscGetXIDRangeCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xc_miscGetXIDRangeRequest(), cookie)
	return Xc_miscGetXIDRangeCookie{cookie}
}

func (c *Conn) Xc_miscGetXIDRangeUnchecked() Xc_miscGetXIDRangeCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xc_miscGetXIDRangeRequest(), cookie)
	return Xc_miscGetXIDRangeCookie{cookie}
}

// Request reply for Xc_miscGetXIDRange
// size: 16
type Xc_miscGetXIDRangeReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	StartId uint32
	Count   uint32
}

// Waits and reads reply data from request Xc_miscGetXIDRange
func (cook Xc_miscGetXIDRangeCookie) Reply() (*Xc_miscGetXIDRangeReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xc_miscGetXIDRangeReply(buf), nil
}

// Read reply into structure from buffer for Xc_miscGetXIDRange
func xc_miscGetXIDRangeReply(buf []byte) *Xc_miscGetXIDRangeReply {
	v := new(Xc_miscGetXIDRangeReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.StartId = Get32(buf[b:])
	b += 4

	v.Count = Get32(buf[b:])
	b += 4

	return v
}

func (cook Xc_miscGetXIDRangeCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xc_miscGetXIDRange
func (c *Conn) xc_miscGetXIDRangeRequest() []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XC-MISC"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Request Xc_miscGetXIDList
// size: 8
type Xc_miscGetXIDListCookie struct {
	*cookie
}

func (c *Conn) Xc_miscGetXIDList(Count uint32) Xc_miscGetXIDListCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.xc_miscGetXIDListRequest(Count), cookie)
	return Xc_miscGetXIDListCookie{cookie}
}

func (c *Conn) Xc_miscGetXIDListUnchecked(Count uint32) Xc_miscGetXIDListCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.xc_miscGetXIDListRequest(Count), cookie)
	return Xc_miscGetXIDListCookie{cookie}
}

// Request reply for Xc_miscGetXIDList
// size: (32 + pad((int(IdsLen) * 4)))
type Xc_miscGetXIDListReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	IdsLen uint32
	// padding: 20 bytes
	Ids []uint32 // size: pad((int(IdsLen) * 4))
}

// Waits and reads reply data from request Xc_miscGetXIDList
func (cook Xc_miscGetXIDListCookie) Reply() (*Xc_miscGetXIDListReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return xc_miscGetXIDListReply(buf), nil
}

// Read reply into structure from buffer for Xc_miscGetXIDList
func xc_miscGetXIDListReply(buf []byte) *Xc_miscGetXIDListReply {
	v := new(Xc_miscGetXIDListReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.IdsLen = Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Ids = make([]uint32, v.IdsLen)
	for i := 0; i < int(v.IdsLen); i++ {
		v.Ids[i] = Get32(buf[b:])
		b += 4
	}
	b = pad(b)

	return v
}

func (cook Xc_miscGetXIDListCookie) Check() error {
	return cook.check()
}

// Write request to wire for Xc_miscGetXIDList
func (c *Conn) xc_miscGetXIDListRequest(Count uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["XC-MISC"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], Count)
	b += 4

	return buf
}
