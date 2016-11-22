package firstelement

import (
	"testing"
)

// FullName A structure to store a full name
type FullName struct {
	FirstName string
	LastName  string
}

// GetSampleData Returns a sample list of FullNames
func GetSampleData() []*FullName {
	samples := []*FullName{
		&FullName{"Danish", "Mujeeb"},
		&FullName{"Mikael", "Danish"},
		&FullName{"Hello", "World"},
	}

	return samples
}

func Convensional(data []*FullName) *FullName {
	if len(data) > 0 {
		return data[0]
	}
	return nil
}

func Classy(data []*FullName) *FullName {
	for _, fn := range data {
		return fn
	}
	return nil
}

func TestConvensional(t *testing.T) {
	samples := GetSampleData()
	fn := Convensional(samples)
	if fn.FirstName != "Danish" {
		t.Error("Expected first name to be 'Danish'", fn.FirstName)
	}
}

func TestClassy(t *testing.T) {
	samples := GetSampleData()
	fn := Classy(samples)
	if fn.FirstName != "Danish" {
		t.Error("Expected first name to be 'Danish'", fn.FirstName)
	}
}

func BenchmarkConvensional(b *testing.B) {
	samples := GetSampleData()
	for i := 0; i < b.N; i++ {
		Convensional(samples)
	}
}

func BenchmarkClassy(b *testing.B) {
	samples := GetSampleData()
	for i := 0; i < b.N; i++ {
		Classy(samples)
	}
}
