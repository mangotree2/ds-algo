package dartac

/**************************************************************\
*                          Constant                            *
\**************************************************************/
const (
	//BufSize = 64 * 4096

	RootState = 1
	FailState = -1

	//int32
	//MaskBegin  = 0xFF000000
	//MaskEnd    = 0x00FF0000
	//MaskType   = 0x0000C000
	//MaskLength = 0x00003FFF
	//

	BitLenType = 15
	BitLenEnd = 17
	BitLenBegin = 29

	//int64 用41位
	MaskBegin   = int64(4096<<BitLenBegin-1)
	MaskEnd   int64  = int64(4096<<BitLenEnd-1)
	MaskType   int64 = int64(3<<BitLenType-1)
	MaskLength  = int64(0x00007FFF)

	TypeInsult = 1
	TypePolitically = 2
	TypeBoth  = 3

	//DictSize = 106264
	//maxRune  = 260
	//maxByte  = 384
)

/**************************************************************\
*                       Global variable                        *
\**************************************************************/
// Cache
//var (
	//Buf   [maxByte]byte // max 383
	//Cache [maxRune]rune // max 257
	//BSP   = 0           // Buf Stack Pointer
//)

// states
//var (
	//Base  []int
	//Check []int
	//Fail  []int
	//Info  []int
//)

// WriteRune put rune to global buffer (extreme ver without err check)
//func WriteRune(r rune) {
//	switch i := uint32(r); {
//	case i <= 127:
//		Buf[BSP] = byte(r)
//		BSP++
//	case i <= 2047:
//		Buf[BSP] = 0xC0 | byte(r>>6)
//		BSP++
//		Buf[BSP] = 0x80 | byte(r)&0x3F
//		BSP++
//	case i <= 65535:
//		Buf[BSP] = 0xE0 | byte(r>>12)
//		BSP++
//		Buf[BSP] = 0x80 | byte(r>>6)&0x3F
//		BSP++
//		Buf[BSP] = 0x80 | byte(r)&0x3F
//		BSP++
//	default:
//		Buf[BSP] = 0xF0 | byte(r>>18)
//		BSP++
//		Buf[BSP] = 0x80 | byte(r>>12)&0x3F
//		BSP++
//		Buf[BSP] = 0x80 | byte(r>>6)&0x3F
//		BSP++
//		Buf[BSP] = 0x80 | byte(r)&0x3F
//		BSP++
//	}
//}
//
//// WriteByType put write target string to buffer via match type
//func WriteByType(match int) {
//	Buf[BSP] = 45
//	BSP++
//	Buf[BSP] = 42
//	BSP++
//	switch (match & MaskType) >> 14 {
//	case TypeInsult:
//		WriteRune(42)
//		WriteRune(42)
//	case TypePolitically:
//		WriteRune(42)
//		WriteRune(42)
//	case TypeBoth:
//		WriteRune(42)
//		WriteRune(42)
//		WriteRune(42)
//		WriteRune(42)
//		WriteRune(42)
//	}
//	Buf[BSP] = 42
//	BSP++
//	Buf[BSP] = 45
//	BSP++
//}

func WriteRune(r rune,buf []byte) ([]byte,int) {
	switch i := uint32(r); {
	case i <= 127:
		buf = append(buf,byte(r))
		return buf,1
	case i <= 2047:
		buf = append(buf,byte(0xC0 | byte(r>>6)))
		buf = append(buf,byte(0x80 | byte(r)&0x3F))
		return buf,2

	case i <= 65535:
		buf = append(buf,byte(0xE0 | byte(r>>12)))
		buf = append(buf,byte(0x80 | byte(r>>6)&0x3F))
		buf = append(buf,byte(0x80 | byte(r)&0x3F))
		return buf,3

	default:
		buf = append(buf,byte(0xF0 | byte(r>>18)))
		buf = append(buf,byte(0x80 | byte(r>>12)&0x3F))
		buf = append(buf,byte(0x80 | byte(r>>6)&0x3F))
		buf = append(buf,byte(0x80 | byte(r)&0x3F))
		return buf,4

	}
}

func WriteByType(match int64,buf []byte) ([]byte,int) {
	return WriteRune(42,buf)
	//WriteRune(42,buf)
	//WriteRune(42,buf)

	switch (match & MaskType) >> 14 {
	case TypeInsult:
		//WriteRune(42)
		//WriteRune(42)
	case TypePolitically:
		//WriteRune(42)
		//WriteRune(42)
	case TypeBoth:
		//WriteRune(42)
		//WriteRune(42)
		//WriteRune(42)
		//WriteRune(42)
		//WriteRune(42)
	}

	return buf,0
}

/**************************************************************\
*                       Line Processor                         *
\**************************************************************/

