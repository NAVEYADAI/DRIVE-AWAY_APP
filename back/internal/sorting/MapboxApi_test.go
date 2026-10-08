package sorting

// White-box tests (package sorting): they swap the unexported mapboxBaseURL /
// caches so FindWay runs against an httptest server, never the real Mapbox.

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func resetMapboxState(t *testing.T, serverURL string) {
	t.Helper()
	t.Setenv("MAPBOX_ACCESS_TOKEN", "test-token")
	origURL := mapboxBaseURL
	mapboxBaseURL = serverURL
	cacheMu.Lock()
	geocodeCache = map[string]coord{}
	durationCache = map[string]time.Duration{}
	cacheMu.Unlock()
	t.Cleanup(func() { mapboxBaseURL = origURL })
}

func TestFindWay_LatLngSkipsGeocodingAndParsesSeconds(t *testing.T) {
	var geocodeCalls, routeCalls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "geocode") {
			atomic.AddInt32(&geocodeCalls, 1)
		}
		if strings.Contains(r.URL.Path, "/directions/v5/") {
			atomic.AddInt32(&routeCalls, 1)
			// lng,lat order: Mapbox expects longitude first.
			assert.Contains(t, r.URL.Path, "34.8,32;34.9,32.1")
		}
		w.Write([]byte(`{"code":"Ok","routes":[{"duration":900.4}]}`))
	}))
	defer srv.Close()
	resetMapboxState(t, srv.URL)

	got := FindWay("32.0,34.8", "32.1,34.9")

	assert.Equal(t, 900*time.Second+400*time.Millisecond, got)
	assert.EqualValues(t, 0, geocodeCalls)
	assert.EqualValues(t, 1, routeCalls)
}

func TestFindWay_GeocodesTextAddressesAndCachesResults(t *testing.T) {
	var geocodeCalls, routeCalls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "geocode") {
			atomic.AddInt32(&geocodeCalls, 1)
			w.Write([]byte(`{"features":[{"geometry":{"coordinates":[34.78,32.08]}}]}`))
			return
		}
		atomic.AddInt32(&routeCalls, 1)
		w.Write([]byte(`{"code":"Ok","routes":[{"duration":600}]}`))
	}))
	defer srv.Close()
	resetMapboxState(t, srv.URL)

	// Sort() re-grades the same pair on every backtrack; each call is billed,
	// so repeats must be served from the cache.
	for i := 0; i < 3; i++ {
		assert.Equal(t, 10*time.Minute, FindWay("דיזנגוף 1 תל אביב", "32.1,34.9"))
	}

	assert.EqualValues(t, 1, geocodeCalls)
	assert.EqualValues(t, 1, routeCalls)
}

func TestFindWay_FailuresReturnZeroAndAreNotCached(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(`{"code":"Ok","routes":[{"duration":60}]}`))
	}))
	defer srv.Close()
	resetMapboxState(t, srv.URL)

	assert.Equal(t, time.Duration(0), FindWay("32.0,34.8", "32.1,34.9"))
	// A transient failure must not poison the cache.
	assert.Equal(t, time.Minute, FindWay("32.0,34.8", "32.1,34.9"))
}

func TestFindWay_MissingTokenOrBadInputReturnsZero(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("no request should be made")
	}))
	defer srv.Close()
	resetMapboxState(t, srv.URL)

	assert.Equal(t, time.Duration(0), FindWay("", "32.1,34.9"))

	t.Setenv("MAPBOX_ACCESS_TOKEN", "")
	assert.Equal(t, time.Duration(0), FindWay("32.0,34.8", "32.1,34.9"))
}

func TestParseDirectionsResponse(t *testing.T) {
	cases := []struct {
		name string
		body string
		want time.Duration
		ok   bool
	}{
		{"valid", `{"code":"Ok","routes":[{"duration":125}]}`, 125 * time.Second, true},
		{"no route", `{"code":"NoRoute","routes":[]}`, 0, false},
		{"empty routes", `{"code":"Ok","routes":[]}`, 0, false},
		{"not json", `<html>`, 0, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, ok := parseDirectionsResponse([]byte(c.body))
			assert.Equal(t, c.ok, ok)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestParseGeocodeResponse(t *testing.T) {
	c, ok := parseGeocodeResponse([]byte(`{"features":[{"geometry":{"coordinates":[34.78,32.08]}}]}`))
	assert.True(t, ok)
	assert.Equal(t, coord{lng: 34.78, lat: 32.08}, c)

	_, ok = parseGeocodeResponse([]byte(`{"features":[]}`))
	assert.False(t, ok)
	_, ok = parseGeocodeResponse([]byte(`{"features":[{"geometry":{"coordinates":[1]}}]}`))
	assert.False(t, ok)
}

func TestParseLatLng(t *testing.T) {
	c, ok := parseLatLng(" 32.0 , 34.8 ")
	assert.True(t, ok)
	assert.Equal(t, coord{lng: 34.8, lat: 32.0}, c)

	for _, bad := range []string{"", "דיזנגוף 1, תל אביב", "91,34", "32,181", "1,2,3"} {
		_, ok := parseLatLng(bad)
		assert.False(t, ok, bad)
	}
}
