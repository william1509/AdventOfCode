import Foundation


// This is insanity
func read_file(numbers: Bool = false) -> [(String, String)] {
    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {

            // Split input line by line
            let entries = savedString.split(separator: "\n").map({ String($0) })
            var input: [(String, String)] = [] 
            for entry in entries {
                let split_entry = entry.split(separator: " ").map({ String($0) })
                let tuple = (split_entry[0], split_entry[1])
                input.append(tuple)
            }
            

            return input
        }
    } catch {
        print("Unable to read the file")
    }
    return [];
}

// Part 1
var input = read_file();

var accumulator = 0
var line_passed: Set<Int> = []
var current_line = 0

while(true) {
    if line_passed.contains(current_line) {
        break;
    }
    
    if current_line == input.count {
        break;
    }

    line_passed.insert(current_line)
    let operation = input[current_line]

    switch operation.0 {
        case "nop":
            current_line += 1
            break
        case "acc":
            current_line += 1
            accumulator += Int(operation.1) ?? 0
            break;
        case "jmp":
            current_line += Int(operation.1) ?? 0

        default:
            throw fatalError()
    }
}
print("Part 1: \(accumulator)")


// Part 2

outer: for ins in 0..<input.count {

    var input_copy = input

    current_line = 0
    line_passed = []
    accumulator = 0

    let instruction = input_copy[ins]

    if instruction.0 == "jmp" {
        input_copy[ins] = ("nop", instruction.1)
    } else if input_copy[ins].0 == "nop" {
        input_copy[ins] = ("jmp", instruction.1)
    }

    while(true) {

        if line_passed.contains(current_line) {
            break
        }
        
        if current_line == input_copy.endIndex {
            break outer
        }

        line_passed.insert(current_line)
        let operation = input_copy[current_line]

        switch operation.0 {
            case "nop":
                current_line += 1
                break
            case "acc":
                current_line += 1
                accumulator += Int(operation.1) ?? 0
                break;
            case "jmp":
                current_line += Int(operation.1) ?? 0

            default:
                throw fatalError()
        }
    }
}

print("Part 2: \(accumulator)")


