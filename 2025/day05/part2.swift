import Foundation

// let filePath = "input_example.txt"
// let filePath = "input_example2.txt"
let filePath = "input_full.txt"

let content = try String(contentsOfFile: filePath, encoding: String.Encoding.utf8)

// print(content)

let lines = content.split(separator: "\n").map(String.init)

let rangesStart = lines.map({ line in line.split(separator: "-") }).compactMap({
    arr -> (Int, Int)? in
    if arr.count >= 2, let first = Int(arr[0]), let second = Int(arr[1]) {
        return (first, second)
    }
    return nil
})

print(rangesStart)

let sortedRanges = rangesStart.sorted { $0.0 < $1.0 }

print(sortedRanges)

var reelRanges: [(Int, Int)] = []
var potentialStart: Int = -1
var potentialEnd: Int = -1
for (i, (start, end)) in sortedRanges.enumerated() {
    // print(start, end, nextStart)
    if potentialStart == -1 {
        potentialStart = start
    }

    if end > potentialEnd {
        potentialEnd = end
    }
    var upcomingStart: Int = -1
    if i < sortedRanges.count - 1 {
        upcomingStart = sortedRanges[i + 1].0

        if upcomingStart <= potentialEnd && i < sortedRanges.count {
            // print("skipped")

            continue
        }
    }

    reelRanges.append((potentialStart, potentialEnd))
    print("added", potentialStart, potentialEnd)

    potentialStart = -1
    potentialEnd = -1
}

var count: Int = 0
for range in reelRanges {
    count += range.1 - range.0 + 1
    print("New count: ", count)
}

print(count)
