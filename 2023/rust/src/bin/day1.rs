//! Day 1 the calibration document for the trebuchet we were loaded into had extra data.
//!
//! Part 1 had us grab the first and last roman numeral from input
//!
//! Part 2 had us grab the first and last roman numeral _or_ spelled out number.

fn main() {
    let input = aoc23::get_or_create_input(1);
    println!("part 1 {}", parse(input.clone(), false));
    println!("part 2 {}", parse(input.clone(), true));
}

fn parse(input: String, words: bool) -> u32 {
    input
    .lines()
    .filter(|line| !line.is_empty())
    .map(|line| {
        if words {
            line.to_string()
                .replace("one", "o1e")
                .replace("two", "t2o")
                .replace("three", "t3e")
                .replace("four", "f4r")
                .replace("five", "f5e")
                .replace("six", "s6x")
                .replace("seven", "s7n")
                .replace("eight", "e8t")
                .replace("nine", "n9e")
        } else {
            line.to_string()
        }
    })
    .map(|line| {
        line.chars().filter_map(|c| c.to_digit(10)).collect::<Vec<u32>>()
    })
    .map(|v| 10 * v.first().unwrap() + v.last().unwrap())
    .sum()
}