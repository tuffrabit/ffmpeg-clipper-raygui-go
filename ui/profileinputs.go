package ui

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/components"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/config"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
)

const (
	PROFILE_INPUTS_START_X = PROFILE_LIST_END_X + MAIN_WIDTH_PADDING
	PROFILE_INPUTS_WIDTH   = float32(WINDOW_WIDTH) - PROFILE_LIST_END_X - (MAIN_WIDTH_PADDING * 2)
	PROFILE_INPUTS_END_X   = PROFILE_INPUTS_START_X + PROFILE_INPUTS_WIDTH
	PROFILE_INPUTS_START_Y = MAIN_HEIGHT_PADDING

	PROFILE_INPUTS_SCALE_DOWN_LABEL_END_Y   = PROFILE_INPUTS_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_SCALE_DOWN_INPUT_START_Y = PROFILE_INPUTS_SCALE_DOWN_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_SCALE_DOWN_INPUT_END_Y   = PROFILE_INPUTS_SCALE_DOWN_INPUT_START_Y + TEXTBOX_HEIGHT

	PROFILE_INPUTS_ENCODERS_LABEL_START_Y = PROFILE_INPUTS_SCALE_DOWN_INPUT_END_Y + MAIN_HEIGHT_PADDING
	PROFILE_INPUTS_ENCODERS_LABEL_END_Y   = PROFILE_INPUTS_ENCODERS_LABEL_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_ENCODERS_INPUT_START_Y = PROFILE_INPUTS_ENCODERS_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_ENCODERS_INPUT_END_Y   = PROFILE_INPUTS_ENCODERS_INPUT_START_Y + DROPDOWNBOX_HEIGHT

	PROFILE_INPUTS_ENCODER_PRESET_LABEL_START_Y = PROFILE_INPUTS_ENCODERS_INPUT_END_Y + MAIN_HEIGHT_PADDING
	PROFILE_INPUTS_ENCODER_PRESET_LABEL_END_Y   = PROFILE_INPUTS_ENCODER_PRESET_LABEL_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_ENCODER_PRESET_INPUT_START_Y = PROFILE_INPUTS_ENCODER_PRESET_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_ENCODER_PRESET_INPUT_END_Y   = PROFILE_INPUTS_ENCODER_PRESET_INPUT_START_Y + DROPDOWNBOX_HEIGHT

	PROFILE_INPUTS_QUALITY_TARGET_LABEL_START_Y = PROFILE_INPUTS_ENCODER_PRESET_INPUT_END_Y + MAIN_HEIGHT_PADDING
	PROFILE_INPUTS_QUALITY_TARGET_LABEL_END_Y   = PROFILE_INPUTS_QUALITY_TARGET_LABEL_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_QUALITY_TARGET_INPUT_START_Y = PROFILE_INPUTS_QUALITY_TARGET_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_QUALITY_TARGET_INPUT_END_Y   = PROFILE_INPUTS_QUALITY_TARGET_INPUT_START_Y + TEXTBOX_HEIGHT

	PROFILE_INPUTS_SATURATION_LABEL_START_Y = PROFILE_INPUTS_QUALITY_TARGET_INPUT_END_Y + MAIN_HEIGHT_PADDING
	PROFILE_INPUTS_SATURATION_LABEL_END_Y   = PROFILE_INPUTS_SATURATION_LABEL_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_SATURATION_INPUT_START_Y = PROFILE_INPUTS_SATURATION_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_SATURATION_INPUT_END_Y   = PROFILE_INPUTS_SATURATION_INPUT_START_Y + TEXTBOX_HEIGHT

	PROFILE_INPUTS_CONTRAST_LABEL_START_Y = PROFILE_INPUTS_SATURATION_INPUT_END_Y + MAIN_HEIGHT_PADDING
	PROFILE_INPUTS_CONTRAST_LABEL_END_Y   = PROFILE_INPUTS_CONTRAST_LABEL_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_CONTRAST_INPUT_START_Y = PROFILE_INPUTS_CONTRAST_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_CONTRAST_INPUT_END_Y   = PROFILE_INPUTS_CONTRAST_INPUT_START_Y + TEXTBOX_HEIGHT

	PROFILE_INPUTS_BRIGHTNESS_LABEL_START_Y = PROFILE_INPUTS_CONTRAST_INPUT_END_Y + MAIN_HEIGHT_PADDING
	PROFILE_INPUTS_BRIGHTNESS_LABEL_END_Y   = PROFILE_INPUTS_BRIGHTNESS_LABEL_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_BRIGHTNESS_INPUT_START_Y = PROFILE_INPUTS_BRIGHTNESS_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_BRIGHTNESS_INPUT_END_Y   = PROFILE_INPUTS_BRIGHTNESS_INPUT_START_Y + TEXTBOX_HEIGHT

	PROFILE_INPUTS_GAMMA_LABEL_START_Y = PROFILE_INPUTS_BRIGHTNESS_INPUT_END_Y + MAIN_HEIGHT_PADDING
	PROFILE_INPUTS_GAMMA_LABEL_END_Y   = PROFILE_INPUTS_GAMMA_LABEL_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_GAMMA_INPUT_START_Y = PROFILE_INPUTS_GAMMA_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_GAMMA_INPUT_END_Y   = PROFILE_INPUTS_GAMMA_INPUT_START_Y + TEXTBOX_HEIGHT

	PROFILE_INPUTS_EXPOSURE_LABEL_START_Y = PROFILE_INPUTS_GAMMA_INPUT_END_Y + MAIN_HEIGHT_PADDING
	PROFILE_INPUTS_EXPOSURE_LABEL_END_Y   = PROFILE_INPUTS_EXPOSURE_LABEL_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_EXPOSURE_INPUT_START_Y = PROFILE_INPUTS_EXPOSURE_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_EXPOSURE_INPUT_END_Y   = PROFILE_INPUTS_EXPOSURE_INPUT_START_Y + TEXTBOX_HEIGHT

	PROFILE_INPUTS_BLACK_LEVEL_LABEL_START_Y = PROFILE_INPUTS_EXPOSURE_INPUT_END_Y + MAIN_HEIGHT_PADDING
	PROFILE_INPUTS_BLACK_LEVEL_LABEL_END_Y   = PROFILE_INPUTS_BLACK_LEVEL_LABEL_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_BLACK_LEVEL_INPUT_START_Y = PROFILE_INPUTS_BLACK_LEVEL_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_BLACK_LEVEL_INPUT_END_Y   = PROFILE_INPUTS_BLACK_LEVEL_INPUT_START_Y + TEXTBOX_HEIGHT
)

