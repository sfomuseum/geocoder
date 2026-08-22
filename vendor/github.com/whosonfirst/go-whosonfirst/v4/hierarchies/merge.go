package hierarchies

// MergeCommonHierarchies merges two or more hierarchies into a single
// hierarchy containing only the common keys (placetypes) and identical
// IDs.  The function delegates to the generic implementation
// MergeCommonHierarchiesGeneric[int64] for the actual work.
//
// This wrapper exists for backward compatibility and to keep the
// original non‑generic signature available.
func MergeCommonHierarchies(hiers []map[string]int64) map[string]int64 {
	return MergeCommonHierarchiesGeneric[int64](hiers)
}

// MergeCommonHierarchiesGeneric merges multiple hierarchies into a single
// hierarchy containing only the keys that appear in every map and have the
// same value.  It is generic over the ID type T, which must satisfy the
// Id constraint (int64 or string).
func MergeCommonHierarchiesGeneric[T Id](hierarchies []map[string]T) map[string]T {
	
	if len(hierarchies) == 0 {
		return nil
	}

	// Keep only the keys that appear in *every* map and that have the same value.
	
	merged := make(map[string]T)
	
	for k, v := range hierarchies[0] {
		
		common := true
		
		for _, h := range hierarchies[1:] {
			if val, ok := h[k]; !ok || val != v {
				common = false
				break
			}
		}
		
		if common {
			merged[k] = v
		}
	}
	
	return merged
}
