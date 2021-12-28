import Foundation

struct entry {
    init(range: String, letter: String, password: String) {
        let r = range.split(separator: "-").map { Int($0) ?? 0 };
        self.range = [r[0], r[1]];

        self.letter = String(letter.dropLast());
        self.password = password;

    }
    let range: [Int];
    let letter: String;
    let password: String;
}

func read_file() -> [entry] {

    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {
            let savedString = savedString.split(separator: "\n").map { $0.split(separator: " ").map({ String($0) }) };
            var result: [entry] = []

            savedString.forEach({ result.append(entry(range: $0[0], letter: $0[1], password: $0[2])) })
            return result
        }
    } catch {
        print("Unable to read the file")
    }
    return [];
}

// Part 1
var input = read_file();

var count = 0

for entry in input {
    var repetitions = 0
    entry.password.forEach( { if $0.lowercased() == entry.letter { repetitions += 1 } } )
    if repetitions >= entry.range[0] && repetitions <= entry.range[1]  {
        count += 1
    }
}
print("Part 1: \(count)")

// Part 2
count = 0
for entry in input {
    let index1 = entry.password.index(entry.password.startIndex, offsetBy: entry.range[0] - 1)
    let index2 = entry.password.index(entry.password.startIndex, offsetBy: entry.range[1] - 1)

    let contains1 = String(entry.password[index1]) == entry.letter
    let contains2 = String(entry.password[index2]) == entry.letter

    if (contains1 && !contains2) || (!contains1 && contains2) {
        count += 1
    }
}
print("Part 2: \(count)")