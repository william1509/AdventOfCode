import Foundation


// This is insanity
func read_file(numbers: Bool = false) -> [Int64] {
    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {

            // Split input line by line
            return savedString.split(separator: "\n").map({ Int64($0) ?? 0 })
            
        }
    } catch {
        print("Unable to read the file")
    }
    return [];
}

func sum_exists(arr: [Int64], goal: Int64) -> Bool {
    let sorted = arr.sorted()

    var left_it = 0
    var right_it = sorted.endIndex - 1

    while left_it < right_it {
        let sum = sorted[left_it] + sorted[right_it]

        if sum < goal {
            left_it += 1
        } else if sum > goal {
            right_it -= 1
        } else {
            return true
        }
    } 
    return false

}

// Part 1
var input = read_file();

var invalid_sum: Int64 = 0;

for i in 25..<input.count {
    if !sum_exists(arr: Array(input[(i - 25)...i]), goal: input[i]) {
        print("Part 1: \(input[i])")
        invalid_sum = input[i]
        break
    }
}

// Part 2
outer: for length in 2..<input.count {
    for offset in 0..<input.count - length {
        let subarray = Array(input[offset..<(offset + length)])

        if subarray.reduce(0, +) == invalid_sum {
            print("Part 2: \((subarray.min() ?? 0) + (subarray.max() ?? 0))")
            break outer
        }
    }
}