var (
	profileInputsScaleDownLabelRect     rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsScaleDownInputRect     rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_SCALE_DOWN_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: TEXTBOX_HEIGHT}
	profileInputsEncodersLabelRect      rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_ENCODERS_LABEL_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsEncodersInputRect      rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_ENCODERS_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: DROPDOWNBOX_HEIGHT}
	profileInputsEncoderPresetLabelRect rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_ENCODER_PRESET_LABEL_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsEncoderPresetInputRect rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_ENCODER_PRESET_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: DROPDOWNBOX_HEIGHT}
	profileInputsQualityTargetLabelRect rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_QUALITY_TARGET_LABEL_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsQualityTargetInputRect rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_QUALITY_TARGET_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: TEXTBOX_HEIGHT}
	profileInputsSaturationLabelRect    rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_SATURATION_LABEL_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsSaturationInputRect    rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_SATURATION_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: TEXTBOX_HEIGHT}
	profileInputsContrastLabelRect      rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_CONTRAST_LABEL_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsContrastInputRect      rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_CONTRAST_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: TEXTBOX_HEIGHT}
	profileInputsBrightnessLabelRect    rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_BRIGHTNESS_LABEL_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsBrightnessInputRect    rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_BRIGHTNESS_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: TEXTBOX_HEIGHT}
	profileInputsGammaLabelRect         rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_GAMMA_LABEL_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsGammaInputRect         rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_GAMMA_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: TEXTBOX_HEIGHT}
	profileInputsExposureLabelRect      rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_EXPOSURE_LABEL_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsExposureInputRect      rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_EXPOSURE_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: TEXTBOX_HEIGHT}
	profileInputsBlackLevelLabelRect    rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_BLACK_LEVEL_LABEL_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsBlackLevelInputRect    rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_BLACK_LEVEL_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: TEXTBOX_HEIGHT}

	scaleFactorEditMode            bool = false
	encoderListEditMode            bool = false
	encoderListDropDownInitialized bool
	encoderListDropDownState       state.DropDownState[config.EncoderType]
	encoderListPreviousActiveName  string
	encoderPresetState             state.EncoderPresetState
	encoderSettingsDropDownState   state.DropDownState[string]
	encoderPresetListEditMode      bool = false
	encoderSettingsValid           bool
	encoderPresetListActiveFake    int32 = 0
	qualityTargetEditMode          bool  = false
	saturationEditMode             bool  = false
	contrastEditMode               bool  = false
	brightnessEditMode             bool  = false
	gammaEditMode                  bool  = false
	exposureEditMode               bool  = false
	blackLevelEditMode             bool  = false
)

