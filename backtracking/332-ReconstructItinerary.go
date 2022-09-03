package backtracking

import "sort"

type pair struct {
	target string
	visited bool
}

type pairs []*pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].target < p[j].target
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func findItinerary(tickets [][]string) []string {
	dict := make(map[string]pairs, 0)
	for _, ticket := range tickets {
		if _, ok := dict[ticket[0]]; ok {
			dict[ticket[0]] = append(dict[ticket[0]], &pair{target: ticket[1], visited: false})
		} else {
			dict[ticket[0]] = make(pairs, 0)
			dict[ticket[0]] = append(dict[ticket[0]], &pair{target: ticket[1], visited: false})
		}
	}

	for k, _ := range dict {
		sort.Sort(dict[k])
	}

	var res []string = []string{"JFK"}
	var backtracking func() bool
	backtracking = func() bool {
		if len(tickets) + 1 == len(res) {
			return true
		}

		for _, pair := range dict[res[len(res)-1]] {
			if pair.visited == false {
				res = append(res, pair.target)
				pair.visited = true
				if backtracking() {
					return true
				}
				res = res[:len(res)-1]
				pair.visited = false
			}
		}
		return false
	}

	backtracking()

	return res
}



