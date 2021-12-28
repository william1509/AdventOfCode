import Foundation

func read_file() -> [[String]] {

    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {
            let input = savedString
                .split(separator: "\n", omittingEmptySubsequences: false)
                .map({ String($0) })
                .split(separator: "")
                .map({ Array($0) })
            
            return input
        }
    } catch {
        print("Unable to read the file")
    }
    return [];
}

// Part 1
var input = read_file();
var result = 0
for entry in input {
    var letters: Set<Character> = []
    for line in entry {
        for letter in line {
            letters.insert(letter)
            
        }
    }
    result += letters.count
}
print(result)

// Part 2

result = 0
for entry in input {
    var letters: Set<Character> = []
    let first_line = entry.first!
    for letter in first_line {
        var contains_letter = true
        for line in entry {
            contains_letter = contains_letter && line.contains(letter)
        }

        if contains_letter {
            letters.insert(letter)
        }
    }
    print(letters)
    result += letters.count
}
print(result)

