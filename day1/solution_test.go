package day1

import "testing"

func TestPart1(t *testing.T) {
    input := loadInput()
    output := part1(input)
    if output != 232 {
        t.Errorf("Expected 232, got %d", output)
    }
}
