package concept

// VowelFinder interface definition
type VowelFinder interface {
	FindVowels() []rune
	// Test() error
}

// MyString defintion
type MyString string

// FindVowels function
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, c := range ms {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels = append(vowels, c)
		}
	}
	return vowels
}

// FindCaps returns only the caps
func (ms MyString) FindCaps() []rune {
	return nil
}
