package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/iamabhishek-dubey/dockerlinter/linter"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK = iota
	ExitCodeParseFlagsError
	ExitCodeNoExistError
	ExitCodeFileError
	ExitCodeAstParseError
	ExitCodeLintCheckError
)

const name = "dockerlinter"

const version = "0.0.1"

const usage = `dockerlinter - Dockerfile Linter written in Golang
Usage: dockerlinter [--ignore RULECODE]
  Lint Dockerfile for errors and best practices
Available options:
  --ignore Provide the rule code which you want to ignore.
Other Commands:
  --help	-h	Help about any command
  --version	-v	Print the version information
`

// CLI represents CLI interface
type CLI struct {
	OutStream, ErrStream io.Writer
}

type sliceString []string

func (ss *sliceString) String() string {
	return fmt.Sprintf("%s", *ss)
}

func (ss *sliceString) Set(value string) error {
	*ss = append(*ss, value)
	return nil
}

// Run it takes Dockerfile as an argument and applies it to analyzer to standard output.
func (cli *CLI) Run(args []string) int {
	var ingnoreRules sliceString
	var isVersion bool

	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.Usage = func() {
		_, _ = fmt.Fprint(cli.OutStream, usage)
	}

	flags.Var(&ingnoreRules, "ignore", "Set ignore strings")
	flags.BoolVar(&isVersion, "version", false, "version")
	flags.BoolVar(&isVersion, "v", false, "version")

	if err := flags.Parse(args[1:]); err != nil {
		_, _ = fmt.Fprintf(cli.ErrStream, "%s\n", err)
		return ExitCodeParseFlagsError
	}

	if isVersion {
		_, _ = fmt.Fprintf(cli.OutStream, "dockerlinter version %s\n", version)
		return ExitCodeOK
	}

	length := len(args)
	// The Dockerfile to be analyzed must be the last.
	if length < 2 {
		_, _ = fmt.Fprintf(cli.ErrStream, "Please provide a Dockerfile\n")
		return ExitCodeNoExistError
	}

	file := args[length-1]
	f, err := os.Open(file)
	if err != nil {
		_, _ = fmt.Fprintf(cli.ErrStream, "%s\n", err)
		return ExitCodeFileError
	}

	r, err := parser.Parse(f)
	if err != nil {
		_, _ = fmt.Fprintf(cli.ErrStream, "%s\n", err)
		return ExitCodeAstParseError
	}

	analyzer := linter.NewAnalyzer(ingnoreRules)
	rst, err := analyzer.Run(r.AST)
	if err != nil {
		_, _ = fmt.Fprintf(cli.ErrStream, "%s\n", err)
		return ExitCodeLintCheckError
	}

	rst = sort.StringSlice(rst)
	var output string
	for _, s := range rst {
		// ends of each strings have "\n"
		output = output + s
	}
	_, _ = fmt.Fprint(cli.OutStream, output)
	return ExitCodeOK
}
