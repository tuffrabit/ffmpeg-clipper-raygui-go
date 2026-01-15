package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"golang.org/x/exp/slices"
)

const CONFIG_FILENAME = "ffmpeg-clipper-config.json"

var configJson ConfigJson

type ConfigJson struct {
	Style        string            `json:"style"`
	ClipProfiles []ClipProfileJson `json:"profiles"`
}

func (cj *ConfigJson) GetProfile(profileName string) (ClipProfileJson, error) {
	for _, clipProfile := range cj.ClipProfiles {
		if clipProfile.ProfileName == profileName {
			return clipProfile, nil
		}
	}

	return ClipProfileJson{}, fmt.Errorf("config.ConfigJson.GetProfile: profile %s does not exist", profileName)
}

type ClipProfileJson struct {
	ProfileName     string                         `json:"profileName"`
	ScaleFactor     float32                        `json:"scaleFactor"`
	Encoder         EncoderType                    `json:"encoder"`
	EncoderSettings ClipProfileJsonEncoderSettings `json:"encoderSettings"`
	Saturation      float32                        `json:"saturation"`
	Contrast        float32                        `json:"contrast"`
	Brightness      float32                        `json:"brightness"`
	Gamma           float32                        `json:"gamma"`
	Exposure        float32                        `json:"exposure"`
	BlackLevel      float32                        `json:"black_level"`
	PlayAfter       bool                           `json:"playAfter"`
}

type ClipProfileJsonEncoderSettings struct {
	Libx264   Libx264EncoderSettings   `json:"libx264"`
	Libx265   Libx265EncoderSettings   `json:"libx265"`
	LibaomAv1 LibaomAv1EncoderSettings `json:"libaom-av1"`
	NvencH264 NvencH264EncoderSettings `json:"h264_nvenc"`
	NvencHevc NvencHevcEncoderSettings `json:"hevc_nvenc"`
	IntelH264 IntelH264EncoderSettings `json:"h264_qsv"`
	IntelHevc IntelHevcEncoderSettings `json:"hevc_qsv"`
	IntelAv1  IntelAv1EncoderSettings  `json:"av1_qsv"`
}

func (s *ClipProfileJsonEncoderSettings) GetEncoderSettings(t string) (EncoderSettingsInterface, error) {
	switch t {
	case Libx264EncoderName:
		return s.Libx264, nil
	case Libx265EncoderName:
		return s.Libx265, nil
	case LibaomAv1EncoderName:
		return s.LibaomAv1, nil
	case NvencH264EncoderName:
		return s.NvencH264, nil
	case NvencHevcEncoderName:
		return s.NvencHevc, nil
	case IntelH264EncoderName:
		return s.IntelH264, nil
	case IntelHevcEncoderName:
		return s.IntelHevc, nil
	case IntelAv1EncoderName:
		return s.IntelAv1, nil
	}

	return nil, fmt.Errorf("config.ClipProfileJsonEncoderSettings.GetEncoderSettings: settings %s does not exist", t)
}

func (s *ClipProfileJsonEncoderSettings) SetEncoderSettings(t string, i EncoderSettingsInterface) {
	switch t {
	case Libx264EncoderName:
		val, ok := i.(*Libx264EncoderSettings)
		if ok {
			s.Libx264.EncodingPreset = val.EncodingPreset
			s.Libx264.QualityTarget = val.QualityTarget
		}
	case Libx265EncoderName:
		val, ok := i.(*Libx265EncoderSettings)
		if ok {
			s.Libx265.EncodingPreset = val.EncodingPreset
			s.Libx265.QualityTarget = val.QualityTarget
		}
	case LibaomAv1EncoderName:
		val, ok := i.(*LibaomAv1EncoderSettings)
		if ok {
			s.LibaomAv1.QualityTarget = val.QualityTarget
		}
	case NvencH264EncoderName:
		val, ok := i.(*NvencH264EncoderSettings)
		if ok {
			s.NvencH264.EncodingPreset = val.EncodingPreset
			s.NvencH264.QualityTarget = val.QualityTarget
		}
	case NvencHevcEncoderName:
		val, ok := i.(*NvencHevcEncoderSettings)
		if ok {
			s.NvencHevc.EncodingPreset = val.EncodingPreset
			s.NvencHevc.QualityTarget = val.QualityTarget
		}
	case IntelH264EncoderName:
		val, ok := i.(*IntelH264EncoderSettings)
		if ok {
			s.IntelH264.EncodingPreset = val.EncodingPreset
			s.IntelH264.QualityTarget = val.QualityTarget
		}
	case IntelHevcEncoderName:
		val, ok := i.(*IntelHevcEncoderSettings)
		if ok {
			s.IntelHevc.EncodingPreset = val.EncodingPreset
			s.IntelHevc.QualityTarget = val.QualityTarget
		}
	case IntelAv1EncoderName:
		val, ok := i.(*IntelAv1EncoderSettings)
		if ok {
			s.IntelAv1.EncodingPreset = val.EncodingPreset
			s.IntelAv1.QualityTarget = val.QualityTarget
		}
	}
}

