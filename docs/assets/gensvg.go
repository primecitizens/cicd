// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build ignore

package main

import (
	_ "embed" // for go:embed
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	generateCoverageBadges(filepath.Join(dir, "coverage"))
}

var (
	//go:embed coverage/100.svg
	tmpl_coverage string
)

func generateCoverageBadges(dir string) {
	const (
		colorPlaceholder   = "#66cc8a"
		percentPlaceholder = "100%"
	)
	for i := 0; i < 100; i++ {
		color := colorPlaceholder
		percent := strconv.FormatInt(int64(i), 10)
		switch {
		case i < 40:
			color = "#ea5234"
		case i < 60: // 40 - 60
			color = "#ce3262"
		case i < 70: // 60 - 70
			color = "#dfb317"
		case i < 80: // 70 - 80
			color = "#fddd00"
		case i < 90: // 80 - 90
			color = "#2dbcaf"
		default:
			color = "#66cc8a"
		}

		filename := strings.Repeat("0", 3-len(percent)) + percent + ".svg"
		file, err := os.OpenFile(filepath.Join(dir, filename), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		func() {
			defer file.Close()

			_, err = strings.NewReplacer(
				colorPlaceholder, color,
				percentPlaceholder, percent+"%",
			).WriteString(file, tmpl_coverage)
			if err != nil {
				panic(err)
			}
		}()
	}
}
