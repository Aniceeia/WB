package config

import (
	"flag"
	"os"
	"testing"
)

func TestParseFlags(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"sort", "-n", "-r", "-k", "2", "test.txt"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	cfg := ParseFlags()
	if !cfg.Numeric {
		t.Error("should parse -n flag")
	}
	if !cfg.Reverse {
		t.Error("should parse -r flag")
	}
	if cfg.KeyField != 2 {
		t.Error("should parse -k flag")
	}
	if cfg.InputFile != "test.txt" {
		t.Error("should parse input file")
	}
}

func TestParseFlagsAll(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"sort", "-n", "-r", "-u", "-M", "-b", "-c", "-h", "-k", "3"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	cfg := ParseFlags()
	if !cfg.Numeric {
		t.Error("should parse -n")
	}
	if !cfg.Reverse {
		t.Error("should parse -r")
	}
	if !cfg.Unique {
		t.Error("should parse -u")
	}
	if !cfg.Month {
		t.Error("should parse -M")
	}
	if !cfg.IgnoreBlanks {
		t.Error("should parse -b")
	}
	if !cfg.Check {
		t.Error("should parse -c")
	}
	if !cfg.Human {
		t.Error("should parse -h")
	}
	if cfg.KeyField != 3 {
		t.Error("should parse -k 3")
	}
}

func TestParseFlagsDefault(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"sort"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	cfg := ParseFlags()
	if cfg.Numeric {
		t.Error("should not have numeric by default")
	}
	if cfg.Reverse {
		t.Error("should not have reverse by default")
	}
	if cfg.KeyField != 0 {
		t.Error("should have keyField 0 by default")
	}
	if cfg.Delimiter != '\t' {
		t.Error("should have tab delimiter by default")
	}
}
