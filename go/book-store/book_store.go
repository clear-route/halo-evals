package bookstore

const basePrice = 800

var setPrices = map[int]int{
	1: basePrice,                                  // 1 * 800 * 1.00 = 800
	2: 2 * basePrice * 95 / 100,                   // 2 * 800 * 0.95 = 1520
	3: 3 * basePrice * 90 / 100,                   // 3 * 800 * 0.90 = 2160
	4: 4 * basePrice * 80 / 100,                   // 4 * 800 * 0.80 = 2560
	5: 5 * basePrice * 75 / 100,                   // 5 * 800 * 0.75 = 3000
}

func Cost(books []int) int {
	if len(books) == 0 {
		return 0
	}

	bookCounts := make(map[int]int)
	for _, bookID := range books {
		bookCounts[bookID]++
	}

	// numGroupsOfSize[s] stores how many discount groups of size 's' are formed.
	// Index 0 is unused, 1 for single books, 2 for 2-book sets etc. up to 5.
	numGroupsOfSize := make([]int, 6)

	for {
		distinctTypesInCurrentPass := 0
		// Iterate through book types 1 to 5 to find distinct available books for this pass
		var currentGroupBooks []int 
		for bookID := 1; bookID <= 5; bookID++ {
			if bookCounts[bookID] > 0 {
				distinctTypesInCurrentPass++
				currentGroupBooks = append(currentGroupBooks, bookID)
			}
		}

		if distinctTypesInCurrentPass == 0 {
			break // No more books to group
		}
        
        // Form a group of size `distinctTypesInCurrentPass`
		groupSize := distinctTypesInCurrentPass
		numGroupsOfSize[groupSize]++
        
        // Decrement counts for the books included in this group
		for _, bookIDInGroup := range currentGroupBooks {
            bookCounts[bookIDInGroup]--
        }
	}
    
	// Optimization: Two groups of 4 are cheaper than one group of 5 and one group of 3.
	// (5 books @ 25% off) + (3 books @ 10% off) = 3000 + 2160 = 5160
	// (4 books @ 20% off) * 2                 = 2560 * 2  = 5120
	// So, if we have groups of 5 and groups of 3, convert them to groups of 4 where possible.
	for numGroupsOfSize[5] > 0 && numGroupsOfSize[3] > 0 {
		numGroupsOfSize[5]--
		numGroupsOfSize[3]--
		numGroupsOfSize[4] += 2 // Add two groups of 4
	}

	totalCost := 0
	for size := 1; size <= 5; size++ {
		if price, ok := setPrices[size]; ok {
			totalCost += numGroupsOfSize[size] * price
		}
	}

	return totalCost
}
