//No prefix set
//problem : https://www.hackerrank.com/challenges/no-prefix-set

package main

import (
	"fmt"
	"errors"
)

const aCode = 97

type TrieNode struct {
	next *Trie
	last bool
}

type Trie struct {
	chars [10]*TrieNode
}

func (t *Trie) Add(word string) (error) {

	size := len(word)
	curr := t
	var err error

	for i:=0;i<size;i++ {

		 curr , err = curr.add( word[i] , i == size-1 )

		 if err != nil {
			return err
		 }

	}

	return nil

}

func (t *Trie) add(ch byte , last bool) (*Trie,error) {

	index := ch - aCode
	if t.chars[index] == nil {
		t.chars[index] = new(TrieNode)

		t.chars[index].next = new(Trie)
		if last {
			t.chars[index].last = true
		}
	} else if t.chars[index].last || last {
			return nil , errors.New("BAD SET")
	}

	return t.chars[index].next , nil
}

func main() {

	var n int
	fmt.Scan(&n)

	t := new(Trie)

	end := false
	var word string
	for i:=1;i<=n;i++ {

  		fmt.Scan(&word)

		 if !end {

			 err := t.Add( word )
			 if err != nil {
				fmt.Println( err.Error() )
				fmt.Println( word )
				end = true
			 }
		 }
	}

	if end==false {
		fmt.Println("GOOD SET")
	}
}
