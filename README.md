
# Arbeperft v1.0

*Perft : PERFormance Test, chess move path enumeration.*<br>
This program is a debugging function to walk the chess move generation tree of strictly legal moves in any valid FEN position to count all leaf nodes of a certain depth. It can also count (ep) captures, castlings and promotions.

*Syntax, command and arguments*

```
-------------------------------------------------------------
> arbeperft [depth] [options] [FENp FENm FENc FENe FENh FENf]
-------------------------------------------------------------
```

All parameters are optional : when none given, default settings are used.<br>
Settings will be loaded by the configuration files, if they exist and have concerning values.<br>
These Settings can be overruled by command arguments.

Option 'z' is the fastest, it counts all leaf nodes of a certain depth, no other counts are done.<br>
When a FEN is given by its (6) parameters, these should be the last on the command line.<br>
You can use the files 'fen.yml' and 'cfg.yml' (in same folder) to define custom settings - when these file(s) are missing, default settings are used.

### Depth: upto 9

### Options string

May contain these characters, in any order:

```
h: show only this Help message, all other options and parameters will be ignored.
z: show just Perft counts, no other info : options x e c p will be ignored:
x: count captures
e: count En Passant captures (only when option x is also given)
c: count Castlings
p: count Promotions
d: show position Diagram
t: show Timer
```

Undefined option characters are ignored.<br>
Omitting the options string will set default values.

### Define the position by FEN notation

```
FENp: Position, eg. r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R
FENm: side-to-Move, "b" or "w"
FENc: Castling rights, eg. "KQkq" or "-"
FENe: En passant square, eg. "e3" or "-"
FENh: Halfmove clock, eg. 0
FENf: Fullmove number, eg. 1
```

Omitting the FEN parameter(s) will set the FEN data pair (fen and description) in 'fen.yml' - if this config file is not found or it contains no FEN data pair, the starting position is loaded.

When a parameter contains a "/" character, it's considered to be the 'FENp' part and the program expects all other 5 FEN parts to be present. No FEN validation is done, an invalid or incomplete FEN may result in an error or unexpected output.

### Examples:

```
$ ./arbeperft
$ ./arbeperft h
$ ./arbeperft 4
$ ./arbeperft 4 z
$ ./arbeperft 4 pd
$ ./arbeperft 4 xce r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1
$ ./arbeperft r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1
```

<hr>

This program is coded in Go.<br>
I used an adapted version of 'dragontoothmg', a fast Go chess library, see <a href="https://github.com/dylhunn/dragontoothmg">https://github.com/dylhunn/dragontoothmg</a>

@

Roelof Berkepeis<br>
Holland<br>
<br>