func handleScaleFactorInput(scaleFactor string, appState *state.AppState) {
	err := appState.ProfileState.ScaleFactor.Set(scaleFactor)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid Scale Down Factor", fmt.Sprintf("Failed to parse scale down factor value %s, error: %v", scaleFactor, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}
}

func handleQualityTargetInput(qualityTarget string, appState *state.AppState) {
	err := appState.ProfileState.QualityTarget.Set(qualityTarget)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid Quality Target", fmt.Sprintf("Failed to parse quality target value %s, error: %v", qualityTarget, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}
}

func handleSaturationInput(saturation string, appState *state.AppState) {
	err := appState.ProfileState.Saturation.Set(saturation)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid Saturation", fmt.Sprintf("Failed to parse saturation value %s, error: %v", saturation, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}
}

func handleContrastInput(contrast string, appState *state.AppState) {
	err := appState.ProfileState.Contrast.Set(contrast)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid Contrast", fmt.Sprintf("Failed to parse contrast value %s, error: %v", contrast, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}
}

func handleBrightnessInput(brightness string, appState *state.AppState) {
	err := appState.ProfileState.Brightness.Set(brightness)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid Brightness", fmt.Sprintf("Failed to parse brightness value %s, error: %v", brightness, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}
}

func handleGammaInput(gamma string, appState *state.AppState) {
	err := appState.ProfileState.Gamma.Set(gamma)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid Gamma", fmt.Sprintf("Failed to parse gamma value %s, error: %v", gamma, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}
}

func handleExposureInput(exposure string, appState *state.AppState) {
	err := appState.ProfileState.Exposure.Set(exposure)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid Exposure", fmt.Sprintf("Failed to parse exposure value %s, error: %v", exposure, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}
}

func handleBlackLevelInput(blackLevel string, appState *state.AppState) {
	err := appState.ProfileState.BlackLevel.Set(blackLevel)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid Black Level", fmt.Sprintf("Failed to parse black level value %s, error: %v", blackLevel, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}
}

func encoderSettingIndex(value string, values []string) int32 {
	for index, v := range values {
		if v == value {
			return int32(index)
		}
	}

	return 0
}

