use std::fs;

fn main() {
    let path = std::env::current_dir().unwrap();
    println!("{}", path.display());
    let input = fs::read_to_string("input").unwrap();
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