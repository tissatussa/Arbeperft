package main

import (
	"fmt"
	// "math"
	"strings"
	"unicode"

	"github.com/tissatussa/dragontoothmg"
)

// Ternary : return integer
func tN( TF bool, int1 int, int2 int ) int {
	if TF {
		return int1
	}
	return int2
}

// Ternary : return string
func tS( TF bool, s1 string, s2 string ) string {
	if TF {
		return s1
	}
	return s2
}

// Ternary : return boolean
func tB( TF bool, b1 bool, b2 bool ) bool {
	if TF {
		return b1
	}
	return b2
}

// sum all array elements
func rtSum(arr []int64) int64 {
	var total int64
	for _, v := range arr {
		total += v
	}
	return total
}

func rtTFallDigits( s string ) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func rtTFallNonDigits( s string ) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func rtTFhas( haystack string, needle string ) bool {
	return strings.Contains(haystack, needle)
}

func rtSetCFG( cfg *CFG, arg string ) string {

	if rtTFhas(arg,"h") {
		fnHelp()
		return "h"
	} else if rtTFhas(arg,"z") {
		return "z"
	} else {
		if rtTFhas(arg,"x") {
			cfg.TFcap = true
		}
		if rtTFhas(arg,"e") {
			cfg.TFep = true
		}
		if rtTFhas(arg,"c") {
			cfg.TFcst = true
		}
		if rtTFhas(arg,"p") {
			cfg.TFprm = true
		}
	}
	return ""
}

func rtDiagram( pBoard *dragontoothmg.Board ) string {
	var c, row, rt string
	var n = 0

	for sq := 0; sq <= 63; sq++ {
		if sq % 8 == 0 {
			if n != 0 {
				rt = fmt.Sprintf(" %d │ %s │ %d\n%s", n, row, n, rt)
				row = ""
			}
			n++
		}

		pc, TFwhite := dragontoothmg.GetPieceType( uint8(sq), pBoard )
		c = string(PieceLetter[pc])
		// ?! @@@ []rune(PieceLetter)
		c = tS( TFwhite, strings.ToUpper(c), c )

		// c = tS( (c==" "), ".", c ) // normal dot
		c = tS( (c==" "), "·", c ) // special dot

		row += tS( (row == ""), "", " " )
		row += c
	}
	rt = fmt.Sprintf(" %d │ %s │ %d\n%s", n, row, n, rt)

	rt = "\n     A B C D E F G H\n   ┌─────────────────┐\n" + rt
	rt += "   └─────────────────┘\n"
	rt += "     A B C D E F G H\n"
	return rt
}

// Each bitboard shall use little-endian rank-file mapping:
// 56  57  58  59  60  61  62  63
// 48  49  50  51  52  53  54  55
// 40  41  42  43  44  45  46  47
// 32  33  34  35  36  37  38  39
// 24  25  26  27  28  29  30  31
// 16  17  18  19  20  21  22  23
// 8   9   10  11  12  13  14  15
// 0   1   2   3   4   5   6   7
// The binary bitboard uint64 thus uses this ordering:
// MSB---------------------------------------------------LSB
// H8 G8 F8 E8 D8 C8 B8 A8 H7 ... A2 H1 G1 F1 E1 D1 C1 B1 A1


//     A B C D E F G H
//   ┌─────────────────┐
// 8 │ r n b q k b n r │ 8
// 7 │ p p p p p p p p │ 7
// 6 │ · · · · · · · · │ 6
// 5 │ · · · · · · · · │ 5
// 4 │ · · · · · · · · │ 4
// 3 │ · · · · · · · · │ 3
// 2 │ P P P P P P P P │ 2
// 1 │ R N B Q K B N R │ 1
//   └─────────────────┘
//     A B C D E F G H


// type Piece uint8

// const (
// 	Nothing = iota
// 	Pawn    = iota
// 	Knight  = iota // list before bishop for promotion loops
// 	Bishop  = iota
// 	Rook    = iota
// 	Queen   = iota
// 	King    = iota
// )

