use std::fs::File;
use std::io::{BufRead, BufReader};
use std::env::args;
use std::char;

fn read() -> Vec<Vec<String>> {
    let file = File::open("/home/william/AdventOfCode/2021/8/main/src/input").unwrap();
    let mut vec = Vec::new();
    let reader = BufReader::new(file);

    for line in reader.lines() {
        let line = line.unwrap();
        let sep_in_two = line.trim().split(" | ");

        let xs: Vec<String> = sep_in_two.map(|s| s.parse().unwrap()).collect();
        vec.push(xs);
    }

    return vec;
}

fn main() {
    let vec = read();
    let mut count = 0;
    for x in vec {
        let output: Vec<&str> = x[1].split(" ").collect();
        for word in output {
            let word_len = word.len();
            if word_len == 7 || word_len == 2 || word_len == 3 || word_len == 4 {
                println!("{}", word);
                count += 1;
            }
        }
        
    }

    println!("{}", count);
}
