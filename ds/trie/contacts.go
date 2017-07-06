//Contacts
//Problem : https://www.hackerrank.com/challenges/contacts

package main

import "fmt"

const aCode = 97

type TrieNode struct {
	next *Trie
	count int
}

type Trie struct {

	nodes [26]*TrieNode

}

func (t *Trie) PrefixCount(prefix string) int {

	curr := t

	for i:=0;i<len(prefix);i++ {

		index := prefix[i]-aCode

		if curr.nodes[index] == nil {
			return 0
		} else {

			if i == len(prefix)-1 {
				return curr.nodes[index].count
			} else {
				curr = curr.nodes[index].next
			}
		}
	}

	return 0

}

func (t *Trie) Add(word string) {

	curr := t

	for i:=0;i<len(word);i++ {

		curr = curr.add( word[i] )

	}

}

func (t *Trie) add(ch byte) *Trie {

	index := ch-aCode

	if t.nodes[index] == nil {
		t.nodes[index] = new(TrieNode)
		t.nodes[index].next = new(Trie)
	}

	t.nodes[index].count++

	return t.nodes[index].next
}

func main() {

	var n int
	fmt.Scan(&n)

	var op,word string
	t := Trie{}

	for i:=1;i<=n;i++ {
		fmt.Scan(&op)
		fmt.Scan(&word)

		if op == "add" {
			t.Add(word)
		} else if op == "find" {
			fmt.Println( t.PrefixCount(word) )
		}
	}

}
