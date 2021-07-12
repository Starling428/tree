package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	startPad   = "│         "
	dirPad     = "├───"
	lastdirPad = "└───"
	longSpace  = "         "
	shortSpace = "   "
	maxSpace   = "            "
	lastPad    = " "
	byteSize   = 12
	hidden     *bool
)

func tree(dir string, pad *string, d int) {
	d++
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !*hidden && file.Name()[0:1] == "." {
			continue
		}
		if file.IsDir() {
			*pad = ((*pad)[:len(*pad)-byteSize]) + dirPad
			if files[len(files)-1] == file {
				*pad = (*pad)[:len(*pad)-byteSize]
				fmt.Println(strings.Replace(*pad, longSpace, shortSpace, -1)+lastdirPad+file.Name(), d)

				subfiles, err := os.ReadDir(path.Join(dir, file.Name()))
				if err != nil {
					fmt.Println(err)
					return
				}
				if len(subfiles) != 0 && len(*pad) > 11 {
					*pad = (*pad)[:len(*pad)-byteSize]
					*pad = (*pad) + maxSpace
				}
				*pad = (*pad) + dirPad
				if d != 0 {
					*pad = startPad + *pad
				} else {
					*pad = longSpace + lastPad + *pad
				}

				tree(path.Join(dir, file.Name()), pad, d)

			} else {
				fmt.Println(strings.Replace(*pad, longSpace, shortSpace, -1)+file.Name(), d)
				*pad = startPad + *pad
				tree(path.Join(dir, file.Name()), pad, d)
			}
		} else {
			if d != 0 {
				*pad = ((*pad)[:len(*pad)-byteSize]) + maxSpace
			} else {
				*pad = ((*pad)[:len(*pad)-byteSize]) + startPad
			}
			specPad := strings.Replace(*pad, longSpace, shortSpace, -1)
			fmt.Println(specPad[:len(specPad)-d]+file.Name(), d)
		}

	}
	if len(*pad) > 11 {
		*pad = (*pad)[byteSize:]
	}
	d--
}

func main() {

	hidden = flag.Bool("hidden", false, "Show hidden files")

	flag.Parse()

	curdir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	padding := startPad
	tree(curdir, &padding, -1)

}
