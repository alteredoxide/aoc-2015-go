package day2

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
)


func fromBytes(data []byte) ([][3]int, error) {
    strData := string(data)
    var out [][3]int
    for _, dimStr := range strings.Split(strData, "\n") {
        if dimStr == "" {
            continue
        }
        var dims [3]int
        parts := strings.Split(dimStr, "x")
        if len(parts) != 3 {
            return nil, fmt.Errorf("malformed dimensions: %s", dimStr)
        }
        for i, dim := range parts {
            d, err := strconv.Atoi(dim)
            if err != nil {
                return nil, err
            }
            dims[i] = d
        }
        out = append(out, dims)
    }
    return out, nil
}

func loadInput() ([][3]int, error) {
    _, file, _, _ := runtime.Caller(0)
    path := filepath.Dir(file)
    data, err := os.ReadFile(filepath.Join(path, "input.txt"))
    if err != nil {
        log.Fatal(err)
    }
    return fromBytes(data)
}


func TestPart1(t *testing.T) {
    input, err := loadInput()
    if err != nil {
        t.Error(err)
    }
    output := part1(input)
    if output != 1606483 {
        t.Errorf("expected 1606483 but got %d", output)
    }
}


func TestPart2(t *testing.T) {
    input, err := loadInput()
    if err != nil {
        t.Error(err)
    }
    output := part2(input)
    expected := 3842356
    if output != expected {
        t.Errorf("expected %d but got %d", expected, output)
    }
}
