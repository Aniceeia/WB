package config

import (
	"flag"
)

type Config struct {
	InputFile    string
	KeyField     int
	Numeric      bool
	Reverse      bool
	Unique       bool
	Month        bool
	IgnoreBlanks bool
	Check        bool
	Human        bool
	Delimiter    rune
}

func ParseFlags() Config {
	var cfg Config

	flag.StringVar(&cfg.InputFile, "input", "", "input file")
	flag.IntVar(&cfg.KeyField, "k", 0, "sort by field")
	flag.BoolVar(&cfg.Numeric, "n", false, "numeric sort")
	flag.BoolVar(&cfg.Reverse, "r", false, "reverse sort")
	flag.BoolVar(&cfg.Unique, "u", false, "unique lines")
	flag.BoolVar(&cfg.Month, "M", false, "month sort")
	flag.BoolVar(&cfg.IgnoreBlanks, "b", false, "ignore blanks")
	flag.BoolVar(&cfg.Check, "c", false, "check sorted")
	flag.BoolVar(&cfg.Human, "h", false, "human numeric")

	flag.Parse()

	if cfg.InputFile == "" && flag.NArg() > 0 {
		cfg.InputFile = flag.Arg(0)
	}

	cfg.Delimiter = '\t'
	return cfg
}
