package constants

import (
	"github.com/gdamore/tcell"
)

// Constants
const (
	CharULCorner                   = '\u250C'
	CharURCorner                   = '\u2510'
	CharLLCorner                   = '\u2514'
	CharLRCorner                   = '\u2518'
	CharHLine                      = '\u2500'
	CharVLine                      = '\u2502'
	CharFaceWhite                  = '\u263A'
	CharFaceBlack                  = '\u263B'
	CharHeart                      = '\u2665'
	CharClub                       = '\u2663'
	CharDiamond                    = '\u2666'
	CharSpades                     = '\u2660'
	CharDot                        = '\u2022'
	CharArrowUp                    = '\u2191'
	CharArrowDown                  = '\u2193'
	CharArrowLeft                  = '\u2190'
	CharBlockSolid                 = '\u2588'
	CharBlockDense                 = '\u2593'
	CharBlockMedium                = '\u2592'
	CharBlockSparce                = '\u2591'
	CharSingleLineHorizontal       = CharHLine
	CharDoubleLineHorizontal       = '\u2550'
	CharSingleLineVertical         = '\u2502'
	CharDoubleLineVertical         = '\u2551'
	CharSingleLineUpLeftCorner     = CharULCorner
	CharDoubleLineUpLeftCorner     = '\u2554'
	CharSingleLineUpRightCorner    = CharURCorner
	CharDoubleLineUpRightCorner    = '\u2557'
	CharSingleLineLowerLeftCorner  = CharLLCorner
	CharDoubleLineLowerLeftCorner  = '\u255A'
	CharSingleLineLowerRightCorner = CharLRCorner
	CharDoubleLineLowerRightCorner = '\u255D'
	CharSingleLineCross            = '\u253C'
	CharDoubleLineCross            = '\u256C'
	CharSingleLineTUp              = '\u2534'
	CharSingleLineTDown            = '\u252C'
	CharSingleLineTLeft            = '\u2524'
	CharSingleLineTRight           = '\u251C'
	CharSingleLineDoubleUp         = '\u256B'
	CharSingleLineDoubleDown       = '\u2565'
	CharSingleLineDoubleLeft       = '\u2561'
	CharSingleLineDoubleRight      = '\u255E'
	CharDoubleLineTUp              = '\u2569'
	CharDoubleLineTDown            = '\u2566'
	CharDoubleLineTLeft            = '\u2563'
	CharDoubleLineTRight           = '\u2560'
	CharDoubleLineTSingleUp        = '\u2567'
	CharDoubleLineTSingleDown      = '\u2564'
	CharDoubleLineTSingleLeft      = '\u2562'
	CharDoubleLineTSingleRight     = '\u255F'
	CharBlockLowerHalf 			   = '\u2584'
	CharBlockUpperHalf  		   = '\u2580'
)

// Black
const (
	ColorBlack = iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
	ColorBrightBlack
	ColorBrightRed
	ColorBrightGreen
	ColorBrightYellow
	ColorBrightBlue
	ColorBrightMagenta
	ColorBrightCyan
	ColorBrightWhite
)

var AnsiColorByIndex = map[int]int32{
	ColorBlack:         int32(tcell.GetColor("black")),
	ColorRed:           int32(tcell.GetColor("maroon")),
	ColorGreen:         int32(tcell.GetColor("green")),
	ColorYellow:        int32(tcell.GetColor("olive")),
	ColorBlue:          int32(tcell.GetColor("navy")),
	ColorMagenta:       int32(tcell.GetColor("purple")),
	ColorCyan:          int32(tcell.GetColor("teal")),
	ColorWhite:         int32(tcell.GetColor("silver")),
	ColorBrightBlack:   int32(tcell.GetColor("gray")),
	ColorBrightRed:     int32(tcell.GetColor("red")),
	ColorBrightGreen:   int32(tcell.GetColor("lime")),
	ColorBrightYellow:  int32(tcell.GetColor("yellow")),
	ColorBrightBlue:    int32(tcell.GetColor("blue")),
	ColorBrightMagenta: int32(tcell.GetColor("fuchsia")),
	ColorBrightCyan:    int32(tcell.GetColor("aqua")),
	ColorBrightWhite:   int32(tcell.GetColor("white")),
}
const (
	AnsiEsc = '\u001b'
)
const NullRune = '\x00'
const NullColor = -1
const NullDataType = -1
const NullTransformValue = -1
const NullSelectionIndex = -1
const NullCellId = -1
const NullCellType = -1
const TransformContrast = 0
const TransformTransparency = 0

const LeftAligned = 0
const RightAligned = 1
const CenterAligned = 2
const FrameStyleNormal = 0
const FrameStyleRaised = 1
const FrameStyleSunken = 2
const CellTypeButton = 1
const CellTypeTextInput = 2

const VirtualFileSystemZip = 1
const VirtualFileSystemRar = 2