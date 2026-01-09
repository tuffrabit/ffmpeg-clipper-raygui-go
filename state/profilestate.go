package state

import (
	"fmt"
	"strconv"

	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/config"
)

type ProfileStateUpdatable interface {
	ScaleFactorUpdated(scaleFactor float32)
}

type ProfileState struct {
	Profile                             config.ClipProfileJson
	ProfileStateUpdatables              []ProfileStateUpdatable
	ScaleFactorInput                    string
	EncoderActive                       int32
	SaturationInput                     string
	ContrastInput                       string
	BrightnessInput                     string
	GammaInput                          string
	ExposureInput                       string
	BlackLevelInput                     string
	Libx264EncodingPresetActive         int32
	Libx264EncodingQualityTargetInput   string
	Libx265EncodingPresetActive         int32
	Libx265EncodingQualityTargetInput   string
	LibaomAv1EncodingQualityTargetInput string
	NvencH264EncodingPresetActive       int32
	NvencH264EncodingQualityTargetInput string
	NvencHevcEncodingPresetActive       int32
	NvencHevcEncodingQualityTargetInput string
	IntelH264EncodingPresetActive       int32
	IntelH264EncodingQualityTargetInput string
	IntelHevcEncodingPresetActive       int32
	IntelHevcEncodingQualityTargetInput string
	IntelAv1EncodingPresetActive        int32
	IntelAv1EncodingQualityTargetInput  string
}

func (ps *ProfileState) Init(profile config.ClipProfileJson) {
	ps.Profile = profile
	ps.ScaleFactorInput = strconv.FormatFloat(float64(profile.ScaleFactor), 'f', -1, 32)
	for _, profileStateUpdatable := range ps.ProfileStateUpdatables {
		profileStateUpdatable.ScaleFactorUpdated(profile.ScaleFactor)
	}
	ps.EncoderActive = int32(config.GetEncoderTypeIndex(profile.Encoder.Name))
	ps.SaturationInput = strconv.FormatFloat(float64(profile.Saturation), 'f', -1, 32)
	ps.ContrastInput = strconv.FormatFloat(float64(profile.Contrast), 'f', -1, 32)
	ps.BrightnessInput = strconv.FormatFloat(float64(profile.Brightness), 'f', -1, 32)
	ps.GammaInput = strconv.FormatFloat(float64(profile.Gamma), 'f', -1, 32)
	ps.ExposureInput = strconv.FormatFloat(float64(profile.Exposure), 'f', -1, 32)
	ps.BlackLevelInput = strconv.FormatFloat(float64(profile.BlackLevel), 'f', -1, 32)
	ps.Libx264EncodingPresetActive = int32(config.GetEncoderPresetIndex(config.Libx264EncoderName, profile.EncoderSettings.Libx264.EncodingPreset))
	ps.Libx264EncodingQualityTargetInput = fmt.Sprintf("%d", profile.EncoderSettings.Libx264.QualityTarget)
	ps.Libx265EncodingPresetActive = int32(config.GetEncoderPresetIndex(config.Libx265EncoderName, profile.EncoderSettings.Libx265.EncodingPreset))
	ps.Libx265EncodingQualityTargetInput = fmt.Sprintf("%d", profile.EncoderSettings.Libx265.QualityTarget)
	ps.LibaomAv1EncodingQualityTargetInput = fmt.Sprintf("%d", profile.EncoderSettings.LibaomAv1.QualityTarget)
	ps.NvencH264EncodingPresetActive = int32(config.GetEncoderPresetIndex(config.NvencH264EncoderName, profile.EncoderSettings.NvencH264.EncodingPreset))
	ps.NvencH264EncodingQualityTargetInput = fmt.Sprintf("%d", profile.EncoderSettings.NvencH264.QualityTarget)
	ps.NvencHevcEncodingPresetActive = int32(config.GetEncoderPresetIndex(config.NvencHevcEncoderName, profile.EncoderSettings.NvencHevc.EncodingPreset))
	ps.NvencHevcEncodingQualityTargetInput = fmt.Sprintf("%d", profile.EncoderSettings.NvencHevc.QualityTarget)
	ps.IntelH264EncodingPresetActive = int32(config.GetEncoderPresetIndex(config.IntelH264EncoderName, profile.EncoderSettings.IntelH264.EncodingPreset))
	ps.IntelH264EncodingQualityTargetInput = fmt.Sprintf("%d", profile.EncoderSettings.IntelH264.QualityTarget)
	ps.IntelHevcEncodingPresetActive = int32(config.GetEncoderPresetIndex(config.IntelHevcEncoderName, profile.EncoderSettings.IntelHevc.EncodingPreset))
	ps.IntelHevcEncodingQualityTargetInput = fmt.Sprintf("%d", profile.EncoderSettings.IntelHevc.QualityTarget)
	ps.IntelAv1EncodingPresetActive = int32(config.GetEncoderPresetIndex(config.IntelAv1EncoderName, profile.EncoderSettings.IntelAv1.EncodingPreset))
	ps.IntelAv1EncodingQualityTargetInput = fmt.Sprintf("%d", profile.EncoderSettings.IntelAv1.QualityTarget)
}

