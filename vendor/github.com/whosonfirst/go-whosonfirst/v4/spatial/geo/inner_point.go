package geo

import (
	"math"

	"github.com/paulmach/orb"
)

// Note: This code was written with the assistance of robots. It has been reviewed and seems fine.
// At least until it is not.

// FindInnerPoint implements Mapshaper's inner point algorithm.
// It tries a horizontal centerline slice first; if that fails or falls into a
// bottleneck, it falls back to the true polygon centroid.
func FindInnerPoint(geom orb.Geometry) (orb.Point, bool) {

	if geom == nil {
		return orb.Point{}, false
	}

	switch g := geom.(type) {
	case orb.Polygon:

		cloned := g.Clone()
		unwrapPolygon(cloned)
		return findMapshaperStylePoint(cloned), true

	case orb.MultiPolygon:

		cloned := g.Clone()
		unwrapMultiPolygon(cloned)

		// Isolate the dominant landmass (ignores small islands like the Farallons or Arctic rocks)
		largestPoly := findLargestPolygon(cloned)

		if len(largestPoly) == 0 {
			return orb.Point{}, false
		}

		return wrapPoint(findMapshaperStylePoint(largestPoly)), true

	default:
		bound := geom.Bound()
		return bound.Center(), !bound.IsEmpty()
	}
}

// findMapshaperStylePoint mirrors Mapshaper's decision matrix for inner anchor generation.
func findMapshaperStylePoint(poly orb.Polygon) orb.Point {

	bound := poly.Bound()

	// Fallback 1: Calculate the true planar centroid of the polygon
	centroid := calculateCentroid(poly)

	// Mapshaper check: If the true centroid is safely inside the polygon shell,
	// use it. This correctly places the point in Downtown SF and Central Russia.
	if pointInPolygon(centroid, poly) {
		return centroid
	}

	// Fallback 2: Mapshaper's scanline fallback
	// Calculate the strict horizontal midpoint of the bounding box
	targetY := bound.Min.Y() + (bound.Max.Y()-bound.Min.Y())/2

	var xIntersections []float64

	for _, ring := range poly {
		xIntersections = append(xIntersections, intersectHorizontalRing(ring, targetY)...)
	}

	// Sort intersections left-to-right
	for i := 0; i < len(xIntersections); i++ {
		for j := i + 1; j < len(xIntersections); j++ {
			if xIntersections[i] > xIntersections[j] {
				xIntersections[i], xIntersections[j] = xIntersections[j], xIntersections[i]
			}
		}
	}

	// Track the longest valid internal horizontal segment
	maxLen := -1.0
	bestX := bound.Center().X()
	inside := false

	for i := 0; i < len(xIntersections)-1; i++ {
		inside = !inside
		if inside {
			x1 := xIntersections[i]
			x2 := xIntersections[i+1]
			length := x2 - x1
			if length > maxLen {
				maxLen = length
				bestX = x1 + length/2
			}
		}
	}

	// If a valid scanline intersection was found, use its midpoint
	if maxLen > 0 {
		return orb.Point{bestX, targetY}
	}

	// Absolute structural fallback
	return bound.Center()
}

// calculateCentroid computes the geographic center of mass using the shoelace method.
func calculateCentroid(poly orb.Polygon) orb.Point {
	if len(poly) == 0 || len(poly[0]) < 3 {
		return orb.Point{0, 0}
	}

	shell := poly[0]
	n := len(shell)

	cx, cy := 0.0, 0.0
	area := 0.0

	for i := range n {
		j := (i + 1) % n
		factor := (shell[i].X() * shell[j].Y()) - (shell[j].X() * shell[i].Y())
		area += factor
		cx += (shell[i].X() + shell[j].X()) * factor
		cy += (shell[i].Y() + shell[j].Y()) * factor
	}

	area /= 2.0
	if area == 0 {
		return shell[0]
	}

	cx /= (6.0 * area)
	cy /= (6.0 * area)

	return orb.Point{cx, cy}
}

// pointInPolygon ensures a point resides inside the exterior shell and outside any holes.
func pointInPolygon(p orb.Point, poly orb.Polygon) bool {
	if !pointInRing(p, poly[0]) {
		return false
	}
	for i := 1; i < len(poly); i++ {
		if pointInRing(p, poly[i]) {
			return false // Inside a cutout hole
		}
	}
	return true
}

func pointInRing(p orb.Point, ring orb.Ring) bool {
	inside := false
	n := len(ring)
	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		if ((ring[i].Y() > p.Y()) != (ring[j].Y() > p.Y())) &&
			(p.X() < (ring[j].X()-ring[i].X())*(p.Y()-ring[i].Y())/(ring[j].Y()-ring[i].Y())+ring[i].X()) {
			inside = !inside
		}
	}
	return inside
}

func intersectHorizontalRing(ring orb.Ring, y float64) []float64 {
	var intersections []float64
	n := len(ring)
	if n < 3 {
		return intersections
	}
	for i := range n {
		p1 := ring[i]
		p2 := ring[(i+1)%n]
		if p1.Y() == p2.Y() {
			continue
		}
		if (p1.Y() <= y && p2.Y() > y) || (p2.Y() <= y && p1.Y() > y) {
			t := (y - p1.Y()) / (p2.Y() - p1.Y())
			x := p1.X() + t*(p2.X()-p1.X())
			intersections = append(intersections, x)
		}
	}
	return intersections
}

func unwrapMultiPolygon(mp orb.MultiPolygon) {
	if len(mp) == 0 || len(mp[0]) == 0 || len(mp[0][0]) == 0 {
		return
	}
	refLng := mp[0][0][0].X()
	for _, poly := range mp {
		for _, ring := range poly {
			for i := range ring {
				ring[i] = orb.Point{unwrapLongitude(ring[i].X(), refLng), ring[i].Y()}
			}
		}
	}
}

func unwrapPolygon(poly orb.Polygon) {
	if len(poly) == 0 || len(poly[0]) == 0 {
		return
	}
	refLng := poly[0][0].X()
	for _, ring := range poly {
		for i := range ring {
			ring[i] = orb.Point{unwrapLongitude(ring[i].X(), refLng), ring[i].Y()}
		}
	}
}

func unwrapLongitude(lng, refLng float64) float64 {
	diff := lng - refLng
	if diff < -180 {
		return lng + 360
	} else if diff > 180 {
		return lng - 360
	}
	return lng
}

func wrapPoint(pt orb.Point) orb.Point {
	lng := pt.X()
	for lng > 180 {
		lng -= 360
	}
	for lng < -180 {
		lng += 360
	}
	return orb.Point{lng, pt.Y()}
}

func findLargestPolygon(mp orb.MultiPolygon) orb.Polygon {
	var largestPoly orb.Polygon
	maxArea := -1.0
	for _, poly := range mp {
		area := calculatePolygonArea(poly)
		if area > maxArea {
			maxArea = area
			largestPoly = poly
		}
	}
	return largestPoly
}

func calculatePolygonArea(poly orb.Polygon) float64 {
	if len(poly) == 0 {
		return 0
	}
	area := math.Abs(calculateRingArea(poly[0]))
	for i := 1; i < len(poly); i++ {
		area -= math.Abs(calculateRingArea(poly[i]))
	}
	return area
}

func calculateRingArea(ring orb.Ring) float64 {
	n := len(ring)
	if n < 3 {
		return 0
	}
	area := 0.0
	for i := range n {
		j := (i + 1) % n
		area += ring[i].X() * ring[j].Y()
		area -= ring[j].X() * ring[i].Y()
	}
	return area / 2.0
}
