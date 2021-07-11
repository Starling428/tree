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
	byteSize   = 12
	hidden     *bool
)

func tree(dir string, pad *string, d int) {
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
				fmt.Println(strings.Replace(*pad, longSpace, shortSpace, -1) + lastdirPad + file.Name())
				*pad = (*pad) + dirPad
			} else {
				fmt.Println(strings.Replace(*pad, longSpace, shortSpace, -1) + file.Name())
			}
		} else {
			if d != 0 {
				*pad = ((*pad)[:len(*pad)-byteSize]) + maxSpace
			} else {
				*pad = ((*pad)[:len(*pad)-byteSize]) + startPad
			}
			fmt.Println(strings.Replace(*pad, longSpace, shortSpace, -1) + file.Name())
		}
		fi, err := os.Stat(path.Join(dir, file.Name()))
		if err != nil {
			fmt.Println(err)
			return
		}
		if fi.IsDir() {
			*pad = startPad + *pad
			d++
			tree(path.Join(dir, file.Name()), pad, d)
		}
	}
	*pad = (*pad)[byteSize:]
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
	tree(curdir, &padding, 0)
	//fi, err := os.Stat("D://Users/moses.sarkisov/go")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//os.Chdir(fi.Name())
	//fmt.Println(fi.Mode().String())
	//if fi.IsDir() {
	//	tree(fi.Name())
	//}

	//fmt.Println(padding)

	/*
		name := ".idea"
		fi, err := os.Stat(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			// do directory stuff
			fmt.Println("directory")
		case mode.IsRegular():
			// do file stuff
			fmt.Println("file")
		}
	*/
}
