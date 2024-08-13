package main

import (
	"flag"
	"fmt"
	"github.com/mihailaz/shadowsocks-utils/pkg/shadowsocks"
	"os"
)

var (
	version = "1.0.0"
	build   = ""
	date    = ""
)

func main() {
	isVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *isVersion {
		fmt.Printf("shadowsocks-utils version: %s.%s date: %s\n", version, build, date)
		os.Exit(0)
	}
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <ss or ssconf url>\n", os.Args[0])
		os.Exit(1)
	}
	settings, err := shadowsocks.Parse(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(settings.String())
	fmt.Println("encryption method: ", settings.EncryptionMethod)
	fmt.Println("host: ", settings.Host)
	fmt.Println("password: ", settings.Password)
}
