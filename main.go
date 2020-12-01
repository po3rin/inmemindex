package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// docid index
var postingsNode struct {
	next *postingsNode
	vals []int // doc id list
}

type term struct {
	id  int
	val string
}

type terms []term

func main() {
	datadir := "testdata"
	files, err := ioutil.ReadDir(datadir)
	if err != nil {
		panic(err)
	}

	ht := NewHashTable()

	for i, file := range files {
		if file.IsDir() {
			continue
		}
		b, err := ioutil.ReadFile(filepath.Join(datadir, file.Name()))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tokens := strings.Split(string(b), " ")
		for _, t := range tokens {
			e, ok := ht.Get(t)
			if !ok {

			}
			pn := postingsNode{
				id:  i,
				val: t,
			}
		}
	}
}
