package osrm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
)

// RouteResponse represents the response structure for route details.
type RouteResponse struct {
	Source string      `json:"source"`
	Routes []RouteData `json:"routes"`
}

// RouteData holds the information about a single route's distance and duration.
type RouteData struct {
	Destination string  `json:"destination"`
	Distance    float64 `json:"distance"`
	Duration    float64 `json:"duration"`
}

// OsrmResponse represents the structure of the response from the OSRM API.
type OsrmResponse struct {
	Code   string `json:"code"`
	Routes []struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
	} `json:"routes"`
}

// getRouteFromOSRM makes a request to the OSRM API to get route details from src to dst.
func getRouteFromOSRM(src, dst string) (*OsrmResponse, error) {
	url := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%s;%s?overview=false", src, dst)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OSRM request failed with status: %s", resp.Status)
	}

	var osrmResponse OsrmResponse
	if err := json.NewDecoder(resp.Body).Decode(&osrmResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return &osrmResponse, nil
}

// SortRoutes sorts the routes in a RouteResponse by duration and distance.
// Routes are sorted first by duration in ascending order. If two routes have
// the same duration, they are then sorted by distance in ascending order.
func SortRoutes(routeResponse *RouteResponse) {
	sort.Slice(routeResponse.Routes, func(i, j int) bool {
		if routeResponse.Routes[i].Duration == routeResponse.Routes[j].Duration {
			return routeResponse.Routes[i].Distance < routeResponse.Routes[j].Distance
		}
		return routeResponse.Routes[i].Duration < routeResponse.Routes[j].Duration
	})
}

// GetDurationAndDistances retrieves the duration and distance for each destination from the source.
func GetDurationAndDistances(src string, dst []string) (*RouteResponse, error) {
	var wg sync.WaitGroup
	results := make(chan RouteData, len(dst))
	errors := make(chan error, len(dst))

	for _, destination := range dst {
		wg.Add(1)
		go func(destination string) {
			defer wg.Done()
			resp, err := getRouteFromOSRM(src, destination)
			if err != nil {
				errors <- err
				return
			}
			// Send RouteData to result channel
			results <- RouteData{Destination: destination, Distance: resp.Routes[0].Distance, Duration: resp.Routes[0].Duration}
		}(destination)
	}

	// wait for each go routine to execute until wg (waitgroup) counter reaches 0.
	wg.Wait()
	close(results)
	close(errors)

	routeResponse := &RouteResponse{
		Source: src,
		Routes: []RouteData{},
	}

	// Create response object by iterating the channel
	for res := range results {
		routeResponse.Routes = append(routeResponse.Routes, res)
	}

	if len(errors) > 0 {
		var errStrings []string
		for err := range errors {
			errStrings = append(errStrings, err.Error())
		}
		return routeResponse, fmt.Errorf("errors occurred: %v", errStrings)
	}

	// sort the routes
	SortRoutes(routeResponse)

	return routeResponse, nil
}
