package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func fnInit(cfg *CFG, fencfg *fenCFG ) {

	// -----------------------------------------------------------
	// read FEN configuration file
	// Find the active (your uncommented) FEN + description
	// -----------------------------------------------------------

	// defaults
	fencfg.Fen     = FENstart
	fencfg.FenDesc = FENstartDesc

	fenData, err := os.ReadFile( FileFenCfg )
	if err != nil {
		fmt.Printf( "@@@ No FEN config file '%s' found, loading standard starting position.\n", FileFenCfg )
	} else {
		err = yaml.Unmarshal(fenData, &fencfg)
		if err != nil {
			fmt.Printf( "@@@ FEN config file '%s' does not contain a valid 'fen' & 'desc' value pair, loading standard starting position.\n", FileFenCfg )
		}
	}

	// -----------------------------------------------------------
	// read general configuration file
	// -----------------------------------------------------------

	// defaults
	cfg.Depth = 3
	cfg.TFcap = false
	cfg.TFep  = false
	cfg.TFcst = false
	cfg.TFprm = false

	cfgData, err := os.ReadFile( FileCfg )
	if err != nil {
		fmt.Printf( "@@@ No config file '%s' found, loading default values:\n", FileCfg )
	} else {
		err = yaml.Unmarshal(cfgData, &cfg)
		if err != nil {
			fmt.Printf( "@@@ config file '%s' found but invalid, loading default values:\n", FileCfg )
		}
	}

}

