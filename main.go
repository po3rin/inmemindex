package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

const limit = 32

var re = regexp.MustCompile("[^a-zA-Z 0-9]+")

func clean(document string) string {
	return re.ReplaceAllString(strings.ToLower(document), "")
}

// schema-independent index
type postings struct {
	ps   []int
	next *postings
}

func min(i, j int) int {
	if i >= j {
		return j
	}
	return i
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func (p *postings) put(pos int, workcnt int) {
	if p.next != nil {
		p.next.put(pos, workcnt+1)
		return
	}

	if cap(p.ps) <= len(p.ps) {
		p.next = &postings{
			ps: make([]int, 0, min(limit, max(2, workcnt*2))),
		}
		p.next.put(pos, workcnt+1)
	}

	p.ps = append(p.ps, pos)
}

func (p *postings) get() []int {
	list := make([]int, 0)
	return p.getTraverse(list)
}

func (p *postings) getTraverse(list []int) []int {
	list = append(list, p.ps...)
	if p.next != nil {
		return p.next.getTraverse(list)
	}
	return list
}

type dict map[string]*postings

type term struct {
	term string
	pos  int
}

type terms []term

// func (t terms) Len() int {
// 	return len(t)
// }

// func (t terms) Swap(i, j int) {
// 	t[i], t[j] = t[j], t[i]
// }

// func (t terms) Less(i, j int) bool {
// 	return t[i].term < t[j].term
// }

func main() {
	datadir := "simpledata"

	files, err := ioutil.ReadDir(datadir)
	if err != nil {
		panic(err)
	}

	terms := make(terms, 0)
	var pos int
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		b, err := ioutil.ReadFile(filepath.Join(datadir, file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		tokens := strings.Split(string(b), " ")
		for _, t := range tokens {
			pos++
			t = clean(t)
			terms = append(terms, term{
				term: t,
				pos:  pos,
			})
		}
	}

	dict := make(dict, 0)
	for _, t := range terms {
		var e *postings
		e, ok := dict[t.term]
		if !ok {
			e = &postings{
				ps: make([]int, 0, 4),
			}
			dict[t.term] = e
		}
		e.put(t.pos, 1)
	}

	for t, e := range dict {
		fmt.Println("--------------")
		fmt.Printf("term: %v\n", t)
		ps := e.get()
		fmt.Printf("pos: %v\n", ps)
	}
}
