package hierarchies

import (
	"fmt"
	"log/slog"
)

type AncestorIdsForLabelOptions struct {
	Hierarchies []map[string]int64
	Placetype   string
	ParentId    int64
}

func AncestorIdsForLabel(opts *AncestorIdsForLabelOptions) []int64 {

	name_ids := make([]int64, 0)

	switch opts.Placetype {
	case "continent", "empire", "country":
		// pass
	default:

		var lineage []string

		switch opts.Placetype {
		case "macroregion", "region":
			lineage = []string{
				"country",
			}
		case "macrocounty", "county":
			lineage = []string{
				"region",
				"country",
			}
		case "metroarea", "localadmin", "locality":

			lineage = []string{
				"region",
				"country",
			}

		case "borough", "campus", "postalcode":

			lineage = []string{
				"localadmin",
				"locality",
				"region",
				"country",
			}

		case "macrohood", "neighbourhood":

			lineage = []string{
				"localadmin",
				"locality",
				"region",
				"country",
			}

		case "microhood":

			lineage = []string{
				// "neighbourhood",
				"localadmin",
				"locality",
				"region",
				"country",
			}

		case "venue":

			lineage = []string{
				"neighbourhood",
				"localadmin",
				"locality",
				"region",
				"country",
			}

		case "building":

			lineage = []string{
				"campus",
				"locality",
				"region",
				"country",
			}

		case "wing":

			lineage = []string{
				"building",
				"campus",
				"locality",
				// "region",
				"country",
			}

		case "concourse":

			lineage = []string{
				"wing",
				"building",
				"campus",
				"locality",
				// "region",
				"country",
			}

		case "arcade":

			lineage = []string{
				"concourse",
				"wing",
				"building",
				"campus",
				// "locality",
				// "region",
				"country",
			}

		case "enclosure":

			lineage = []string{
				"arcade",
				"concourse",
				"wing",
				// "building",
				"campus",
				// "locality",
				// "region",
				"country",
			}

		default:
			slog.Warn("Unsupported placetype", "placetype", opts.Placetype)
		}

		var use_hier map[string]int64

		switch len(opts.Hierarchies) {
		case 0:
			// pass - len(hiers) checked above
		case 1:
			use_hier = opts.Hierarchies[0]
		default:

			switch opts.ParentId >= 0 {
			case true:

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

			default:
				use_hier = MergeCommonHierarchies(opts.Hierarchies)
			}
		}

		if use_hier == nil {
			slog.Warn("Failed to determine principal (or merged) hierarchy, simply pulling the first one (which may be incorrect)")
			use_hier = opts.Hierarchies[0]
		}

		for _, pt := range lineage {

			hier_k := fmt.Sprintf("%s_id", pt)
			name_id, exists := use_hier[hier_k]

			if exists {
				name_ids = append(name_ids, name_id)
			}
		}
	}

	return name_ids
}