func (ps *ProfileState) SetScaleFactor(input string) error {
	if input == ps.ScaleFactorInput {
		return nil
	}

	var scaleFactor float32
	if input != "" {
		v, err := strconv.ParseFloat(input, 32)
		if err != nil {
			return fmt.Errorf("state.ProfileState.SetScaleFactor: failed to convert input to float, error: %w", err)
		}
		scaleFactor = float32(v)
	}

	ps.Profile.ScaleFactor = scaleFactor
	ps.ScaleFactorInput = input

	for _, profileStateUpdatable := range ps.ProfileStateUpdatables {
		profileStateUpdatable.ScaleFactorUpdated(scaleFactor)
	}

	return nil
}

func (ps *ProfileState) SetEncoder(active int32) {
	if active == ps.EncoderActive {
		return
	}

	for i, encoderType := range config.GetEncoderTypeValues() {
		encoderActive := int32(i)
		if active == encoderActive {
			ps.Profile.Encoder = encoderType
			ps.EncoderActive = encoderActive
			break
		}
	}
}

func (ps *ProfileState) SetSaturation(input string) error {
	if input == ps.SaturationInput {
		return nil
	}

	if input == "" {
		ps.Profile.Saturation = 0
		ps.SaturationInput = input
		return nil
	}

	v, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetSaturation: failed to convert input to float, error: %w", err)
	}
	ps.Profile.Saturation = float32(v)
	ps.SaturationInput = input

	return nil
}

func (ps *ProfileState) SetContrast(input string) error {
	if input == ps.ContrastInput {
		return nil
	}

	if input == "" {
		ps.Profile.Contrast = 0
		ps.ContrastInput = input
		return nil
	}

	v, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetContrast: failed to convert input to float, error: %w", err)
	}
	ps.Profile.Contrast = float32(v)
	ps.ContrastInput = input

	return nil
}

func (ps *ProfileState) SetBrightness(input string) error {
	if input == ps.BrightnessInput {
		return nil
	}

	if input == "" {
		ps.Profile.Brightness = 0
		ps.BrightnessInput = input
		return nil
	}

	v, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetBrightness: failed to convert input to float, error: %w", err)
	}
	ps.Profile.Brightness = float32(v)
	ps.BrightnessInput = input

	return nil
}

func (ps *ProfileState) SetGamma(input string) error {
	if input == ps.GammaInput {
		return nil
	}

	if input == "" {
		ps.Profile.Gamma = 0
		ps.GammaInput = input
		return nil
	}

	v, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetGamma: failed to convert input to float, error: %w", err)
	}
	ps.Profile.Gamma = float32(v)
	ps.GammaInput = input

	return nil
}

