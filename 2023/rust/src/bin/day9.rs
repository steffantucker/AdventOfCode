use itertools::Itertools;

fn main() {
    let input = aoc23::get_or_create_input(9);
    println!("part 1: {}", part1(&input));
    println!("part 2: {}", part2(&input));
}

fn part1(input: &String) -> isize {
    let lines = input.lines();
    let mut sum = 0;
    for l in lines {
        sum += OASISPredictor::new(l).predict();
    }
    sum
}

fn part2(input: &String) -> isize {
    let lines = input.lines();
    let mut sum = 0;
    for l in lines {
        sum += OASISPredictor::new(l).lookback();
    }
    sum
}

#[derive(Debug, Clone)]
struct OASISPredictor(Vec<Vec<isize>>);

impl OASISPredictor {
    fn new(initial_history: &str) -> Self {
        let i = initial_history
            .split(" ")
            .map(|n| n.parse::<isize>().expect("this should be a number"))
            .collect();

        let mut s = Self(vec![i]);
        s.generate_history();
        s
    }

    fn generate_history(&mut self) {
        while !self.generate_differences() {};
    }

    fn generate_differences(&mut self) -> bool {
        let last = self.0.last().cloned().expect("there should be something here");
        let mut diffs = Vec::new();
        let mut is_all_zeroes = true;
        for (a, b) in last.iter().tuple_windows() {
            let diff = b - a;
            if diff != 0 {
                is_all_zeroes = false;
            }
            diffs.push(b - a);
        }
        self.0.push(diffs);
        is_all_zeroes
    }

    fn predict(&self) -> isize {
        self.0
            .iter()
            .fold(0, |acc, v| v.last().unwrap() + acc)
    }

    fn lookback(&self) -> isize {
        self.0
            .iter()
            .cloned()
            .rev()
            .fold(0, |p, h| h.first().expect("there for sure should be a first value") - p)
    }
}

#[cfg(test)]
mod test {
    use crate::{part1, part2};

    const EXAMPLE: &str = r"0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45";

    #[test]
    fn test_part1() {
        assert_eq!(part1(&EXAMPLE.to_string()), 114);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&EXAMPLE.to_string()), 2);
    }
}