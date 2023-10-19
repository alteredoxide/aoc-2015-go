// --- Day 4: The Ideal Stocking Stuffer ---
// Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for
// all the economically forward-thinking little girls and boys.
//
// To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five
// zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below)
// followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest
// positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.
//
// For example:
//
// - If your secret key is abcdef, the answer is `609043`, because the MD5 hash of
//   `abcdef609043` starts with five zeroes (`000001dbbfa...`), and it is the lowest such
//   number to do so.
// - If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash
//   starting with five zeroes is `1048970`; that is, the MD5 hash of `pqrstuv1048970`
//   looks like `000006136ef...`.

package day4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)


func partOneOrTwo(input string, prefix string) int {
    input = strings.Trim(input, "\n")
    var i int = 0
    for {
        hash := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
        encoded := hex.EncodeToString(hash[:])
        if strings.HasPrefix(encoded, prefix) {
            return i
        }
        i += 1
    }
}


func worker(input string, prefix string, done *int32, tx chan<- int) {
    input = strings.Trim(input, "\n")
    var i int = 0
    for {
        if atomic.LoadInt32(done) == 1 {
            break
        }
        hash := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
        encoded := hex.EncodeToString(hash[:])
        if strings.HasPrefix(encoded, prefix) {
            tx <- i
        }
        i += 1
    }
}


func parallelSolution(input string, prefix string, nThreads int) int {
    results := make(chan int)
    var done int32
    atomic.StoreInt32(&done, 0)
    var wg sync.WaitGroup
    // start workers
    for i := 0; i < nThreads; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            worker(input, prefix, &done, results)
        }()
    }
    answer := <-results
    atomic.StoreInt32(&done, 1)
    return answer
}
