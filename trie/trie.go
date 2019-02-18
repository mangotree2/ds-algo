package trie

type Node struct {
	val      rune
	term     bool
	depth    int
	meta     interface{}
	parent   *Node
	children map[rune]*Node
}

type Trie struct {
	root *Node
	size int
}

const nul = 0x0

// 使用初始化的根节点创建新的Trie。
func New() *Trie {
	return &Trie{
		root: &Node{children: make(map[rune]*Node), depth: 0},
		size: 0,
	}
}

// 返回Trie的根节点。
func (t *Trie) Root() *Node {
	return t.root
}

//将key添加到Trie，包括value。
func (t *Trie) Add(key string, meta interface{}) *Node {
	//fmt.Println("Add begin key:",key)
	t.size++
	runes := []rune(key)
	node := t.root

	for i := range runes {
		r := runes[i]

		if n, ok := node.children[r]; ok {
			node = n
		} else {
			node = node.NewChild(r, nil, false)
		}
	}
	node = node.NewChild(nul,  meta, true)
	return node
}

//完整查找
func (t *Trie) Find(key string) (*Node, bool) {
	node := findNode(t.Root(), []rune(key))
	if node == nil {
		return nil, false
	}

	node, ok := node.Children()[nul]
	if !ok || !node.term {
		return nil, false
	}

	return node, true
}

//
func (t *Trie) HasKeysWithPrefix(key string) bool {
	node := findNode(t.Root(), []rune(key))
	return node != nil
}

//删除字符串
func (t *Trie) Remove(key string) {
	var (
		i    int
		rs   = []rune(key)
		node = findNode(t.Root(), []rune(key))
	)

	if node == nil {
		return
	}

	t.size--
	for n := node.Parent(); n != nil; n = n.Parent() {
		i++
		if len(n.Children()) > 1 {
			r := rs[len(rs)-i]
			n.RemoveChild(r)
			break
		}
	}
}

// 返回所有字符串
func (t *Trie) Keys() []string {
	return t.PrefixSearch("")
}

// 执行前缀搜索。
func (t Trie) PrefixSearch(pre string) []string {
	node := findNode(t.Root(), []rune(pre))
	if node == nil {
		return nil
	}

	return collect(node)
}

// 创建并返回指向该节点的新子节点的指针。
func (n *Node) NewChild(val rune, meta interface{}, term bool) *Node {
	node := &Node{
		val:      val,
		term:     term,
		meta:     meta,
		parent:   n,
		children: make(map[rune]*Node),
		depth:    n.depth + 1,
	}
	n.children[val] = node
	return node
}

func (n *Node) RemoveChild(r rune) {
	delete(n.children, r)

}

// 返回此节点的父节点。
func (n Node) Parent() *Node {
	return n.parent
}

// value
func (n Node) Meta() interface{} {
	return n.meta
}

func (n Node) Children() map[rune]*Node {
	return n.children
}

func (n Node) Terminating() bool {
	return n.term
}

func (n Node) Val() rune {
	return n.val
}

func (n Node) Depth() int {
	return n.depth
}



func findNode(node *Node, runes []rune) *Node {
	if node == nil {
		return nil
	}

	if len(runes) == 0 {
		return node
	}

	n, ok := node.Children()[runes[0]]
	if !ok {
		return nil
	}

	var nrunes []rune
	if len(runes) > 1 {
		nrunes = runes[1:]
	} else {
		nrunes = runes[0:0]
	}

	return findNode(n, nrunes)
}

func collect(node *Node) []string {
	var (
		keys []string
		n    *Node
		i    int
	)
	nodes := []*Node{node}
	for l := len(nodes); l != 0; l = len(nodes) {
		i = l - 1
		n = nodes[i]
		nodes = nodes[:i]
		for _, c := range n.children {
			nodes = append(nodes, c)
		}
		if n.term {
			word := ""
			for p := n.parent; p.depth != 0; p = p.parent {
				word = string(p.val) + word
			}
			keys = append(keys, word)
		}
	}
	return keys
}