type EncoderSettingsInterface interface {
	Validate() bool
	GetEncodingPreset() string
	GetQualityTarget() int
}

type Libx264EncoderSettings struct {
	EncodingPreset string `json:"encodingPreset"`
	QualityTarget  int    `json:"qualityTarget"`
}

func (s Libx264EncoderSettings) Validate() bool {
	if s.QualityTarget < 0 || s.QualityTarget > 51 {
		return false
	}

	return true
}

func (s Libx264EncoderSettings) GetEncodingPreset() string {
	return s.EncodingPreset
}

func (s Libx264EncoderSettings) GetQualityTarget() int {
	return s.QualityTarget
}

type Libx265EncoderSettings struct {
	EncodingPreset string `json:"encodingPreset"`
	QualityTarget  int    `json:"qualityTarget"`
}

func (s Libx265EncoderSettings) Validate() bool {
	if s.QualityTarget < 0 || s.QualityTarget > 51 {
		return false
	}

	return true
}

func (s Libx265EncoderSettings) GetEncodingPreset() string {
	return s.EncodingPreset
}

func (s Libx265EncoderSettings) GetQualityTarget() int {
	return s.QualityTarget
}

type LibaomAv1EncoderSettings struct {
	QualityTarget int `json:"qualityTarget"`
}

func (s LibaomAv1EncoderSettings) Validate() bool {
	if s.QualityTarget < 0 || s.QualityTarget > 63 {
		return false
	}

	return true
}

func (s LibaomAv1EncoderSettings) GetEncodingPreset() string {
	return ""
}

func (s LibaomAv1EncoderSettings) GetQualityTarget() int {
	return s.QualityTarget
}

type NvencH264EncoderSettings struct {
	EncodingPreset string `json:"encodingPreset"`
	QualityTarget  int    `json:"qualityTarget"`
}

func (s NvencH264EncoderSettings) Validate() bool {
	if s.QualityTarget < 0 || s.QualityTarget > 51 {
		return false
	}

	return true
}

func (s NvencH264EncoderSettings) GetEncodingPreset() string {
	return s.EncodingPreset
}

func (s NvencH264EncoderSettings) GetQualityTarget() int {
	return s.QualityTarget
}

type NvencHevcEncoderSettings struct {
	EncodingPreset string `json:"encodingPreset"`
	QualityTarget  int    `json:"qualityTarget"`
}

func (s NvencHevcEncoderSettings) Validate() bool {
	if s.QualityTarget < 0 || s.QualityTarget > 51 {
		return false
	}

	return true
}

func (s NvencHevcEncoderSettings) GetEncodingPreset() string {
	return s.EncodingPreset
}

func (s NvencHevcEncoderSettings) GetQualityTarget() int {
	return s.QualityTarget
}

type IntelH264EncoderSettings struct {
	EncodingPreset string `json:"encodingPreset"`
	QualityTarget  int    `json:"qualityTarget"`
}

func (s IntelH264EncoderSettings) Validate() bool {
	if s.QualityTarget < 1 || s.QualityTarget > 51 {
		return false
	}

	return true
}

func (s IntelH264EncoderSettings) GetEncodingPreset() string {
	return s.EncodingPreset
}

func (s IntelH264EncoderSettings) GetQualityTarget() int {
	return s.QualityTarget
}

type IntelHevcEncoderSettings struct {
	EncodingPreset string `json:"encodingPreset"`
	QualityTarget  int    `json:"qualityTarget"`
}

func (s IntelHevcEncoderSettings) Validate() bool {
	if s.QualityTarget < 1 || s.QualityTarget > 51 {
		return false
	}

	return true
}

