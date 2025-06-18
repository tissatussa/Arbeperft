package main

import (
	"strings"

	"github.com/tissatussa/dragontoothmg"
)

func rtPerft(pBoard *dragontoothmg.Board, dp int, cfg CFG, currentDepth int, pStats *PerftStats) {
	if dp == 0 {
		return
	}

	moves := pBoard.GenerateLegalMoves()
	for _, move := range moves {
		pStats.Nodes[currentDepth+1]++

		fromBitboard := uint64(1) << move.From()

		TFcapture := false
		TFep := false
		if cfg.TFcap {
			TFcapture = dragontoothmg.IsCapture(move,pBoard)
			if cfg.TFep {
				aFEN := strings.Split(pBoard.ToFen(), " ")
				if (aFEN[3] != "-") {
					// @@@ This move is a pawn move
					if (fromBitboard & pBoard.White.Pawns) != 0 || (fromBitboard & pBoard.Black.Pawns) != 0 {
						epIndex, err := dragontoothmg.AlgebraicToIndex(aFEN[3])
						if err == nil && move.To() == epIndex {
							TFep = true
						}
					}
				}
			}
		}

		TFcastle := false
		isWhiteKing := (fromBitboard & pBoard.White.Kings) != 0
		isBlackKing := (fromBitboard & pBoard.Black.Kings) != 0
		if isWhiteKing && move.From() == 4 && (move.To() == 6 || move.To() == 2) {
			TFcastle = true
		}
		if isBlackKing && move.From() == 60 && (move.To() == 62 || move.To() == 58) {
			TFcastle = true
		}

		TFpromo := false
		if move.Promote() != 0 {
			TFpromo = true
		}

		fnUnApply := pBoard.Apply(move)
		rtPerft(pBoard, dp-1, cfg, currentDepth+1, pStats)
		fnUnApply()

		if cfg.TFcap && TFcapture {
			pStats.Captures[currentDepth+1]++
		}
		if cfg.TFep && TFep {
			pStats.EPs[currentDepth+1]++
		}
		if cfg.TFcst && TFcastle {
			pStats.Castlings[currentDepth+1]++
		}
		if cfg.TFprm && TFpromo {
			pStats.Promotions[currentDepth+1]++
		}

	}
}