// HandleLine take one line , process and write it
func (ac *AC)Handle(input string) (output []byte) {
	var Matches []int64
	var nMatch, match, rCursor, wCursor, info, mLength, mBegin int64
	var overlap bool
	state := RootState

	//Cache := [len([]rune(input))]rune{}
	Cache := make([]rune,len([]rune(input)))
	Buf := make([]byte,0,len([]byte(input)))

	// stage 1 : find all match
	//for _, c := range *(*string)(unsafe.Pointer(&input)) {
	for _, c := range input {

	transfer:
		Cache[rCursor] = c
		//fmt.Println(string(c),int(c))

		// state transfer
		if t := state + int(c) + RootState; t < len(ac.Check) {
			if state == ac.Check[t] { // 53870341
				state = ac.Base[t]
				goto match
			}
			if state == RootState { // 117354923
				state = RootState
				goto match
			}
		}

		if state == RootState {
			goto done
		}

		// reach fail state
		state = ac.Failure[state]
		goto transfer

	match:
		if info = ac.Output[state]; info != 0 {
			// there's a match , check it out
			mLength = info & MaskLength
			match = (rCursor-mLength+1)<<BitLenBegin | (rCursor << BitLenEnd) | info

			// lastMatch.End < match.Begin : add new match
			if nMatch == 0 || ((Matches[nMatch-1]&MaskEnd)>>BitLenEnd ) < (match&MaskBegin)>>BitLenBegin {
				//Matches[nMatch] = match
				Matches = append(Matches,match)
				match = 0
				nMatch ++
				goto done
			}

			// newMatch.Begin <= lastMatch.Begin : abandon old match
			overlap = false
			for ; nMatch > 0 && ((match&MaskBegin)>>BitLenBegin) <= ((Matches[nMatch-1])&MaskBegin)>>BitLenBegin; nMatch-- {
				overlap = true
			}
			if overlap {
				Matches[nMatch] = match
				//Matches = append(Matches,match)
				match = 0
				nMatch++
				goto done
			}
			// miss: omit new match
		}

	done:
		rCursor ++
	}

	// stage 2 : replace all match
	if nMatch == 0 {
		// no match, write line back
		return []byte(input)
	} else {
		// reset buf
		BSP := 0
		add := 0
		matchIndex := 0
		match = Matches[matchIndex]
		mBegin = (match & MaskBegin) >> BitLenBegin
		mLength = match & MaskLength

		for wCursor = 0; wCursor < rCursor; {
			if wCursor < mBegin {
				Buf,add = WriteRune(Cache[wCursor],Buf)
				BSP+= add
				wCursor ++
				continue
			}
			if wCursor == mBegin {
				Buf,add = WriteByType(match,Buf)
				BSP+= add
				wCursor += match & MaskLength
				if int64(matchIndex+1) < nMatch {
					matchIndex++
					match = Matches[matchIndex]
					mBegin = (match & MaskBegin) >> BitLenBegin
					mLength = match & MaskLength
				}
				continue
			} else {
				Buf,add = WriteRune(Cache[wCursor],Buf)
				BSP+= add
				wCursor ++
			}
		}
		return Buf[:BSP]
	}
}


//func (ac *AC)HandleLine1(input string) (output []byte) {
//	var Matches [5]int
//	var nMatch, match, rCursor, wCursor, info, mLength, mBegin int
//	var overlap bool
//	state := RootState
//
//	//Cache := [len([]rune(input))]rune{}
//	Cache := make([]rune,len([]rune(input)))
//	Buf := make([]byte,0,len([]byte(input)))
//
//	// stage 1 : find all match
//	//for _, c := range *(*string)(unsafe.Pointer(&input)) {
//	for _, c := range input {
//
//	transfer:
//		Cache[rCursor] = c
//		//fmt.Println(string(c),int(c))
//
//		// state transfer
//		if t := state + int(c) + RootState; t < len(ac.Check) { // 219169484
//			if state == ac.Check[t] { // 53870341
//				state = ac.Base[t]
//				goto match
//			}
//			if state == RootState { // 117354923
//				state = RootState
//				goto match
//			}
//		}
//		// reach fail state
//		state = ac.Failure[state]
//		goto transfer
//
//	match:
//		if info = ac.Output[state]; info != 0 {
//			// there's a match , check it out
//			mLength = info & MaskLength
//			match = (rCursor-mLength+1)<<24 | (rCursor << 16) | info
//
//			// lastMatch.End < match.Begin : add new match
//			if nMatch == 0 || ((Matches[nMatch-1]&MaskEnd)>>16 ) < (match&MaskBegin)>>24 {
//				Matches[nMatch] = match
//				match = 0
//				nMatch ++
//				goto done
//			}
//
//			// newMatch.Begin <= lastMatch.Begin : abandon old match
//			overlap = false
//			for ; nMatch > 0 && ((match&MaskBegin)>>24) <= ((Matches[nMatch-1])&MaskBegin)>>24; nMatch-- {
//				overlap = true
//			}
//			if overlap {
//				Matches[nMatch] = match
//				match = 0
//				nMatch++
//				goto done
//			}
//			// miss: omit new match
//		}
//
//	done:
//		rCursor ++
//	}
//
//	// stage 2 : replace all match
//	if nMatch == 0 {
//		// no match, write line back
//		return []byte(input)
//	} else {
//		// reset buf
//		BSP := 0
//		add := 0
//		matchIndex := 0
//		match = Matches[matchIndex]
//		mBegin = (match & MaskBegin) >> 24
//		mLength = match & MaskLength
//
//		for wCursor = 0; wCursor < rCursor; {
//			if wCursor < mBegin {
//				Buf,add = WriteRune(Cache[wCursor],Buf)
//				BSP+= add
//				wCursor ++
//				continue
//			}
//			if wCursor == mBegin {
//				Buf,add = WriteByType(match,Buf)
//				BSP+= add
//				wCursor += match & MaskLength
//				if matchIndex+1 < nMatch {
//					matchIndex++
//					match = Matches[matchIndex]
//					mBegin = (match & MaskBegin) >> 24
//					mLength = match & MaskLength
//				}
//				continue
//			} else {
//				Buf,add = WriteRune(Cache[wCursor],Buf)
//				BSP+= add
//				wCursor ++
//			}
//		}
//		return Buf[:BSP]
//	}
//}

