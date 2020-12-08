package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "hexcode",
		Usage: "hex encoding/decoding cli tool",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "encode",
				Aliases: []string{"e"},
			},
			&cli.BoolFlag{
				Name:    "decode",
				Aliases: []string{"d"},
			},
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i", "in"},
			},
		},
		Action: func(ctx *cli.Context) error {
			var (
				encodeFlag = ctx.Bool("encode")
				decodeFlag = ctx.Bool("decode")
				input      = ctx.String("input")
				in         io.Reader
			)
			if arg := ctx.Args().First(); arg != "" {
				input = arg
			}

			if input == "" {
				in = os.Stdin
			}

			if input != "" {
				in = strings.NewReader(input)
			}

			switch {
			case encodeFlag:
				err := encode(in, os.Stdout)
				if err != nil {
					return fmt.Errorf("cannot encode input: %s", err)
				}

				break
			case decodeFlag:
				err := decode(in, os.Stdout)
				if err != nil {
					return fmt.Errorf("cannot dencode input: %s", err)
				}

				break
			default:
				return nil
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func encode(input io.Reader, output io.Writer) error {
	b, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}

	w := hex.NewEncoder(output)

	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func decode(input io.Reader, output io.Writer) error {
	r := hex.NewDecoder(input)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	_, err = output.Write(b)
	if err != nil {
		return err
	}

	return nil
}