func (s IntelHevcEncoderSettings) GetEncodingPreset() string {
	return s.EncodingPreset
}

func (s IntelHevcEncoderSettings) GetQualityTarget() int {
	return s.QualityTarget
}

type IntelAv1EncoderSettings struct {
	EncodingPreset string `json:"encodingPreset"`
	QualityTarget  int    `json:"qualityTarget"`
}

func (s IntelAv1EncoderSettings) Validate() bool {
	if s.QualityTarget < 1 || s.QualityTarget > 51 {
		return false
	}

	return true
}

func (s IntelAv1EncoderSettings) GetEncodingPreset() string {
	return s.EncodingPreset
}

func (s IntelAv1EncoderSettings) GetQualityTarget() int {
	return s.QualityTarget
}

type EncoderType struct {
	Name  string `json:"name"`
	Title string `json:"type"`
}

func (et EncoderType) String() string {
	return et.Title
}

const (
	Libx264EncoderName   string = "libx264"
	Libx265EncoderName   string = "libx265"
	LibaomAv1EncoderName string = "libaom-av1"
	NvencH264EncoderName string = "h264_nvenc"
	NvencHevcEncoderName string = "hevc_nvenc"
	IntelH264EncoderName string = "h264_qsv"
	IntelHevcEncoderName string = "hevc_qsv"
	IntelAv1EncoderName  string = "av1_qsv"
)

type EncoderPresetSeparated struct {
	Names  []string
	Titles []string
}

var (
	Libx264EncoderType   EncoderType = EncoderType{Name: Libx264EncoderName, Title: "CPU H.264"}
	Libx265EncoderType   EncoderType = EncoderType{Name: Libx265EncoderName, Title: "CPU H.265"}
	LibaomAv1EncoderType EncoderType = EncoderType{Name: LibaomAv1EncoderName, Title: "CPU AV1"}
	NvencH264EncoderType EncoderType = EncoderType{Name: NvencH264EncoderName, Title: "Nvidia H.264"}
	NvencHevcEncoderType EncoderType = EncoderType{Name: NvencHevcEncoderName, Title: "Nvidia H.265"}
	IntelH264EncoderType EncoderType = EncoderType{Name: IntelH264EncoderName, Title: "Intel H.264"}
	IntelHevcEncoderType EncoderType = EncoderType{Name: IntelHevcEncoderName, Title: "Intel H.265"}
	IntelAv1EncoderType  EncoderType = EncoderType{Name: IntelAv1EncoderName, Title: "Intel AV1"}

	Libx264Presets = EncoderPresetSeparated{
		Names:  []string{"ultrafast", "superfast", "veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow"},
		Titles: []string{"ultrafast", "superfast", "veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow"},
	}
	Libx265Presets = EncoderPresetSeparated{
		Names:  []string{"ultrafast", "superfast", "veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow"},
		Titles: []string{"ultrafast", "superfast", "veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow"},
	}
	LibaomAv1Presets = EncoderPresetSeparated{
		Names:  []string{},
		Titles: []string{},
	}
	Nvenc264Presets = EncoderPresetSeparated{
		Names:  []string{"p1", "p2", "p3", "p4", "p5", "p6", "p7"},
		Titles: []string{"fastest", "faster", "fast", "medium", "slow", "slower", "slowest"},
	}
	NvencHevcPresets = EncoderPresetSeparated{
		Names:  []string{"p1", "p2", "p3", "p4", "p5", "p6", "p7"},
		Titles: []string{"fastest", "faster", "fast", "medium", "slow", "slower", "slowest"},
	}
	IntelH264Presets = EncoderPresetSeparated{
		Names:  []string{"veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow"},
		Titles: []string{"veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow"},
	}
	IntelHevcPresets = EncoderPresetSeparated{
		Names:  []string{"veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow"},
		Titles: []string{"veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow"},
	}
	IntelAv1Presets = EncoderPresetSeparated{
		Names:  []string{"veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow"},
		Titles: []string{"veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow"},
	}
)

func GetEncoderTypeIndex(encoderType string) int {
	encoderTypeValues := GetEncoderTypeValues()
	for i, encoderTypeValue := range encoderTypeValues {
		if encoderTypeValue.Name == encoderType {
			return i
		}
	}

	return 0
}

