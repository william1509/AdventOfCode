use std::fs::File;
use std::io::{BufRead, BufReader};
use std::env::args;
use std::char;

fn read(path: String) -> Vec<String> {
    let file = File::open(path).unwrap();
    let mut vec = Vec::new();
    let reader = BufReader::new(file);

    for line in reader.lines() {
        let line = line.unwrap();

        vec.push(line.to_string());

    }
    return vec;
}

fn get_gamma_rate(vec: Vec<String>) -> u32 {

    let mut result: u32 = 0;
    for i in 0..vec[0].len() {

        let mut zeros: u32 = 0;
        let mut ones: u32 = 0;

        for j in vec.iter() {
            match j.chars().nth(i).unwrap() {
                '0' => {
                    zeros += 1;
                }
                '1' => {
                    ones += 1;
                }
                _ => {
                    println!("Unexpected symbol")
                }
            }
        }
        result <<= 1;
        if ones > zeros {
            result +=1
        }
    }
    return result;
}

fn get_epsilon_rate(vec: Vec<String>) -> u32 {

    let mut result: u32 = 0;
    for i in 0..vec[0].len() {

        let mut zeros: u32 = 0;
        let mut ones: u32 = 0;

        for j in vec.iter() {
            match j.chars().nth(i).unwrap() {
                '0' => {
                    zeros += 1;
                }
                '1' => {
                    ones += 1;
                }
                _ => {
                    println!("Unexpected symbol")
                }
            }
        }
        result <<= 1;
        if zeros > ones {
            result +=1
        }
    }
    return result;
}

fn get_oxygen_rate(vec: Vec<String>) -> u32 {

    let mut remaining = vec.clone();
    let mut result = 0;
    for i in 0..vec[0].len() {

        let mut zeros: u32 = 0;
        let mut ones: u32 = 0;

        for j in remaining.iter() {
            match j.chars().nth(i).unwrap() {
                '0' => {
                    zeros += 1;
                }
                '1' => {
                    ones += 1;
                }
                _ => {
                    println!("Unexpected symbol")
                }
            }
        }
        if ones >= zeros {
            remaining.retain(|x| x.chars().nth(i).unwrap() == '1');
        } else {
            remaining.retain(|x| x.chars().nth(i).unwrap() == '0');
        }
    }
    println!("Result {}", remaining.first().unwrap());
    for i in remaining.first().unwrap().chars() {
        result <<= 1;
        result += i.to_digit(32).unwrap();
    }
    return result;
}

fn get_co2_rate(vec: Vec<String>) -> u32 {

    let mut remaining = vec.clone();
    let mut result = 0;
    for i in 0..vec[0].len() {

        if remaining.len() == 1 {
            println!("Result {}", remaining.first().unwrap());
            for i in remaining.first().unwrap().chars() {
                result <<= 1;
                result += i.to_digit(32).unwrap();
            }
            return result;
        }

        let mut zeros: u32 = 0;
        let mut ones: u32 = 0;

        for j in remaining.iter() {
            match j.chars().nth(i).unwrap() {
                '0' => {
                    zeros += 1;
                }
                '1' => {
                    ones += 1;
                }
                _ => {
                    println!("Unexpected symbol")
                }
            }
        }
        if zeros <= ones {
            remaining.retain(|x| x.chars().nth(i).unwrap() == '0');
        } else {
            remaining.retain(|x| x.chars().nth(i).unwrap() == '1');
        }
    }
    println!("Result {}", remaining.first().unwrap());
    for i in remaining.first().unwrap().chars() {
        result <<= 1;
        result += i.to_digit(32).unwrap();
    }
    return result;
}

fn main() {
    let path = "input.txt".to_string();
    let vec = read(path);

    let oxygen_rate = get_oxygen_rate(vec.clone());
    let co2_rate = get_co2_rate(vec.clone());
    //let epsilon_rate = get_epsilon_rate(vec);


    println!("Oxygen rate {} Co2 rate {} and total {}", oxygen_rate, co2_rate, oxygen_rate * co2_rate);

}