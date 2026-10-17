package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadJSONRejectsInvalidRequests(t *testing.T) {
	cases := []struct {
		name, body string
		valid      bool
	}{
		{"valid", `{"action":"log"}`, true},
		{"empty", "", false},
		{"malformed", `{"action":`, false},
		{"multiple values", `{"action":"log"} {"action":"mail"}`, false},
		{"over size limit", `{"action":"` + strings.Repeat("x", 1048576) + `"}`, false},
	}
	app := &Config{}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/handle", strings.NewReader(tc.body))
			var payload RequestPayload
			err := app.readJSON(httptest.NewRecorder(), r, &payload)
			if (err == nil) != tc.valid {
				t.Fatalf("valid=%v, error=%v", tc.valid, err)
			}
			if tc.valid && payload.Action != "log" {
				t.Fatalf("unexpected action %q", payload.Action)
			}
		})
	}
}

func TestHandleSubmissionRejectsUnknownAction(t *testing.T) {
	app := &Config{}
	r := httptest.NewRequest(http.MethodPost, "/handle", strings.NewReader(`{"action":"unsupported"}`))
	w := httptest.NewRecorder()
	app.HandleSubmission(w, r)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("status=%d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "unknown action") {
		t.Fatalf("response=%s", w.Body.String())
	}
}

func BenchmarkReadJSON(b *testing.B) {
	app := &Config{}
	body := `{"action":"log","log":{"name":"benchmark","data":"` + strings.Repeat("x", 1024) + `"}}`
	b.ReportAllocs()
	b.SetBytes(int64(len(body)))
	for i := 0; i < b.N; i++ {
		r := httptest.NewRequest(http.MethodPost, "/handle", strings.NewReader(body))
		var payload RequestPayload
		if err := app.readJSON(httptest.NewRecorder(), r, &payload); err != nil {
			b.Fatal(err)
		}
	}
}
