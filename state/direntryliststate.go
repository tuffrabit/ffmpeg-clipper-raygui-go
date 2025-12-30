package state

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type DirEntryListState struct {
	ScrollIndex int32
	Active      int32
	EntryList   []os.DirEntry
	ListEntries string
}

func (dels *DirEntryListState) Reset() {
	dels.ScrollIndex = 0
	dels.Active = 0
	dels.EntryList = nil
	dels.ListEntries = ""
}

func (dels *DirEntryListState) InitFileList(directory string, includeExtensions []string) error {
	if len(dels.EntryList) > 0 {
		return nil
	}

	if directory == "" {
		return nil
	}

	directoryEntries, err := os.ReadDir(directory)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %v", directory, err)
	}

	for _, directoryEntry := range directoryEntries {
		if directoryEntry.IsDir() || directoryEntry.Name() == "" {
			continue
		}

		if len(includeExtensions) > 0 {
			extension := filepath.Ext(directoryEntry.Name())

			if !slices.Contains(includeExtensions, extension) {
				continue
			}
		}

		dels.EntryList = append(dels.EntryList, directoryEntry)
		dels.ListEntries = fmt.Sprintf("%s%s;", dels.ListEntries, directoryEntry.Name())
	}

	dels.ListEntries = strings.TrimSuffix(dels.ListEntries, ";")

	return nil
}
