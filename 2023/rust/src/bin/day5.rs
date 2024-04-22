use itertools::Itertools;
use regex::Regex;
use std::collections::HashMap;
use std::fmt;
use std::fmt::Formatter;
use std::ops::{Index, Range};

#[derive(Debug, Clone)]
struct RangeSet {
    ranges: Vec<Range<i64>>
}

impl RangeSet {
    fn new() -> Self {
        Self {
            ranges: Vec::new(),
        }
    }

    fn push(&mut self, r: Range<i64>) {
        match self.ranges.binary_search_by(|a| a.start.cmp(&r.start)) {
            Ok(v) => self.ranges.insert(v, r),
            Err(v) => self.ranges.insert(v, r),
        }
    }

    fn append(&mut self, mut other: Self) {
        self.ranges.append(&mut other.ranges);
        self.ranges.sort_by(|a, b| a.start.cmp(&b.start))
    }

    fn contains(&self, value: i64) -> bool {
        for r in &self.ranges {
            if r.contains(&value) {
                return true;
            }
        }
        false
    }
    
    fn increase_all(&mut self, increase_value: i64) {
        self.ranges = self
            .ranges
            .iter()
            .map(|value| (value.start+increase_value)..(value.end+increase_value))
            .collect();
    }

    fn inner(&self, other_ranges: &Self) -> Self {
        let mut result_ranges = Vec::new();
        for result in self.ranges.iter() {
            for other in other_ranges.ranges.iter() {
                if result.start < other.end && result.end > other.start {
                    let start = result.start.max(other.start);
                    let end = result.end.min(other.end);
                    result_ranges.push(start..end);
                }
            }
        }

        Self {
            ranges: result_ranges
        }
    }

    fn outer(&self, other_ranges: &Self) -> Self {
        let mut result_ranges = self.ranges.clone();
        for other in &other_ranges.ranges {
            let mut new_result: Vec<Range<i64>> = Vec::new();
            for result in result_ranges.iter() {
                if result.start >= other.end || result.end <= other.start {
                    new_result.push(result.clone());
                    continue;
                }
                if result.start < other.start {
                    new_result.push(result.start..other.start);
                }
                if result.end > other.end {
                    new_result.push(other.end..result.end)
                }
            }
            result_ranges = new_result;
        }

        Self {
            ranges: result_ranges
        }
    }
}

#[derive(Debug)]
struct MapRange {
    source_start: u64,
    dest_start: u64,
    range: u64,
    source_end: u64,
    _dest_end: u64,
}

#[derive(Debug)]
struct XtoYMap {
    map: Vec<MapRange>,
}

impl XtoYMap {
    fn new() -> Self {
        Self {
            map: Vec::new(),
        }
    }
    fn insert(&mut self, dest: u64, source: u64, range: u64) {
        self.map.push(MapRange {
            source_start: source,
            dest_start: dest,
            range,
            source_end: (source + range) - 1,
            _dest_end: (dest + range) - 1,
        });
    }
    fn get(&self, num: &u64) -> u64 {
        for MapRange{ source_start, dest_start, source_end , .. } in &self.map {
            if num >= source_start && num < source_end {
                return dest_start + (num - source_start)
            }
        }
        *num
    }
}

impl fmt::Display for XtoYMap {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        write!(f, "{:?}", self.map)
    }
}

struct SeedsAndMap {
    seeds: Vec<u64>,
    maps: Maps,
}

