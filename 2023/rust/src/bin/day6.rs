//! After the ferry ride, turns out there's no sand where the sand is supposed to be. The sand comes from
//! another island in the sky, the aptly-named Desert Island. Luckily, a local boat race is offering and all-expense
//! paid trip to Desert Island. The boats are toys though, and you power them by holding a button to increase the
//! acceleration of the boat, and must win by going the furthest by a given time.

use itertools::Itertools;

fn main() {
    let input = aoc23::get_or_create_input(6);
    println!("part 1: {}", part1(&input));
    println!("part 2: {}", part2(&input));
}

fn get_nums(v: &str) -> Vec<u64> {
    v
        .split_once(":")
        .unwrap().1
        .trim()
        .split(" ")
        .filter_map(|v| v.parse::<u64>().ok())
        .collect()
}

/// There are multiple time-values you can hold the button and win the race.
///
/// # Output
/// The margin-of-error for our calculations is the product of the number of time-values for each race.
fn part1(input: &String) -> u64 {
    let mut lines = input.lines();
    let times = get_nums(lines.next().unwrap());
    let dists = get_nums(lines.next().unwrap());
    times
        .iter()
        .zip(dists.iter())
        .map(|(time,dist)| (1..*time).map(|t| t*(time-t)).filter(|d| d > dist).count())
        .fold(1, |acc, v| acc * v as u64 )
}

fn line_to_number(l: &str) -> u64 {
    l.split_once(":").unwrap().1.split(" ").join("").parse::<u64>().expect("not a number??")
}

/// Oops, we've once again misread our input. Turns out there's only 1 race and not multiple, our paper just
/// has bad kerning.
///
/// # Output
/// Output is still the number of time-values we can use to win, but now just for the 1 race.
fn part2(input: &String) -> u64 {
    let mut lines = input.lines();
    let time = line_to_number(lines.next().unwrap());
    let dist = line_to_number(lines.next().unwrap());
    (1..time).map(|t| t*(time-t)).filter(|&d| d > dist).count() as u64
}

#[cfg(test)]
mod test {
    use crate::{part1, part2};

    const EXAMPLE: &str = r"Time:      7  15   30
Distance:  9  40  200";
    #[test]
    fn test_part1() {
        assert_eq!(part1(&EXAMPLE.to_string()), 288);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&EXAMPLE.to_string()), 71503);
    }
}