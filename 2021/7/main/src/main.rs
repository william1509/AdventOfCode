use std::fs;

fn main() {
    let path = "/home/william/AdventOfCode/2021/7/main/src/input";
    let mut crabs = fs::read_to_string(path).unwrap().split(',').map(|x| x.parse::<i32>().unwrap()).collect::<Vec<i32>>();

    crabs.sort();

    // Part 1
    // let median = crabs[crabs.len() / 2];
    let mut fuel = 0;
    // for crab in crabs {
    //     fuel += (crab - median).abs();
    // }

    let sum: i32 = Iterator::sum(crabs.iter());
    let mean = sum / crabs.len() as i32;
    for crab in crabs {
        for i in 1..=(crab - mean).abs() {
            fuel += i;
        }
    }
    println!("{}", mean);
    println!("{}", fuel);
}
