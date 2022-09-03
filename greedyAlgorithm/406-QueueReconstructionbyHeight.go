package greedyAlgorithm

import (
	"container/list"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	//res := make([][]int, 0)
	//
	//for _, info := range people {
	//	res = append(res, info)
	//	copy(res[info[1]+1:], res[info[1]:])
	//	res[info[1]] = info
	//}
	//
	//return res
	l := list.New()
	for _, info := range people {
		idx := info[1]
		mark := l.PushBack(info)
		target := l.Front()
		for idx != 0 {
			idx--
			target = target.Next()
		}
		l.MoveBefore(mark, target)
	}

	res := make([][]int, 0)
	for e := l.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.([]int))
	}
	return res
}
