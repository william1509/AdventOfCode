use std::fs::File;
use std::io::{BufRead, BufReader, Error, ErrorKind};
use std::env::args;
use std::io;

fn read(path: &str) -> Result<Vec<i64>, io::Error> {
    let file = File::open(path)?;

    let br = BufReader::new(file);
    let mut v = Vec::new();
    for line in br.lines() {
        let line = line?;
        let n = line
            .trim() 
            .parse()
            .map_err(|e| Error::new(ErrorKind::InvalidData, e))?;
        v.push(n);
    }
    Ok(v)
}

fn count_increases(vec: Vec<i64>) -> i64 {
    let mut count = 0;
    let mut previous_elem = vec.first().unwrap();
    for elem in vec.iter() {
        if elem > previous_elem {
            count += 1;
        }
        previous_elem = elem;
    }

    return count;
}

fn count_window_increases(vec: Vec<i64>) -> i64 {
    let mut count = 0;
    let mut previous_window = vec[0] + vec[1] + vec[2];
    for (i, elem) in vec.iter().enumerate() {
        if i == vec.len() - 2 {
            break;
        }
        let current_window = elem + vec[i + 1] + vec[i + 2];
        if current_window > previous_window {
            count += 1;
        }
        // println!("The current window is {} and the previous is {}", current_window, previous_window);
        previous_window = current_window;
    }

    return count;
}

fn main() {
    let path: String = args().nth(1).unwrap();
    let vec = read(&path).unwrap();

    let count = count_window_increases(vec);

    println!("{}", count);

}