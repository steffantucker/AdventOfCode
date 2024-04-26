//! We make our way to the metal island and there is no one waiting. We start following signs for a hot
//! spring, but get distracted by an animal ducking into and out of pipes that are part of the landscape.
//! We make a quick sketch of the pipes we can see.

use aoc23::MapLocation;

fn main() {
    let input = aoc23::get_or_create_input(10);
    println!("part 1: {}", part1(&input))
}

/// The animal seems to be going into a loop of pipes, wonder what the farthest location is from the
/// start of the loop?
///
/// # Output
/// How many steps from the start is the farthest point.
fn part1(_input: &String) -> usize {
    0
}

#[derive(Debug, Default)]
struct Pipe {
    north: bool,
    south: bool,
    east: bool,
    west: bool,
}

impl MapLocation for Pipe {
    fn get_tile(c: char) -> Self {
        match c {
            _ => Default::default(),
        }
    }
}

#[cfg(test)]
mod test {
    use crate::part1;
    const EXAMPLE1: &str = r"-L|F7
7S-7|
L|7||
-L-J|
L|-JF";

    const EXAMPLE2: &str = r"7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ";

    #[test]
    fn test_part1() {
        assert_eq!(part1(&EXAMPLE1.to_string()), 4);
        assert_eq!(part1(&EXAMPLE2.to_string()), 8);
    }
}