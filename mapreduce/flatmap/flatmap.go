package main

import (
	"fmt"
	"sync"

	"github.com/zeromicro/go-zero/core/mr"
)

var (
	persons = []string{"john", "mary", "alice", "bob"}
	friends = map[string][]string{
		"john":  {"harry", "hermione", "ron"},
		"mary":  {"sam", "frodo"},
		"alice": {},
		"bob":   {"jamie", "tyrion", "cersei"},
	}
)

func main() {
	var (
		allFriends []string
		lock       sync.Mutex
	)
	mr.ForEach(func(source chan<- interface{}) {
		for _, each := range persons {
			source <- each
		}
	}, func(item interface{}) {
		lock.Lock()
		defer lock.Unlock()
		allFriends = append(allFriends, friends[item.(string)]...)
	}, mr.WithWorkers(100))
	fmt.Println(allFriends)
}
