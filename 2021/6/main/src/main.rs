use std::fs;

fn part_1() {
    let path = "/home/william/AdventOfCode/2021/6/main/src/input.txt";
    let mut fishes = fs::read_to_string(path).unwrap().split(',').map(|x| x.parse::<i32>().unwrap()).collect::<Vec<i32>>();
    for day in 0..256 {
        let mut new_fishes = 0;
        for fish in fishes.iter_mut() {
            if *fish - 1 == -1 {
                new_fishes += 1;
                *fish = 6;
            } else {
                *fish -= 1;
            }
        }

        fishes.append(&mut vec![8; new_fishes]);
        println!("{}", day);
    }
    println!("{}", fishes.len());
}

// Using the same algorithm as part 1 is way too slow.
fn part_2() {
    let path = "/home/william/AdventOfCode/2021/6/main/src/input.txt";
    let fishes = fs::read_to_string(path).unwrap().split(',').map(|x| x.parse::<i32>().unwrap()).collect::<Vec<i32>>();

    let mut fish_ages = vec![0; 8 + 256 + 1];

    let mut number_fish = fishes.len();

    for fish in fishes {
        fish_ages[fish as usize] += 1;
    }

    for day in 0..256  {
        // Fish that are reproducing today
        let current_fishes = fish_ages[day as usize];

        // Fish that are reproducing in 6 days
        fish_ages[day + 7] += current_fishes;

        // Fish that are reproducing in 8 days
        fish_ages[day + 9] += current_fishes;

        number_fish += current_fishes;
    }

    println!("{}", number_fish);
}

fn main() {
    part_2();
}