func (ac *AC) Filter(input string) string {
	return string(ac.Handle(input))
}



/**************************************************************\
*                          Driver                              *
\**************************************************************/
// Run program with given parameter
//func Run(inputPath, outputPath, dictPath string) {
//	// dict
//	ac := FromFile(dictPath)
//	Base = ac.Base
//	Check = ac.Check
//	Fail = ac.Failure
//	Info = ac.Output
//
//	// input
//	inputFile, err := os.Open(inputPath)
//	if err != nil {
//		panic(err)
//	}
//	defer inputFile.Close()
//	reader := bufio.NewReaderSize(inputFile, BufSize)
//
//	// output
//	outputFile, err := os.Create(outputPath)
//	if err != nil {
//		panic(err)
//	}
//	defer outputFile.Close()
//	writer := bufio.NewWriterSize(outputFile, BufSize)
//
//	// process loop
//	var line []byte
//	for err = nil; err != io.EOF; line, err = reader.ReadSlice('\n') {
//		writer.Write(HandleLine(line))
//	}
//	writer.Flush()
//}


//func HandleLine(input []byte) (output []byte) {
//	var Matches [5]int
//	var nMatch, match, rCursor, wCursor, info, mLength, mBegin int
//	var overlap bool
//	state := RootState
//
//	// stage 1 : find all match
//	for _, c := range *(*string)(unsafe.Pointer(&input)) {
//	transfer:
//		Cache[rCursor] = c
//
//		// state transfer
//		if t := state + int(c) + RootState; t < DictSize { // 219169484
//			if state == Check[t] { // 53870341
//				state = Base[t]
//				goto match
//			}
//			if state == RootState { // 117354923
//				state = RootState
//				goto match
//			}
//		}
//		// reach fail state
//		state = Fail[state]
//		goto transfer
//
//	match:
//		if info = Info[state]; info != 0 {
//			// there's a match , check it out
//			mLength = info & MaskLength
//			match = (rCursor-mLength+1)<<24 | (rCursor << 16) | info
//
//			// lastMatch.End < match.Begin : add new match
//			if nMatch == 0 || ((Matches[nMatch-1]&MaskEnd)>>16 ) < (match&MaskBegin)>>24 {
//				Matches[nMatch] = match
//				match = 0
//				nMatch ++
//				goto done
//			}
//
//			// newMatch.Begin <= lastMatch.Begin : abandon old match
//			overlap = false
//			for ; nMatch > 0 && ((match&MaskBegin)>>24) <= ((Matches[nMatch-1])&MaskBegin)>>24; nMatch-- {
//				overlap = true
//			}
//			if overlap {
//				Matches[nMatch] = match
//				match = 0
//				nMatch++
//				goto done
//			}
//			// miss: omit new match
//		}
//
//	done:
//		rCursor ++
//	}
//
//	// stage 2 : replace all match
//	if nMatch == 0 {
//		// no match, write line back
//		return input
//	} else {
//		// reset buf
//		BSP = 0
//		matchIndex := 0
//		match = Matches[matchIndex]
//		mBegin = (match & MaskBegin) >> 24
//		mLength = match & MaskLength
//
//		for wCursor = 0; wCursor < rCursor; {
//			if wCursor < mBegin {
//				WriteRune(Cache[wCursor])
//				wCursor ++
//				continue
//			}
//			if wCursor == mBegin {
//				WriteByType(match)
//				wCursor += match & MaskLength
//				if matchIndex+1 < nMatch {
//					matchIndex++
//					match = Matches[matchIndex]
//					mBegin = (match & MaskBegin) >> 24
//					mLength = match & MaskLength
//				}
//				continue
//			} else {
//				WriteRune(Cache[wCursor])
//				wCursor ++
//			}
//		}
//		return Buf[:BSP]
//	}
//}
