package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/steffantucker/AdventOfCode/2023/day1"
	"github.com/steffantucker/AdventOfCode/2023/day2"
	"github.com/steffantucker/AdventOfCode/2023/day3"
	"github.com/steffantucker/AdventOfCode/2023/utils"
)

func main() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
	dayFuncs := []func(){
		day1.Run,
		day2.Run,
		day3.Run,
		// day4.Run,
		// day5.Run,
		// day6.Run,
		// day7.Run,
		// day8.Run,
		// day9.Run,
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
		fmt.Print("input flag")
		_, err := utils.GetInputFile(2023, *dayToRun)
		if err != nil {
			log.Fatalf("Can't get input %e", err)
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
