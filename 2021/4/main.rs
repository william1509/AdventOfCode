use std::fs::File;
use std::io::BufReader;
use std::io::BufRead;
use std::process::exit;
use std::collections::HashSet;
fn check_if_win(vec: Vec<bool>) -> bool {
    let mut win = false;
    let mut num = 0;
    for i in 0..5 {
        if vec[num] && vec[num + 1] && vec[num + 2] && vec[num + 3] && vec[num + 4] {
            win = true;
        }
        num += 5;
    }
    for i in 0..5 {
        if vec[0 + i] && vec[5 + i] && vec[10 + i] && vec[15 + i] && vec[20 + i] {
            win = true;
        }
    }
    return win;
}

fn main() {
    
    let path = "input.txt";
    let mut numbers = Vec::new();
    let mut file = File::open(path).expect("file not found");
    let mut reader = BufReader::new(file);

    let mut buf = String::new();

    reader.read_line(&mut buf);
    numbers = buf.trim().split(',').map(|x| x.parse::<u32>().unwrap()).collect();

    // skip a line
    reader.read_line(&mut buf);

    let mut boards: Vec<Vec<u32>> = Vec::new();
    boards.resize(100, Vec::new());
    for i in 0..100 {
        for _j in 0..5 {
            buf = String::new();

            // read a line ex: 22 13 17 11 0
            reader.read_line(&mut buf);

            let mut dd: Vec<u32> = buf.trim().split_whitespace().map(|x| x.parse::<u32>().unwrap()).collect();

            boards[i].append(&mut dd);
        }
        // skip a line
        reader.read_line(&mut buf);
    }

    let mut picked_numbers: Vec<Vec<bool>> = Vec::new();
    picked_numbers.resize(100, Vec::new());

    picked_numbers.iter_mut().for_each(|x| x.resize(25, false));

    let mut last_sum = 0;

    let mut won = HashSet::new();

    for num in numbers {
        for i in 0..boards.len() {
            if won.contains(&i) {
                continue;
            }
            for j in 0..boards[i].len() {
                if boards[i][j] == num {
                    picked_numbers[i][j] = true;
                    break;
                }
            }
            if check_if_win(picked_numbers[i].clone()) {
                let mut sum = 0;
                for j in 0..25 {
                    if !picked_numbers[i][j] {
                        sum += boards[i][j];
                    }
                }

                won.insert(i);
                sum *= num;
                
                last_sum = sum;

            }
        }

    }
    println!("{}", last_sum);

}