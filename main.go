package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/tj/mmake/help"
)

var (
	version = "dev"
	commit  = "HEAD"
	date    = "unknown"
)

func main() {
	var (
		makefile     string
		printversion bool
	)

	flag.BoolVar(&printversion, "version", false, "Display version info and exit program")
	flag.StringVar(
		&makefile,
		"makefile",
		"",
		`The path for the make file. If not provided it will try to look for
"Makefile" then "makefile" in the current working directory`,
	)
	flag.Parse()

	if printversion {
		fmt.Printf("Version: %v, Build commit: %v, Released on: %v\n", version, commit, date)
		os.Exit(0)
	}

	r, err := newReader(makefile)

	if err != nil {
		panic(err)
	}

	if err := help.OutputAllLong(r, os.Stdout, []string{}); err != nil {
		panic(err)
	}
}

func newReader(path string) (io.Reader, error) {
	if path == "" {
		return findDefault()
	}

	if f := tryRead(path); f != nil {
		return f, nil
	}

	return nil, errors.New("failed to find makefile")
}

func findDefault() (io.Reader, error) {
	cwd, err := os.Getwd()
	if err != nil {
		panic("cannot read current working directory")
	}

	p := path.Join(cwd, "Makefile")
	if f := tryRead(p); f != nil {
		return f, nil
	}

	p = path.Join(cwd, "makefile")
	if f := tryRead(p); f != nil {
		return f, nil
	}

	return nil, errors.New("failed to find default makefile and non was given")
}

func tryRead(path string) io.Reader {
	f, err := os.Open(path)

	if err == nil {
		return bufio.NewReader(f)
	}

	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	panic(err)
}
