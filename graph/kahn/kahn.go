package kahn

import (
	"fmt"
	"github.com/mangotree2/ds-algo/graph/nwgraph"
)

//s=>t ：s先于t被执行
//拓扑都是有向的
func TopologySortByKahn(g *nwgraph.Graph) {
	inDegree := make([]int,g.VCnt())
	for i:=0;i<g.VCnt();i++{
		for e := g.VList(i).Front();e != nil;e.Next() {
			k := e.Value.(int)
			inDegree[k]++
		}
	}

	queue := []int{}
	for i := 0; i < g.VCnt();i++{
		if inDegree[i] == 0 {
			queue = append(queue,i)
		}
	}

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		fmt.Printf("->%d",i)
		for e := g.VList(i).Front();e != nil;e.Next(){
			k := e.Value.(int)
			inDegree[k]--
			if inDegree[k] == 0 {
				queue = append(queue,k)
			}
		}
	}


}

func TopologySortByDFS(g *nwgraph.Graph) {
	//构建逆邻接表，边s=>t，表示 s依赖于t,t先执行
	inverseG := nwgraph.NewGraph(g.VCnt())

	//通过邻接表生成逆邻接表
	for i := 0; i <g.VCnt();i++ {
		for e := g.VList(i).Front();e != nil; e.Next() {
			s := e.Value.(int)
			inverseG.AddEdge(s,i)
		}
	}


	visited := make([]bool,g.VCnt())

	//深度遍历图
	for i :=0;i<g.VCnt();i++ {
		if visited[i] == false{
			visited[i] = true
			dfs(i,inverseG,visited)
		}
	}

}

func dfs(vertex int, inverse *nwgraph.Graph, visited []bool) {

	for e :=inverse.VList(vertex).Front();e!=nil;e.Next() {
		s := e.Value.(int)
		if visited[s] == true {
			continue
		}
		visited[s] = true
		dfs(s,inverse,visited)
	}

	fmt.Printf("->%d",vertex)
}

