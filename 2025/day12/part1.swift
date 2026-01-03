import Foundation

let filePath = "input_example.txt"
// let filePath = "input_full.txt"

let content = try String(contentsOfFile: filePath, encoding: String.Encoding.utf8)

// print(content)

let contentGroups = content.split(separator: "\n\n").map(String.init)

guard let layouts = contentGroups.last else {
    exit(1)
}

var shapes: [[[Character]]] = [[[]], [[]], [[]], [[]], [[]], [[]]]
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
        for chr in line {
            // print("i", i, "j", j )
            // print(shapes)
            shapes[i][j - 1].append(chr)
        }
    }
}

print(shapes)
