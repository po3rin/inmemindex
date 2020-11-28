package main

import "fmt"

type dic struct {
	entries []interface{}
}

func (d *dic) getOrCreate(token string) interface{} {
	return nil
}

func (d *dic) sort() {
}

type postingList struct{}

func (p *postingList) add() {
}

func main() {
	var pos int
	dic := &dic{}
	pl := &postingList{}

	tokens := []string{}
	for _, t := range tokens {
		e := dic.getOrCreate(t)
		_ = e
		pl.add()
		pos++
	}

	dic.sort()

	for _, e := range dic.entries {
		_ = e
		fmt.Println("write")
	}
}
