package main

import (
	"github.com/lecuong04/compressembed/lib"
	"github.com/lecuong04/compressembed/lib/flag"
)

var cfg = lib.Config{
	Pkg:    "main",
	Func:   lib.StrGen(6),
	Input:  "",
	Key:    lib.KeyGen(),
	Output: "resource.dat",
	Var:    "",
	Src:    "compressed.go",
}

func main() {
	flag.StringVar(&cfg.Input, "in", cfg.Input, "Input file (Require)")
	flag.StringVar(&cfg.Output, "out", cfg.Output, "Compressed output file")
	flag.StringVar(&cfg.Src, "src", cfg.Src, "Source file name to create")
	flag.StringVar(&cfg.Pkg, "pkg", cfg.Pkg, "Name of package for source file to output")
	flag.StringVar(&cfg.Var, "var", cfg.Var, "Variable name for decompressed resource (Require)")
	flag.Parse()
	lib.Run(cfg)
}
