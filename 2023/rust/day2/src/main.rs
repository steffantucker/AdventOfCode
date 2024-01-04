use std::{fs, collections::HashMap, cmp};

fn main() {
    let input = fs::read_to_string("input").unwrap();
    println!("part 1 {}", part1(input.clone()));
    println!("part 2 {}", part2(input.clone()))
}

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