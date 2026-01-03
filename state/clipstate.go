package state

type ClipState struct {
	Start         string
	StartEditMode bool
	End           string
	EndEditMode   bool
}

func (cs *ClipState) Reset() {
	cs.Start = ""
	cs.End = ""
}