func GetEncoderPresetIndex(encoderType string, encoderPresetName string) int {
	encoderPresetsSeparated := GetEncoderPresetsSeparated()
	encoderPresetSeparated, ok := encoderPresetsSeparated[encoderType]
	if !ok {
		return 0
	}

	for i, encoderPresetSeparatedName := range encoderPresetSeparated.Names {
		if encoderPresetSeparatedName == encoderPresetName {
			return i
		}
	}

	return 0
}

func GetEncoderPresetsSeparated() map[string]EncoderPresetSeparated {
	m := make(map[string]EncoderPresetSeparated)

	m[Libx264EncoderName] = Libx264Presets
	m[Libx265EncoderName] = Libx265Presets
	m[LibaomAv1EncoderName] = LibaomAv1Presets
	m[NvencH264EncoderName] = Nvenc264Presets
	m[NvencHevcEncoderName] = NvencHevcPresets
	m[IntelH264EncoderName] = IntelH264Presets
	m[IntelHevcEncoderName] = IntelHevcPresets
	m[IntelAv1EncoderName] = IntelAv1Presets

	return m
}

func NewProfile(name string) *ClipProfileJson {
	libx264EncoderSettings := Libx264EncoderSettings{
		EncodingPreset: "slow",
		QualityTarget:  24,
	}

	libx265EncoderSettings := Libx265EncoderSettings{
		EncodingPreset: "slow",
		QualityTarget:  29,
	}

	libaomAv1EncoderSettings := LibaomAv1EncoderSettings{
		QualityTarget: 29,
	}

	nvencH264EncoderSettings := NvencH264EncoderSettings{
		EncodingPreset: "p4",
		QualityTarget:  26,
	}

	nvencHevcEncoderSettings := NvencHevcEncoderSettings{
		EncodingPreset: "p4",
		QualityTarget:  31,
	}

	intelH264EncoderSettings := IntelH264EncoderSettings{
		EncodingPreset: "medium",
		QualityTarget:  29,
	}

	intelHevcEncoderSettings := IntelHevcEncoderSettings{
		EncodingPreset: "medium",
		QualityTarget:  31,
	}

	intelAv1EncoderSettings := IntelAv1EncoderSettings{
		EncodingPreset: "medium",
		QualityTarget:  33,
	}

	encoderSettings := ClipProfileJsonEncoderSettings{
		Libx264:   libx264EncoderSettings,
		Libx265:   libx265EncoderSettings,
		LibaomAv1: libaomAv1EncoderSettings,
		NvencH264: nvencH264EncoderSettings,
		NvencHevc: nvencHevcEncoderSettings,
		IntelH264: intelH264EncoderSettings,
		IntelHevc: intelHevcEncoderSettings,
		IntelAv1:  intelAv1EncoderSettings,
	}

	return &ClipProfileJson{
		ProfileName:     name,
		ScaleFactor:     2.666,
		Encoder:         Libx264EncoderType,
		EncoderSettings: encoderSettings,
		Saturation:      1,
		Contrast:        1,
		Brightness:      0,
		Gamma:           1,
		Exposure:        0,
		BlackLevel:      0,
		PlayAfter:       true,
	}
}

func GetEncoderTypes() map[string]EncoderType {
	m := make(map[string]EncoderType)

	m[Libx264EncoderName] = Libx264EncoderType
	m[Libx265EncoderName] = Libx265EncoderType
	m[LibaomAv1EncoderName] = LibaomAv1EncoderType
	m[NvencH264EncoderName] = NvencH264EncoderType
	m[NvencHevcEncoderName] = NvencHevcEncoderType
	m[IntelH264EncoderName] = IntelH264EncoderType
	m[IntelHevcEncoderName] = IntelHevcEncoderType
	m[IntelAv1EncoderName] = IntelAv1EncoderType

	return m
}

func GetEncoderTypesByTitle() map[string]EncoderType {
	m := make(map[string]EncoderType)

	m[Libx264EncoderType.Title] = Libx264EncoderType
	m[Libx265EncoderType.Title] = Libx265EncoderType
	m[LibaomAv1EncoderType.Title] = LibaomAv1EncoderType
	m[NvencH264EncoderType.Title] = NvencH264EncoderType
	m[NvencHevcEncoderType.Title] = NvencHevcEncoderType
	m[IntelH264EncoderType.Title] = IntelH264EncoderType
	m[IntelHevcEncoderType.Title] = IntelHevcEncoderType
	m[IntelAv1EncoderType.Title] = IntelAv1EncoderType

	return m
}

