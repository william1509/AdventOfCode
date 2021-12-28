import Foundation

func read_file_as_ints() -> [UInt] {

    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {
            return savedString.split(separator: "\n").map { UInt($0) ?? 0 }
        }
    } catch {
        print("Unable to read the file")
    }
    return [];
}

// Part 1
var input = read_file_as_ints();
var begin = 0;
var end = input.count - 1;

var number = 2020;

input.sort();

for _ in 0...input.count {
    let sum = input[begin] + input[end];
    if sum == number {
        print("Part 1: \(input[begin] * input[end])");
        break;
    } else if sum < number {
        begin += 1;
    } else if sum > number {
        end -= 1;
    }

    if begin >= end {
        print("No solution found")
        break;
    }
}

// Part 2
// Probably a better way to do this

outer: for i in input {
    for j in input {
        for k in input {
            if i + j + k == number {
                print("Part 2: \(i * j * k)");
                break outer;
            }
        }
    }
}