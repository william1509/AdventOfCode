import Foundation

func read_file() -> [String] {

    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {
            let savedString = savedString.split(separator: "\n").map({String($0)});
            
            return savedString;
        }
    } catch {
        print("Unable to read the file")
    }
    return [];
}

// Part 1
var input = read_file();

var posX = 0
var count = 0
let length = input.first?.count ?? 0

for line in input {
    let index = line.index(line.startIndex, offsetBy: posX % length)

    if String(line[index]) == "#" {
        count += 1
    }
    posX += 3
}
print("Part 1: \(count)")

// Part 2
var posX1 = 0
var posX2 = 0
var posX3 = 0
var posX4 = 0

count = 0
var count1 = 0
var count2 = 0
var count3 = 0
var count4 = 0
var count5 = 0

for line in input {

    var index = line.index(line.startIndex, offsetBy: posX1 % length)
    if String(line[index]) == "#" {
        count1 += 1
    }
    posX1 += 1

    index = line.index(line.startIndex, offsetBy: posX2 % length)
    if String(line[index]) == "#" {
        count2 += 1
    }
    posX2 += 3 

    index = line.index(line.startIndex, offsetBy: posX3 % length)
    if String(line[index]) == "#" {
        count3 += 1
    }
    posX3 += 5

    index = line.index(line.startIndex, offsetBy: posX4 % length)
    if String(line[index]) == "#" {
        count4 += 1
    }
    posX4 += 7
}
var posX5 = 0
var posY5 = 0

while posY5 < input.count {
    let index = input[posY5].index(input[posY5].startIndex, offsetBy: posX5 % length)
    if String(input[posY5][index]) == "#" {
        count5 += 1
    }
    posX5 += 1
    posY5 += 2
}

count = count1 * count2 * count3 * count4 * count5
print("Part 2: \(count)")
