use std::collections::HashMap;
use std::fs::{File, read_to_string};
use std::env;
use std::io::prelude::*;
use std::io::{BufReader};

fn main() {
    let args: Vec<String> = env::args().collect();
    let choice = if args.len() > 1 {&args[1]} else {""};
    let debug_input = "test";
    let real_input = "input";
    let input = if choice == "test" {debug_input} else {real_input};

    part1(input);
    part2(input);
}

fn itemvalues() -> HashMap<char, i32> {
    let mut item_values: HashMap<char, i32> = HashMap::new();
    for (i, c) in "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ".to_string().char_indices() {
        item_values.insert(c, (i as i32)+1);
    }
    item_values
}

fn part1(input :&str) {
    let item_values = itemvalues();
    let f = File::open(input).unwrap();
    let reader = BufReader::new(f);

    let mut priority_sum = 0;
    for line in reader.lines() {
        let packs = line.unwrap();
        let (pack1, pack2) = packs.split_at(packs.len()/2);

        for item in pack1.chars() {
            if pack2.contains(item) {
                println!("matched {:?}", item);
                priority_sum += item_values.get(&item).unwrap();
                break;
            }
        }
    }
    println!("sum {}", priority_sum)
}

fn part2(input: &str) {
    let item_values = itemvalues();
    
    let file = read_to_string(input).expect("read file error");
    let mut badge_sum = 0;

    let mut elves = file.split("\n");

    loop { 
        let elf1 = match elves.next(){
            Some(i) => i,
            None => break
        };
        println!("{}", elf1);
        let elf2 = elves.next().expect("bah1");
        println!("{}", elf2);
        let elf3 = elves.next().expect("bah2");
        println!("{}", elf3);

        for b in elf1.chars() {
            if elf2.contains(b) && elf3.contains(b) {
                badge_sum += item_values.get(&b).unwrap();
                break;
            }
        }
    }
    println!("badge sum {}", badge_sum)
}