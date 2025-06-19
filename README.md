
# Arbeperft v1.0

Jump to:<br>
- [Example output](#example-output)<br>
- [How to compile](#how-to-compile)<br>
- [Configuration files](#configuration-files)<br>
- [Info on Perft](#info-on-perft)<br>
- [My OS specifications](#my-os-specifications)

<hr>

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

## Example output

Depth 5, using the 'z' option : this is fastest, it doesn't count any (ep) captures, castlings and promotions :

```
$ ./arbeperft 5 zt

Arbeperft v1.0 - give argument 'h' for Help

fen: r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1
description: Kiwipete @@@ 1:48 @@@ 2:2039 @@@ 3:97862 @@@ 4:4085603 @@@ 5:193690690

Perft(5) : 193690690

Elapsed time: 3s (2706 ms)
```

Depth 5 and count all (ep) captures, castlings and promotions (option 'd' shows a position diagram) :

```
$ ./arbeperft 5 xcepdt

Arbeperft v1.0 - give argument 'h' for Help

fen: r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1
description: Kiwipete @@@ 1:48 @@@ 2:2039 @@@ 3:97862 @@@ 4:4085603 @@@ 5:193690690

     A B C D E F G H
   ┌─────────────────┐
 8 │ r · · · k · · r │ 8
 7 │ p · p p q p b · │ 7
 6 │ b n · · p n p · │ 6
 5 │ · · · P N · · · │ 5
 4 │ · p · · P · · · │ 4
 3 │ · · N · · Q · p │ 3
 2 │ P P P B B P P P │ 2
 1 │ R · · · K · · R │ 1
   └─────────────────┘
     A B C D E F G H

┌─────────┬───────────────┬───────────────┬─────────────┬─────────────┬─────────────┐
│ Perft # │         Nodes │      Captures │         EPs │   Castlings │  Promotions │
└─────────┴───────────────┴───────────────┴─────────────┴─────────────┴─────────────┘
        1              48               8             0             2             0
        2            2039             351             1            91             0
        3           97862           17102            45          3162             0
        4         4085603          757163          1929        128013         15172
        5       193690690        35043416         73365       4993637          8392

Elapsed time: 14m55s (895222 ms)
```

It seems especially counting all ep captures takes a long time.
Btw. this is output of my notebook, running Xubuntu 22.04 - specifications see page bottom.

## How to compile

This program is coded in Go.<br>
I used an adapted version of 'dragontoothmg', a fast Go chess library, see <a href="https://github.com/dylhunn/dragontoothmg">https://github.com/dylhunn/dragontoothmg</a>

The minimum required Go version is 1.22<br>
NOTE: higher versions are generally backward compatible, Go maintains strong backward compatibility.

Download and install Go 1.22 from the official website:<br>
- Go to <a href="https://golang.org/dl">https://golang.org/dl</a><br>
- Download the installer for your OS (Linux, Windows, macOS).<br>
- Follow the installation instructions.<br>

Set up your environment.<br>
After installation, make sure Go is in your PATH.<br>
You can verify the installed version with this terminal command:

```bash
$ go version
```

It should output something like:

```
go version go1.22.x linux/amd64
```

Then `cd` into the Arbeperft project root and compile as follows:

```bash
$ cd Arbeperft
$ go build
```

Optionally you can use `GVM`, an interface to Go Version Management, see <a href="https://github.com/moovweb/gvm">https://github.com/moovweb/gvm</a> . Installing `gvm` needs these prerequisites:<br>
- `Git`<br>
- `Bash` or `Zsh`<br>
- `Curl`

When using `Bash` do:

```bash
$ bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

Or, for `Zsh`:

```bash
$ zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

Then restart your terminal, or `source` your shell config file e.g.<br>
`$ source ~/.bashrc`

Installing and using a certain Go version is now easy:

```bash
$ gvm install go1.22
$ gvm use go1.22
```

## Configuration files

Arbeperft uses 2 config files : 'cfg.yml' for general settings and 'fen.yml' to set a FEN position, both should be in the same folder as the executable. Their format is YAML, a human-readable data serialization language. It is commonly used for configuration files and in applications where data is being stored or transmitted, see https://en.wikipedia.org/wiki/YAML<br>
Another commonly used format for config files is JSON, but i use YAML because those files can contain comment lines (starting with a '#' character), so you can easily disable alternative settings.

Here are these 2 .yml files with my default content.<br>
These settings can be overruled by the options argument.

**cfg.yml**

```yml
depth: 3
captures: false
eps: false
castlings: false
promotions: false
```

When the file 'cfg.yml' does not exist, the shown values are used.

**fen.yml**

```yml
# fen: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
# desc: Standard Start position @@@ 1:20 @@@ 2:400 @@@ 3:8902 @@@ 4:197281

fen: r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1
desc: Kiwipete @@@ 1:48 @@@ 2:2039 @@@ 3:97862 @@@ 4:4085603 @@@ 5:193690690
```

When the file 'fen.yml' does not exist, the standard starting position is loaded.

## Info on Perft

Perft Results : <a href="https://www.chessprogramming.org/Perft_Results">https://www.chessprogramming.org/Perft_Results</a>

## My OS specifications

```
$ neofetch
           `-/osyhddddhyso/-`              roelof@roelof-HP-Elite-x2-1012-G2 
        .+yddddddddddddddddddy+.           --------------------------------- 
      :yddddddddddddddddddddddddy:         OS: Xubuntu 22.04.5 LTS x86_64 
    -yddddddddddddddddddddhdddddddy-       Host: HP Elite x2 1012 G2 
   odddddddddddyshdddddddh`dddd+ydddo      Kernel: 5.15.0-141-generic 
 `yddddddhshdd-   ydddddd+`ddh.:dddddy`    Uptime: 4 days, 5 hours, 17 mins 
 sddddddy   /d.   :dddddd-:dy`-ddddddds    Packages: 4474 (dpkg), 10 (flatpak), 23 (snap) 
:ddddddds    /+   .dddddd`yy`:ddddddddd:   Shell: bash 5.1.16 
sdddddddd`    .    .-:/+ssdyodddddddddds   Resolution: 1920x1080 
ddddddddy                  `:ohddddddddd   DE: Xfce 4.16 
dddddddd.                      +dddddddd   WM: Xfwm4 
sddddddy                        ydddddds   WM Theme: Default 
:dddddd+                      .oddddddd:   Theme: Adwaita-dark [GTK2/3] 
 sdddddo                   ./ydddddddds    Icons: elementary-xfce-darker [GTK2/3] 
 `yddddd.              `:ohddddddddddy`    Terminal: konsole 
   oddddh/`      `.:+shdddddddddddddo      CPU: Intel i5-7200U (4) @ 3.100GHz 
    -ydddddhyssyhdddddddddddddddddy-       GPU: Intel HD Graphics 620 
      :yddddddddddddddddddddddddy:         Memory: 6079MiB / 7816MiB 
        .+yddddddddddddddddddy+.
           `-/osyhddddhyso/-`                                      
```

```
$ cat /proc/cpuinfo
processor       : 0
vendor_id       : GenuineIntel
cpu family      : 6
model           : 142
model name      : Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
stepping        : 9
microcode       : 0xf6
cpu MHz         : 911.391
cache size      : 3072 KB
physical id     : 0
siblings        : 4
core id         : 0
cpu cores       : 2
apicid          : 0
initial apicid  : 0
fpu             : yes
fpu_exception   : yes
cpuid level     : 22
wp              : yes
flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc art arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf pni pclmulqdq dtes64 monitor ds_cpl est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm 3dnowprefetch cpuid_fault epb invpcid_single pti ssbd ibrs ibpb stibp fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid mpx rdseed adx smap clflushopt intel_pt xsaveopt xsavec xgetbv1 xsaves dtherm ida arat pln pts hwp hwp_notify hwp_act_window hwp_epp md_clear flush_l1d arch_capabilities
bugs            : cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit srbds mmio_stale_data retbleed gds
bogomips        : 5399.81
clflush size    : 64
cache_alignment : 64
address sizes   : 39 bits physical, 48 bits virtual

( ... )
```

@

Roelof Berkepeis<br>
Holland<br>
<br>