func ProfileInputs(appState *state.AppState) error {
	if appState.LocalDirectory == "" {
		return nil
	}

	var err error
	profile := appState.ProfileListState.SelectedProfile()
	gui.Label(profileInputsScaleDownLabelRect, "Scale Down Factor")
	scaleFactor := fmt.Sprintf("%f", profile.ScaleFactor)
	if gui.TextBox(profileInputsScaleDownInputRect, &scaleFactor, 30, scaleFactorEditMode) {
		scaleFactorEditMode = !scaleFactorEditMode
	}
	handleScaleFactorInput(scaleFactor, appState)

	if !encoderListDropDownInitialized {
		encoderListDropDownState, err = state.CreateDropDownState(config.GetEncoderTypeTitles(), config.GetEncoderTypeValues())
		if err != nil {
			appState.GlobalMessageModalState.Init("Encoder List Error", fmt.Sprintf("Failed to init encoder list values, error: %v", err), components.MESSAGE_MODAL_TYPE_ERROR)
		} else {
			encoderListDropDownInitialized = true
		}
	}

	encoderType := encoderListDropDownState.Selected().Name
	if encoderListPreviousActiveName != encoderType {
		encoderPresetState, err = appState.GetEncoderPresetState(encoderType)
		if err == nil {
			encoderSettingsDropDownState, err = state.CreateDropDownState(encoderPresetState.EncoderPresetSeparated.Titles, encoderPresetState.EncoderPresetSeparated.Names)
			if err != nil {
				encoderSettingsValid = false
			} else {
				encoderSettingsDropDownState.Active = encoderSettingIndex(encoderPresetState.SelectedPreset, encoderPresetState.EncoderPresetSeparated.Names)
				encoderSettingsValid = true
			}
		} else {
			encoderSettingsValid = false
		}
	}
	encoderListPreviousActiveName = encoderType
	appState.ProfileState.Encoder = encoderType

	qualityTarget := fmt.Sprintf("%d", encoderPresetState.QualityTarget)
	gui.Label(profileInputsQualityTargetLabelRect, "Quality Target (0 to 51)")
	if gui.TextBox(profileInputsQualityTargetInputRect, &qualityTarget, 30, qualityTargetEditMode) {
		qualityTargetEditMode = !qualityTargetEditMode
	}
	handleQualityTargetInput(qualityTarget, appState)

	saturation := fmt.Sprintf("%f", profile.Saturation)
	gui.Label(profileInputsSaturationLabelRect, "Saturation (Default 1.0)")
	if gui.TextBox(profileInputsSaturationInputRect, &saturation, 30, saturationEditMode) {
		saturationEditMode = !saturationEditMode
	}
	handleSaturationInput(saturation, appState)

	contrast := fmt.Sprintf("%f", profile.Contrast)
	gui.Label(profileInputsContrastLabelRect, "Contrast (Default 1.0)")
	if gui.TextBox(profileInputsContrastInputRect, &contrast, 30, contrastEditMode) {
		contrastEditMode = !contrastEditMode
	}
	handleContrastInput(contrast, appState)

	brightness := fmt.Sprintf("%f", profile.Brightness)
	gui.Label(profileInputsBrightnessLabelRect, "Brightness (Default 0.0)")
	if gui.TextBox(profileInputsBrightnessInputRect, &brightness, 30, brightnessEditMode) {
		brightnessEditMode = !brightnessEditMode
	}
	handleBrightnessInput(brightness, appState)

	gamma := fmt.Sprintf("%f", profile.Gamma)
	gui.Label(profileInputsGammaLabelRect, "Gamma (Default 1.0)")
	if gui.TextBox(profileInputsGammaInputRect, &gamma, 30, gammaEditMode) {
		gammaEditMode = !gammaEditMode
	}
	handleGammaInput(gamma, appState)

	exposure := fmt.Sprintf("%f", profile.Exposure)
	gui.Label(profileInputsExposureLabelRect, "Exposure (Default 0.0)")
	if gui.TextBox(profileInputsExposureInputRect, &exposure, 30, exposureEditMode) {
		exposureEditMode = !exposureEditMode
	}
	handleExposureInput(exposure, appState)

	blackLevel := fmt.Sprintf("%f", profile.BlackLevel)
	gui.Label(profileInputsBlackLevelLabelRect, "Black Level (Default 0.0)")
	if gui.TextBox(profileInputsBlackLevelInputRect, &blackLevel, 30, blackLevelEditMode) {
		blackLevelEditMode = !blackLevelEditMode
	}
	handleBlackLevelInput(blackLevel, appState)

	gui.Label(profileInputsEncoderPresetLabelRect, "Encoding Preset")
	if encoderSettingsValid {
		if gui.DropdownBox(profileInputsEncoderPresetInputRect, encoderSettingsDropDownState.ListEntries, &encoderSettingsDropDownState.Active, encoderPresetListEditMode) {
			encoderPresetListEditMode = !encoderPresetListEditMode
		}
		appState.ProfileState.EncoderPreset = encoderSettingsDropDownState.Selected()
	} else {
		appState.ProfileState.EncoderPreset = ""
		gui.SetState(gui.STATE_DISABLED)
		gui.DropdownBox(profileInputsEncoderPresetInputRect, ";", &encoderPresetListActiveFake, false)
		gui.SetState(gui.STATE_NORMAL)
	}

	gui.Label(profileInputsEncodersLabelRect, "Encoder")
	if gui.DropdownBox(profileInputsEncodersInputRect, encoderListDropDownState.ListEntries, &encoderListDropDownState.Active, encoderListEditMode) {
		encoderListEditMode = !encoderListEditMode
	}

	appState.ProfileListState.UpdateSelectedProfile(profile)

	return nil
}
