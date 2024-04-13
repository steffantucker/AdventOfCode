use std::{collections::HashMap, fs};

struct Grid {
    grid: Vec<Vec<char>>,
    width: usize,
    height: usize,
    parts: Vec<PartNumber>,
    parts_map: HashMap<Coord, u32>,
    asterisks: Vec<Coord>,
}

impl Grid {
    fn new(lines: String) -> Self {
        let grid: Vec<Vec<char>> = lines.lines().map(|line| line.chars().collect()).collect();
        let mut g = Grid { 
            grid: grid.clone(), 
            width: grid[0].len(), 
            height: grid.len(), 
            parts: Vec::new(),
            parts_map: HashMap::new(),
            asterisks: Vec::new(),
        };
        g.collect_parts_and_gears();
        g.map_parts();
        g
    }

    fn get(&self, c: Coord) -> char {
        self.grid[c.y][c.x]
    }

    fn find_gears(&self) -> u32 {
        let mut sum = 0;
        let offsets = [-1, 0, 1];
        for c in self.asterisks.as_slice() {
            let mut parts: Vec<u32> = Vec::new();
            for y_off in offsets {
                let y = c.y as isize + y_off;
                if y < 0 || y >= self.height as isize {
                    continue;
                }
                for x_off in offsets {
                    let x = c.x as isize + x_off;
                    if x < 0 || x >= self.width as isize {
                        continue;
                    }
                    if self.parts_map.contains_key(&Coord { x: x as usize, y: y as usize }) {
                        let p = self.parts_map.get(&Coord { x: x as usize, y: y as usize }).unwrap();
                        parts.push(*p);
                    }
                }
            }
            parts.dedup();
            if parts.len() == 2 {
                sum += parts[0] * parts[1];
            }
        }
        sum
    }

    fn map_parts(&mut self) {
        for part in self.parts.as_slice() {
            for x in part.coord.x..(part.coord.x+part.len) {
                self.parts_map.insert(Coord { x, y: part.coord.y }, part.number);
            }
        }
    }

    fn collect_parts_and_gears(&mut self) {
        for (y, row) in self.grid.iter().enumerate() {
            let mut part: PartNumber = PartNumber::new();
            for (x, c) in row.iter().enumerate() {
                if c.is_digit(10) {
                    part.push(Coord {x, y}, c.to_digit(10).unwrap())
                } else if !part.is_empty() {
                    self.parts.push(part);
                    part = PartNumber::new();
                }
                if c == &'*' {
                    self.asterisks.push(Coord{x,y});
                }
            }
            if !part.is_empty() {
                self.parts.push(part);
            }
        }
    }

    fn bounded_coord(&self, x: isize, y: isize) -> Coord {
        Coord{ 
            x: if x < 0 {
                0
            } else if x >= self.height as isize {
                self.height - 1
            } else {
                x as usize
            },
            y: if y < 0 {
                0
            } else if y >= self.width as isize {
                self.width - 1
            } else {
                y as usize
            }
        }
    }
}

#[derive(Clone, Copy, PartialEq, Eq, Hash)]
struct Coord {
    x: usize,
    y: usize,
}

struct PartNumber {
    coord: Coord,
    len: usize,
    number: u32,
}

impl PartNumber {
    fn new() -> Self {
        PartNumber {
            coord: Coord { x: 0, y: 0 },
            len: 0,
            number: 0,
        }
    }

    fn push(&mut self, c: Coord, n: u32) {
        if self.len == 0 {
            self.coord = c;
            self.len = 1;
            self.number = n;
            return;
        }
        self.len += 1;
        self.number = self.number * 10 + n;
    }

    fn is_empty(&self) -> bool {
        self.len == 0
    }
}

fn main() {
    let input = aoc23::get_or_create_input(3);
    println!("part 1 {}", part1(input.clone()));
    println!("part 2 {}", part2(input));
}

fn part1(input: String) -> u32 {
    let schematic: Grid = Grid::new(input);
    let mut sum = 0;
    'parts: for part in schematic.parts.as_slice() {
        for y_off in [-1, 0, 1] {
            let y = (part.coord.y as isize) + y_off;
            if y < 0 {
                continue;
            }
            for x in ((part.coord.x as isize)-1)..=((part.coord.x+part.len) as isize) {
                if x < 0 {
                    continue;
                }
                let coord = schematic.bounded_coord(x as isize, y);
                let char = schematic.get(coord);
                if char.is_digit(10) {
                    continue;
                }
                if char != '.' {
                    sum += part.number;
                    continue 'parts;
                }
            }
        }
    }

    sum
}

fn part2(input: String) -> u32 {
    let schematic = Grid::new(input);
    schematic.find_gears()
}

#[cfg(test)]
mod tests {
    use crate::{part1, part2};
        const EXAMPLE: &str = 
r"467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..";

    #[test]
    fn test_part1() {
        assert_eq!(part1(EXAMPLE.to_string()), 4361);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(EXAMPLE.to_string()), 467835);
    }
}