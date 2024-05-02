/*
 * This file was last modified at 2024.01.04 15:07 by Victor N. Skurikhin.
 * This is free and unencumbered software released into the public domain.
 * For more information, please refer to <http://unlicense.org>
 * main_test.go
 * $Id$
 */
//!+
package main

import (
	"bufio"
	"bytes"
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{name: "Test case #0",
			input: "0\n0\n0\n",
			want:  "0\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var b bytes.Buffer
			w := bufio.NewWriter(&b)
			wrap(strings.NewReader(test.input), w, do)
			got := b.String()
			assert.Equal(t, test.want, got)
		})
	}
}
func TestMain(m *testing.M) {
	flag.Parse()
	exitCode := m.Run()
	// Exit
	os.Exit(exitCode)
}

func TestWithoutArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"main"}
	main()
}

func BenchmarkTimeSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(8 * time.Nanosecond)
	}
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
