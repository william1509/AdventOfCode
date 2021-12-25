import Foundation

func read_file() -> [ArraySlice<[[Substring.SubSequence]]>] {

    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {
            let input = savedString
                .split(separator: "\n", omittingEmptySubsequences: false)
                .map({ $0.split(separator: " ").map({ $0.split(separator: ":") }) })
                .split(separator: [])

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
    var count = 0
    var has_cid = false
    for line in entry {
        for field in line {
            count += 1
            if field[0] == "cid" {
                has_cid = true
            }
        }
    }
    if !has_cid && count == 7 {
        result += 1
    } else if count == 8 {
        result += 1
    }
}
print("Part 1: \(result)")

// Part 2

result = 0
for entry in input {
    var count = 0
    var has_cid = false
    var is_valid = true
    for line in entry {
        for field in line {
            count += 1
            switch field[0] {
                case "byr":
                    let byr_int = Int(field[1]) ?? 0
                
                    is_valid = is_valid && (1920...2002).contains(byr_int)
                    
                    break;
                case "iyr":
                    let iyr_int = Int(field[1]) ?? 0
                    is_valid = is_valid && (2010...2020).contains(iyr_int)
                    
                    break;
                case "eyr":
                    let eyr_int = Int(field[1]) ?? 0
                    is_valid = is_valid && (2020...2030).contains(eyr_int)
                     
                    break;
                case "hgt":
                    if field[1].suffix(2) == "cm" {

                        let hgt_int = Int(field[1].prefix(3)) ?? 0
                        is_valid = is_valid && (150...193).contains(hgt_int)
                        
                    } else if field[1].suffix(2) == "in" {
                        let hgt_int = Int(field[1].prefix(2)) ?? 0
                        is_valid = is_valid && (59...76).contains(hgt_int)
                        
                    } else { 
                        is_valid = false
                    }
                    
                    break;
                case "hcl":
                    let range = NSRange(location: 0, length: field[1].utf16.count)
                    let regex = try! NSRegularExpression(pattern: "^#([A-Fa-f0-9]{6})")
                    is_valid = is_valid && (regex.firstMatch(in: String(field[1]), options: [], range: range) != nil)
                    
                case "ecl":
                    let colors = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]

                    is_valid = is_valid && colors.contains(String(field[1]))
                    
                    break;
                case "pid":
                    is_valid = is_valid && field[1].count == 9
                    
                    break;
                case "cid" :
                    has_cid = true
                    break;
                default:
                    print("Weird")
            }
        }
    }
    if !has_cid && count == 7 && is_valid {
        print(entry)
        result += 1
    } else if count == 8 && is_valid {
        print(entry)

        result += 1
    }
}
print("Part 2: \(result)")