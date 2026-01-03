import Foundation

let filePath = "input_example.txt"
// let filePath = "input_full.txt"

let content = try String(contentsOfFile: filePath, encoding: String.Encoding.utf8)

// print(content)

let contentGroups = content.split(separator: "\n\n").map(String.init)

let layouts = contentGroups.last!

var shapes: [[[Character]]] = [[[]], [[]], [[]], [[]], [[]], [[]]]

var shapeCounts: [Int] = []
for (i, group) in contentGroups.enumerated() {
    if i > 5 {
        break
    }
    let lines = group.split(separator: "\n")
    for (j, line) in lines.enumerated() {
        if j == 0 {
            continue
        }
        if j > 1 {
            shapes[i].append([])
        }
        var hashCount = 0
        for chr in line {
            // print("i", i, "j", j )
            // print(shapes)
            shapes[i][j - 1].append(chr)
            if chr == "#" {
                hashCount += 1
            }
        }
        shapeCounts.append(hashCount)
    }
}

print(shapes)

// print(layouts)

var successCount = 0
for layout in layouts.split(separator: "\n").map(String.init) {
    let parts = layout.split(separator: ":")

    let areaString = parts[0]
    let height = Int(areaString.split(separator: "x")[0])!
    let width = Int(areaString.split(separator: "x")[1])!

    let shapeCounts = parts[1].split(separator: " ").compactMap({ a in Int(a)! })

    if canFit(width, height, shapeCounts) {
        successCount += 1
    }
}

print("Res", successCount)

func canFit(_ width: Int, _ height: Int, _ shapeCounts: [Int]) -> Bool {
    var matrix: [[Character]] = Array(repeating: Array(repeating: ".", count: width), count: height)

    return false
}
