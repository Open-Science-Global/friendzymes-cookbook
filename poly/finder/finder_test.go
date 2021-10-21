package finder

import (
	"fmt"
)

func ExampleForbiddenSequencesFinder() {
	sequence := "AAAAAATCGGTCGTAAAAAATT"
	var functions []func(string) []Match
	functions = append(functions, ForbiddenSequence([]string{"AAAAAA"}))

	problems := Find(sequence, functions)
	fmt.Println(problems)
	// Output: [{0 6 Forbidden sequence: AAAAAA} {14 20 Forbidden sequence: AAAAAA}]
}

func ExampleMatchSequences() {
	sequence := "AAAAAATCGGTCGTAAGGTCTCAAAATTGAGACC"
	var functions []func(string) []Match
	functions = append(functions, MatchSequences(map[string]string{"GGTCTC": "BsaI restriction binding site"}))

	problems := Find(sequence, functions)
	fmt.Println(problems)
	// Output: [{16 22 BsaI restriction binding site | GGTCTC} {28 34 BsaI restriction binding site | GAGACC}]
}

func ExampleRemoveRepeat() {
	sequence := "AAAAAATCGGTCGTAAGGTCTCAAAATTGAGACC"
	var functions []func(string) []Match
	functions = append(functions, RemoveRepeat(5))

	problems := Find(sequence, functions)
	fmt.Println(problems)
	// Output: [{1 6 Repeated sequence | AAAAA} {22 27 Repeated sequence | AAAAT}]

}

func ExampleGlobalRemoveRepeat() {
	sequence := "ATGAGTATTCAACATTTCCGTGTCGCCCTTATTCCCTTTTTTGCGGCATTTTGCCTTCCTGTTTTTGCTCACCCAGAAACGCTGGTGAAAGTAAAAGATGCTGAAGATCAGTTGGGTGCACGAGTGGGTTACATCGAACTGGATCTCAACAGCGGTAAGATCCTTGAGAGTTTTCGCCCCGAAGAACGTTTTCCAATGATGAGCACTTTTAAAGTTCTGCTATGTGGCGCGGTATTATCCCGTATTGACGCCGGGCAAGAGCAACTCGGTCGCCGCATACACTATTCTCAGAATGACTTGGTTGAGTACTCACCAGTCACAGAAAAGCATCTTACGGATGGCATGACAGTAAGAGAATTATGCAGTGCTGCCATAACCATGAGTGATAACACTGCGGCCAACTTACTTCTGACAACGATCGGAGGACCGAAGGAGCTAACCGCTTTTTTGCACAACATGGGGGATCATGTAACTCGCCTTGATCGTTGGGAACCGGAGCTGAATGAAGCCATACCAAACGACGAGCGTGACACCACGATGCCTGTAGCAATGGCAACAACGTTGCGCAAACTATTAACTGGCGAACTACTTACTCTAGCTTCCCGGCAACAATTAATAGACTGGATGGAGGCGGATAAAGTTGCAGGACCACTTCTGCGCTCGGCCCTTCCGGCTGGCTGGTTTATTGCTGATAAATCTGGAGCCGGTGAGCGTGGGTCTCGCGGTATCATTGCAGCACTGGGGCCAGATGGTAAGCCCTCCCGTATCGTAGTTATCTACACGACGGGGAGTCAGGCAACTATGGATGAACGAAATAGACAGATCGCTGAGATAGGTGCCTCACTGATTAAGCATTGGTAA"
	var functions []func(string) []Match
	functions = append(functions, GlobalRemoveRepeat(33, "ATGAGTATTCAACATTTCCGTGTCGCCCTTATT"))

	problems := Find(sequence, functions)
	fmt.Println(problems)
	// Output: [{0 33 Global repeated sequence | ATGAGTATTCAACATTTCCGTGTCGCCCTTATT}]
}

func ExampleHarpinSequencesFinder() {
	sequence := "ATGAGTATTCAACATTTCCGTGTCGCCCTTATTCCCTTTTTTGCGGCATACGGAAATGTTGAATACTCATTTTGCCTTCCTGTTTTTGCTCACCCAGAAACGCTGGTGAAAGTAAAAGATGCTGAAGATCAGTTGGGTGCACGAGTGGGTTACATCGAACTGGATCTCAACAGCGGTAAGATCCTTGAGAGTTTTCGCCCCGAAGAACGTTTTCCAATGATGAGCACTTTTAAAGTTCTGCTATGTGGCGCGGTATTATCCCGTATTGACGCCGGGCAAGAGCAACTCGGTCGCCGCATACACTATTCTCAGAATGACTTGGTTGAGTACTCACCAGTCACAGAAAAGCATCTTACGGATGGCATGACAGTAAGAGAATTATGCAGTGCTGCCATAACCATGAGTGATAACACTGCGGCCAACTTACTTCTGACAACGATCGGAGGACCGAAGGAGCTAACCGCTTTTTTGCACAACATGGGGGATCATGTAACTCGCCTTGATCGTTGGGAACCGGAGCTGAATGAAGCCATACCAAACGACGAGCGTGACACCACGATGCCTGTAGCAATGGCAACAACGTTGCGCAAACTATTAACTGGCGAACTACTTACTCTAGCTTCCCGGCAACAATTAATAGACTGGATGGAGGCGGATAAAGTTGCAGGACCACTTCTGCGCTCGGCCCTTCCGGCTGGCTGGTTTATTGCTGATAAATCTGGAGCCGGTGAGCGTGGGTCTCGCGGTATCATTGCAGCACTGGGGCCAGATGGTAAGCCCTCCCGTATCGTAGTTATCTACACGACGGGGAGTCAGGCAACTATGGATGAACGAAATAGACAGATCGCTGAGATAGGTGCCTCACTGATTAAGCATTGGTAA"
	var functions []func(string) []Match
	functions = append(functions, AvoidHairpin(20, 200))

	problems := Find(sequence, functions)
	fmt.Println(problems)

	// Output: Hello
}