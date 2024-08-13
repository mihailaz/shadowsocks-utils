package main

import (
	"flag"
	"fmt"
	"github.com/mihailaz/shadowsocks-utils/pkg/shadowsocks"
	"log/slog"
	"os"
)

var (
	version = "1.0.0"
	build   = ""
	date    = ""
)

func main() {
	verFlag := flag.Bool("version", false, "print version and exit")
	debugFlag := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	if *verFlag {
		fmt.Printf("shadowsocks-utils version: %s.%s date: %s\n", version, build, date)
		os.Exit(0)
	}
	if *debugFlag {
		logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
		slog.SetDefault(logger)
		slog.Debug("debug mode: on")
	}
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <ss or ssconf url>\n", os.Args[0])
		os.Exit(1)
	}
	uri := os.Args[len(os.Args)-1]
	info, err := shadowsocks.Parse(uri)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("encryption method: ", info.EncryptionMethod)
	fmt.Println("host: ", info.Host)
	fmt.Println("password: ", info.Password)
}
