package projetred

func isdead(c Character) bool {
	if c.PV <= 0 {
		return true
	}
	return false
}
