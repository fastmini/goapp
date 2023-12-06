package enum

type State int

const (
	CPYDY State = iota
	CPYKS
	CPYSPH
	CYTDYN
	CYTKSN
	CPYDY_EXTEND
	CPYKS_EXTEND
	CPYSPH_EXTEND
)

func GetCytEnum() []string {
	return []string{CYTDYN.String(), CYTKSN.String()}
}

func (s State) String() string {
	switch s {
	case CPYDY:
		return "CPYDY"
	case CPYKS:
		return "CPYKS"
	case CPYSPH:
		return "CPYSPH"
	case CYTDYN:
		return "CYTDYN"
	case CYTKSN:
		return "CYTKSN"
	case CPYDY_EXTEND:
		return "CPYDY_EXTEND"
	case CPYKS_EXTEND:
		return "CPYKS_EXTEND"
	case CPYSPH_EXTEND:
		return "CPYSPH_EXTEND"
	default:
		return "Unknown"
	}
}
