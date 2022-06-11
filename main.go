package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pschlump/filelib"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "mk-vcard: Usage: %s files...\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	fns := flag.Args()

	for _, fn := range fns {

		file, err := filelib.Fopen(fn, "r")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open %s error:%s\n", fn, err)
			continue
		}
		defer file.Close()

		b, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read %s error:%s\n", fn, err)
			continue
		}

		str := string(b)
		qrFname := fn + ".png"

		err = qrcode.WriteFile(str, qrcode.Medium, 256, qrFname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to write %s .png file, error:%s\n", fn, err)
		}

	}

}
