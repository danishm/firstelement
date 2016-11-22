package firstelement

import (
	"testing"
)

// GetSampleData Returns a sample list of FullNames
func GetSampleData() []FullName {
	samples := []FullName{
		FullName{"Danish", "Mujeeb"},
		FullName{"Mikael", "Danish"},
		FullName{"Hello", "World"},
	}

	return samples
}

// GetSampleData Returns a sample list of FullNames
func GetSampleDataP() []*FullName {
	samples := []*FullName{
		&FullName{"Danish", "Mujeeb"},
		&FullName{"Mikael", "Danish"},
		&FullName{"Hello", "World"},
	}

	return samples
}

// Testing functions

func TestGetFirstTraditional(t *testing.T) {
	samples := GetSampleData()
	fn, err := GetFirstTraditional(samples)
	if err != nil {
		t.Error("Error retrieving first element")
	}
	if fn.FirstName != "Danish" {
		t.Error("Expected first name to be 'Danish'", fn.FirstName)
	}
}

func TestGetFirstRangeLoop(t *testing.T) {
	samples := GetSampleData()
	fn, err := GetFirstRangeLoop(samples)
	if err != nil {
		t.Error("Error retrieving first element")
	}
	if fn.FirstName != "Danish" {
		t.Error("Expected first name to be 'Danish'", fn.FirstName)
	}
}

// Benchmarking functions

func BenchmarkGetFirstTraditional(b *testing.B) {
	samples := GetSampleData()
	for i := 0; i < b.N; i++ {
		GetFirstTraditional(samples)
	}
}

func BenchmarkGetFirstRangeLoop(b *testing.B) {
	samples := GetSampleData()
	for i := 0; i < b.N; i++ {
		GetFirstRangeLoop(samples)
	}
}

func BenchmarkGetFirstTraditionalLoop(b *testing.B) {
	samples := GetSampleData()
	for i := 0; i < b.N; i++ {
		GetFirstTraditional(samples)
	}
}

// Benchmarking functions: Pointer based

func BenchmarkGetFirstTraditionalP(b *testing.B) {
	samples := GetSampleDataP()
	for i := 0; i < b.N; i++ {
		GetFirstTraditionalP(samples)
	}
}

func BenchmarkGetFirstRangeLoopP(b *testing.B) {
	samples := GetSampleDataP()
	for i := 0; i < b.N; i++ {
		GetFirstRangeLoopP(samples)
	}
}

func BenchmarkGetFirstTraditionalLoopP(b *testing.B) {
	samples := GetSampleDataP()
	for i := 0; i < b.N; i++ {
		GetFirstTraditionalP(samples)
	}
}
