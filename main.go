package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func tree(dir string, pad *string, d bool) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name()[0:1] == "." {
			continue
		}
		if file.IsDir() {
			*pad = ((*pad)[:len(*pad)-12]) + "├───"
			if files[len(files)-1] == file {
				*pad = (*pad)[:len(*pad)-12]
				fmt.Println(strings.Replace(*pad, "         ", "   ", -1) + "└───" + file.Name())
				*pad = (*pad) + "├───"
			} else {
				fmt.Println(strings.Replace(*pad, "         ", "   ", -1) + file.Name())
			}
		} else {
			if d {
				*pad = ((*pad)[:len(*pad)-12]) + "            "
			} else {
				*pad = ((*pad)[:len(*pad)-12]) + "│         "
			}
			fmt.Println(strings.Replace(*pad, "         ", "   ", -1) + file.Name())
		}
		fi, err := os.Stat(path.Join(dir, file.Name()))
		if err != nil {
			fmt.Println(err)
			return
		}
		if fi.IsDir() {
			*pad = "│         " + *pad
			tree(path.Join(dir, file.Name()), pad, true)
		}
	}
	*pad = (*pad)[12:]
}

func main() {

	padding := "│         "
	fmt.Println("D://Users/moses.sarkisov/go")
	tree("D://Users/moses.sarkisov/go", &padding, false)

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
