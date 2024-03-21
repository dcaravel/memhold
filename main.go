package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("missing size in MiB")
	}

	sizeStr := os.Args[1]

	size, err := strconv.ParseUint(sizeStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	sizeBytes := size * 1024 * 1024
	memBuf := make([]byte, sizeBytes)

	log.Printf("Filling %d MiB of mem..", size)
	rand.Read(memBuf)
	log.Print("Done filling mem, holding...")

	go startReportStatus(memBuf)

	exitOnSig()
}

// memBuf is passed here and logged to avoid GC from cleaning it up.
func startReportStatus(memBuf []byte) {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		log.Printf("Alloc: %v MB\tSys: %v MB\tAt %p", m.Alloc/1024/1024, m.Sys/1024/1024, &memBuf)
		time.Sleep(5 * time.Second)
	}
}

func exitOnSig() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println()
	log.Printf("Exiting...")
	os.Exit(1)
}