func (ps *ProfileState) SetExposure(input string) error {
	if input == ps.ExposureInput {
		return nil
	}

	if input == "" {
		ps.Profile.Exposure = 0
		ps.ExposureInput = input
		return nil
	}

	v, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetExposure: failed to convert input to float, error: %w", err)
	}
	ps.Profile.Exposure = float32(v)
	ps.ExposureInput = input

	return nil
}

func (ps *ProfileState) SetBlackLevel(input string) error {
	if input == ps.BlackLevelInput {
		return nil
	}

	if input == "" {
		ps.Profile.BlackLevel = 0
		ps.BlackLevelInput = input
		return nil
	}

	v, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetBlackLevel: failed to convert input to float, error: %w", err)
	}
	ps.Profile.BlackLevel = float32(v)
	ps.BlackLevelInput = input

	return nil
}

func (ps *ProfileState) SetLibx264EncodingPreset(active int32) {
	if active == ps.Libx264EncodingPresetActive {
		return
	}

	for i, presetName := range config.Libx264Presets.Names {
		presetActive := int32(i)
		if active == presetActive {
			ps.Profile.EncoderSettings.Libx264.EncodingPreset = presetName
			ps.Libx264EncodingPresetActive = presetActive
			break
		}
	}
}

func (ps *ProfileState) SetLibx264QualityTarget(input string) error {
	if input == ps.Libx264EncodingQualityTargetInput {
		return nil
	}

	if input == "" {
		ps.Profile.EncoderSettings.Libx264.QualityTarget = 0
		ps.Libx264EncodingQualityTargetInput = input
		return nil
	}

	v, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetLibx264QualityTarget: failed to convert input to int, error: %w", err)
	}
	ps.Profile.EncoderSettings.Libx264.QualityTarget = v
	ps.Libx264EncodingQualityTargetInput = input

	return nil
}

func (ps *ProfileState) SetLibx265EncodingPreset(active int32) {
	if active == ps.Libx265EncodingPresetActive {
		return
	}

	for i, presetName := range config.Libx265Presets.Names {
		presetActive := int32(i)
		if active == presetActive {
			ps.Profile.EncoderSettings.Libx265.EncodingPreset = presetName
			ps.Libx265EncodingPresetActive = presetActive
			break
		}
	}
}

func (ps *ProfileState) SetLibx265QualityTarget(input string) error {
	if input == ps.Libx265EncodingQualityTargetInput {
		return nil
	}

	if input == "" {
		ps.Profile.EncoderSettings.Libx265.QualityTarget = 0
		ps.Libx265EncodingQualityTargetInput = input
		return nil
	}

	v, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetLibx265QualityTarget: failed to convert input to int, error: %w", err)
	}
	ps.Profile.EncoderSettings.Libx265.QualityTarget = v
	ps.Libx265EncodingQualityTargetInput = input

	return nil
}

func (ps *ProfileState) SetLibaomAv1QualityTarget(input string) error {
	if input == ps.LibaomAv1EncodingQualityTargetInput {
		return nil
	}

	if input == "" {
		ps.Profile.EncoderSettings.LibaomAv1.QualityTarget = 0
		ps.LibaomAv1EncodingQualityTargetInput = input
		return nil
	}

	v, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetLibaomAv1QualityTarget: failed to convert input to int, error: %w", err)
	}
	ps.Profile.EncoderSettings.LibaomAv1.QualityTarget = v
	ps.LibaomAv1EncodingQualityTargetInput = input

	return nil
}

func (ps *ProfileState) SetNvencH264EncodingPreset(active int32) {
	if active == ps.NvencH264EncodingPresetActive {
		return
	}

	for i, presetName := range config.Nvenc264Presets.Names {
		presetActive := int32(i)
		if active == presetActive {
			ps.Profile.EncoderSettings.NvencH264.EncodingPreset = presetName
			ps.NvencH264EncodingPresetActive = presetActive
			break
		}
	}
}

