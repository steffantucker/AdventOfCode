use std::collections::HashMap;
use std::str::FromStr;
use anyhow::anyhow;
use num_integer::Integer;

fn main() {
    let input = aoc23::get_or_create_input(8);
    println!("part 1: {}", part1(&input));
    println!("part 2: {}", part2(&input));
}

fn part1(input: &String) -> usize {
    let m: Map = input.parse().expect("can't parse input");
    m.steps(Node::START, Node::is_end)
}

#[derive(Copy, Clone)]
enum Directions {
    Left,
    Right,
}

#[derive(Clone, Copy, Eq, Hash, PartialEq)]
struct Node(u8, u8, u8);

impl Node {
    const START: Node = Node(b'A',b'A',b'A');
    const END: Node = Node(b'Z',b'Z',b'Z');

    fn is_end(n: Node) -> bool {
        n == Self::END
    }

    fn is_ghost_start(n: Node) -> bool {
        n.2 == b'A'
    }

    fn is_ghost_end(n: Node) -> bool {
        n.2 == b'Z'
    }
}

impl FromStr for Node {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let [a, b, c] = s.as_bytes() else { return Err(anyhow!("not the right amount of bytes")); };
        Ok(Self(*a, *b, *c))
    }
}

struct Instructions(Vec<Directions>);

impl Instructions {
    fn cycle(&self) -> impl Iterator<Item=Directions> + '_ {
        self.0.iter().copied().cycle()
    }
}

impl FromStr for Instructions {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let i = s
            .chars()
            .map(|c| match c {
                    'L' => Ok(Directions::Left),
                    'R' => Ok(Directions::Right),
                    _ => Err(anyhow!("unknown instruction {}", c))
            })
            .collect::<Result<Vec<_>, Self::Err>>()?;
        Ok(Self(i))
    }
}

struct Map {
    instructions: Instructions,
    nodes: HashMap<Node, (Node, Node)>,
}

impl Map {
    fn steps<F>(&self, start: Node, end: F) -> usize
    where
        F: Fn(Node) -> bool
    {
        let mut steps = 0;
        let mut cycle = self.instructions.cycle();
        let mut current_node = start;
        while !end(current_node) {
            let (left, right) = self.nodes[&current_node];
            match cycle.next() {
                Some(Directions::Left) => current_node = left,
                Some(Directions::Right) => current_node = right,
                _ => panic!("bad instruction"),
            }
            steps += 1;
        }

        steps
    }

    fn ghost_steps(&self) -> usize {
        self
            .nodes
            .keys()
            .copied()
            .filter_map(|n| {
                if Node::is_ghost_start(n) {
                    Some(self.steps(n, Node::is_ghost_end))
                } else {
                    None
                }
            })
            .fold(1, |acc, b| b.lcm(&acc))
    }
}

impl FromStr for Map {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut lines = s.lines();

        let instructions = lines
            .next()
            .ok_or(anyhow!("no instructions"))?
            .parse::<Instructions>()?;
        lines.next();

        let mut n = HashMap::new();
        for l in lines {
            n.insert(l[0..3].parse()?, (l[7..10].parse()?, l[12..15].parse()?));
        }

        Ok(Self{ instructions, nodes: n })
    }
}

fn part2(input: &String) -> usize {
    let m: Map = input.parse().expect("failed parsing input");
    m.ghost_steps()
}

#[cfg(test)]
mod test {
    use crate::{part1, part2};

    const EXAMPLE1: &str = r"RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)";

    const EXAMPLE2: &str = r"LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)";

    const EXAMPLE3: &str = r"LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)";

    #[test]
    fn test_part1() {
        assert_eq!(part1(&EXAMPLE1.to_string()), 2);
        assert_eq!(part1(&EXAMPLE2.to_string()), 6);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&EXAMPLE3.to_string()), 6);
    }
}