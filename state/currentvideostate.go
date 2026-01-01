package state

type CurrentVideoState struct {
	Name     string
	FullPath string
	Width    int
	Height   int
}

func (cvs *CurrentVideoState) Reset() {
	cvs.Name = ""
	cvs.FullPath = ""
	cvs.Width = 0
	cvs.Height = 0
}
