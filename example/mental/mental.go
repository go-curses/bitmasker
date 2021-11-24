package mental

//go:generate bitmasker -type=State
type State uint

const (
	Unconscious State = 0
	Conscious   State = 1 << iota
	Meditative
	Distracted
	Entertained = Distracted
)
