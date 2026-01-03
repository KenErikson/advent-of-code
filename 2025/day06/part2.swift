import Foundation

// let filePath = "input_example.txt"
let filePath = "input_full.txt"

let content = try String(contentsOfFile: filePath, encoding: String.Encoding.utf8)

// print(content)

let lines = content.split(separator: "\n").map(String.init)

guard let lastLine = lines.last else {
    fatalError()
}

let signs = lastLine.split(separator: " ")

guard let firstLine = lines.first else { fatalError() }

let columnCount = firstLine.count

var numbers: [[Int]] = Array(repeating: [Int](), count: signs.count)
var j = 0
for i in 0..<columnCount{
    // if (i+1) % 4 == 0 && i != 0 {
    //     j += 1
    //     continue
    // }
    var numb = ""
    for line in lines {
        if line.contains("+"){
            break
        }
        // print(line)
        if Array(line)[i] == " "{
            continue
        }
        print("added",Array(line)[i])
        numb.append(Array(line)[i])
    }
    print("numb:",numb)
    // if numbers.count < j + 1{
    //     numbers.append([])
    // }
    guard let numbInt = Int(numb) else {
        j += 1
        continue
    }

    // print(numbInt!)

    print(j)
    numbers[j].append(numbInt)
    print(numbers)
}

var res = 0
for (i,sign) in signs.enumerated() {
    switch sign {
        case "+":
            var tmpRes = 0
            for nr in numbers[i] {
                tmpRes += nr
            }
            res += tmpRes
            print(tmpRes)
        default:
            var tmpRes = 1
            for nr in numbers[i]{
                tmpRes *= nr
            }
            res += tmpRes
            print(tmpRes)
    }
    print(res)
}
print(res)
