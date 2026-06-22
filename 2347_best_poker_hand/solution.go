package bestpokerhand

func bestHand(ranks []int, suits []byte) string {
	// check for flush
	for i := 1; i < len(suits); i++ {
		if suits[i] != suits[0] {
			break
		}

		if i == len(suits)-1 {
			return "Flush"
		}
	}

	// count ranks
	rankCounts := make(map[int]int)
	for _, rank := range ranks {
		rankCounts[rank]++
	}

	hasPair := false
	for _, count := range rankCounts {
		if count == 2 {
			hasPair = true
		}

		if count >= 3 {
			return "Three of a Kind"
		}
	}

	if hasPair {
		return "Pair"
	}

	return "High Card"
}
