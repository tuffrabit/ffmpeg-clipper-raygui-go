package state

import (
	"fmt"
	"strings"
)

func CreateDropDownState[T any](entries map[string]T) DropDownState[T] {
	dropDownState := DropDownState[T]{}
	dropDownState.Init(entries)

	return dropDownState
}

type DropDownState[T any] struct {
	Active      int32
	EntryList   map[string]T
	ListEntries string
}

func (dds *DropDownState[T]) Reset() {
	dds.Active = 0
	dds.EntryList = nil
	dds.ListEntries = ""
}

func (dds *DropDownState[T]) Init(entries map[string]T) error {
	if len(dds.EntryList) > 0 {
		return nil
	}

	dds.EntryList = make(map[string]T)

	for title, entry := range entries {
		dds.EntryList[title] = entry
		dds.ListEntries = fmt.Sprintf("%s%s;", dds.ListEntries, title)
	}

	dds.ListEntries = strings.TrimSuffix(dds.ListEntries, ";")

	return nil
}
