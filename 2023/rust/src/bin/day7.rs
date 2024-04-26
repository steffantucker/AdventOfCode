//! Our all-expense-paid trip to Desert Island is a 5 minute airship ride. We ride some camels with an
//! elf to Desert Island's problem (borked machines), and play a game of Camel Cards along the way. Instead of cards, we're
//! given a list of poker hands and antes and calculate the total winnings. Total winnings is the sum
//! of a hands score, which is the hand's ante times the hand's rank (after ordering the hands from best to worst).

use std::cmp::Ordering;
use std::fmt::{Display, Formatter};
use itertools::Itertools;
use crate::HandStrength::{Five, Four, Full, High, Pair, Three, TwoPair};

fn main() {
    let input = aoc23::get_or_create_input(7);
    println!("part 1: {}", solution(&input, false));
    println!("part 2: {}", solution(&input, true));
}

#[derive(Debug, Eq, PartialEq, Ord, PartialOrd)]
enum HandStrength {
    Five,
    Four,
    Full,
    Three,
    TwoPair,
    Pair,
    High,
}

fn hand_to_numbers(h: &str, wilds: bool) -> Vec<usize> {
    h
        .split("")
        .filter_map(|c| {
        if let Ok(v) = c.parse::<usize>() {
            return Some(v);
        }
        match c {
            "A" => return Some(14),
            "K" => return Some(13),
            "Q" => return Some(12),
            "J" => return if wilds { Some(1) } else { Some(11) } ,
            "T" => return Some(10),
            _ => None,
        }
    }).collect()
}

#[derive(Debug)]
struct Hand {
    hand: Vec<usize>,
    hand_letters: String,
    ante: usize,
    strength: HandStrength,
}

impl Hand {
    fn new(h: &str, a: &str, wilds: bool) -> Self {
        let strength = if wilds {
            Self::calculate_wild_strength(h)
        } else {
            Self::calculate_hand_strength(h)
        };

        Self {
            hand: hand_to_numbers(h, wilds),
            hand_letters: h.to_string(),
            ante: a.parse::<usize>().expect("ante isn't a number"),
            strength,
        }
    }

    fn calculate_wild_strength(h: &str) -> HandStrength {
        let counts_map = h
            .chars()
            .counts();
        let wild_count = *counts_map.get(&'J').or(Some(&0_usize)).unwrap();
        if wild_count == 0 {
            return Self::calculate_hand_strength(h);
        }
        let counts = counts_map
            .iter()
            .filter_map(|(k, v)| {
                if k.eq_ignore_ascii_case(&'j') {
                    return None;
                }
                Some(*v)
            })
            .sorted()
            .collect_vec();
        return match counts[..] {
            [1, 1, 1, 1] => Pair,
            [1, 1, 2] | [1,1,1] => Three,
            [2, 2] => Full,
            [1,3] | [1,2] | [1,1] => Four,
            _ => Five,
        }
    }

    fn calculate_hand_strength(h: &str) -> HandStrength {
        let counts: Vec<_> = h
            .chars()
            .counts()
            .into_values()
            .collect();
        if counts.contains(&5) {
            return Five;
        }
        if counts.contains(&4) {
            return Four;
        }
        if counts.contains(&3) && counts.contains(&2) {
            return Full;
        }
        if counts.contains(&3) {
            return Three;
        }
        if counts.iter().counts().contains_key(&2) && counts.iter().counts()[&2] == 2 {
            return TwoPair;
        }
        if counts.contains(&2) {
            return Pair;
        }
        High
    }
}

impl Display for Hand {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "Hand: {}\nStrength: {:?}\nAnte: {}", self.hand_letters, self.strength, self.ante)
    }
}

impl PartialEq<Self> for Hand {
    fn eq(&self, other: &Self) -> bool {
        self.hand == other.hand
    }
}

impl Eq for Hand {}

impl Ord for Hand {
    fn cmp(&self, other: &Self) -> Ordering {
        if self.eq(other) {
            return Ordering::Equal;
        }
        if self.strength < other.strength {
            return Ordering::Greater;
        }
        if self.strength > other.strength {
            return Ordering::Less;
        }
        for (l, r) in self.hand.iter().zip(other.hand.iter()) {
            if l < r {
                return Ordering::Less;
            }
            if l > r {
                return Ordering::Greater;
            }
        }
        Ordering::Equal
    }
}

impl PartialOrd for Hand {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        if self.eq(other) {
            return Some(Ordering::Equal);
        }
        if self.strength < other.strength {
            return Some(Ordering::Greater);
        }
        if self.strength > other.strength {
            return Some(Ordering::Less);
        }
        for (l, r) in self.hand.iter().zip(other.hand.iter()) {
            if l < r {
                return Some(Ordering::Less);
            }
            if l > r {
                return Some(Ordering::Greater);
            }
        }
        None
    }
}

/// Part 1 we want to know the total winnings.
/// Part 2 we want to know the total winnings, but using Jacks as wild (and worth the least).
fn solution(input: &String, wilds: bool) -> usize {
    let mut hands: Vec<Hand> = Vec::new();
    for l in input.lines() {
        let (hand, ante) = l.split_once(" ").expect("no space");
        hands.push(Hand::new(hand, ante, wilds));
    }
    hands.sort();
    let mut winnings = 0;
    for (i, h) in hands.iter().enumerate() {
        winnings += (i+1)*h.ante;
    }
    winnings
}

#[cfg(test)]
mod test {
    use crate::{solution};

    const EXAMPLE: &str = r"32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483";

    #[test]
    fn test_part1() {
        assert_eq!(solution(&EXAMPLE.to_string(), false), 6440);
    }

    #[test]
    fn test_part2() {
        assert_eq!(solution(&EXAMPLE.to_string(), true), 5905);
    }
}