func GetEncoderTypeTitles() []string {
	return []string{
		Libx264EncoderType.Title,
		Libx265EncoderType.Title,
		LibaomAv1EncoderType.Title,
		NvencH264EncoderType.Title,
		NvencHevcEncoderType.Title,
		IntelH264EncoderType.Title,
		IntelHevcEncoderType.Title,
		IntelAv1EncoderType.Title,
	}
}

func GetEncoderTypeValues() []EncoderType {
	return []EncoderType{
		Libx264EncoderType,
		Libx265EncoderType,
		LibaomAv1EncoderType,
		NvencH264EncoderType,
		NvencHevcEncoderType,
		IntelH264EncoderType,
		IntelHevcEncoderType,
		IntelAv1EncoderType,
	}
}

func ValidateEncoderType(encoderName string) bool {
	for t := range GetEncoderTypes() {
		if encoderName == string(t) {
			return true
		}
	}

	return false
}

func LoadConfig() error {
	_, err := os.Stat(CONFIG_FILENAME)
	if err != nil {
		err = createDefaultConfigFile()
		if err != nil {
			return fmt.Errorf("config.LoadConfig: could not create default config: %w", err)
		}
	}

	_, err = os.Stat(CONFIG_FILENAME)
	if err != nil {
		return errors.New("config.LoadConfig: config file does not exist")
	}

	file, err := os.Open(CONFIG_FILENAME)
	if err != nil {
		return fmt.Errorf("config.LoadConfig: could not open config file: %w", err)
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("config.LoadConfig: could not read config file: %w", err)
	}

	var tempConfigJson ConfigJson
	err = json.Unmarshal(fileBytes, &tempConfigJson)
	if err != nil {
		return fmt.Errorf("config.LoadConfig: could not marshal json: %w", err)
	}

	configJson = tempConfigJson

	return nil
}

func GetConfig() *ConfigJson {
	return &configJson
}

func GetConfigWithLoad() (*ConfigJson, error) {
	err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("config.GetConfigWithLoad: could not load config: %w", err)
	}

	return &configJson, nil
}

func GetProfile(name string) *ClipProfileJson {
	for _, p := range configJson.ClipProfiles {
		if name == p.ProfileName {
			return &p
		}
	}

	return nil
}

func GetEncoderSettingsFromProfile(name string, encoder EncoderType) (ClipProfileJsonEncoderSettings, error) {
	profile := GetProfile(name)
	if profile == nil {
		return ClipProfileJsonEncoderSettings{}, fmt.Errorf("config.GetEncoderSettingsFromProfile: profile %v does not exist", name)
	}

	return profile.EncoderSettings, nil
}

func SaveStyle(name string) error {
	configJson, err := GetConfigWithLoad()
	if err != nil {
		return fmt.Errorf("config.SaveStyle: could not get config: %w", err)
	}

	configJson.Style = name
	err = writeConfigFile(configJson)
	if err != nil {
		return fmt.Errorf("config.SaveStyle: could not write config: %w", err)
	}

	err = LoadConfig()
	if err != nil {
		return fmt.Errorf("config.SaveStyle: could not load updated config: %w", err)
	}

	return nil
}

func SaveProfile(profileJson *ClipProfileJson) error {
	configJson, err := GetConfigWithLoad()
	if err != nil {
		return fmt.Errorf("config.SaveProfile: could not get config: %w", err)
	}

	profileIndex := -1
	for index, profile := range configJson.ClipProfiles {
		if profile.ProfileName == profileJson.ProfileName {
			profileIndex = index
			break
		}
	}

	if profileIndex != -1 {
		configJson.ClipProfiles[profileIndex] = *profileJson
	} else {
		configJson.ClipProfiles = append(configJson.ClipProfiles, *profileJson)
	}

	err = writeConfigFile(configJson)
	if err != nil {
		return fmt.Errorf("config.SaveProfile: could not write config: %w", err)
	}

	err = LoadConfig()
	if err != nil {
		return fmt.Errorf("config.SaveProfile: could not load updated config: %w", err)
	}

	return nil
}

