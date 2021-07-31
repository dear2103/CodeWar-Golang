package kata

func DNAStrand(dna string) string {
	// your code here
	dnaS := ""
	for i := 0; i < len(dna); i++ {
		switch string(dna[i]) {
		case "A":
			dnaS += "T"
		case "T":
			dnaS += "A"
		case "G":
			dnaS += "C"
		case "C":
			dnaS += "G"
		default:
			dnaS += ""
		}
	}
	return dnaS
}
