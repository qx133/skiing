package skiing

type tuple struct {
	r, c int //row and column representing the cell of a 2d array
}

//BestSkiPath returns the length and height of the ski path that is longest and has the largest drop
func BestSkiPath(input [][]int) (int, int) {
	rowLen := len(input)
	colLen := len(input[0])
	maxLen := 0
	maxHeight := 0
	length := make([][]int, rowLen)
	height := make([][]int, rowLen)
	q := []tuple{}           //queue to keep track of cells that needs to be updated
	qMap := map[tuple]bool{} //map to keep track of cells already in queue

	for i := range length { //initialize arrays
		length[i] = make([]int, colLen)
		height[i] = make([]int, colLen)
		copy(height[i], input[i])
	}

	for r, rows := range input {
		for c := range rows {
			if length[r][c] == 0 { //add to q if this cell has not been visited yet
				q = append(q, tuple{r, c})
				qMap[tuple{r, c}] = true
			}
			for len(q) > 0 {
				qFirst := q[0]
				q = q[1:]
				qMap[qFirst] = false
				changed := false //track if there changes to this cell, if yes, the neighbors are added to the q
				if length[qFirst.r][qFirst.c] == 0 {
					length[qFirst.r][qFirst.c] = 1
					changed = true
				}
				rcChanges := [][]int{
					{-1, 0},
					{0, -1},
					{1, 0},
					{0, 1},
				}
				for _, rows := range rcChanges { //check against neighbors for better path
					rNew := qFirst.r + rows[0]
					cNew := qFirst.c + rows[1]
					if validPair(rNew, cNew, input) && input[rNew][cNew] < input[qFirst.r][qFirst.c] {
						if length[rNew][cNew]+1 > length[qFirst.r][qFirst.c] {
							length[qFirst.r][qFirst.c] = length[rNew][cNew] + 1
							height[qFirst.r][qFirst.c] = height[rNew][cNew]
							changed = true
						} else if length[rNew][cNew]+1 == length[qFirst.r][qFirst.c] {
							height[qFirst.r][qFirst.c] = min(height[qFirst.r][qFirst.c], height[rNew][cNew])
							changed = true
						}
					}
				}
				if length[qFirst.r][qFirst.c] > maxLen { //longer path found
					maxLen = length[qFirst.r][qFirst.c]
					maxHeight = input[qFirst.r][qFirst.c] - height[qFirst.r][qFirst.c]
				} else if length[qFirst.r][qFirst.c] == maxLen { //higher path might be found
					maxHeight = max(maxHeight, input[qFirst.r][qFirst.c]-height[qFirst.r][qFirst.c])
				}
				if changed {
					for _, rows := range rcChanges { //changed cell means neighbors' values could change too
						rNew := qFirst.r + rows[0]
						cNew := qFirst.c + rows[1]
						if validPair(rNew, cNew, input) && input[rNew][cNew] > input[qFirst.r][qFirst.c] && !qMap[tuple{rNew, cNew}] {
							q = append(q, tuple{rNew, cNew})
							qMap[tuple{rNew, cNew}] = true
						}
					}
				}
			}
		}
	}
	return maxLen, maxHeight
}

func validPair(r, c int, input [][]int) bool {
	if r >= 0 && c >= 0 && r < len(input) && c < len(input[0]) {
		return true
	}
	return false
}

func max(a ...int) int {
	length := a[0]
	for _, v := range a {
		if v > length {
			length = v
		}
	}
	return length
}

func min(a ...int) int {
	length := a[0]
	for _, v := range a {
		if v < length {
			length = v
		}
	}
	return length
}
