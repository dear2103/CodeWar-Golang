package kata

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	// Implement me
	var x int
	if dadYearsOld < 2*sonYearsOld {
		x = 2*sonYearsOld - dadYearsOld
	} else {
		x = dadYearsOld - 2*sonYearsOld
	}
	return x
}
