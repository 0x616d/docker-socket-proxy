package main

import (
	"strings"
)

type pathTrie struct {
	child map[string]*pathTrie
	end   bool
}

func newPathTrie() *pathTrie {
	return &pathTrie{child: make(map[string]*pathTrie)}
}

func (p *pathTrie) insert(path string) {
	segments := strings.Split(strings.Trim(path, "/"), "/")
	node := p
	for _, segment := range segments {
		if childNode, exist := node.child[segment]; exist {
			node = childNode
		} else {
			node.child[segment] = newPathTrie()
			node = node.child[segment]
		}
	}
	node.end = true
}

func (p *pathTrie) search(path string) bool {
	segments := strings.Split(strings.Trim(path, "/"), "/")
	node := p
	for _, segment := range segments {
		if n, ok := node.child[segment]; ok {
			node = n
			continue
		}
		if n, ok := node.child["*"]; ok {
			node = n
			continue
		}
		return false
	}
	return node.end
}
