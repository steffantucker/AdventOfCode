//! The gondola lifts us up to Island Island, another island in the sky. Now, we must make our way to a normal
//! island on a lake on Island Island. We need a boat to get there, and an elf at the gondola platform will
//! let us borrow theirs if we help them figure out what they won from their mountain of scratch cards.

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

/// Each scratch cards has a list of winning numbers and a list of numbers we have. First match between the
/// 2 lists is 1 point, subsequent matches double the points.
///
/// # Output
/// Outputs the sum of the points of all cards
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

/// Turns out the actual winning instructions were on the back of the cards and are weird and obtuse.
/// You win more cards based on how many matches you have on the cards, e.g. if card 10 had 5 matches, you
/// win 1 copy of each card: 11, 12, 13, 14, 15. The same rule applies for every copy
///
/// # Output
/// Since each card only wins more cards, output is the total number of cards we have after calculating all wins.
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

    const EXAMPLE: Vec<(&str, &str)> = r"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11".
        lines().
        map(|l| l.split_once(":").unwrap().1).
        map(|l| l.trim().split_once(" | ").unwrap()).
        collect();
    #[test]
    fn test_part1() {
        assert_eq!(part1(EXAMPLE), 13);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(EXAMPLE), 30);
    }
}