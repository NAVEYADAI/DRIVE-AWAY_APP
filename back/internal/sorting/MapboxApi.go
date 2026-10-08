package sorting

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func mapboxToken() string {
	return os.Getenv("MAPBOX_ACCESS_TOKEN")
}

// mapsHTTPClient is swappable in tests so FindWay can be exercised without
// network access.
var mapsHTTPClient = &http.Client{Timeout: 10 * time.Second}

// mapboxBaseURL is swappable in tests (httptest server).
var mapboxBaseURL = "https://api.mapbox.com"

// Process-lifetime caches. Sort() asks for the same (driver, drive) pair many
// times (every backtrack re-grades), and each Mapbox call is billed.
// Failures are never cached.
var (
	cacheMu       sync.Mutex
	geocodeCache  = map[string]coord{}
	durationCache = map[string]time.Duration{}
)

type coord struct{ lng, lat float64 }

func (c coord) String() string {
	return strconv.FormatFloat(c.lng, 'f', -1, 64) + "," + strconv.FormatFloat(c.lat, 'f', -1, 64)
}

// FindWay returns the driving time (with typical traffic) between two places.
// Each place is either "lat,lng" or a free-text address. Returns 0 on any
// failure (see specs/03-sorting.md §4: a failed lookup reads as "0 minutes").
func FindWay(origin, destination string) time.Duration {
	key := origin + "\x00" + destination
	cacheMu.Lock()
	if d, ok := durationCache[key]; ok {
		cacheMu.Unlock()
		return d
	}
	cacheMu.Unlock()

	from, ok := resolveCoord(origin)
	if !ok {
		return 0
	}
	to, ok := resolveCoord(destination)
	if !ok {
		return 0
	}

	d, ok := routeDuration(from, to)
	if !ok {
		return 0
	}
	cacheMu.Lock()
	durationCache[key] = d
	cacheMu.Unlock()
	return d
}

// parseLatLng accepts "lat,lng" (the format the UI uses for Gps fields).
func parseLatLng(s string) (coord, bool) {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return coord{}, false
	}
	lat, err1 := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	lng, err2 := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err1 != nil || err2 != nil || lat < -90 || lat > 90 || lng < -180 || lng > 180 {
		return coord{}, false
	}
	return coord{lng: lng, lat: lat}, true
}

func resolveCoord(place string) (coord, bool) {
	place = strings.TrimSpace(place)
	if place == "" {
		return coord{}, false
	}
	if c, ok := parseLatLng(place); ok {
		return c, true
	}
	cacheMu.Lock()
	if c, ok := geocodeCache[place]; ok {
		cacheMu.Unlock()
		return c, true
	}
	cacheMu.Unlock()

	body, ok := mapboxGet(fmt.Sprintf("%s/search/geocode/v6/forward?q=%s&country=il&language=he&limit=1",
		mapboxBaseURL, url.QueryEscape(place)))
	if !ok {
		return coord{}, false
	}
	c, ok := parseGeocodeResponse(body)
	if !ok {
		return coord{}, false
	}
	cacheMu.Lock()
	geocodeCache[place] = c
	cacheMu.Unlock()
	return c, true
}

// routeDuration uses the Directions API (one route, one billed request).
// The Matrix API rejects a single from/to pair (it needs >= 2 matrix elements).
func routeDuration(from, to coord) (time.Duration, bool) {
	body, ok := mapboxGet(fmt.Sprintf("%s/directions/v5/mapbox/driving-traffic/%s;%s?overview=false",
		mapboxBaseURL, from, to))
	if !ok {
		return 0, false
	}
	return parseDirectionsResponse(body)
}

// mapboxGet appends the access token and returns the body of a 200 response.
func mapboxGet(rawURL string) ([]byte, bool) {
	token := mapboxToken()
	if token == "" {
		return nil, false
	}
	resp, err := mapsHTTPClient.Get(rawURL + "&access_token=" + url.QueryEscape(token))
	if err != nil {
		return nil, false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, false
	}
	return body, true
}

// parseGeocodeResponse extracts features[0].geometry.coordinates ([lng, lat]).
func parseGeocodeResponse(body []byte) (coord, bool) {
	var r struct {
		Features []struct {
			Geometry struct {
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
		} `json:"features"`
	}
	if err := json.Unmarshal(body, &r); err != nil || len(r.Features) == 0 {
		return coord{}, false
	}
	cs := r.Features[0].Geometry.Coordinates
	if len(cs) < 2 {
		return coord{}, false
	}
	return coord{lng: cs[0], lat: cs[1]}, true
}

// parseDirectionsResponse extracts routes[0].duration (seconds).
func parseDirectionsResponse(body []byte) (time.Duration, bool) {
	var r struct {
		Code   string `json:"code"`
		Routes []struct {
			Duration float64 `json:"duration"`
		} `json:"routes"`
	}
	if err := json.Unmarshal(body, &r); err != nil || r.Code != "Ok" || len(r.Routes) == 0 {
		return 0, false
	}
	return time.Duration(r.Routes[0].Duration * float64(time.Second)), true
}
