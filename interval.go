package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const usage = `usage: interval duration

interval prints the current time and an ASCII BEL character at intervals that
are duration long. See https://golang.org/pkg/time/#ParseDuration for the syntax
of the duration.`

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
	}
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	d, err := time.ParseDuration(args[0])
	if err != nil {
		log.Fatal(err)
	}
	for {
		now := time.Now().Format(time.ANSIC)
		fmt.Println(now + "\x07")
		time.Sleep(d)
	}
}
