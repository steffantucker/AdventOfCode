fn main() {
    let input = aoc23::get_or_create_input(4);
    let cards: Vec<(&str, &str)> = input.
        lines().
        map(|l| l.split_once(":").unwrap().1).
        map(|l| l.trim().split_once(" | ").unwrap()).
        collect();
    println!("part 1 {}", part1(cards.clone()));
    println!("part 2 {}", part2(cards.clone()));
}

fn part1(cards: Vec<(&str, &str)>) -> u32 {
    let mut sum: u32 = 0;

    for (a, b) in cards {
        let matches = compute_matches(a, b) as i32 - 1;
        if matches >= 0 {
            sum += 2u32.pow(matches as u32);
        }
    }
    sum
}

fn compute_matches(a: &str, b: &str) -> u32 {
    let win: Vec<&str> = a.trim().split_ascii_whitespace().collect();
    let card: Vec<&str> = b.trim().split_ascii_whitespace().collect();
    let mut matches: u32 = 0;
    for n in win {
        if card.contains(&n) {
            matches += 1;
        }
    }
    matches
}

fn part2(cards: Vec<(&str, &str)>) -> u32 {
    let mut card_counts = vec![1; cards.len()];
    for i in 1..=cards.len() {
        let (a, b) = cards[i-1];
        let self_count = card_counts[i-1];
        let matches = compute_matches(a, b);
        for j in 0..(matches as usize) {
            card_counts[i+j] += self_count;
        }
    }
    card_counts.iter().sum()
}

#[cfg(test)]
mod test {
    use crate::{part1, part2};

    const EXAMPLE: &str = r"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11";
    #[test]
    fn test_part1() {
        assert_eq!(part1(EXAMPLE.to_string()), 13);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(EXAMPLE.to_string()), 30);
    }
}