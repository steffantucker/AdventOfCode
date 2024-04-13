use std::{env, fs};
use minreq;

pub fn get_or_create_input(day: usize) -> String {
    let path = format!("inputs/day{}input", day);
    if let Ok(data) = fs::read_to_string(path.as_str()) {
        return data;
    }
    let session_key = env::var("adventofcodecookie").expect("AoC env var not set/found");
    let body = minreq::get(format!("https://adventofcode.com/2023/day/{}/input", day))
        .with_header("Cookie", format!("session={}",session_key))
        .send()
        .expect("problem GETting");
    if body.status_code == 200 {
        let input = body.as_str().expect("body not string??");
        let mkdir = fs::create_dir("inputs");
        let _ = fs::write(path.as_str(), input).expect("error writing to file");
        return input.to_string();
    } else {
        panic!("GET return status code: {}", body.status_code);
    }
}
