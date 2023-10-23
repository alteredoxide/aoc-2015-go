// --- Day 6: Probably a Fire Hazard ---
// Because your neighbors keep defeating you in the holiday house decorating contest year
// after year, you've decided to deploy one million lights in a 1000x1000 grid.
//
// Furthermore, because you've been especially nice this year, Santa has mailed you
// instructions on how to display the ideal lighting configuration.
//
// Lights in your grid are numbered from 0 to 999 in each direction; the lights at each
// corner are at `0,0`, `0,999`, `999,999`, and `999,0`. The instructions include whether
// to `turn on`, `turn off`, or `toggle` various inclusive ranges given as coordinate
// pairs. Each coordinate pair represents opposite corners of a rectangle, inclusive; a
// coordinate pair like `0,0` through `2,2` therefore refers to 9 lights in a 3x3 square.
// The lights all start turned off.
//
// To defeat your neighbors this year, all you have to do is set up your lights by doing
// the instructions Santa sent you in order.
//
// For example:
//
// - `turn on 0,0 through 999,999` would turn on (or leave on) every light.
// - `toggle 0,0 through 999,0` would toggle the first line of 1000 lights, turning off
//    the ones that were on, and turning on the ones that were off.
// - `turn off 499,499 through 500,500` would turn off (or leave off) the middle four
//    lights.
//
// After following the instructions, how many lights are lit?

package day6

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)


type Action int
const (
    Off Action = iota
    On
    Toggle
)


type Instruction struct {
    action Action
    coords [2][2]int
}


func instructionFromLine(line string) (Instruction, error) {
    var action Action
    var out = Instruction{}
    // action
    switch {
    case strings.HasPrefix(line, "turn on"):
        action = On
    case strings.HasPrefix(line, "turn off"):
        action = Off
    case strings.HasPrefix(line, "toggle"):
        action = Toggle
    default:
        return Instruction{}, errors.New("malformed string: no action")
    }
    out.action = action
    // coords
    re := regexp.MustCompile(`.*\s(\d+),(\d+) through (\d+),(\d+).*`)
    matches := re.FindStringSubmatch(line)
    if len(matches) < 5 {
        return out, errors.New("malformed string: not enough coords")
    }
    var coords = [2][2]int{}
    for n, x := range matches[1:] {
        var i int
        if n > 1 {
            i = 1
        } else {
            i = 0
        }
        j := n % 2
        value, err := strconv.Atoi(x)
        if err != nil {
            return out, errors.New("malformed string: unable to parse coord to int")
        }
        coords[i][j] = value
    }
    out.coords = coords
    return out, nil
}


func instructionsFromInput(lines []string) []Instruction {
    var out = []Instruction{}
    for i, line := range lines {
        if  line == "" { continue }
        instruct, err := instructionFromLine(line)
        if err != nil {
            fmt.Printf("%s on line num: %d\n", err.Error(), i+1)
        }
        out = append(out, instruct)
    }
    return out
}


func partOne(input []string) int {
    instructions := instructionsFromInput(input)
    // init the lights array with all off
    var lights = [1000][1000]bool{}
    for i := 0; i < 999; i++ {
        for j := 0; j < 999; j++ {
            lights[i][j] = false
        }
    }
    // ops
    for _, inst := range instructions {
        rowI := inst.coords[0][0]
        rowF := inst.coords[1][0]
        colI := inst.coords[0][1]
        colF := inst.coords[1][1]
        for r := rowI; r <= rowF; r++ {
            for c := colI; c <= colF; c++ {
                switch inst.action {
                case Off:
                    lights[r][c] = false
                case On:
                    lights[r][c] = true
                case Toggle:
                    lights[r][c] = !lights[r][c]
                }
            }
        }
    }
    count := 0
    for _, row := range lights {
        for _, x := range row {
           if x { count += 1 } 
        }
    }
    return count
}
