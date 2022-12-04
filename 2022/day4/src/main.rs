use std::env;
use std::fs::read_to_string;

fn main() {
    let args: Vec<String> = env::args().collect();
    let choice = if args.len() > 1 {&args[1]} else {""};
    let debug_input = "test";
    let real_input = "input";
    let input = if choice == "test" {debug_input} else {real_input};

    part1(input);
    part2(input);
}

fn part1(input: &str) {
    let reader = read_to_string(input).expect("can't read file");

    let elves = reader.split('\n')
                                            .map(Pairs::from_input_string)
                                            .map(|pair| pair.complete_overlap())
                                            .fold(0, |sum, o| sum + 1 * (o as i32));
    println!("overlaps {elves}");
}

fn part2(input: &str) {
    let reader = read_to_string(input).expect("can't read file");

    let elves = reader.split('\n')
                                            .map(Pairs::from_input_string)
                                            .map(|pair| pair.any_overlap())
                                            .fold(0, |sum, o| sum + 1 * (o as i32));
    println!("any overlaps {elves}");
}

#[derive(Debug, Clone, Copy)]
struct Pairs {
    a: Range,
    b: Range,
}

#[derive(Debug, Clone, Copy)]
struct Range {
    low: i32,
    high: i32,
}

impl Range {
    pub fn from_split_string(input: &str) -> Self {
        let (l, h) = input.split_once('-').expect("bad split of h/l");
        let low = l.trim().parse::<i32>().expect("bad conv of l");
        let high = h.trim().parse::<i32>().expect("bad conv of h");
        Self {
            low,
            high,
        }
    }
}

impl Pairs {
    pub fn from_input_string(input: &str) -> Self {
        let (pairs1, pairs2) = input.split_once(',').expect("bad split of pair");
        Self {
            a: Range::from_split_string(pairs1),
            b: Range::from_split_string(pairs2),
        }
    }
    pub fn complete_overlap(self) -> bool {
        (self.a.low <= self.b.low && self.a.high >= self.b.high) || (self.b.low <= self.a.low && self.b.high >= self.a.high)
    }

    pub fn any_overlap(self) -> bool {
        (self.a.low <= self.b.low && self.a.high >= self.b.low) || (self.b.low <= self.a.low && self.b.high >= self.a.low)
        || (self.a.low <= self.b.high && self.a.high >= self.b.high) || (self.b.low <= self.a.high && self.b.high >= self.a.high)
    }
}