package util

type SecretString struct {
	passed map[rune]struct{}
	secret string
}

func NewSecretString(word string) *SecretString {
	return &SecretString{
		secret: word,
		passed: make(map[rune]struct{}),
	}
}

func (s *SecretString) Guess(r rune) {
	s.passed[r] = struct{}{}
}

func (s *SecretString) Get() string {
	ans := ""
	for _, c := range s.secret {
		_, in := s.passed[c]
		if !in {
			ans += "*"
		} else {
			ans += string(c)
		}
	}
	return ans
}
