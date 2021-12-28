import Foundation

func read_file(numbers: Bool = false) -> [[Character]] {
    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {

            // Split input line by line
            let input = savedString.split(separator: "\n").map({ Array($0) })
            return input
            
        }
    } catch {
        print("Unable to read the file")
    }
    return [];
}

func get_occupied_around(input: [[Character]], row: Int, letter: Int) -> Int {
    var count = 0

    let pos_to_check = [
        (row, letter - 1),
        (row, letter + 1),
        (row - 1, letter),
        (row + 1, letter),
        (row - 1, letter - 1),
        (row - 1, letter + 1),
        (row + 1, letter - 1),
        (row + 1, letter + 1),
    ]

    for pos in pos_to_check {
        if (0..<input.count).contains(pos.0) && (0..<input[pos.0].count).contains(pos.1) {
            if input[pos.0][pos.1] == "#" {
                count += 1
            }
        }
    }

    return count
}

// Part 1
var input = read_file();

var input_copy = input

var changes = 1

while changes != 0 {
    var new_input = input_copy
    changes = 0
    for row in 0..<input.count {
        for letter in 0..<input[row].count {
            switch input_copy[row][letter] {
                case "L":
                    if get_occupied_around(input: input_copy, row: row, letter: letter) == 0 {
                        new_input[row][letter] = "#"
                        changes += 1
                    }
                    break
                case "#":
                    if get_occupied_around(input: input_copy, row: row, letter: letter) >= 4 {
                        new_input[row][letter] = "L"
                        changes += 1
                    }
                    break
                case ".":
                    break
                default:
                    fatalError()
            }
        }
    }
    input_copy = new_input
}

var count = 0

for row in input_copy {
    for seat in row {
        if seat == "#" {
            count += 1
        }
    }
} 

print("Part 1: \(count)")

// Part 2
func get_occupied_in_directions(input: [[Character]], row: Int, letter: Int) -> Int {
    var count = 0

    var direction_done = [Bool](repeating: false, count: 8)

    var i = 1

    while direction_done.contains(false) {
        let pos_to_check = [
            (row, letter - i),
            (row, letter + i),
            (row - i, letter),
            (row + i, letter),
            (row - i, letter - i),
            (row - i, letter + i),
            (row + i, letter - i),
            (row + i, letter + i),
        ]
        for pos in pos_to_check {
            guard let index = pos_to_check.firstIndex(where: { $0 == pos }) else {
                continue
            }

            if direction_done[index] {
                continue
            }
            if (0..<input.count).contains(pos.0) && (0..<input[pos.0].count).contains(pos.1) {
                if input[pos.0][pos.1] == "#" {
                    count += 1
                    direction_done[index] = true
                    
                } else if input[pos.0][pos.1] == "L" {
                    direction_done[index] = true

                }
            } else {
                direction_done[index] = true
            }
        }
        i += 1
    }
    return count
}

input_copy = input

changes = 1

while changes != 0 {
    var new_input = input_copy
    changes = 0
    for row in 0..<input.count {
        for letter in 0..<input[row].count {
            switch input_copy[row][letter] {
                case "L":
                    if get_occupied_in_directions(input: input_copy, row: row, letter: letter) == 0 {
                        new_input[row][letter] = "#"
                        changes += 1
                    }
                    break
                case "#":
                    if get_occupied_in_directions(input: input_copy, row: row, letter: letter) >= 5 {
                        new_input[row][letter] = "L"
                        changes += 1
                    }
                    break
                case ".":
                    break
                default:
                    fatalError()
            }
        }
    }
    input_copy = new_input
}

count = 0

for row in input_copy {
    for seat in row {
        if seat == "#" {
            count += 1
        }
    }
} 

print("Part 2: \(count)")