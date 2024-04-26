//! After landing on a snow-less Snow Island in the sky, and elf challenges us to a game of cubes in a bag.
use std::{cmp, collections::HashMap};

fn main() {
    let input = aoc23::get_or_create_input(2);
    println!("part 1 {}", part1(input.clone()));
    println!("part 2 {}", part2(input.clone()))
}

/// We're given a list of games played, and asked if there are only 13 green, 12 red, and 14 blue cubes in the bag,
/// how many of the games from the list could have been played with that bag.
///
/// # Output
/// Sum of the IDs of games that are possible to have been played with the given bag.
fn part1(input: String) -> u32 {
    let green_max = 13;
    let red_max = 12;
    let blue_max = 14;
    input
    .lines()
    .filter(|line| !line.is_empty())
    .map(line_reduce)
    .filter(|game| game["green"] <= green_max)
    .filter(|game| game["red"] <= red_max)
    .filter(|game| game["blue"] <= blue_max)
    .fold(0, |acc, game| acc + game["id"])
}

/// Using the same list of games played, the elf asks us what the minimum of each colored cube is that could have
/// been used to play each game.
///
/// # Output
/// The `"power"` of a game is `(# red) * (# blue) * (# green)`, and the output is the sum of all powers of the games on the list.
fn part2(input: String) -> u32 {
    input
    .lines()
    .filter(|line| !line.is_empty())
    .map(line_reduce)
    .fold(0, |acc, game| acc + (game["green"] * game["red"] * game["blue"]))
}

fn line_reduce(line: &str) -> HashMap<&str, u32> {
    let (id, games) = line.split_once(":").unwrap();
    let mut h = HashMap::from([("id", id[5..].parse().unwrap()),("green", 0), ("red", 0), ("blue", 0)]);
    for game in games.split(";"){
        for draw in game.split(",") {
            let (c, color) = draw.trim().split_once(" ").unwrap();
            let count: u32 = c.parse().unwrap();
            match color {
                "green" => h.entry("green").and_modify(|v| *v = cmp::max(*v, count)),
                "red" => h.entry("red").and_modify(|v| *v = cmp::max(*v, count)),
                "blue" => h.entry("blue").and_modify(|v| *v = cmp::max(*v, count)),
                &_ => panic!("unknown color {color}")
            };
        }
    };
    h
}