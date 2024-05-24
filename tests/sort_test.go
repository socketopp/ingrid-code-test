package tests

import (
	"home/osrm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortRoutesByDurationAndDistance(t *testing.T) {
	// Mock input data
	resp := &osrm.RouteResponse{
		Source: "13.388860,52.517037",
		Routes: []osrm.RouteData{
			{Destination: "13.397634,52.529407", Distance: 1886.8, Duration: 260.3},
			{Destination: "13.428555,52.523219", Distance: 500.5, Duration: 100.0},
		},
	}

	// Call the sorting function
	osrm.SortRoutes(resp)

	if len(resp.Routes) > 1 {
		for i := 1; i < len(resp.Routes); i++ {
			assert.LessOrEqual(t, resp.Routes[i-1].Duration, resp.Routes[i].Duration, "Routes not sorted by duration")
			assert.LessOrEqual(t, resp.Routes[i-1].Distance, resp.Routes[i].Distance, "Routes not sorted by distnace")

		}
	}
}
