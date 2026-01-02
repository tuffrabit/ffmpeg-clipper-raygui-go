package state

import (
	"fmt"
	"strings"

	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/config"
)

type ProfileListState struct {
	ScrollIndex int32
	Active      int32
	EntryList   []config.ClipProfileJson
	ListEntries string
}

func (pls *ProfileListState) Reset() {
	pls.ScrollIndex = 0
	pls.Active = 0
	pls.EntryList = nil
	pls.ListEntries = ""
}

func (pls *ProfileListState) SelectedProfile() config.ClipProfileJson {
	if len(pls.EntryList) > 0 {
		return pls.EntryList[pls.Active]
	}

	return config.ClipProfileJson{}
}

func (pls *ProfileListState) UpdateSelectedProfile(profile config.ClipProfileJson) {
	if len(pls.EntryList) == 0 {
		return
	}

	pls.EntryList[pls.Active] = profile
}

func (pls *ProfileListState) InitProfileList() error {
	if len(pls.EntryList) > 0 {
		return nil
	}

	for _, profile := range config.GetConfig().ClipProfiles {
		pls.EntryList = append(pls.EntryList, profile)
		pls.ListEntries = fmt.Sprintf("%s%s;", pls.ListEntries, profile.ProfileName)
	}

	pls.ListEntries = strings.TrimSuffix(pls.ListEntries, ";")

	return nil
}
