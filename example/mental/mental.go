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

//go:generate bitmasker -kebab -type=ThingFlags
type ThingFlags uint

const (
	ThingZeroFlag ThingFlags = 0
	ThingOneFlag  ThingFlags = 1 << iota
	ThingTwoFlag
	ThingNopFlag = ThingTwoFlag
)
