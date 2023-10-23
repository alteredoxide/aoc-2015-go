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
//
// --- Part Two ---
// Realizing the error of his ways, Santa has switched to a better model of determining
// whether a string is naughty or nice. None of the old rules apply, as they are all
// clearly ridiculous.
// 
// Now, a nice string is one with all of the following properties:
// 
// - It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
// - It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
// 
// For example:
// 
// - `qjhvhtzxzqqjkmpb` is nice because is has a pair that appears twice (`qj`) and a letter
//   that repeats with exactly one letter between them (`zxz`).
// - `xxyxx` is nice because it has a pair that appears twice and a letter that repeats
//   with one between, even though the letters used by each rule overlap.
// - `uurcxstgmygtbstg` is naughty because it has a pair (`tg`) but no repeat with a
//   single letter between them.
// - `ieodomkazucvgmuy` is naughty because it has a repeating letter with one between
//   (`odo`), but no pair that appears twice.
//
// How many strings are nice under these new rules?

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


// ==== Part Two

func pairTwiceWithoutOverlap(s string) bool {
    pairs := map[string]bool{}
    for i := 0; i < len(s) - 1; i++ {
        pair := s[i:i+2]
        if i > 0 {
            if s[i-1:i+1] == pair {
                continue
            }
        }
        if pairs[pair] {
            return true
        }
        pairs[pair] = true
    }
    return false
}


func hasSandwich(s string) bool {
    for i := 0; i < len(s) - 2; i++ {
        triple := s[i:i+3]
        if triple[0] == triple[2] {
            return true
        }
    }
    return false
}


func partTwo(input []string) int {
    var nNice int = 0
    for _, s := range input {
        if pairTwiceWithoutOverlap(s) && hasSandwich(s) {
            nNice += 1
        }
    }
    return nNice
}
