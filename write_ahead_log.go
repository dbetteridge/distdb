package main

import (
	"bufio"
	"os"
	"time"
)

type WAL struct {
	modified time.Time
	writes   []string
}

func (w *WAL) new() {
	w.writes = make([]string, 0)
	w.modified = time.Now()
}

func (w *WAL) write(data string) {
	w.writes = append(w.writes, data)
	w.modified = time.Now()

	f, err := os.OpenFile("./wal.db", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(data); err != nil {
		panic(err)
	}
}

func (w *WAL) flush() []string {
	f, err := os.OpenFile("./wal.db", os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = f.Truncate(0); err != nil {
		panic(err)
	}

	if _, err = f.Seek(0, 0); err != nil {
		panic(err)
	}

	return lines
}
