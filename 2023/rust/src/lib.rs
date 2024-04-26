//! Helper and common-use functions for solving Advent of Code problems.
//! Only get_or_create_input() has been used or tested in any puzzle solution yet.

use std::{env, fs};
use anyhow::anyhow;
use minreq;

/// Downloads the input for the given day from the Advent of Code website.
///
/// Relies on the `adventofcodecookie` environment variable, which is the session cookie from the website.
/// First checks to see if the input has already been downloaded, and returns that. Otherwise, downloads then saves the
/// input and returns it.
/// Currently hardcoded for 2023 puzzles.
///
/// # Panics
/// Panics if the return code is not `200 OK`, probably means the env variable is incorrect of formatted improperly.
pub fn get_or_create_input(day: usize) -> String {
    let path = format!("inputs/day{}input", day);
    if let Ok(data) = fs::read_to_string(path.as_str()) {
        return data;
    }
    let session_key = env::var("adventofcodecookie").expect("AoC env var not set/found");
    let body = minreq::get(format!("https://adventofcode.com/2023/day/{}/input", day))
        .with_header("Cookie", format!("session={}", session_key))
        .send()
        .expect("problem GETting");
    if body.status_code == 200 {
        let input = body.as_str().expect("body not string??");
        let mkdir = fs::create_dir("inputs");
        let err = mkdir.err().unwrap();
        if err.to_string().contains("exits") {
            panic!("unexpected create directory error: {}", err);
        }
        let _ = fs::write(path.as_str(), input).expect("error writing to file");
        return input.to_string();
    } else {
        panic!("GET return status code: {}", body.status_code);
    }
}

/// Implement MapLocation trait to transform a character in the input into a user-defined data type.
pub trait MapLocation {
    fn get_tile(c: char) -> Self;
}

/// Cardinal directions for traversing or referencing direction on a Map
pub enum MapDirections {
    North,
    South,
    East,
    West,
}

/// Simple location reference for a Map
#[derive(Debug, Copy, Clone)]
pub struct MapPosition {
    x: usize,
    y: usize,
}

/// Provides a way to traverse a Map
#[derive(Debug)]
pub struct MapWalker<T> {
    map: Map<T>,
    location: MapPosition,
}

impl<T> MapWalker<T>
    where T: MapLocation
{
    /// Creates a new MapWalker starting at a given position, creating the map from input.
    pub fn new(input: &str, start: MapPosition) -> MapWalker<T> {
        let map = Map::new(input);
        Self {
            map,
            location: start,
        }
    }

    /// Move 1 step in a cardinal direction
    pub fn move_direction(&mut self, d: MapDirections) -> Result<&T, anyhow::Error> {
        match d {
            MapDirections::North => self.move_north(),
            MapDirections::South => self.move_south(),
            MapDirections::East => self.move_east(),
            MapDirections::West => self.move_west(),
        }
    }

    pub fn move_north(&mut self) -> Result<&T, anyhow::Error> {
        let new_pos = MapPosition{x: self.location.x, y: self.location.y-1};
        let new_tile = self.map.get(&new_pos);
        if new_tile.is_ok() {
            self.location = new_pos;
        }
        new_tile
    }
    pub fn move_south(&mut self) -> Result<&T, anyhow::Error> {
        let new_pos = MapPosition{x: self.location.x, y: self.location.y+1};
        let new_tile = self.map.get(&new_pos);
        if new_tile.is_ok() {
            self.location = new_pos;
        }
        new_tile
    }
    pub fn move_east(&mut self) -> Result<&T, anyhow::Error> {
        let new_pos = MapPosition{x: self.location.x+1, y: self.location.y};
        let new_tile = self.map.get(&new_pos);
        if new_tile.is_ok() {
            self.location = new_pos;
        }
        new_tile
    }
    pub fn move_west(&mut self) -> Result<&T, anyhow::Error> {
        let new_pos = MapPosition{x: self.location.x-1, y: self.location.y};
        let new_tile = self.map.get(&new_pos);
        if new_tile.is_ok() {
            self.location = new_pos;
        }
        new_tile
    }

    /// Returns data in the cardinal direction without moving to it.
    pub fn peek_direction(&mut self, d: MapDirections) -> Result<&T, anyhow::Error> {
        match d {
            MapDirections::North => self.peek_north(),
            MapDirections::South => self.peek_south(),
            MapDirections::East => self.peek_east(),
            MapDirections::West => self.peek_west(),
        }
    }

    pub fn peek_north(&self) -> Result<&T, anyhow::Error> {
        let new_pos = MapPosition{x: self.location.x, y: self.location.y-1};
        self.map.get(&new_pos)
    }
    pub fn peek_south(&self) -> Result<&T, anyhow::Error> {
        let new_pos = MapPosition{x: self.location.x, y: self.location.y+1};
        self.map.get(&new_pos)
    }
    pub fn peek_east(&self) -> Result<&T, anyhow::Error> {
        let new_pos = MapPosition{x: self.location.x+1, y: self.location.y};
        self.map.get(&new_pos)
    }
    pub fn peek_west(&self) -> Result<&T, anyhow::Error> {
        let new_pos = MapPosition{x: self.location.x-1, y: self.location.y};
        self.map.get(&new_pos)
    }
}

#[derive(Debug)]
pub struct Map<T> {
    pub tiles: Vec<Vec<T>>,
    pub width: usize,
    pub height: usize,
}

/// A map created from given input. Map is a 2d map of data. `T` is a MapLocation that provides the function to
/// transform the characters in the input into the required data.
impl<T> Map<T>
    where T: MapLocation {
    fn new(input: &str) -> Map<T> {
        let tiles = input
            .lines()
            .map(|l| l.chars().map(|c| T::get_tile(c)).collect())
            .collect::<Vec<Vec<T>>>();
        let width = tiles.len();
        let height = tiles.first().unwrap().len();
        Self {
            tiles,
            width,
            height,
        }
    }

    pub fn get(&self, l: &MapPosition) -> Result<&T, anyhow::Error> {
        if l.x > self.width || l.y > self.height {
            Err(anyhow!("position out-of-bounds: {:?}", l))
        } else {
            Ok(self.tiles.get(l.y).unwrap().get(l.x).unwrap())
        }
    }

    pub fn update_tile(&mut self, l: &MapPosition, new_tile: T) {
        self.tiles[l.y][l.x] = new_tile;
    }
}
