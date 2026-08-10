package hierarchies

// Merge two or hierarchies in to a single hierarchy containing only common keys (placetypes) AND ids.
func MergeCommonHierarchies(hiers []map[string]int64) map[string]int64 {

	result := make(map[string]int64)

	switch len(hiers) {
	case 0:
		return result
	case 1:
		return hiers[0]
	default:
		// do the work below
	}

	baseline := hiers[0]

	for key, baseline_v := range baseline {

		match := true

		for i := 1; i < len(hiers); i++ {

			current_v, exists := hiers[i][key]

			// If key is missing or value doesn't match, invalidate it
			if !exists || current_v != baseline_v {
				match = false
				break
			}
		}

		if match {
			result[key] = baseline_v
		}
	}

	return result
}
