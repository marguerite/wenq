package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/marguerite/go-stdlib/dir"
	"github.com/marguerite/wenq/glyphutils"
	"github.com/marguerite/wenq/ufo3"
)

func main() {
	cwd, _ := os.Getwd()
	directories, _ := dir.Glob(filepath.Dir(filepath.Dir(cwd)) + "/WenQuanYiZenHei*.ufo3")

	for _, v := range directories {
		files, _ := dir.Ls(filepath.Join(v, "glyphs"), true, true)

		for _, f := range files {
			if !strings.HasSuffix(f, ".glif") {
				continue
			}
			glyph := ufo3.NewGlyphFromFile(f)

			for i, v := range glyph.Outline.Contours {
				j, p := v.FindPointByX("4")
				if p.IsNil() {
					continue
				}
				if j != len(v.Points)-1 {
					continue
				}
				fmt.Printf("fixing %s\n", glyphutils.CodepointFromGlifFileName(f))
				glyph.DeletePoint(i, j)
			}

			ioutil.WriteFile(f, glyph.Bytes(), 0644)
		}
	}
}
