package firstelement

import (
	"errors"
)

// FullName A structure to store a full name
type FullName struct {
	FirstName string
	LastName  string
}

// GetFirstTraditional returns the first element by doing
// a length check and using the 0 index of an array
func GetFirstTraditional(data []FullName) (FullName, error) {
	if len(data) > 0 {
		return data[0], nil
	}
	return FullName{}, errors.New("Empty list")
}

// GetFirstRangeLoop returns the first element by using
// a range based loop and returning element in 1st loop
func GetFirstRangeLoop(data []FullName) (FullName, error) {
	for _, fn := range data {
		return fn, nil
	}
	return FullName{}, errors.New("Empty list")
}

// GetFirstTraditionalLoop returns the first element by using
// a traditional loop and returning element in 1st loop
func GetFirstTraditionalLoop(data []FullName) (FullName, error) {
	for i := 0; i < len(data); i++ {
		return data[i], nil
	}
	return FullName{}, errors.New("Empty list")
}

// GetFirstTraditionalP returns the first element by doing
// a length check and using the 0 index of an array
func GetFirstTraditionalP(data []*FullName) (*FullName, error) {
	if len(data) > 0 {
		return data[0], nil
	}
	return nil, errors.New("Empty list")
}

// GetFirstRangeLoopP returns the first element by using
// a range based loop and returning element in 1st loop
func GetFirstRangeLoopP(data []*FullName) (*FullName, error) {
	for _, fn := range data {
		return fn, nil
	}
	return nil, errors.New("Empty list")
}

// GetFirstTraditionalLoopP returns the first element by using
// a traditional loop and returning element in 1st loop
func GetFirstTraditionalLoopP(data []*FullName) (*FullName, error) {
	for i := 0; i < len(data); i++ {
		return data[i], nil
	}
	return nil, errors.New("Empty list")
}
