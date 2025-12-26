use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    // let filepath = "src/bin/input5.txt";
    let filepath = "src/bin/input5_full.txt";

    let path = Path::new(filepath);
    let file = File::open(path).expect("Could not open file");
    let reader = io::BufReader::new(file);

    let mut fresh_ranges: Vec<String> = Vec::new();
    let mut ingredient_ids: Vec<u64> = Vec::new();
    let mut ranges_done = false;
    for line in reader.lines() {
        let mut line = line.expect("Could not read line");
        line = line.trim().to_string();
        if line == "" {
            ranges_done = true;
            continue;
        }

        if !ranges_done {
            fresh_ranges.push(line)
        } else {
            ingredient_ids.push(line.parse().unwrap())
        }
    }

    let mut good_ids: Vec<(u64, u64)> = Vec::new();
    for line in fresh_ranges {
        println!("{line}");
        let splits: Vec<u64> = line
            .split('-')
            .map(|nr| nr.trim().parse().unwrap())
            .collect();
        good_ids.push((splits[0],splits[1]));
    }
    // println!("{}", good_ids.len().to_string());

    let good_ids = good_ids;

    let mut count = 0;
    for id in ingredient_ids {
        for (a,b) in &good_ids {
            if id >= *a && id <= *b {
                count += 1;
                break;
            }
        }
    }

    println!("{}", count)
}
