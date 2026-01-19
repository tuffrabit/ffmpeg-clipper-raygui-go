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
	ActiveName  string
	EntryList   []os.DirEntry
	InitAttempt bool
	ListEntries string
}

func (dels *DirEntryListState) Reset() {
	dels.ScrollIndex = 0
	dels.Active = 0
	dels.ActiveName = ""
	dels.EntryList = nil
	dels.InitAttempt = false
	dels.ListEntries = ""
}

func (dels *DirEntryListState) ResetWithSelection() {
	activeName := dels.ActiveName
	dels.Reset()
	dels.ActiveName = activeName
}

func (dels *DirEntryListState) SetActive(index int32) {
	if len(dels.EntryList) == 0 {
		return
	}

	dels.Active = index
	dels.ActiveName = dels.EntryList[index].Name()
}

func (dels *DirEntryListState) InitFileList(directory string, includeExtensions []string) error {
	if dels.InitAttempt {
		return nil
	}

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

	activePreserved := false

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

		if dels.ActiveName != "" && dels.ActiveName == directoryEntry.Name() {
			dels.Active = int32(len(dels.EntryList) - 1)
			activePreserved = true
		}
	}

	if !activePreserved {
		dels.SetActive(0)
	}

	dels.ListEntries = strings.TrimSuffix(dels.ListEntries, ";")
	dels.InitAttempt = true

	return nil
}
