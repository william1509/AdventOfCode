import Foundation

func read_file() -> [String] {

    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {
            let input = savedString
                .split(separator: "\n")
                .map({ String($0) })
            

            return input
        }
    } catch {
        print("Unable to read the file")
    }
    return [];
}

// Part 1
var input = read_file();

var max_ID = 0.0

for entry in input {
    var upper_bound = 127.0
    var lower_bound = 0.0;

    for c in entry.prefix(7) {
        if c == "F" {
            upper_bound = floor(lower_bound + (upper_bound - lower_bound) / 2)
        } else if c == "B" {
            lower_bound = ceil(lower_bound + (upper_bound - lower_bound) / 2)
        } else {
            print("Weird")
        }
    }

    let row = lower_bound

    upper_bound = 7.0
    lower_bound = 0.0;

    for c in entry.suffix(3) {
        if c == "L" {
            upper_bound = floor(lower_bound + (upper_bound - lower_bound) / 2)
        } else if c == "R" {
            lower_bound = ceil(lower_bound + (upper_bound - lower_bound) / 2)
        } else {
            print("Weird")
        }
    }

    let column = lower_bound

    max_ID = max(max_ID, row * 8 + column)
}
print("Part 1: \(max_ID)")

// Part 2
var seats: [Int] = []

for entry in input {
    var upper_bound = 127.0
    var lower_bound = 0.0;

    for c in entry.prefix(7) {
        if c == "F" {
            upper_bound = floor(lower_bound + (upper_bound - lower_bound) / 2)
        } else if c == "B" {
            lower_bound = ceil(lower_bound + (upper_bound - lower_bound) / 2)
        } else {
            print("Weird")
        }
    }

    let row = lower_bound

    upper_bound = 7.0
    lower_bound = 0.0;

    for c in entry.suffix(3) {
        if c == "L" {
            upper_bound = floor(lower_bound + (upper_bound - lower_bound) / 2)
        } else if c == "R" {
            lower_bound = ceil(lower_bound + (upper_bound - lower_bound) / 2)
        } else {
            print("Weird")
        }
    }

    let column = lower_bound

    seats.append(Int(row * 8 + column))
}
seats.sort()
for i in 1..<seats.count - 1 {
    if seats[i + 1] != seats[i] + 1 {
        print("Part 2: \(seats[i] + 1)")
        break;
    }
}