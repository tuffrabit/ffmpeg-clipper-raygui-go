package ui

import (
	"fmt"
	"log"

	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/components"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PROFILE_LIST_START_X         = VIDEO_LIST_END_X + MAIN_WIDTH_PADDING
	PROFILE_LIST_WIDTH   float32 = 350
	PROFILE_LIST_END_X           = PROFILE_LIST_START_X + PROFILE_LIST_WIDTH
	PROFILE_LIST_START_Y         = MAIN_HEIGHT_PADDING
	PROFILE_LIST_HEIGHT  float32 = 250
	PROFILE_LIST_END_Y           = PROFILE_LIST_START_Y + PROFILE_LIST_HEIGHT

	PROFILE_LIST_BUTTON_WIDTH   = (PROFILE_LIST_WIDTH - (MAIN_WIDTH_PADDING * 2)) / 3
	PROFILE_LIST_BUTTON_START_Y = PROFILE_LIST_END_Y + MAIN_HEIGHT_PADDING

	PROFILE_LIST_UPDATE_BUTTON_START_X = PROFILE_LIST_START_X
	PROFILE_LIST_UPDATE_BUTTON_END_X   = PROFILE_LIST_START_X + PROFILE_LIST_BUTTON_WIDTH

	PROFILE_LIST_CREATE_BUTTON_START_X = PROFILE_LIST_UPDATE_BUTTON_END_X + MAIN_WIDTH_PADDING
	PROFILE_LIST_CREATE_BUTTON_END_X   = PROFILE_LIST_CREATE_BUTTON_START_X + PROFILE_LIST_BUTTON_WIDTH

	PROFILE_LIST_DELETE_BUTTON_START_X = PROFILE_LIST_CREATE_BUTTON_END_X + MAIN_WIDTH_PADDING

	PROFILE_LIST_UPDATE_CONFIRM_CONTEXT = "profile_list_update"
	PROFILE_LIST_CREATE_INPUT_CONTEXT   = "profile_list_create"
	PROFILE_LIST_UPDATE_DELETE_CONTEXT  = "profile_list_delete"
)

var (
	profileListRect             rl.Rectangle = rl.Rectangle{X: PROFILE_LIST_START_X, Y: PROFILE_LIST_START_Y, Width: PROFILE_LIST_WIDTH, Height: PROFILE_LIST_HEIGHT}
	profileListUpdateButtonRect rl.Rectangle = rl.Rectangle{X: PROFILE_LIST_UPDATE_BUTTON_START_X, Y: PROFILE_LIST_BUTTON_START_Y, Width: PROFILE_LIST_BUTTON_WIDTH, Height: BUTTON_HEIGHT}
	profileListCreateButtonRect rl.Rectangle = rl.Rectangle{X: PROFILE_LIST_CREATE_BUTTON_START_X, Y: PROFILE_LIST_BUTTON_START_Y, Width: PROFILE_LIST_BUTTON_WIDTH, Height: BUTTON_HEIGHT}
	profileListDeleteButtonRect rl.Rectangle = rl.Rectangle{X: PROFILE_LIST_DELETE_BUTTON_START_X, Y: PROFILE_LIST_BUTTON_START_Y, Width: PROFILE_LIST_BUTTON_WIDTH, Height: BUTTON_HEIGHT}
)

func handleProfileUpdateClick(clicked bool, appState *state.AppState) {
	if clicked {
		profile := appState.ProfileListState.EntryList[appState.ProfileListState.Active]
		appState.GlobalConfirmModalState.Init("Update Profile?", fmt.Sprintf("Are you sure you want to update %s?", profile.ProfileName), PROFILE_LIST_UPDATE_CONFIRM_CONTEXT)
	}
}

func handleProfileUpdateGlobalConfirm(appState *state.AppState) {
	if appState.GlobalConfirmModalState.Completed(PROFILE_LIST_UPDATE_CONFIRM_CONTEXT) {
		if appState.GlobalConfirmModalState.Result {
			profile := appState.ProfileListState.EntryList[appState.ProfileListState.Active]
			log.Println("Updated profile:", profile.ProfileName)
			var err error
			if err != nil {
				appState.GlobalMessageModalState.Init("Update Profile Error", fmt.Sprintf("Failed to update %s, error: %v", profile.ProfileName, err), components.MESSAGE_MODAL_TYPE_ERROR)
			}
			appState.ProfileListState.Reset()
		}

		appState.GlobalConfirmModalState.Reset()
	}
}

func handleProfileCreateClick(clicked bool, appState *state.AppState) {
	if clicked {
		appState.GlobalInputModalState.Init("Create New Profile", "New profile name:", PROFILE_LIST_CREATE_INPUT_CONTEXT)
	}
}

func handleProfileCreateGlobalInput(appState *state.AppState) {
	if appState.GlobalInputModalState.Completed(PROFILE_LIST_CREATE_INPUT_CONTEXT) {
		if appState.GlobalInputModalState.Result != "" {
			log.Println("Created profile:", appState.GlobalInputModalState.Result)
			var err error
			if err != nil {
				appState.GlobalMessageModalState.Init("Create Profile Error", fmt.Sprintf("Failed to create %s, error: %v", appState.GlobalInputModalState.Result, err), components.MESSAGE_MODAL_TYPE_ERROR)
			}
			appState.ProfileListState.Reset()
		}

		appState.GlobalInputModalState.Reset()
	}
}

func handleProfileDeleteClick(clicked bool, appState *state.AppState) {
	if clicked {
		profile := appState.ProfileListState.EntryList[appState.ProfileListState.Active]
		appState.GlobalConfirmModalState.Init("Delete Profile?", fmt.Sprintf("Are you sure you want to delete %s?", profile.ProfileName), PROFILE_LIST_UPDATE_DELETE_CONTEXT)
	}
}

func handleProfileDeleteGlobalConfirm(appState *state.AppState) {
	if appState.GlobalConfirmModalState.Completed(PROFILE_LIST_UPDATE_DELETE_CONTEXT) {
		if appState.GlobalConfirmModalState.Result {
			profile := appState.ProfileListState.EntryList[appState.ProfileListState.Active]
			log.Println("Deleted profile:", profile.ProfileName)
			var err error
			if err != nil {
				appState.GlobalMessageModalState.Init("Delete Profile Error", fmt.Sprintf("Failed to delete %s, error: %v", profile.ProfileName, err), components.MESSAGE_MODAL_TYPE_ERROR)
			}
			appState.ProfileListState.Reset()
		}

		appState.GlobalConfirmModalState.Reset()
	}
}

func ProfileList(appState *state.AppState) error {
	if appState.LocalDirectory == "" {
		return nil
	}

	appState.ProfileListState.InitProfileList()
	appState.ProfileListState.Active = gui.ListView(profileListRect, appState.ProfileListState.ListEntries, &appState.ProfileListState.ScrollIndex, appState.ProfileListState.Active)
	updateButton := gui.Button(profileListUpdateButtonRect, gui.IconText(gui.ICON_FILE_SAVE, "Update"))
	createButton := gui.Button(profileListCreateButtonRect, gui.IconText(gui.ICON_FILE_ADD, "Create"))
	deleteButton := gui.Button(profileListDeleteButtonRect, gui.IconText(gui.ICON_FILE_DELETE, "Delete"))

	handleProfileUpdateClick(updateButton, appState)
	handleProfileUpdateGlobalConfirm(appState)
	handleProfileCreateClick(createButton, appState)
	handleProfileCreateGlobalInput(appState)
	handleProfileDeleteClick(deleteButton, appState)
	handleProfileDeleteGlobalConfirm(appState)

	return nil
}
