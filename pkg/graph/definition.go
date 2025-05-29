package graph

import (
	"fmt"
)

// Edge
//
//	@Description: 定义边
type Edge struct {
	From, To string
	Weight   float64
}

func (e Edge) ToString() string {
	return fmt.Sprintf("from: %s , to: %s ;weight : %g", e.From, e.To, e.Weight)
}

// Graph
//
//	@Description: 定义图的接口
type Graph interface {
	//
	// IsDirect
	//  @Description: 是否为有向图
	//  @return int
	//
	IsDirect() bool

	//
	// EdgesNum
	//  @Description: 边的个数
	//  @return int
	//
	EdgesNum() int

	//
	// VerticesNum
	//  @Description: 顶点的个数
	//  @return int
	//
	VerticesNum() int

	//
	// Vertices
	//  @Description: 获取所有的顶点
	//  @return Set
	//
	Vertexes() Set

	//
	// Contains
	//  @Description: 是否包含名称为v的顶点
	//  @param string
	//  @return bool
	//
	Contains(v string) bool

	// Connect
	//  @Description: 两点直接创建连接
	//  @param from
	//  @param to
	//  @param weight
	//
	Connect(from, to string, weight float64)

	//
	// Adj
	//  @Description: 返回一个点的所有邻边
	//  @param v
	//  @return map[string]float64
	//
	Adj(v string) map[string]float64

	//
	// HasEdge
	//  @Description: 判断两个点是否有边
	//  @param from
	//  @param to
	//  @return bool
	//
	HasEdge(from, to string) bool

	//
	// GetEdgeWeight
	//  @Description: 获取边的权重
	//  @param from
	//  @param string
	//  @return float64 如果存在，返回权重值
	//  @return bool 是否存在边
	//
	GetEdgeWeight(from, to string) (float64, bool)

	//
	// Show
	//  @Description: 打印图
	//
	Show()
}

// Set 邻边迭代器使用
type Set map[string]bool

func CreateSet() Set {
	var s Set
	s = make(map[string]bool)
	return s
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) Contains(e string) bool {
	v := s[e]
	return v
}

func (s Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s Set) Add(e string) {
	s[e] = true
}

func (s Set) Delete(e string) {
	delete(s, e)
}

// GetRandomOne
//
//	@Description: 随机拿一个
//	@receiver s
func (s Set) GetRandomOne() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	var v string
	for tmpV := range s {
		v = tmpV
		break
	}
	return v, true
}
