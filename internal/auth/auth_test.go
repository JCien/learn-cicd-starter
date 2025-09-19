package auth

import (
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header        http.Header
		expectedKey   string
		expectedError string
	}{
		"Valid API Key":         {header: http.Header{"Authorization": []string{"ApiKey secret-key-123"}}, expectedKey: "secret-key-123"},
		"No Auth Header":        {header: http.Header{}, expectedKey: "", expectedError: "no authorization header included"},
		"Malformed Auth Header": {header: http.Header{"Authorization": []string{"-"}}, expectedKey: "", expectedError: "malformed authorization header"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.header)

			if err != nil {
				if strings.Contains(err.Error(), tc.expectedError) {
					return
				}
				t.Errorf("Unexpected from test: %v: %v\n", name, err)
			}

			diff := cmp.Diff(tc.expectedKey, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
