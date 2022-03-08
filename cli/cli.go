package cli

import (
	"fmt"
	"io"
	"runtime"

	"github.com/itchyny/base58-go"
	"github.com/jessevdk/go-flags"
)

const name = "base58"
const version = "0.0.1"

var revision = "HEAD"

const (
	exitCodeOK = iota
	exitCodeErr
)

type cli struct {
	inStream  io.Reader
	outStream io.Writer
	errStream io.Writer
}

type flagopts struct {
	Decode   bool             `short:"D" long:"decode" description:"decodes input"`
	Encoding *base58.Encoding `short:"e" long:"encoding" default:"flickr" choice:"flickr" choice:"ripple" choice:"bitcoin" description:"encoding name"`
	Input    []string         `short:"i" long:"input" default:"-" description:"input file"`
	Output   string           `short:"o" long:"output" default:"-" description:"output file"`
	Version  bool             `short:"v" long:"version" description:"print version"`
}

func (cli *cli) run(args []string) int {
	var opts flagopts
	args, err := flags.NewParser(
		&opts, flags.HelpFlag|flags.PassDoubleDash,
	).ParseArgs(args)
	if err != nil {
		if err, ok := err.(*flags.Error); ok && err.Type == flags.ErrHelp {
			fmt.Fprintln(cli.outStream, err.Error())
			return exitCodeOK
		}
		fmt.Fprintln(cli.errStream, err.Error())
		return exitCodeErr
	}
	if opts.Version {
		fmt.Fprintf(cli.outStream, "%s %s (rev: %s/%s)\n", name, version, revision, runtime.Version())
		return exitCodeOK
	}
	status := exitCodeOK
	return status
}
