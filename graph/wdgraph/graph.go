package wgraph

import (
	"container/list"
	"fmt"
	"github.com/mangotree2/ds-algo/queue/pqueue"
	"math"
)

//图的 邻接表法存储， 无向图
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

type Edge struct {
	sid int //边的起始顶点号
	tid int //边的终止顶点号
	w int	//权重
}

func (g *Graph) AddEdge(s, t, w int) {
	g.adj[s].PushBack(Edge{s,t,w})

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

type Vertex struct {
	id int //顶点编号
	dist int //从起始顶点到这个顶点的距离

}

func (v *Vertex) CompareTo(o Vertex)int {
	if o.dist > v.dist {
		return -1
	} else {
		return 1
	}
}


//从顶点s到顶点t的最短路径
//这里没测试
func (g *Graph)Dijkstra(s, t int) {
	predecessor := make([]int,g.v)//用来还原最短路径
	vs := make([]Vertex,g.v)//记录起始顶点到这个顶点的距离
	for i:= 0; i < g.v;i++ {
		vs[i].id = i
		vs[i].dist = math.MaxUint64
	}

	pq := pqueue.New(g.v) //小顶堆实现的优先队列
	inQueue := make([]bool,g.v)

	pq.Push(pqueue.Item{
		Value:    vs[s],
		Priority: int64(vs[s].dist),
	})
	vs[s].dist = 0
	inQueue[s] = true

	for pq.Len() > 0 {
		 minV := pq.Pop().(pqueue.Item).Value.(Vertex) //取dist 最小的
		 if minV.id == t {
		 	break
		 }

		 for e := g.adj[minV.id].Front();e != nil ;e.Next() {// 取出一条minVetex相连的边
		 	e := e.Value.(Edge)
		 	nextV := vs[e.tid] // minV==> nextV
		 	//找到一条到nextV的更短路径
		 	//todo minV.dist ==> vs[minV.dist].dist测试点 还没测
			if minV.dist + e.w < nextV.dist {
				vs[e.tid].dist = minV.dist + e.w //更新dist
				predecessor[nextV.id] = minV.id //更新前驱节点
				if inQueue[nextV.id] == false {
					//放入队列
					pq.Push(pqueue.Item{
						Value:    vs[nextV.id],
						Priority: int64(vs[nextV.id].dist),
					})
					inQueue[nextV.id] = true
				}

			}

		 }

	}

	fmt.Printf("%d",s)
	printPrev(predecessor,s,t)


}


