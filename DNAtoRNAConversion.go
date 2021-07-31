package kata

func DNAtoRNA(dna string) string {
	// your code here
	rna := ""
	for i := 0; i < len(dna); i++ {
		switch string(dna[i]) {
		case "G":
			rna = rna + "G"
		case "C":
			rna = rna + "C"
		case "A":
			rna = rna + "A"
		case "T":
			rna = rna + "U"
		default:
			rna = rna + ""
		}
	}
	return rna
}
