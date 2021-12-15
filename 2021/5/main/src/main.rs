// For some reason part 2 answer is 1 too low
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::cmp;

fn parse_to_vec(input: &str) -> Vec<i32> {
    let test = input.replace("->", ",");
    let it = test.split(',');

    return it.map(|x| x.trim().parse::<i32>().unwrap())
    .collect()

}

fn main() {
    let file = File::open("/home/william/AdventOfCode/2021/5/main/src/input.txt").unwrap();
    let reader = BufReader::new(file);

    let mut vents:[[i32;1000];1000] = [[0;1000];1000];

    for line in reader.lines() {
        let line = line.unwrap();
        let vec = parse_to_vec(&line);

        // vertical line
        if vec[0] == vec[2] {
            let min = cmp::min(vec[1], vec[3]);
            let max = cmp::max(vec[1], vec[3]);
            for i in min..=max {
                vents[i as usize][vec[0] as usize] += 1;
            }
        }

        // horizontal line
        else if vec[1] == vec[3] {
            let min = cmp::min(vec[0], vec[2]);
            let max = cmp::max(vec[0], vec[2]);
            for i in min..=max {
                vents[vec[1] as usize][i as usize] += 1;
            }
        } 
        
        // diagonal
        // part 2
        else {
            let x_inc;
            let y_inc;

            if vec[0] < vec[2] {
                x_inc = 1; 
            } else {
                x_inc = -1;
            }
            if vec[1] < vec[3] {
                y_inc = 1; 
            } else {
                y_inc = -1;
            }

            let mut x = vec[0];
            let mut y = vec[1];
            for _ in 0..(vec[2] - vec[0]).abs() {
                vents[x as usize][y as usize] += 1;
                x += x_inc;
                y += y_inc;
            }

        }
    }
    let mut count = 0;
    for x in vents.iter() {
        for y in x.iter() {
            if *y > 1 {
                count += 1;
            }
        }
    }

    println!("{}", count);
}