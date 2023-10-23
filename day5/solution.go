// --- Day 5: Doesn't He Have Intern-Elves For This? ---
// Santa needs help figuring out which strings in his text file are naughty or nice.
//
// A nice string is one with all of the following properties:
//
// - It contains at least three vowels (`aeiou` only), like `aei`, `xazegov`, or
//   `aeiouaeiouaeiou`.
// - It contains at least one letter that appears twice in a row, like `xx`, `abcdde`
//   (`dd`), or `aabbccdd` (`aa`, `bb`, `cc`, or `dd`).
// - It does not contain the strings `ab`, `cd`, `pq`, or `xy`, even if they are part of
//   one of the other requirements.
//
// For example:
//
// - `ugknbfddgicrmopn` is nice because it has at least three vowels (`u...i...o...`), a
//   double letter (`...dd...`), and none of the disallowed substrings.
// - `aaa` is nice because it has at least three vowels and a double letter, even though
//   the letters used by different rules overlap.
// - `jchzalrnumimnmhp` is naughty because it has no double letter.
// - `haegwjzuvuyypxyu` is naughty because it contains the string `xy`.
// - `dvszwmarrgswjxmb` is naughty because it contains only one vowel.
//
// How many strings are nice?

package day5

import (
	"strings"
)


func hasThreeVowels(s string) bool {
    var n_vowels int = 0
    const VOWELS = "aeiou"
    for _, c := range s {
        if strings.ContainsRune(VOWELS, c) {
            n_vowels += 1
            if n_vowels == 3 {
                return true
            }
        }
    }
    return false
}


func hasDouble(s string) bool {
    var previous rune
    for i, c := range s {
        if i == 0 {
            previous = c
            continue
        }
        if c == previous {
            return true
        }
        previous = c
    }
    return false
}


func hasNaughtySeq(s string) bool {
    naughty := []string{"ab", "cd", "pq", "xy"}
    for i := 0; i < len(s) - 1; i++ {
        substr := s[i:i+2]
        for _, double := range naughty {
            if substr == double {
                return true
            }
        }
     }
     return false
}


func partOne(input []string) int {
    var nNice int = 0
    for _, s := range input {
        if hasThreeVowels(s) && hasDouble(s) && !hasNaughtySeq(s) {
            nNice += 1
        }
    }
    return nNice
}
