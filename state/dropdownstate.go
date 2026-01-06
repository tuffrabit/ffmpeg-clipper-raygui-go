package state

import (
	"fmt"
	"strings"
)

func CreateDropDownState[T any](titles []string, entries []T) (DropDownState[T], error) {
	dropDownState := DropDownState[T]{}
	err := dropDownState.Init(titles, entries)
	if err != nil {
		return dropDownState, fmt.Errorf("state.CreateDropDownState: failed to init, error: %w", err)
	}

	return dropDownState, nil
}

type DropDownState[T any] struct {
	Active      int32
	EntryList   []T
	ListEntries string
}

func (dds *DropDownState[T]) Reset() {
	dds.Active = 0
	dds.EntryList = nil
	dds.ListEntries = ""
}

func (dds *DropDownState[T]) Selected() T {
	if len(dds.EntryList) > 0 {
		return dds.EntryList[dds.Active]
	}

	var thing T
	return thing
}

func (dds *DropDownState[T]) Init(titles []string, entries []T) error {
	if len(dds.EntryList) > 0 {
		return nil
	}

	titlesLen := len(titles)
	entriesLen := len(entries)

	if titlesLen == 0 || entriesLen == 0 {
		return fmt.Errorf("state.DropDownState.Init: zero length options are not allowed")
	}

	if titlesLen != entriesLen {
		return fmt.Errorf("state.DropDownState.Init: title length does not equal entries length")
	}

	for index, entry := range entries {
		dds.EntryList = append(dds.EntryList, entry)
		dds.ListEntries = fmt.Sprintf("%s%s;", dds.ListEntries, titles[index])
	}

	dds.ListEntries = strings.TrimSuffix(dds.ListEntries, ";")

	return nil
}
