package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/steffantucker/AdventOfCode/2024/go/day1"
	"github.com/steffantucker/AdventOfCode/2024/go/day2"
	"github.com/steffantucker/AdventOfCode/2024/go/day3"
	"github.com/steffantucker/AdventOfCode/2024/go/day4"
	"github.com/steffantucker/AdventOfCode/2024/go/day5"
	"github.com/steffantucker/AdventOfCode/2024/go/day6"
	"github.com/steffantucker/AdventOfCode/2024/go/day7"
	"github.com/steffantucker/AdventOfCode/2024/go/day9"
	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func main() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
	dayFuncs := []func(){
		day1.Run,
		day2.Run,
		day3.Run,
		day4.Run,
		day5.Run,
		day6.Run,
		day7.Run,
		// day8.Run,
		day9.Run,
		// day10.Run,
		// day11.Run,
		// day12.Run,
		// day13.Run,
		// day14.Run,
		// day15.Run,
		// day16.Run,
		// day17.Run,
		// day18.Run,
		// day19.Run,
		// day20.Run,
		// day21.Run,
		// day22.Run,
		// day23.Run,
		// day24.Run,
		// day25.Run,
	}
	dayToRun := flag.Int("day", len(dayFuncs), "--day 1")
	runAll := flag.Bool("all", false, "--all")
	input := flag.Bool("input", false, "get input for --day")

	flag.Parse()

	if *input {
		fmt.Printf("Getting input for day %d", dayToRun)
		_, err := utils.GetInputFile(2024, *dayToRun)
		if err != nil {
			log.Fatalf("Can't get input: %e", err)
		}
		return
	}

	if *runAll {
		for _, fn := range dayFuncs {
			fn()
		}
	} else {
		dayFuncs[*dayToRun-1]()
	}
}
