use std::fs::File;
use std::env;
use std::io::prelude::*;
use std::io::BufReader;

fn main() {
    let args: Vec<String> = env::args().collect();
    let choice = if args.len() > 1 {&args[1]} else {""};
    let debug_input = "test";
    let real_input = "input";
    let input = if choice == "test" {debug_input} else {real_input};

    part1(input);
    part2(input);
}

fn part1(input :&str) {
    let f = File::open(input).unwrap();
    let reader = BufReader::new(f);

    let mut score = 0;

    for line in reader.lines() {
        let guide = line.unwrap();
        let (theirs, mut mine) = guide.split_at(1);
        mine = mine.trim();
        match theirs {
            "A" => score += rock(mine),
            "B" => score += paper(mine),
            "C" => score += scissors(mine),
            _ => print!("bad input")
        }
    }
    println!("score {}", score);
}

fn part2(input :&str) {
    let f = File::open(input).unwrap();
    let reader = BufReader::new(f);

    let mut score = 0;

    for line in reader.lines() {
        let guide = line.unwrap();
        let (theirs, mut mine) = guide.split_at(1);
        mine = mine.trim();
        match theirs {
            "A" => score += true_rock(mine),
            "B" => score += true_paper(mine),
            "C" => score += true_scissors(mine),
            _ => print!("bad input")
        }
    }
    println!("true score {}", score);
}

fn rock(mine: &str) -> i32 {
    match mine {
        "X" => {
            return 4
        },
        "Y" => {
            return 8
        },
        "Z" => {
            return 3
        },
        _ => println!(" bad input")
    }
    return 0
}

fn paper(mine: &str) -> i32 {
    match mine {
        "X" => {
            return 1
        },
        "Y" => {
            return 5
        },
        "Z" => {
            return 9
        },
        _ => println!(" bad input")
    }
    return 0
}

fn scissors(mine: &str) -> i32 {
    match mine {
        "X" => {
            return 7
        },
        "Y" => {
            return 2
        },
        "Z" => {
            return 6
        },
        _ => println!(" bad input")
    }
    return 0
}
fn true_rock(mine: &str) -> i32 {
    match mine {
        "X" => {
            return 3
        },
        "Y" => {
            return 4
        },
        "Z" => {
            return 8
        },
        _ => println!(" bad input")
    }
    return 0
}

fn true_paper(mine: &str) -> i32 {
    match mine {
        "X" => {
            return 1
        },
        "Y" => {
            return 5
        },
        "Z" => {
            return 9
        },
        _ => println!(" bad input")
    }
    return 0
}

fn true_scissors(mine: &str) -> i32 {
    match mine {
        "X" => {
            return 2
        },
        "Y" => {
            return 6
        },
        "Z" => {
            return 7
        },
        _ => println!(" bad input")
    }
    return 0
}
