package main

import (
	"fmt"
	"os"

	"birc.au.dk/gsa"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}
	genomeFile := os.Args[1]
	readsFile := os.Args[2]

	genome := gsa.LoadFasta(genomeFile)
	genomeSa := map[string][]int32{}
	for name, seq := range genome {
		genomeSa[name] = gsa.Sais(seq)
	}

	gsa.ScanFastq(readsFile, func(rec *gsa.FastqRecord) {
		for chrName, chrSeq := range genome {
			gsa.BSearch(rec.Read, chrSeq, genomeSa[chrName], func(i int32) {
				cigar := fmt.Sprintf("%d%s", len(rec.Read), "M")
				gsa.PrintSam(rec.Name, chrName, i, cigar, rec.Read)
			})
		}
	})
}