impl SeedsAndMap {
    fn new(input: &String) -> Self {
        let mut s = Self {
            seeds: Vec::new(),
            maps: Maps::new(),
        };
        s.maps.insert("seed".to_string(), XtoYMap::new());
        s.maps.insert("soil".to_string(), XtoYMap::new());
        s.maps.insert("fertilizer".to_string(), XtoYMap::new());
        s.maps.insert("water".to_string(), XtoYMap::new());
        s.maps.insert("light".to_string(), XtoYMap::new());
        s.maps.insert("temperature".to_string(), XtoYMap::new());
        s.maps.insert("humidity".to_string(), XtoYMap::new());

        let mut lines = input.lines();
        s.seeds = lines
            .next()
            .unwrap()
            .replace("seeds: ", "")
            .split(' ')
            .map(|n| n.parse::<u64>().expect("seed not a number"))
            .collect();
        let mut working_map: String = String::from("");
        for l in lines {
            let header_regex = Regex::new(r"(?<source>\w+)-to-(?<dest>\w+) map:").expect("failed to compile regex");
            if let Some(m) = header_regex.captures(l) {
                working_map = m["source"].to_string();
                continue;
            }
            if l.trim().is_empty() {
                continue;
            }
            let (dest, source, range) = l
                .split(' ')
                .map(|n| n.parse::<u64>().expect("number isn't number"))
                .collect_tuple()
                .expect("more or less numbers than expected");
            s.maps.entry(working_map.clone())
                .and_modify(|m: &mut XtoYMap| m.insert(dest, source, range));
        }
        s
    }

    fn seed_to_location(&self, seed: &u64) -> u64 {
        let soil = self.maps.index("seed").get(seed);
        let fertilizer = self.maps.index("soil").get(&soil);
        let water = self.maps.index("fertilizer").get(&fertilizer);
        let light = self.maps.index("water").get(&water);
        let temperature = self.maps.index("light").get(&light);
        let humidity = self.maps.index("temperature").get(&temperature);
        self.maps.index("humidity").get(&humidity)
    }
}
type Maps = HashMap<String, XtoYMap>;

fn part1(sm: &SeedsAndMap) -> u64 {
    /*for (k, v) in &maps {
        println!("{}\n{}", k, v);
    }*/
    let mut lowest_location = u64::MAX;
    // seed->soil->fertilizer->water->light->temperature->humidity->location
    for seed in &sm.seeds {
        let location = sm.seed_to_location(seed);
        if location < lowest_location {
            lowest_location = location;
        }
    }
    lowest_location
}

fn part2(input: &String) -> u64 {
    let mut lines = input.lines();
    // collect seed ranges
    let seeds: Vec<(i64, i64)> = lines
        .next()
        .unwrap()
        .replace("seeds: ", "")
        .split(' ')
        .map(|n| n.parse::<i64>().expect("seed not a number"))
        .tuples()
        .collect();
    let mut locations = RangeSet::new();
    for (start, range) in seeds {
        locations.push(start..(start+range));
    }
    let map_value_regex = Regex::new(r"(?<dest>\d+) (?<source>\d+) (?<range>\d+)").expect("regex failed to compile");
    // iterate through map vals
    let mut new_locations = RangeSet::new();
    for line in lines {
        if line.is_empty() {
            continue;
        }
        if let Some(caps) = map_value_regex.captures(line) {
            let dest = caps["dest"].parse::<i64>().expect("number not parsed");
            let source = caps["source"].parse::<i64>().expect("number not parsed");
            let range = caps["range"].parse::<i64>().expect("number not parsed");
            //  calc diff = dest - source
            let diff: i64 = dest - source;
            let map = RangeSet {ranges: vec![source..(source+range)]};
            //  get intersect
            let mut inner = locations.inner(&map);
            inner.increase_all(diff);
            //  get outer
            let outer = locations.outer(&map);
            //  new vals = (intersect + diff) append outer
            new_locations.append(inner);
            locations = outer;

            continue;
        } //  repeat
        locations.append(new_locations);
        new_locations = RangeSet::new();
    }
    locations.append(new_locations);
    locations.ranges.first().expect("somehow empty").start as u64
}

fn main() {
    let input = aoc23::get_or_create_input(5);
    let sm = SeedsAndMap::new(&input);
    println!("part 1: {}", part1(&sm));
    println!("part 2: {}", part2(&input));
}

#[cfg(test)]
mod test {
    use crate::{part1, part2, SeedsAndMap};
    const EXAMPLE: &str = r"seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4";
    #[test]
    fn test_part1() {
        let sm = SeedsAndMap::new(&EXAMPLE.to_string());
        assert_eq!(part1(&sm), 35);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&EXAMPLE.to_string()), 46)
    }
}