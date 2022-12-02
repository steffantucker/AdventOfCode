use std::env;
use std::io::prelude::*;
use std::io::BufReader;
use std::fs::File;


fn main() {
    let args: Vec<String> = env::args().collect();
    let choice = if args.len() > 1 {&args[1]} else {""};
    let debug_input = "test";
    let real_input = "input";
    let input = if choice == "test" {debug_input} else {real_input};

    let f = File::open(input).unwrap();
    let reader = BufReader::new(f);

    let mut sum = 0;
    let mut calorie_count = Vec::new();
    for line in reader.lines() {
        let calories = line.unwrap();

        if calories == "" {
            calorie_count.push(sum);
            sum = 0;
            continue;
        }
        sum += calories.parse::<i32>().unwrap();
    }

    calorie_count.sort();
    calorie_count.reverse();
    sum = calorie_count[0] + calorie_count[1] + calorie_count[2];
    println!("Max calories {}", calorie_count[0]);
    println!("Top 3: {} {} {}\n sum: {}", calorie_count[0], calorie_count[1], calorie_count[2], sum)
    
}