package problem_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"go.followtheprocess.codes/problem"
	"go.followtheprocess.codes/test"
)

func TestProblemJSON(t *testing.T) {
	test.ColorEnabled(true) // Force colour for diffs

	// TODO(@FollowTheProcess): Use snapshot here
	tests := []struct {
		name    string
		golden  string
		problem problem.Problem
	}{
		{
			name:    "empty",
			problem: problem.Problem{},
			golden:  "empty.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.MarshalIndent(tt.problem, "", "  ")
			test.Ok(t, err)

			// MarshalIndent doesn't add a newline at the end
			got = append(got, '\n')

			want, err := os.ReadFile(filepath.Join("testdata", tt.golden))
			test.Ok(t, err)

			test.DiffBytes(t, got, want)
		})
	}
}
