package problem_test

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"flag"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"go.followtheprocess.codes/problem"
	"go.followtheprocess.codes/test"
)

var update = flag.Bool("update", false, "Update golden files")

func TestProblemJSON(t *testing.T) {
	test.ColorEnabled(os.Getenv("CI") == "") // Force colour for diffs locally

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
		{
			name:    "empty new",
			problem: problem.New(),
			golden:  "empty-new.json", // Not the same as empty as `type` defaults to "about:blank"
		},
		{
			name: "no extra",
			problem: problem.Problem{
				Type:     "https://example.com/probs/out-of-credit",
				Title:    "Not enough credit",
				Detail:   "Your current balance is 30, but that costs 50",
				Instance: "/account/12345/msgs/abc",
				Status:   http.StatusBadRequest,
			},
			golden: "no-extra.json",
		},
		{
			name: "no extra new",
			problem: problem.New(
				problem.Type("https://example.com/probs/out-of-credit"),
				problem.Title("Not enough credit"),
				problem.Detail("Your current balance is 30, but that costs 50"),
				problem.Instance("/account/12345/msgs/abc"),
				problem.Status(http.StatusBadRequest),
			),
			golden: "no-extra.json", // Should serialize exactly the same
		},
		{
			name: "with extra",
			problem: problem.Problem{
				Type:     "https://example.com/probs/out-of-credit",
				Title:    "Not enough credit",
				Detail:   "Your current balance is 30, but that costs 50",
				Instance: "/account/12345/msgs/abc",
				Status:   http.StatusBadRequest,
				Extra: map[string]any{
					"balance":  30,
					"accounts": []string{"/accounts/12345", "/accounts/67890"},
				},
			},
			golden: "with-extra.json",
		},
		{
			name: "with extra new",
			problem: problem.New(
				problem.Type("https://example.com/probs/out-of-credit"),
				problem.Title("Not enough credit"),
				problem.Detail("Your current balance is 30, but that costs 50"),
				problem.Instance("/account/12345/msgs/abc"),
				problem.Status(http.StatusBadRequest),
				problem.Extra("balance", 30),
				problem.Extra("accounts", []string{"/accounts/12345", "/accounts/67890"}),
			),
			golden: "with-extra.json", // Should serialize exactly the same
		},
		{
			name: "with extra map new",
			problem: problem.New(
				problem.Type("https://example.com/probs/out-of-credit"),
				problem.Title("Not enough credit"),
				problem.Detail("Your current balance is 30, but that costs 50"),
				problem.Instance("/account/12345/msgs/abc"),
				problem.Status(http.StatusBadRequest),
				problem.ExtraMap(map[string]any{
					"balance":  30,
					"accounts": []string{"/accounts/12345", "/accounts/67890"},
				}),
			),
			golden: "with-extra.json", // Should serialize exactly the same
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.problem, jsontext.WithIndent("  "), json.Deterministic(true))
			test.Ok(t, err)

			golden := filepath.Join("testdata", tt.golden)

			if *update {
				test.Ok(t, os.WriteFile(golden, got, 0o644))
			}

			want, err := os.ReadFile(golden)
			test.Ok(t, err)

			test.DiffBytes(t, got, want)
		})
	}
}
