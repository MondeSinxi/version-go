package main

import (
	"regexp"
	"testing"
)

func TestBumpMajor(t *testing.T) {
	want := regexp.MustCompile("1.0.0")

	version, err := BumpVersion("major", "VERSION")
	if !want.MatchString(version) || err != nil {
		t.Fatalf(`Expected 1.0.0, got %s instead`, version)

	}
}

func TestBumpMinor(t *testing.T) {
	want := regexp.MustCompile("0.2.0")

	version, err := BumpVersion("minor", "VERSION")
	if !want.MatchString(version) || err != nil {
		t.Fatalf(`Expected 0.2.0, got %s instead`, version)

	}
}

func TestBumpPatch(t *testing.T) {
	want := regexp.MustCompile("0.1.1")

	version, err := BumpVersion("patch", "VERSION")
	if !want.MatchString(version) || err != nil {
		t.Fatalf(`Expected 0.1.1, got %s instead`, version)

	}
}
