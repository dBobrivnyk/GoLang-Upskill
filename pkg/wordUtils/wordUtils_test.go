package wordUtils_test

import (
	"github.com/dbobrivnykObj/GoLang-Upskill/pkg/wordUtils"
	"testing"
)

func TestAreAnagram(t *testing.T) {
	res1, res2 :=
		wordUtils.AreAnagram("listen", "silent"),
		wordUtils.AreAnagram("hand", "hans")

	if res1 != true {
		t.Errorf("Anagram == %t, want %t", res1, !res1)
	}
	if res2 != false {
		t.Errorf("Anagram == %t, want %t", res2, !res2)
	}

}

func TestIsPalindrome(t *testing.T) {
	res1, res2 :=
		wordUtils.IsPalindrome("kążąk"),
		wordUtils.IsPalindrome("samolot")

	if res1 != true {
		t.Errorf("Palindorme == %t, want %t", res1, !res1)
	}
	if res2 != false {
		t.Errorf("Palindorme == %t, want %t", res2, !res2)
	}
}