func DeleteProfile(profileName string) error {
	configJson, err := GetConfigWithLoad()
	if err != nil {
		return fmt.Errorf("config.DeleteProfile: could not get config: %w", err)
	}

	profileIndex := -1
	for index, profile := range configJson.ClipProfiles {
		if profile.ProfileName == profileName {
			profileIndex = index
			break
		}
	}

	configJson.ClipProfiles = slices.Delete(configJson.ClipProfiles, profileIndex, profileIndex+1)
	err = writeConfigFile(configJson)
	if err != nil {
		return fmt.Errorf("config.DeleteProfile: could not write config: %w", err)
	}

	err = LoadConfig()
	if err != nil {
		return fmt.Errorf("config.DeleteProfile: could not load updated config: %w", err)
	}

	return nil
}

func writeConfigFile(configJson *ConfigJson) error {
	jsonBytes, err := json.MarshalIndent(configJson, "", " ")
	if err != nil {
		return fmt.Errorf("config.writeConfigFile: could not marshal json: %w", err)
	}

	err = os.WriteFile(CONFIG_FILENAME, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("config.writeConfigFile: could not write config: %w", err)
	}

	return nil
}

func createDefaultConfigFile() error {
	defaultConfigJson := generateDefaultConfigJson()
	err := writeConfigFile(&defaultConfigJson)
	if err != nil {
		return fmt.Errorf("config.createDefaultConfigFile: could not write file: %w", err)
	}

	return nil
}

func generateDefaultConfigJson() ConfigJson {
	libx264EncoderSettings := Libx264EncoderSettings{
		EncodingPreset: "slow",
		QualityTarget:  24,
	}

	libx265EncoderSettings := Libx265EncoderSettings{
		EncodingPreset: "slow",
		QualityTarget:  29,
	}

	libaomAv1EncoderSettings := LibaomAv1EncoderSettings{
		QualityTarget: 29,
	}

	nvencH264EncoderSettings := NvencH264EncoderSettings{
		EncodingPreset: "p4",
		QualityTarget:  26,
	}

	nvencHevcEncoderSettings := NvencHevcEncoderSettings{
		EncodingPreset: "p4",
		QualityTarget:  31,
	}

	intelH264EncoderSettings := IntelH264EncoderSettings{
		EncodingPreset: "medium",
		QualityTarget:  29,
	}

	intelHevcEncoderSettings := IntelHevcEncoderSettings{
		EncodingPreset: "medium",
		QualityTarget:  31,
	}

	intelAv1EncoderSettings := IntelAv1EncoderSettings{
		EncodingPreset: "medium",
		QualityTarget:  33,
	}

	encoderSettings := ClipProfileJsonEncoderSettings{
		Libx264:   libx264EncoderSettings,
		Libx265:   libx265EncoderSettings,
		LibaomAv1: libaomAv1EncoderSettings,
		NvencH264: nvencH264EncoderSettings,
		NvencHevc: nvencHevcEncoderSettings,
		IntelH264: intelH264EncoderSettings,
		IntelHevc: intelHevcEncoderSettings,
		IntelAv1:  intelAv1EncoderSettings,
	}

	huntDayClipProfile := ClipProfileJson{
		ProfileName:     "Hunt - Day",
		ScaleFactor:     2.666,
		Encoder:         Libx264EncoderType,
		EncoderSettings: encoderSettings,
		Saturation:      2,
		Contrast:        1.1,
		Brightness:      0,
		Gamma:           1,
		Exposure:        0,
		BlackLevel:      0,
		PlayAfter:       true,
	}

	huntNightClipProfile := ClipProfileJson{
		ProfileName:     "Hunt - Night",
		ScaleFactor:     2.666,
		Encoder:         Libx264EncoderType,
		EncoderSettings: encoderSettings,
		Saturation:      2,
		Contrast:        1.1,
		Brightness:      0.1,
		Gamma:           1,
		Exposure:        0,
		BlackLevel:      0,
		PlayAfter:       true,
	}

	destinyClipProfile := ClipProfileJson{
		ProfileName:     "Destiny",
		ScaleFactor:     2.666,
		Encoder:         Libx264EncoderType,
		EncoderSettings: encoderSettings,
		Saturation:      1,
		Contrast:        1,
		Brightness:      0,
		Gamma:           1,
		Exposure:        0,
		BlackLevel:      0,
		PlayAfter:       true,
	}

	return ConfigJson{
		Style: "default",
		ClipProfiles: []ClipProfileJson{
			huntDayClipProfile,
			huntNightClipProfile,
			destinyClipProfile,
		},
	}
}
