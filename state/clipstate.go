package state

type ClipState struct {
	Start string
	End   string
}

func (cs *ClipState) Reset() {
	cs.Start = ""
	cs.End = ""
}
