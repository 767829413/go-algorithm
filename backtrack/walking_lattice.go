package backtrack

var minDest = int(^uint(0) >> 1)

func minDistBT(i, j, dist, n int, w [][]int) {
	dist += w[i][j]
	if i == n && j == n {
		if minDest > dist {
			minDest = dist
		}
		return
	}
	if i < n {
		minDistBT(i+1, j, dist, n, w)
	}
	if j < n {
		minDistBT(i, j+1, dist, n, w)
	}

}

func GetMinDest(i, j, dist, n int, w [][]int) int {
	minDistBT(i, j, dist, n-1, w)
	return minDest
}
