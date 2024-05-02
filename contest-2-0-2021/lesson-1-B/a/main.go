/*
 * This file was last modified at 2024-05-02 10:50 by Victor N. Skurikhin.
 * main.go
 * $Id$
 */
//!+
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

func flush()                   { out.Flush() }
func scan(a ...interface{})    { fmt.Fscan(in, a...) }
func println(a ...interface{}) { fmt.Fprintln(out, a...) }

func do() {

	var r, i, c int
	scan(&r, &i, &c)

	result := solution(r, i, c)

	println(result)
	flush()
}

func solution(r int, i int, c int) int {

	switch {
	case i == 0:
		if r != 0 {
			return 3
		} else {
			return c
		}
	case i == 1:
		return c
	case i == 4:
		if r != 0 {
			return 3
		} else {
			return 4
		}
	case i == 6:
		return 0
	case i == 7:
		return 1
	default:
		return i
	}
}

func wrap(i io.Reader, o io.Writer, do func()) {
	in = bufio.NewReader(i)
	out = bufio.NewWriter(o)
	do()
}

func main() {
	wrap(os.Stdin, os.Stdout, do)
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
