package hashing

import "sort"

type numFreq struct {
	Num  int
	Freq int
}

type numFreList []numFreq

var _ sort.Interface = &numFreList{}

func (n numFreList) Len() int           { return len(n) }
func (n numFreList) Less(i, j int) bool { return n[i].Freq > n[j].Freq }
func (n numFreList) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

// TopKFrequentElementsSol1 returns the top k frequent elements in nums, https://leetcode.com/problems/top-k-frequent-elements/solutions/
func TopKFrequentElementsSol1(nums []int, k int) []int {
	// count the frequency of each number
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// create a slice of numFreq
	numFreqs := make([]numFreq, 0)
	for num, freq := range freqMap {
		numFreqs = append(numFreqs, numFreq{num, freq})
	}

	// sort the slice by frequency
	sort.Sort(numFreList(numFreqs)) // can use sort.Slice(numFreqs, func(i, j int) bool { return numFreqs[i].Freq > numFreqs[j].Freq })

	// get the top k frequent elements
	output := make([]int, 0)
	for _, freq := range numFreqs[:k] {
		output = append(output, freq.Num)
	}

	return output
}

func TopKFrequentElementsSol2(nums []int, k int) []int {
	// count the frequency of each number
	freqMap := make(map[int]int)
	maxOfFreq := -1
	for _, num := range nums {
		freqMap[num]++
		if val, ok := freqMap[num]; ok && val > maxOfFreq {
			maxOfFreq = val
		}
	}

	// create a slice of num frequencies
	numsFreqs := make([][]int, maxOfFreq+1)
	for num, freq := range freqMap {
		// index represents the frequency of the numbers stored in its values
		numsFreqs[freq] = append(numsFreqs[freq], num)
	}

	// get the top k frequent elements
	totalNums := len(numsFreqs)
	output := make([]int, 0)
	for i := totalNums - 1; i > 0; i-- {
		if len(numsFreqs[i]) > 0 {
			output = append(output, numsFreqs[i]...)
			k = k - len(numsFreqs[i])
		}
		if k == 0 {
			break
		}
	}
	return output
}
