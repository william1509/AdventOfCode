import Foundation


// This is insanity
func read_file(numbers: Bool = false) -> [String: [(Int, String)]] {
    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {

            // Split input line by line
            let entries = savedString.split(separator: "\n");

            var input: [String: [(Int, String)]] = [:]
            for entry in entries {

                // Split line in two, the bag and the child bags
                let entry_split = entry.components(separatedBy: "bags contain")

                // Split childs bags
                let child_split = entry_split[1].components(separatedBy: ", ")

                // Separate words in parent bag name
                let parent = entry_split[0].split(separator: " ").map({ String($0) })

                // Separate words in child bags names
                let childs: [[String]] = child_split.map({ $0.split(separator: " ").map({ String($0) }) })

                // Number of bags, name of the bag
                var result_child: [(Int, String)] = []

                for c in childs {
                    // Create tuple of child bags
                    result_child.append((Int(c[0]) ?? 0, "\(c[1]) \(c[2])"))
                }

                // Insert into map
                input["\(parent[0]) \(parent[1])"] = result_child

            }
            return input
        }
    } catch {
        print("Unable to read the file")
    }
    return [:];
}

// Part 1
var input = read_file();

var bags : Set<String> = []

for entry in input {
    // current, parent
    var queue: [(String , String)] = [(entry.key , entry.key)]

    while !queue.isEmpty {
        guard let current = queue.popLast() else {
            continue
        }

        guard let childs = input[current.0] else {
            continue
        }

        for c in childs {
            if c.1.contains("shiny gold") {
                bags.insert(current.1);
            }
            
            queue.append((c.1, current.1))
        }
    }
}
print("Part 1: \(bags.count)")

// Part 2

// name of the bag, number of bags
var queue: [(String, Int)] = [("shiny gold", 1)]

var count = 0;

while !queue.isEmpty {
    guard let current = queue.popLast() else {
        continue
    }

    count += current.1

    guard let childs = input[current.0] else {
        continue
    }

    for c in childs {
        
        
        queue.append((c.1, c.0 * current.1))
    }
}
print("Part 2: \(count - 1)")
