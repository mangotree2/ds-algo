package ac


type AcNode struct {
	val      rune
	term     bool
	depth    int
	meta     interface{}
	parent   *AcNode
	children map[rune]*AcNode
	fail	 *AcNode
}

type Ac struct {
	root *AcNode
	size int
}

const nul = 0x0

func New() *Ac {
	return &Ac{
		root: &AcNode{children: make(map[rune]*AcNode), depth: 0},
		size: 0,
	}
}

func (t *Ac) Root() *AcNode {
	return t.root
}

func (t *Ac) Add(key string, meta interface{}) *AcNode {
	t.size++
	runes := []rune(key)
	acNode := t.root

	for i := range runes {
		r := runes[i]

		if n, ok := acNode.children[r]; ok {
			acNode = n
		} else {
			acNode = acNode.NewChild(r, nil, false)
		}
	}
	acNode = acNode.NewChild(nul,  meta, true)
	return acNode
}

func (t *Ac)BuildFailPointer() {

	queue := make([]*AcNode,0)
	t.root.fail = nil

	queue = append(queue,t.root)

	for {
		if len(queue) == 0 {
			break
		}

		node := queue[0]
		queue = queue[1:]

		for _,nc := range node.children {

			if nc.term == true {
				continue
			}

			//root 得子节点得fail指针均指向root
			if node == t.root {
				nc.fail = t.root
			} else {
				q := node.fail
				for q != nil {
					qc,ok := q.children[nc.val]
					if ok {
						nc.fail = qc
						break
					}
					q = q.fail
				}

				if q == nil {
					nc.fail = t.root
				}
			}

			queue = append(queue,nc)

		}

	}

}

func (t *Ac) Match(input string) map[int]string {

	node := t.root
	var ok bool
	var temp *AcNode
	result := make(map[int]string)
	for i,r := range input {

		if _,ok = node.children[r];!ok && node != t.root {
			node = node.fail
		}

		if temp,ok = node.children[r];!ok {
			node = t.root
			temp = node
		} else {
			node = temp
		}

		for temp != t.root && !temp.term {
			if _,ok = temp.children[nul];ok {
				word := ""
				for p := temp; p.depth != 0; p = p.parent {
					word = string(p.val) + word
				}
				//fmt.Println(word)
				result[i-len(word)+1] = word
			}
			temp = temp.fail
		}


	}
	return result
}

func (t *Ac) Find(key string) (*AcNode, bool) {
	acNode := findNode(t.Root(), []rune(key))
	if acNode == nil {
		return nil, false
	}

	acNode, ok := acNode.Children()[nul]
	if !ok || !acNode.term {
		return nil, false
	}

	return acNode, true
}

func (t *Ac) HasKeysWithPrefix(key string) bool {
	acNode := findNode(t.Root(), []rune(key))
	return acNode != nil
}

func (t *Ac) Remove(key string) {
	var (
		i    int
		rs   = []rune(key)
		acNode = findNode(t.Root(), []rune(key))
	)

	if acNode == nil {
		return
	}

	t.size--
	for n := acNode.Parent(); n != nil; n = n.Parent() {
		i++
		if len(n.Children()) > 1 {
			r := rs[len(rs)-i]
			n.RemoveChild(r)
			break
		}
	}
}

func (t *Ac) Keys() []string {
	return t.PrefixSearch("")
}

func (t Ac) PrefixSearch(pre string) []string {
	acNode := findNode(t.Root(), []rune(pre))
	if acNode == nil {
		return nil
	}

	return collect(acNode)
}

func (n *AcNode) NewChild(val rune, meta interface{}, term bool) *AcNode {
	acNode := &AcNode{
		val:      val,
		term:     term,
		meta:     meta,
		parent:   n,
		children: make(map[rune]*AcNode),
		depth:    n.depth + 1,
	}
	n.children[val] = acNode
	return acNode
}

func (n *AcNode) RemoveChild(r rune) {
	delete(n.children, r)

}

func (n AcNode) Parent() *AcNode {
	return n.parent
}

func (n AcNode) Meta() interface{} {
	return n.meta
}

func (n AcNode) Children() map[rune]*AcNode {
	return n.children
}

func (n AcNode) Terminating() bool {
	return n.term
}

func (n AcNode) Val() rune {
	return n.val
}

func (n AcNode) Depth() int {
	return n.depth
}

func findNode(acNode *AcNode, runes []rune) *AcNode {
	if acNode == nil {
		return nil
	}

	if len(runes) == 0 {
		return acNode
	}

	n, ok := acNode.Children()[runes[0]]
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

func collect(node *AcNode) []string {
	var (
		keys []string
		n    *AcNode
		i    int
	)
	nodes := []*AcNode{node}
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