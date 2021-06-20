package main

import (
	"encoding/xml"
	"fmt"
	"go/build"
	"io"
	"os"
	"time"

	"github.com/bvobart/gocover-cobertura/cobertura"
	"github.com/bvobart/gocover-cobertura/gocover"
)

const coberturaDTDDecl = "<!DOCTYPE coverage SYSTEM \"http://cobertura.sourceforge.net/xml/coverage-04.dtd\">\n"

func main() {
	convert(os.Stdin, os.Stdout)
}

func convert(in io.Reader, out io.Writer) {
	profiles, err := gocover.ParseProfiles(in)
	if err != nil {
		panic("Can't parse profiles")
	}

	srcDirs := build.Default.SrcDirs()
	sources := make([]*cobertura.Source, len(srcDirs))
	for i, dir := range srcDirs {
		sources[i] = &cobertura.Source{dir}
	}

	coverage := cobertura.Coverage{Sources: sources, Packages: nil, Timestamp: time.Now().UnixNano() / int64(time.Millisecond)}
	coverage.FromProfiles(profiles)

	fmt.Fprintf(out, xml.Header)
	fmt.Fprintf(out, coberturaDTDDecl)

	encoder := xml.NewEncoder(out)
	encoder.Indent("", "\t")
	err = encoder.Encode(coverage)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(out)
}
