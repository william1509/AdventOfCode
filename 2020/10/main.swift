import Foundation


func read_file(numbers: Bool = false) -> [Int] {
    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {

            // Split input line by line
            let input = savedString.split(separator: "\n").map({ Int($0) ?? 0 })
            return input.sorted()
            
        }
    } catch {
        print("Unable to read the file")
    }
    return [];
}

// Part 1
var input = read_file();

// Charging outlet has 0
var input_copy = input
input_copy.insert(0, at: 0)

var jolt3 = 0
var jolt1 = 0

for i in 0..<input_copy.count - 1 {
    switch input_copy[i + 1] - input_copy[i] {
        case 1:
            jolt1 += 1
            break
        case 3:
            jolt3 += 1
            break

        default:
            fatalError()
    }
}

// Device adapter always 3 higher
jolt3 += 1

print("Part 1: \(jolt1 * jolt3)")


// Part 2

jolt3 = 0
jolt1 = 0

var queue: [Int] = []

var n_paths = 0

var paths: [Int : Int] = [0: 1]

input.append((input.last ?? 0) + 3)

for i in input {
    paths[i] = (paths[i - 3] ?? 0) + (paths[i - 2] ?? 0) + (paths[i - 1] ?? 0)
}

print("Part 2: \(paths[input.last!]!)")
