package state

type CurrentVideoState struct {
	Name        string
	FullPath    string
	Width       int
	Height      int
	ScaleFactor float32
	NewWidth    int
	NewHeight   int
}

func (cvs *CurrentVideoState) Reset() {
	cvs.Name = ""
	cvs.FullPath = ""
	cvs.Width = 0
	cvs.Height = 0
	cvs.ScaleFactor = 0
	cvs.NewWidth = 0
	cvs.NewHeight = 0
}

func (cvs *CurrentVideoState) Update(width int, height int) {
	cvs.Width = width
	cvs.Height = height

	cvs.ScaleFactorUpdated(cvs.ScaleFactor)
}

func (cvs *CurrentVideoState) ScaleFactorUpdated(scaleFactor float32) {
	if scaleFactor == 0 {
		return
	}

	cvs.NewWidth = int(float32(cvs.Width) / scaleFactor)
	cvs.NewHeight = int(float32(cvs.Height) / scaleFactor)
	cvs.ScaleFactor = scaleFactor
}
