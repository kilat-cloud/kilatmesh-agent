package version

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const httpAgent = "pkg/test"

func TestNewUpdate(t *testing.T) {
	t.Setenv(EnvAllowUpstreamUpdate, "1")
	version = "1.0.0"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "10.0.0")
	}))
	defer svr.Close()
	versionURL = svr.URL

	wg := &sync.WaitGroup{}
	wg.Add(1)

	onUpdate := false
	u := NewUpdateAndStart(httpAgent)
	defer u.StopWatch()
	u.SetOnUpdateListener(func() {
		onUpdate = true
		wg.Done()
	})

	waitTimeout(wg)
	if onUpdate != true {
		t.Errorf("update not found")
	}
}

func TestDoNotUpdate(t *testing.T) {
	t.Setenv(EnvAllowUpstreamUpdate, "1")
	version = "11.0.0"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "10.0.0")
	}))
	defer svr.Close()
	versionURL = svr.URL

	wg := &sync.WaitGroup{}
	wg.Add(1)

	onUpdate := false
	u := NewUpdateAndStart(httpAgent)
	defer u.StopWatch()
	u.SetOnUpdateListener(func() {
		onUpdate = true
		wg.Done()
	})

	waitTimeout(wg)
	if onUpdate == true {
		t.Errorf("invalid update")
	}
}

func TestDaemonUpdate(t *testing.T) {
	t.Setenv(EnvAllowUpstreamUpdate, "1")
	version = "11.0.0"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "11.0.0")
	}))
	defer svr.Close()
	versionURL = svr.URL

	wg := &sync.WaitGroup{}
	wg.Add(1)

	onUpdate := false
	u := NewUpdateAndStart(httpAgent)
	defer u.StopWatch()
	u.SetOnUpdateListener(func() {
		onUpdate = true
		wg.Done()
	})

	u.SetDaemonVersion("10.0.0")

	waitTimeout(wg)
	if onUpdate != true {
		t.Errorf("invalid daemon version check")
	}
}

// TestFetcherDisabledByDefault proves the KilatMesh default: without the
// explicit operator opt-in the fetcher never contacts the version endpoint.
func TestFetcherDisabledByDefault(t *testing.T) {
	t.Setenv(EnvAllowUpstreamUpdate, "")
	if err := os.Unsetenv(EnvAllowUpstreamUpdate); err != nil {
		t.Fatalf("unset %s: %v", EnvAllowUpstreamUpdate, err)
	}
	version = "1.0.0"

	var hits atomic.Int32
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		fmt.Fprintf(w, "10.0.0")
	}))
	defer svr.Close()
	versionURL = svr.URL

	u := NewUpdateAndStart(httpAgent)
	defer u.StopWatch()
	u.SetOnUpdateListener(func() {
		t.Error("update listener must not fire while the fetcher is disabled")
	})

	time.Sleep(200 * time.Millisecond)
	if got := hits.Load(); got != 0 {
		t.Errorf("fetcher contacted the version endpoint %d times while disabled", got)
	}
	if u.LatestVersion() != nil {
		t.Errorf("latest version must stay unknown while the fetcher is disabled, got %s", u.LatestVersion())
	}
}

func TestUpstreamUpdateAllowedParsing(t *testing.T) {
	t.Setenv(EnvAllowUpstreamUpdate, "")
	if err := os.Unsetenv(EnvAllowUpstreamUpdate); err != nil {
		t.Fatalf("unset %s: %v", EnvAllowUpstreamUpdate, err)
	}
	if UpstreamUpdateAllowed() {
		t.Error("unset env must default to disabled")
	}

	for value, want := range map[string]bool{
		"1":     true,
		"true":  true,
		"0":     false,
		"false": false,
		"bogus": false,
	} {
		t.Setenv(EnvAllowUpstreamUpdate, value)
		if got := UpstreamUpdateAllowed(); got != want {
			t.Errorf("%s=%q: got %v, want %v", EnvAllowUpstreamUpdate, value, got, want)
		}
	}
}

func waitTimeout(wg *sync.WaitGroup) {
	c := make(chan struct{})
	go func() {
		wg.Wait()
		close(c)
	}()
	select {
	case <-c:
		return
	case <-time.After(time.Second):
		return
	}
}
