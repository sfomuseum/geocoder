package hierarchies

import (
	"fmt"
	"log/slog"
)

// AncestorIdsForLabelOptions holds the options for retrieving ancestor IDs.
// Hierarchies is a slice of maps that represent the hierarchical data,
// where each map key is of the form "<placetype>_id" (e.g. "country_id") and
// the value is the corresponding ID.  Placetype is the place type for
// which ancestors are being requested, and ParentId is the ID of the
// current place.
type AncestorIdsForLabelOptions struct {
	Hierarchies []map[string]int64
	Placetype   string
	ParentId    int64
}

// Id may be an int64 or a string – both are acceptable for the map values.
type Id interface{ ~int64 | ~string }

// AncestorIdsForLabelOptionsGeneric is the generic version of
// AncestorIdsForLabelOptions.  The type parameter T must satisfy the Id
// constraint and is used consistently for the map values and the
// ParentId field.
type AncestorIdsForLabelOptionsGeneric[T Id] struct {
	Hierarchies []map[string]T // one or more hierarchy maps
	Placetype   string
	ParentId    T // must be the same type that is used in the hierarchy maps
}

// AncestorIdsForLabel retrieves the ancestor IDs for the given options
// and returns them as a slice of int64.  Internally it creates a
// generic options struct and delegates to AncestorIdsForLabelGeneric.
func AncestorIdsForLabel(opts *AncestorIdsForLabelOptions) []int64 {

	genericOpts := &AncestorIdsForLabelOptionsGeneric[int64]{
		Hierarchies: opts.Hierarchies,
		Placetype:   opts.Placetype,
		ParentId:    opts.ParentId,
	}

	return AncestorIdsForLabelGeneric[int64](genericOpts)
}

// AncestorIdsForLabelGeneric returns the ancestor IDs for the given
// options, where the ID type is specified by the type parameter T.
// The function determines the lineage for the requested placetype,
// selects the appropriate hierarchy map, and collects the ancestor
// IDs in order from the nearest ancestor to the farthest.
func AncestorIdsForLabelGeneric[T Id](opts *AncestorIdsForLabelOptionsGeneric[T]) []T {
	
	name_ids := make([]T, 0)

	// “continent”, “empire” and “country” are leaf nodes – they have no
	// ancestors to return, so we simply return an empty list
	
	switch opts.Placetype {
	case "continent", "empire", "country":
		// nothing to do
	default:

		var lineage []string
		
		switch opts.Placetype {
		case "macroregion", "region":
			lineage = []string{"country"}
		case "macrocounty", "county":
			lineage = []string{"region", "country"}
		case "metroarea", "localadmin", "locality":
			lineage = []string{"region", "country"}
		case "borough", "campus", "postalcode":
			lineage = []string{"localadmin", "locality", "region", "country"}
		case "macrohood", "neighbourhood":
			lineage = []string{"localadmin", "locality", "region", "country"}
		case "microhood":
			lineage = []string{"localadmin", "locality", "region", "country"}
		case "venue":
			lineage = []string{"neighbourhood", "localadmin", "locality", "region", "country"}
		case "building":
			lineage = []string{"campus", "locality", "region", "country"}
		case "wing":
			lineage = []string{"building", "campus", "locality", "country"}
		case "concourse":
			lineage = []string{"wing", "building", "campus", "locality", "country"}
		case "arcade":
			lineage = []string{"concourse", "wing", "building", "campus", "country"}
		case "enclosure":
			lineage = []string{"arcade", "concourse", "wing", "campus", "country"}
		case "marinearea":
			lineage = []string{ "country" }
		default:
			slog.Debug("Unsupported placetype", "placetype", opts.Placetype)
		}

		var use_hier map[string]T

		switch len(opts.Hierarchies) {
		case 0:
			// pass
		case 1:
			use_hier = opts.Hierarchies[0]
		default:

			for _, h := range opts.Hierarchies {
				
				for _, hid := range h {
					if hid == opts.ParentId {
						use_hier = h
						break
					}
				}
				
				if use_hier != nil {
					break
				}
			}

			if use_hier == nil {
				// No hierarchy contains ParentId – fall back to the merged view.
				use_hier = MergeCommonHierarchiesGeneric(opts.Hierarchies)
			}
		}

		if use_hier == nil {
			
			slog.Warn("Failed to determine principal (or merged) hierarchy, falling back to the first one")
			
			if len(opts.Hierarchies) > 0 {
				use_hier = opts.Hierarchies[0]
			}
		}

		for _, pt := range lineage {
			
			k := fmt.Sprintf("%s_id", pt)
			
			id, exists := use_hier[k]

			if exists {
				name_ids = append(name_ids, id)
			}
		}
	}

	return name_ids
}
