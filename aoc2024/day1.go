package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
    "sort"
    "math"
)

func day2() {
    // Reads
    file, err := os.Open("inputs/1.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Scans
    scanner := bufio.NewScanner(file)

    var leftParts []int
    var rightParts []int

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, "   ")

        if len(parts) == 2 {
            leftNum, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
            rightNum, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

            if err1 == nil && err2 == nil {
                leftParts = append(leftParts, leftNum)
                rightParts = append(rightParts, rightNum)
            } else {
                // Error print in case cast goes wrong
                fmt.Println("Skipping malformed line:", line)
            }
        } else {
            // Error print in case splitting goes wrong
            fmt.Println("Skipping malformed line:", line)
        }
    }

    // Error handling for scanner
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    // Sort left and right arrays
    sort.Ints(leftParts)
    sort.Ints(rightParts)

    totalDistance := 0
    if len(leftParts) == len(rightParts) {
        fmt.Println("lengths match")
        for i := 0; i < len(leftParts) && i < len(rightParts); i++ {
            referenceValue := leftParts[i]
            referenceCount := 0
            for j := 0; j < len(rightParts); j++ {
                if rightParts[j] == referenceValue {
                    referenceCount += 1
                }
            }
            partialDistance := referenceValue * referenceCount
            totalDistance += partialDistance
        }
        
    } else {
        // Error print in case lengths does not match
        fmt.Println("Skipping distance comparison. Lenghts do not match")
    }

    // Print totalDistance
    fmt.Println("Total distance:", totalDistance)
}

func day1() {
    // Opens input file and handles error
    // Defers so it always closes file to avoid memory leaks
    file, err := os.Open("inputs/1.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Scans. Goes line by line using the NewScanner instance
    // Fun fact: Removes \n automatically
    // Another fact: scanner.Scan() returns false when there are no more lines
    scanner := bufio.NewScanner(file)

    var leftParts []int
    var rightParts []int

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, "   ")

        if len(parts) == 2 {
            leftNum, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
            rightNum, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

            if err1 == nil && err2 == nil {
                leftParts = append(leftParts, leftNum)
                rightParts = append(rightParts, rightNum)
            } else {
                // Error print in case cast goes wrong
                fmt.Println("Skipping malformed line:", line)
            }
        } else {
            // Error print in case splitting goes wrong
            fmt.Println("Skipping malformed line:", line)
        }
    }

    // Error handling for scanner
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    // Sort left and right arrays
    sort.Ints(leftParts)
    sort.Ints(rightParts)

    totalDistance := 0
    if len(leftParts) == len(rightParts) {
        fmt.Println("lengths match")
        for i := 0; i < len(leftParts) && i < len(rightParts); i++ {
            diff := int(math.Abs(float64(leftParts[i] - rightParts[i])))
            totalDistance += diff
        }
        
    } else {
        // Error print in case lengths does not match
        fmt.Println("Skipping distance comparison. Lenghts do not match")
    }

    // Print totalDistance
    fmt.Println("Total distance:", totalDistance)
}

func main() {
    day1()
    day2()
}
