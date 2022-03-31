package util

import "os"

func GetOrDefault(key, defaultValue string) string {
	ans := os.Getenv(key)
	if ans == "" {
		ans = defaultValue
	}
	return ans
}

func UniqueCharacters(s string) map[rune]struct{} {
	ans := make(map[rune]struct{})
	for _, chr := range s {
		ans[chr] = struct{}{}
	}
	return ans
}