func (ps *ProfileState) SetNvencH264QualityTarget(input string) error {
	if input == ps.NvencH264EncodingQualityTargetInput {
		return nil
	}

	if input == "" {
		ps.Profile.EncoderSettings.NvencH264.QualityTarget = 0
		ps.NvencH264EncodingQualityTargetInput = input
		return nil
	}

	v, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetNvencH264QualityTarget: failed to convert input to int, error: %w", err)
	}
	ps.Profile.EncoderSettings.NvencH264.QualityTarget = v
	ps.NvencH264EncodingQualityTargetInput = input

	return nil
}

func (ps *ProfileState) SetNvencHevcEncodingPreset(active int32) {
	if active == ps.NvencHevcEncodingPresetActive {
		return
	}

	for i, presetName := range config.NvencHevcPresets.Names {
		presetActive := int32(i)
		if active == presetActive {
			ps.Profile.EncoderSettings.NvencHevc.EncodingPreset = presetName
			ps.NvencHevcEncodingPresetActive = presetActive
			break
		}
	}
}

func (ps *ProfileState) SetNvencHevcQualityTarget(input string) error {
	if input == ps.NvencHevcEncodingQualityTargetInput {
		return nil
	}

	if input == "" {
		ps.Profile.EncoderSettings.NvencHevc.QualityTarget = 0
		ps.NvencHevcEncodingQualityTargetInput = input
		return nil
	}

	v, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetNvencHevcQualityTarget: failed to convert input to int, error: %w", err)
	}
	ps.Profile.EncoderSettings.NvencHevc.QualityTarget = v
	ps.NvencHevcEncodingQualityTargetInput = input

	return nil
}

func (ps *ProfileState) SetIntelH264EncodingPreset(active int32) {
	if active == ps.IntelH264EncodingPresetActive {
		return
	}

	for i, presetName := range config.IntelH264Presets.Names {
		presetActive := int32(i)
		if active == presetActive {
			ps.Profile.EncoderSettings.IntelH264.EncodingPreset = presetName
			ps.IntelH264EncodingPresetActive = presetActive
			break
		}
	}
}

func (ps *ProfileState) SetIntelH264QualityTarget(input string) error {
	if input == ps.IntelH264EncodingQualityTargetInput {
		return nil
	}

	if input == "" {
		ps.Profile.EncoderSettings.IntelH264.QualityTarget = 0
		ps.IntelH264EncodingQualityTargetInput = input
		return nil
	}

	v, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetIntelH264QualityTarget: failed to convert input to int, error: %w", err)
	}
	ps.Profile.EncoderSettings.IntelH264.QualityTarget = v
	ps.IntelH264EncodingQualityTargetInput = input

	return nil
}

func (ps *ProfileState) SetIntelHevcEncodingPreset(active int32) {
	if active == ps.IntelHevcEncodingPresetActive {
		return
	}

	for i, presetName := range config.IntelHevcPresets.Names {
		presetActive := int32(i)
		if active == presetActive {
			ps.Profile.EncoderSettings.IntelHevc.EncodingPreset = presetName
			ps.IntelHevcEncodingPresetActive = presetActive
			break
		}
	}
}

func (ps *ProfileState) SetIntelHevcQualityTarget(input string) error {
	if input == ps.IntelHevcEncodingQualityTargetInput {
		return nil
	}

	if input == "" {
		ps.Profile.EncoderSettings.IntelHevc.QualityTarget = 0
		ps.IntelHevcEncodingQualityTargetInput = input
		return nil
	}

	v, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetIntelHevcQualityTarget: failed to convert input to int, error: %w", err)
	}
	ps.Profile.EncoderSettings.IntelHevc.QualityTarget = v
	ps.IntelHevcEncodingQualityTargetInput = input

	return nil
}

func (ps *ProfileState) SetIntelAv1EncodingPreset(active int32) {
	if active == ps.IntelAv1EncodingPresetActive {
		return
	}

	for i, presetName := range config.IntelAv1Presets.Names {
		presetActive := int32(i)
		if active == presetActive {
			ps.Profile.EncoderSettings.IntelAv1.EncodingPreset = presetName
			ps.IntelAv1EncodingPresetActive = presetActive
			break
		}
	}
}

