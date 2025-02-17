package slidingwindow

/*
HackerLand Sports Club wants to send a team for a relay race. There are n racers in the group indexed from 0 to n. 1. The th racer has a speed of speed[i] units.

The coach decided to send some contiguous subsegments of racers for the race i.e. racers with index i, i+1, +2...,
such that each racer has the same speed in the group to ensure smooth baton transfer.
To achieve the goal, the coach decided to remove some racers from the group such that the number of racers with the same speed in some contiguous segment is maximum.

Given the array, racers, and an integer k, find the maximum possible number of racers in some contiguous segment of racers with the same speed after at most kracers are removed.


Example Suppose n = 6, k=2, and speed = [1, 4, 4, 2, 2, 4].


It is optimal to remove the two racers with speed 2 to get the racers [1, 4, 4, 4]. Each racer with speed 4 can now be sent as they are in a contiguous segment. A maximum of 3 racers can be sent for the relay race.

Can anyone solve it?

// Maximum subarray lenght with same value
*/

func maxRacers(speed []int, k int) int {
	freqMap := make(map[int]int)
	var maxFreqSoFar, left, maxCount int

	for right := 0; right < len(speed); right++ {
		freqMap[speed[right]]++

		maxFreqSoFar = max(maxFreqSoFar, freqMap[speed[right]])

		windowSize := right - left + 1
		if windowSize-maxFreqSoFar > k {
			freqMap[speed[left]]--
			left++
		}

		// store the max freq of most occuring speed which can exist in the window with at most k removal
		maxCount = max(maxCount, maxFreqSoFar)
	}

	return maxCount
}
