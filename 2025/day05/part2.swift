import Foundation

let filePath = "input_example.txt"
// let filePath = "input_full.txt"

let content = try String(contentsOfFile: filePath, encoding: String.Encoding.utf8)

// print(content)

let lines = content.split(separator: "\n").map(String.init)

let rangesStart = lines.map({ line in line.split(separator: "-") }).compactMap({ arr in
    if arr.count >= 2, let first = Int(arr[0]), let second = Int(arr[1]) {
        return (first, second)
    }
    return nil
})

print(rangesStart)

let sortedRanges = rangesStart.sorted { $0.0 < $1.0 }

print(sortedRanges)

var reelRanges: [(Int, Int)] = []
var nextStart = -1
for (i, (start, end)) in sortedRanges.enumerated() {
    if nextStart == -1 {
        nextStart = start
    }

    var nextNextStart = -1
    if i < sortedRanges.count - 1 {
        nextNextStart = sortedRanges[i + 1].0
    }

    if nextNextStart <= end {
        continue
    }

    reelRanges.append((nextStart, end))

    nextStart = -1
}

for range in reelRanges {
    print("\(range)")
}
// rangesStart.forEach({ Int in print(Int) })
