package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

// isTerminal reports whether f is an interactive terminal rather than a pipe or
// a redirected file, so `remoji` alone can warn instead of hanging on stdin.
func isTerminal(f *os.File) bool {
	info, err := f.Stat()
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeCharDevice != 0
}

const usage = `remoji - remove emoji from text files

usage:
  remoji input.txt            write the cleaned text to stdout
  remoji -i input.txt ...     rewrite the files in place
  remoji < input.txt          read stdin
  remoji --help               show this help

tips:
  remoji notes.md > clean.md         save the result to a new file
  remoji -i *.md                     clean every markdown file in place
  pbpaste | remoji | pbcopy          clean the clipboard (macOS)
`

// errNoFile is a usage problem, not a runtime failure: it prints the tips.
var errNoFile = errors.New("no file selected")

func main() {
	inPlace := flag.Bool("i", false, "edit files in place")
	help := flag.Bool("help", false, "show usage")
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	flag.Parse()

	if *help {
		fmt.Print(usage)
		return
	}

	err := run(flag.Args(), *inPlace)
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "remoji: %v\n", err)
	if errors.Is(err, errNoFile) {
		fmt.Fprint(os.Stderr, "\n"+usage)
		os.Exit(2)
	}
	os.Exit(1)
}

func run(paths []string, inPlace bool) error {
	if len(paths) == 0 {
		if inPlace {
			return fmt.Errorf("-i needs at least one file")
		}
		if isTerminal(os.Stdin) {
			// Nothing piped in and no file named: reading stdin would just hang.
			return errNoFile
		}
		src, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		return writeStdout("stdin", src)
	}

	for _, path := range paths {
		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := checkText(path, src); err != nil {
			return err
		}
		if !inPlace {
			if err := writeStdout(path, src); err != nil {
				return err
			}
			continue
		}
		out := []byte(Strip(string(src)))
		if bytes.Equal(out, src) {
			continue
		}
		info, err := os.Stat(path)
		if err != nil {
			return err
		}
		if err := os.WriteFile(path, out, info.Mode().Perm()); err != nil {
			return err
		}
	}
	return nil
}

func writeStdout(name string, src []byte) error {
	if err := checkText(name, src); err != nil {
		return err
	}
	_, err := io.WriteString(os.Stdout, Strip(string(src)))
	return err
}

// checkText rejects anything that isn't UTF-8 text, whatever the extension says.
func checkText(name string, src []byte) error {
	if bytes.IndexByte(src, 0) >= 0 {
		return fmt.Errorf("%s: not a text file (contains NUL bytes)", name)
	}
	if !utf8.Valid(src) {
		return fmt.Errorf("%s: not valid UTF-8 text", name)
	}
	return nil
}
