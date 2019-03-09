package nwdgraph

import (
	"container/list"
	"fmt"
)

//图的 邻接表法存储， 有向图 无权图
type Graph struct {
	adj []*list.List
	v int
}

func NewGraph(v int) *Graph {
	g := &Graph{
		adj: make([]*list.List,v),
		v:   v,
	}
	for i := range g.adj {
		g.adj[i] = list.New()
	}

	return g
}

func (g *Graph) AddEdge(s, t int) {
	g.adj[s].PushBack(t)

}
//print path recursively
func printPrev(prev []int, s int, t int) {

	if t == s || prev[t] == -1 {
		fmt.Printf("%d ", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf("%d ", t)
	}

}

func (g *Graph) VCnt() int {
	return g.v
}

func (g *Graph) VList(v int) *list.List {
	return g.adj[v]
}


//广度优先搜索 最短路径
func (g *Graph) BFS(s, t int) {
	if s == t {
		return
	}

	visited := make([]bool,g.v) //用来记录已经被访问的顶点，用来避免顶点被重复访问。
	queue := []int{} // 用来存储已经被访问、但相连的顶点还没有被访问的顶点。因为广度优先搜索是逐层访问的，也就是说，我们只有把第k层的顶点都访问完成之后，才能访问第k+1层的顶点。当我们访问到第k层的顶点的时候，我们需要把第k层的顶点记录下来，稍后才能通过第k层的顶点来找第k+1层的顶点。所以，我们用这个队列来实现记录的功能。
	prev := []int{} //用来记录搜索路径。当我们从顶点s开始，广度优先搜索到顶点t后，prev数组中存储的就是搜索的路径。不过，这个路径是反向存储的。prev[w]存储的是，顶点w是从哪个前驱顶点遍历过来的。比如，我们通过顶点2的邻接表访问到顶点3，那prev[3]就等于2。为了正向打印出路径，我们需要递归地来打印

	queue = append(queue,s)
	visited[s] = true
	for i := range prev {
		prev[i] = -1
	}
	isFound := false

	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		for e := g.adj[top].Front();e!=nil;e=e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue,k)
				visited[k] = true
			}

		}

	}

	if isFound {
		printPrev(prev,s,t)
	}

}

//深度优先搜索
func (g *Graph) DFS(s , t int)  {

	prev := make([]int,g.v)
	for i := range prev {
		prev[i] = -1
	}

	visited := make([]bool,g.v)
	visited[s] = true
	g.recurse(s,t,prev,visited,false)
	printPrev(prev, s, t)

}

func (g *Graph) recurse(s, t int, prev []int, visited []bool, isFound bool) {
	if isFound {
		return
	}
	visited[s] = true

	if s == t {
		isFound = true
		return
	}

	for e := g.adj[s].Front();e != nil; e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			g.recurse(k,t,prev,visited,isFound)
		}
	}

}


