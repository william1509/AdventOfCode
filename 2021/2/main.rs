use std::fs::File;
use std::io::{BufRead, BufReader};
use std::env::args;

struct Action {
    direction: String,
    magnitude: i64,
}

fn read(path: String) -> Vec<Action> {
    let file = File::open(path).unwrap();
    let mut vec = Vec::new();
    let reader = BufReader::new(file);

    for line in reader.lines() {
        let line = line.unwrap();
        let words: Vec<&str> = line.split_whitespace().collect();

        let action = Action {
            direction: words[0].to_string(),
            magnitude: words[1].parse::<i64>().unwrap(),
        };


        vec.push(action);

    }
    return vec;
}

fn main() {
    let path: String = args().nth(1).unwrap();
    let vec = read(path);
    let mut x = 0;
    let mut y = 0;
    let mut aim = 0;
    
    for action in vec {
        match action.direction.as_str() {
            "forward" => {
                x += action.magnitude;
                y += action.magnitude * aim;
            },
            "down" => {
                aim += action.magnitude;
            }
            "up" => {
                aim -= action.magnitude;
            }
            _ => {
                println!("An error has occured");
            }
        }
    }


    println!("{} times {} is {}", x, y, x * y);

}