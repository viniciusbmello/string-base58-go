package main

import (
	"fmt"
	"io"

	"github.com/itchyny/base58-go"
	"github.com/jessevdk/go-flags"
)

type cli struct {
	inStream  io.Reader
	outStream io.Writer
	errStream io.Writer
}

const (
	exitCodeOK = iota
	exitCodeErr
)

func main() {

}

func (cli *cli) run(args []string) int {
	var opts struct {
		Decode   bool             `short:"D" long:"decode" description:"decodes input"`
		Encoding *base58.Encoding `short:"e" long:"encoding" default:"flickr" choice:"flickr" choice:"ripple" choice:"bitcoin" description:"encoding name"`
		Version  bool             `short:"v" long:"version" description:"print version"`
	}

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
	status := exitCodeOK
	return status
}
