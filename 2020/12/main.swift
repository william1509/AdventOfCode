import Foundation

func read_file() -> [(Character, Int)] {
    let fileURL = URL(fileURLWithPath: "input", relativeTo: URL(fileURLWithPath: FileManager.default.currentDirectoryPath))
    do {
        let savedData = try Data(contentsOf: fileURL)
        if let savedString = String(data: savedData, encoding: .utf8) {

            // Split input line by line
            let input = savedString.split(separator: "\n").map({ (str) -> (Character, Int) in
                let a = Array(str)
                return (a[0], Int(String(a[1..<a.count])) ?? 0) 
            })
            return input
            
        }
    } catch {
        print("Unable to read the file")
    }
    return [];
}

func mod(_ a: Int, _ n: Int) -> Int {
    precondition(n > 0, "modulus must be positive")
    let r = a % n
    return r >= 0 ? r : r + n
}

let input = read_file()

// 0 = EAST, 1 = SOUTH, 2 = WEST, 3 = NORTH
var direction = 0

var posX = 0
var posY = 0

for move in input {
    switch move.0 {
        case "N":
            posY += move.1
            break
        case "S":
            posY -= move.1
            break
        case "E":
            posX += move.1
            break
        case "W":
            posX -= move.1
            break
        case "L":
            direction = mod((direction - move.1 / 90), 4)
            break
        case "R":
            direction = mod((direction + move.1 / 90), 4)
            break
        case "F":
            switch direction {
                case 0:
                    posX += move.1
                    break
                case 1:
                    posY -= move.1
                    break
                case 2:
                    posX -= move.1
                    break
                case 3:
                    posY += move.1
                    break
                default:
                    fatalError()
            }
        default:
            fatalError()

    }
}

print("Part 1: \(abs(posX) + abs(posY))")

// 0 = EAST, 1 = SOUTH, 2 = WEST, 3 = NORTH
direction = 0

var ship = (0, 0)
var waypoint = (10, 1)

for move in input {
    switch move.0 {
        case "N":
            waypoint.1 += move.1
            break
        case "S":
            waypoint.1 -= move.1
            break
        case "E":
            waypoint.0 += move.1
            break
        case "W":
            waypoint.0 -= move.1
            break
        case "L", "R":
            for _ in 0..<move.1 / 90 {
                let x = waypoint.0
                let y = waypoint.1
                
                if move.0 == "L" {
                    waypoint = (-y, x)
                } else if move.0 == "R" {
                    waypoint = (y, -x)
                }
            }
        case "F":
            ship.0 += move.1 * waypoint.0
            ship.1 += move.1 * waypoint.1
            
        default:
            fatalError()

    }
}

print("Part 2: \(abs(ship.0) + abs(ship.1))")