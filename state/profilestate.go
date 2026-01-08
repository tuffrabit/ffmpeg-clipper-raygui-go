package state

type ProfileState struct {
	ProfileName   string
	ScaleFactor   Float32StringValue
	Encoder       string
	EncoderPreset string
	QualityTarget IntStringValue
	Saturation    Float32StringValue
	Contrast      Float32StringValue
	Brightness    Float32StringValue
	Gamma         Float32StringValue
	Exposure      Float32StringValue
	BlackLevel    Float32StringValue
	PlayAfter     bool
}
