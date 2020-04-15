//+build ignore

package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/shabbyrobe/go-fxhash"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var rust bool
	flag.BoolVar(&rust, "rust", false, "Use rust mode (prepend slice size)")
	flag.Parse()

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	var hash uint64
	if rust {
		var num = make([]byte, 8)
		binary.LittleEndian.PutUint64(num, uint64(len(buf)))
		hash = fxhash.Append64(hash, num)
		hash = fxhash.Append64(hash, buf)
	} else {
		hash = fxhash.Sum64(buf)
	}

	fmt.Println(hash)

	return nil
}
