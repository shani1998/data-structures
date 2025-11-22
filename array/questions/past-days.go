package questions

import "fmt"

func getFutureLake(rains []int, pastDaysMp map[int]int, index int) int {
	n := len(rains)
	for i := index + 1; i < n; i++ {
		if val, ok := pastDaysMp[rains[i]]; ok && val > 0 && rains[i] > 0 {
			return rains[i]
		}
	}
	return -1
}

func solve(rains []int) []int {
	prefixArr, n := make([]int, 0), len(rains)
	pastDaysMp := make(map[int]int)
	ans := make([]int, n)
	dryDays := []int{}

	for i := 0; i < n; i++ {
		if rains[i] > 0 {
			prefixArr = append(prefixArr, rains[i])
			if pastDaysMp[rains[i]] > 0 {
				if len(dryDays) == 0 {
					return []int{}
				}
				d := dryDays[0]
				dryDays = dryDays[1:]
				ans[d] = rains[i]
			}

			pastDaysMp[rains[i]] = 1
			ans[i] = -1
			continue
		}

		ans[i] = 1
		dryDays = append(dryDays, i)
	}
	return ans
}

func solve(rains []int) []int { // 2,1,3,0,1,9
	prefixArr, n := make([]int, 0), len(rains)
	pastDaysMp := make(map[int]int)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if rains[i] > 0 {
			prefixArr = append(prefixArr, rains[i])
			pastDaysMp[rains[i]]++
			ans[i] = -1
			continue
		}
		key := getFutureLake(rains, pastDaysMp, i)
		if key != -1 {
			pastDaysMp[rains[i]]--
			ans[i] = key
			continue
		}

		// not found in future
		for _, val := range pastDaysMp {
			if val > 0 {
				ans[i] = val
				pastDaysMp[rains[i]]--
			}
		}

	}

	return ans
}

func main() {
	fmt.Println(solve([]int{2, 1, 3, 0, 1, 9}))
	fmt.Println(solve([]int{2, 1, 0, 0, 9}))
	fmt.Println(solve([]int{69, 0, 0, 0, 69}))

}