func (ps *ProfileState) SetIntelAv1QualityTarget(input string) error {
	if input == ps.IntelAv1EncodingQualityTargetInput {
		return nil
	}

	if input == "" {
		ps.Profile.EncoderSettings.IntelAv1.QualityTarget = 0
		ps.IntelAv1EncodingQualityTargetInput = input
		return nil
	}

	v, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("state.ProfileState.SetIntelAv1QualityTarget: failed to convert input to int, error: %w", err)
	}
	ps.Profile.EncoderSettings.IntelAv1.QualityTarget = v
	ps.IntelAv1EncodingQualityTargetInput = input

	return nil
}

func (ps *ProfileState) SetEncodingPreset(encoderName string, active int32) {
	switch encoderName {
	case config.Libx264EncoderName:
		ps.SetLibx264EncodingPreset(active)
	case config.Libx265EncoderName:
		ps.SetLibx265EncodingPreset(active)
	case config.NvencH264EncoderName:
		ps.SetNvencH264EncodingPreset(active)
	case config.NvencHevcEncoderName:
		ps.SetNvencHevcEncodingPreset(active)
	case config.IntelH264EncoderName:
		ps.SetIntelH264EncodingPreset(active)
	case config.IntelHevcEncoderName:
		ps.SetIntelHevcEncodingPreset(active)
	case config.IntelAv1EncoderName:
		ps.SetIntelAv1EncodingPreset(active)
	}
}

func (ps *ProfileState) SetQualityTarget(encoderName string, input string) error {
	switch encoderName {
	case config.Libx264EncoderName:
		return ps.SetLibx264QualityTarget(input)
	case config.Libx265EncoderName:
		return ps.SetLibx265QualityTarget(input)
	case config.LibaomAv1EncoderName:
		return ps.SetLibaomAv1QualityTarget(input)
	case config.NvencH264EncoderName:
		return ps.SetNvencH264QualityTarget(input)
	case config.NvencHevcEncoderName:
		return ps.SetNvencHevcQualityTarget(input)
	case config.IntelH264EncoderName:
		return ps.SetIntelH264QualityTarget(input)
	case config.IntelHevcEncoderName:
		return ps.SetIntelHevcQualityTarget(input)
	case config.IntelAv1EncoderName:
		return ps.SetIntelAv1QualityTarget(input)
	}

	return nil
}

func (ps *ProfileState) GetEncodingPresetActive(encoderName string) int32 {
	switch encoderName {
	case config.Libx264EncoderName:
		return ps.Libx264EncodingPresetActive
	case config.Libx265EncoderName:
		return ps.Libx265EncodingPresetActive
	case config.NvencH264EncoderName:
		return ps.NvencH264EncodingPresetActive
	case config.NvencHevcEncoderName:
		return ps.NvencHevcEncodingPresetActive
	case config.IntelH264EncoderName:
		return ps.IntelH264EncodingPresetActive
	case config.IntelHevcEncoderName:
		return ps.IntelHevcEncodingPresetActive
	case config.IntelAv1EncoderName:
		return ps.IntelAv1EncodingPresetActive
	}

	return 0
}

func (ps *ProfileState) GetQualityTargetInput(encoderName string) string {
	switch encoderName {
	case config.Libx264EncoderName:
		return ps.Libx264EncodingQualityTargetInput
	case config.Libx265EncoderName:
		return ps.Libx265EncodingQualityTargetInput
	case config.LibaomAv1EncoderName:
		return ps.LibaomAv1EncodingQualityTargetInput
	case config.NvencH264EncoderName:
		return ps.NvencH264EncodingQualityTargetInput
	case config.NvencHevcEncoderName:
		return ps.NvencHevcEncodingQualityTargetInput
	case config.IntelH264EncoderName:
		return ps.IntelH264EncodingQualityTargetInput
	case config.IntelHevcEncoderName:
		return ps.IntelHevcEncodingQualityTargetInput
	case config.IntelAv1EncoderName:
		return ps.IntelAv1EncodingQualityTargetInput
	}

	return ""
}
