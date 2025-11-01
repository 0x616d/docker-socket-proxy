package main

type router struct {
	routes map[string]*pathTrie
}

func newRouter() *router {
	return &router{routes: make(map[string]*pathTrie)}
}

func (r *router) insert(method, path string) {
	root, ok := r.routes[method]
	if !ok {
		root = newPathTrie()
		r.routes[method] = root
	}
	root.insert(path)
}

func (r *router) search(method, path string) bool {
	root, ok := r.routes[method]
	if !ok {
		return false
	}
	return root.search(path)
}
