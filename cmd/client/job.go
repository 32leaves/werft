package cmd

// Copyright © 2019 Christian Weichel

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"os"

	"github.com/32leaves/werft/pkg/prettyprint"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var (
	outputFormat   string
	outputTemplate string
)

// jobCmd represents the job command
var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Interacts with currently running or previously run jobs",
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(jobCmd)

	jobCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "o", "template", "selects the output format: string, json, yaml, template")
	jobCmd.PersistentFlags().StringVar(&outputTemplate, "output-template", "", "template to use in combination with --output-format template")
}

func prettyPrint(obj proto.Message, defaultTpl string) error {
	format := prettyprint.Format(outputFormat)
	if !prettyprint.HasFormat(format) {
		return xerrors.Errorf("format %s is not supported", format)
	}

	tpl := outputTemplate
	if tpl == "" {
		tpl = defaultTpl
	}

	ctnt := &prettyprint.Content{
		Obj:      obj,
		Format:   format,
		Writer:   os.Stdout,
		Template: tpl,
	}
	return ctnt.Print()
}
