package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	//
	"github.com/gsiems/go-read-wrap/srw"
)

var tr = map[byte]string{
	byte(0x00): " nul ",
	byte(0x01): " soh ",
	byte(0x02): " stx ",
	byte(0x03): " etx ",
	byte(0x04): " eot ",
	byte(0x05): " enq ",
	byte(0x06): " ack ",
	byte(0x07): " bel ",
	byte(0x08): "  bs ",
	byte(0x09): "  ht ",
	byte(0x0a): "  lf ",
	byte(0x0b): "  vt ",
	byte(0x0c): "  ff ",
	byte(0x0d): "  cr ",
	byte(0x0e): "  so ",
	byte(0x0f): "  si ",
	byte(0x10): " dle ",
	byte(0x11): " dc1 ",
	byte(0x12): " dc2 ",
	byte(0x13): " dc3 ",
	byte(0x14): " dc4 ",
	byte(0x15): " nak ",
	byte(0x16): " syn ",
	byte(0x17): " etb ",
	byte(0x18): " can ",
	byte(0x19): "  em ",
	byte(0x1a): " sub ",
	byte(0x1b): " esc ",
	byte(0x1c): "  fs ",
	byte(0x1d): "  gs ",
	byte(0x1e): "  rs ",
	byte(0x1f): "  us ",
	byte(0x7f): " del ",
}

func main() {

	var alNum bool
	var esChr bool
	var byteCount int
	var source string
	var target string
	var width int

	flag.BoolVar(&alNum, "a", false, "Output ASCII letters, numbers, and symbols.")
	flag.BoolVar(&esChr, "e", false, "Output codes for single character escape sequences.")
	flag.IntVar(&byteCount, "c", 0, "The number of bytes to dump. The default is to dump the entire file.")
	flag.StringVar(&source, "i", "", "The input file to hexdump.")
	flag.StringVar(&target, "o", "", "The output file to write to.")
	flag.IntVar(&width, "w", 16, "The number of bytes to output per line.")
	flag.Parse()

	r := openInput(source)
	defer deferredClose(r)
	s := srw.BuffReader(0, r)
	f := openOutput(target)
	defer deferredClose(f)
	w := bufio.NewWriter(f)

	i := 0
	c := 0

	for {

		if byteCount > 0 && c >= byteCount {
			break
		}
		c++

		if i >= width && width > 0 {
			i = 0
			writeStr(w, "\n")
		}
		i++

		b := make([]byte, 1)
		_, err := s.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		writeByte(w, b[0], esChr, alNum)
	}

	writeStr(w, "\n")
	w.Flush()
}

//
func deferredClose(f *os.File) {
	if cerr := f.Close(); cerr != nil {
		log.Fatal(fmt.Printf("File close failed: %q", cerr))
	}
}

// openInput opens the appropriate input source of bytes to read
func openInput(source string) (r *os.File) {

	var err error

	if source == "" || source == "-" {
		r = os.Stdin
	} else {
		r, err = os.Open(source)
		if err != nil {
			log.Fatal(fmt.Printf("File open failed: %q", err))
		}
	}
	return r
}

// openInput opens the appropriate target for writing output
func openOutput(target string) (f *os.File) {

	var err error

	if target == "" || target == "-" {
		f = os.Stdout
	} else {
		f, err = os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(fmt.Printf("File open failed: %q", err))
		}
	}
	return f
}

// writeByte formats a byte and writes the result
func writeByte(f io.Writer, b byte, esChr bool, alNum bool) {

	str := ""

	ok := false
	if esChr {
		str, ok = tr[b]
	}

	if esChr && ok {
		// should already have str, so do nothing
	} else if int(b) == 0 {
		str = "0x00 "
	} else if int(b) < 0x10 {
		str = fmt.Sprintf("0x0%x ", b)
	} else if alNum && int(b) >= 0x21 && int(b) <= 0x7e {
		str = fmt.Sprintf("   %s ", string(b))
	} else {
		str = fmt.Sprintf("%#v ", b)
	}

	writeStr(f, str)

}

// writeStr writes the supplied string to the output
func writeStr(f io.Writer, str string) {
	_, err := f.Write([]byte(str))
	if err != nil {
		log.Fatal(fmt.Printf("Failed to write to output: %q", err))
	}
}
