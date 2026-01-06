package state

import (
	"fmt"

	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/config"
)

type EncoderPresetState struct {
	EncoderPresetSeparated config.EncoderPresetSeparated
	QualityTarget          int
	SelectedPreset         string
}

func InitPresetStatesFromConfig(profileName string) (map[string]EncoderPresetState, error) {
	configJson, err := config.GetConfigWithLoad()
	if err != nil {
		return nil, fmt.Errorf("state.InitPresetStatesFromConfig: failed to load config, error: %w", err)
	}
	clipProfile, err := configJson.GetProfile(profileName)
	if err != nil {
		return nil, fmt.Errorf("state.InitPresetStatesFromConfig: failed to get profile %s, error: %w", profileName, err)
	}
	encoderPresetStates := make(map[string]EncoderPresetState)
	encoderPresetsSeparated := config.GetEncoderPresetsSeparated()

	for encoderTypeName, encoderPresetSeparated := range encoderPresetsSeparated {
		encoderSettings, err := clipProfile.EncoderSettings.GetEncoderSettings(encoderTypeName)
		if err != nil {
			return nil, fmt.Errorf("state.InitPresetStatesFromConfig: failed to get encoder settings %s, error: %w", encoderTypeName, err)
		}

		encoderPresetStates[encoderTypeName] = EncoderPresetState{
			EncoderPresetSeparated: encoderPresetSeparated,
			QualityTarget:          encoderSettings.GetQualityTarget(),
			SelectedPreset:         encoderSettings.GetEncodingPreset(),
		}
	}

	return encoderPresetStates, nil
}
