package grpc

import (
	"strings"
	"testing"

	"github.com/mholt/caddy"
)

func TestSetupPolicy(t *testing.T) {
	tests := []struct {
		input          string
		shouldErr      bool
		expectedPolicy string
		expectedErr    string
	}{
		// positive
		{"grpc . 127.0.0.1 {\npolicy random\n}\n", false, "random", ""},
		{"grpc . 127.0.0.1 {\npolicy round_robin\n}\n", false, "round_robin", ""},
		{"grpc . 127.0.0.1 {\npolicy sequential\n}\n", false, "sequential", ""},
		// negative
		{"grpc . 127.0.0.1 {\npolicy random2\n}\n", true, "random", "unknown policy"},
	}

	for i, test := range tests {
		c := caddy.NewTestController("dns", test.input)
		g, err := parseGRPC(c)

		if test.shouldErr && err == nil {
			t.Errorf("Test %d: expected error but found %s for input %s", i, err, test.input)
		}

		if err != nil {
			if !test.shouldErr {
				t.Errorf("Test %d: expected no error but found one for input %s, got: %v", i, test.input, err)
			}

			if !strings.Contains(err.Error(), test.expectedErr) {
				t.Errorf("Test %d: expected error to contain: %v, found error: %v, input: %s", i, test.expectedErr, err, test.input)
			}
		}

		if !test.shouldErr && g.p.String() != test.expectedPolicy {
			t.Errorf("Test %d: expected: %s, got: %s", i, test.expectedPolicy, g.p.String())
		}
	}
}
