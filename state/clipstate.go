package state

type ClipState struct {
	Start         string
	StartEditMode bool
	End           string
	EndEditMode   bool
}

func (cs *ClipState) Reset() {
	cs.Start = ""
	cs.StartEditMode = false
	cs.End = ""
	cs.EndEditMode = false
}
