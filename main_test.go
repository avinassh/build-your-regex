package main

import "testing"

// tests regexes containing just the string literals
func TestMatchLiteral(t *testing.T) {
	tests := []struct{
		regex string
		text string
		expected bool
	}{
		{"a", "a", true},
		{"a", "ab", true},
		{"aa", "aa", true},
		{"goo", "google", true},
		{"le", "google", true},
		{"a", "b", false},
		{"aa", "ab", false},
		{"ba", "ab", false},
	}
	for _, tt := range tests {
		if got := match(tt.regex, tt.text); got != tt.expected {
			t.Fatalf("failed for regex %s in %s, expected=%t got=%t", tt.regex, tt.text, tt.expected, got)
		}
	}
}

// tests '.' regex
func TestMatchSingleCharacter(t *testing.T) {
	tests := []struct{
		regex string
		text string
		expected bool
	}{
		{".", "a", true},
		{"a.", "ab", true},
		{"a.a", "aaa", true},
		{"..ng", "king", true},
		{".", ".", true},
		{"a.", "b", false},
	}
	for _, tt := range tests {
		if got := match(tt.regex, tt.text); got != tt.expected {
			t.Fatalf("failed for regex %s in %s, expected=%t got=%t", tt.regex, tt.text, tt.expected, got)
		}
	}
}

// tests '^' regex
func TestMatchBeginnings(t *testing.T) {
	tests := []struct{
		regex string
		text string
		expected bool
	}{
		{"^a", "a", true},
		{"^ab", "ab", true},
		{"^aa", "aaa", true},
		{"^le", "google", false},
		{"^abc", "acabc", false},
	}
	for _, tt := range tests {
		if got := match(tt.regex, tt.text); got != tt.expected {
			t.Fatalf("failed for regex %s in %s, expected=%t got=%t", tt.regex, tt.text, tt.expected, got)
		}
	}
}

// tests '$' regex
func TestMatchEnds(t *testing.T) {
	tests := []struct{
		regex string
		text string
		expected bool
	}{
		{"a$", "a", true},
		{"ab$", "ab", true},
		{"aa$", "aaa", true},
		{"le$", "google", true},
		{"go$", "google", false},
		{"abc$", "abcac", false},
	}
	for _, tt := range tests {
		if got := match(tt.regex, tt.text); got != tt.expected {
			t.Fatalf("failed for regex %s in %s, expected=%t got=%t", tt.regex, tt.text, tt.expected, got)
		}
	}
}

// tests '*' regex
func TestMatchStar(t *testing.T) {
	tests := []struct{
		regex string
		text string
		expected bool
	}{
		{"a*", "aaaaaa", true},
		{"ab*", "ab", true},
		{"ab*", "a", true},
		{"aa*", "aaaa", true},
		{"go*gle", "google", true},
		{"go*gle", "gooooooogle", true},
	}
	for _, tt := range tests {
		if got := match(tt.regex, tt.text); got != tt.expected {
			t.Fatalf("failed for regex %s in %s, expected=%t got=%t", tt.regex, tt.text, tt.expected, got)
		}
	}
}