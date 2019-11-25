package main

import (
	"fmt"
	"math"
)

func main() {
	//测试数据
	var (
		graph   = make(map[string]interface{})
		costs   = make(map[string]int64)
		parents = make(map[string]string)
		presson []string
	)
	//构建路径花费
	start := make(map[string]int64)
	start["a"] = 6
	start["b"] = 2
	graph["start"] = start
	a := make(map[string]int64)
	a["fin"] = 1
	graph["a"] = a
	b := make(map[string]int64)
	b["a"] = 3
	b["fin"] = 5
	graph["b"] = b

	//节点开销
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = math.MaxInt64

	//父节点存储
	presson = []string{}
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	//找开销最小节点
	node := findLowestCostNode(costs, presson)
	for node != "" {
		cost := costs[node]
		if graph[node] == nil {
			break
		}
		neighbors := graph[node].(map[string]int64)
		for negNode, negCost := range neighbors {
			newCost := cost + negCost
			if costs[negNode] > newCost {
				parents[negNode] = node
				costs[negNode] = newCost
				presson = append(presson, node)
			}
		}
		node = findLowestCostNode(costs, presson)
	}
	//输出结果
	fmt.Println(parents)
	fmt.Println(costs)
}

func findLowestCostNode(costs map[string]int64, presson []string) string {
	var (
		lowCost = int64(math.MaxInt64)
		lowNode string
	)
	for node, cost := range costs {
		if cost < lowCost && !in(node, presson) {
			lowCost = cost
			lowNode = node
		}
	}
	return lowNode
}

func in(value string, presson []string) bool {
	for _, val := range presson {
		if value == val {
			return true
		}
	}
	return false
}
