package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	outfile := flag.String("o", "concat.txt", "output filename")
	filter := flag.String("f", "*.*", "regex for file selection")
	flag.Parse()

	wd, err := os.Getwd()
	fmt.Println(wd)
	files, err := filepath.Glob(filepath.Join(wd, *filter))

	if err != nil {
		fmt.Print("Filepath.Glob(*filepath)", err)
		return
	}

	sort.Strings(files)
	var output []byte

	out_buf := bytes.NewBuffer(output)

	for _, f := range files {
		p, e := filepath.Rel(wd, f)
		if e != nil {
			return
		}

		contents, err := ioutil.ReadFile(f)
		if err == nil {
			fmt.Println("reading contents of ", p)
			out_buf.WriteString("// ")
			out_buf.WriteString(p)
			out_buf.WriteString("\r\n")
			out_buf.Write(contents)
		} else {
			fmt.Println("ioutil.ReadFile(f):", err)
			return
		}
	}
	fmt.Println("writing contents of files to ", *outfile)
	o_fh, err := os.Create(*outfile)
	if err != nil {
		fmt.Println(err)
	} else {
		o_fh.Write(out_buf.Bytes())
	}
	o_fh.Close()
}
