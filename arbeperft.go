package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dylhunn/dragontoothmg"
)

// !!! struct field names must have Capital first letter

type CFG struct {
	Depth   int     `yaml:"depth"`
	TFcap   bool    `yaml:"captures"`
	TFep    bool    `yaml:"eps"`
	TFcst   bool    `yaml:"castlings"`
	TFprm   bool    `yaml:"promotions"`
}

type fenCFG struct {
	FenDesc string `yaml:"desc"` // FEN DESCription
	Fen     string `yaml:"fen"`
}

type PerftStats struct {
	Nodes      [10]int64
	Captures   [10]int64
	EPs        [10]int64
	Castlings  [10]int64
	Promotions [10]int64
}

const Version      = "v1.0"
const FENstart     = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
const FENstartDesc = "Standard Start position @@@ 1:20 @@@ 2:400 @@@ 3:8902 @@@ 4:197281"
const FileFenCfg   = "fen.yml" // name of FEN ConFiGuration file
const FileCfg      = "cfg.yml" // name of general ConFiGuration file
const DepthMax     = 9 // Perft max depth
const PieceLetter  = " pnbrqk" // for diagram print

// --------------------------
// @@@ dragontoothmg code @@@
// --------------------------
// type Piece uint8
//
//const (
//0  Nothing = iota
//1  Pawn    = iota
//2  Knight  = iota // list before bishop for promotion loops
//3  Bishop  = iota
//4  Rook    = iota
//5  Queen   = iota
//6  King    = iota
//)
// --------------------------


var args = os.Args  // command line arguments (parameters) 
var timerDone chan bool


func main() {

	var fencfg fenCFG
	var cfg CFG

	// detect OPTions argument IndeX
	var ixopt int // options argument

	var rt string // ReTurn string
	var TFperft = false // do just normal Perft (option 'z')
	var TFdesc  = true  // show FEN DESCription
	var TFtimer = false // show timer

	fmt.Printf("\nArbeperft %s - give argument 'h' for Help\n", Version)

	// -----------------------------------------
	// LOAD CONFIGS
	// arguments overwrite our default values
	// -----------------------------------------

	// load default config values from (2) existing .yml files
	fnInit( &cfg, &fencfg )

	if len(args) > 1 { // OK : minimal one argument

		// this must be the options part
		if rtTFallNonDigits(args[1]) {
			ixopt = 1
			rt = rtSetCFG( &cfg, args[ixopt] )
			if rt == "h" { // Help text was printed
				return // quit program
			}
			TFperft = ( rt == "z" )

			// FEN parts
			if len(args) > 2 {
				fencfg.Fen = strings.TrimSpace( strings.Join(args[2:], " ") )
				if fencfg.Fen == "" {
					fencfg.Fen = FENstart
				}
				TFdesc = false
			}

		// this must be the depth or FEN parts
		} else {

			// depth argument
			if rtTFallDigits(args[1]) {
				cfg.Depth, _ = strconv.Atoi(args[1])
				cfg.Depth = tN( (cfg.Depth > DepthMax), DepthMax, cfg.Depth )

				if len(args) > 2 {

					// this must be the first FEN part
					if rtTFhas(args[2],"/") {
						fencfg.Fen = strings.TrimSpace( strings.Join(args[2:], " ") )

					// this must be the options part
					} else if rtTFallNonDigits(args[2]) {
						ixopt = 2
						rt = rtSetCFG( &cfg, args[ixopt] )
						if rt == "h" { // Help text was printed
							return // quit program
						}
						TFperft = ( rt == "z" )

						// FEN parts
						if len(args) > 3 {
							fencfg.Fen = strings.TrimSpace( strings.Join(args[3:], " ") )
							if fencfg.Fen == "" {
								fencfg.Fen = FENstart
							}
							TFdesc = false
						}
					}
				}

			// this must be the first FEN part
			} else if rtTFhas(args[1],"/") {
				fencfg.Fen = strings.TrimSpace( strings.Join(args[1:], " ") )
				TFdesc = false
			}
		}

	} else { // use default settings
		rt = rtSetCFG( &cfg, "" )
	}

	// -----------------------------------------
	// SCRIPT
	// -----------------------------------------

	fmt.Printf("\nfen: %s\n", fencfg.Fen)
	if TFdesc {
		fmt.Printf("description: %s\n", fencfg.FenDesc)
	}

	board := dragontoothmg.ParseFen(fencfg.Fen)

	if ixopt != 0 {
		if rtTFhas( args[ixopt], "d" ) {
			fmt.Println( rtDiagram(&board) )
		} else {
			fmt.Println()
		}
		TFtimer = rtTFhas( args[ixopt], "t" )
	}

	timerStart := time.Now()

	if TFtimer {
		timerDone = make(chan bool)
		go func() { // timer goroutine
			timerTicker := time.NewTicker(time.Second)
			defer timerTicker.Stop()
			for {
				select {
					case <-timerDone:
						return
					case <-timerTicker.C:
						elapsed := time.Since(timerStart)
						fmt.Printf("\r")
						fmt.Printf("Counting... (%v)", elapsed.Truncate(time.Second)) // Truncate to whole seconds
				}
			}
		}()
	}

	if TFperft {
		zPerft := dragontoothmg.Perft(&board, cfg.Depth)
		fmt.Printf("\r") // clear timer line
		fmt.Printf("Perft(%d) : %d\n", cfg.Depth, zPerft )
	} else {

		pStats := &PerftStats{}
		rtPerft(&board, cfg.Depth, cfg, 0, pStats)

		fmt.Printf("\r") // clear timer line
		fmt.Println("┌─────────┬───────────────┬───────────────┬─────────────┬─────────────┬─────────────┐")
		fmt.Println("│ Perft # │         Nodes │      Captures │         EPs │   Castlings │  Promotions │")
		fmt.Println("└─────────┴───────────────┴───────────────┴─────────────┴─────────────┴─────────────┘")

		for d := 1; d <= cfg.Depth; d++ {
			fmt.Printf(
				"%9d %15d %s %s %s %s\n",
				d,
				pStats.Nodes[d],
				tS( cfg.TFcap,             fmt.Sprintf("%15d", pStats.Captures[d]),   fmt.Sprintf("%15s", " ") ),
				tS( cfg.TFep && cfg.TFcap, fmt.Sprintf("%13d", pStats.EPs[d]),        fmt.Sprintf("%13s", " ") ),
				tS( cfg.TFcst,             fmt.Sprintf("%13d", pStats.Castlings[d]),  fmt.Sprintf("%13s", " ") ),
				tS( cfg.TFprm,             fmt.Sprintf("%13d", pStats.Promotions[d]), fmt.Sprintf("%13s", " ") ),
			)
		}
	}

	if TFtimer {
		timerDone <- true // signal timer to stop
	}

	// 1e6 = 1,000,000 (nanoseconds in a millisecond)
	timerElapsedMs := float64(time.Since(timerStart).Nanoseconds()) / 1e6
	fmt.Printf("\nElapsed time: %v (%.0f ms)\n\n", time.Since(timerStart).Round(time.Second), timerElapsedMs)

}

