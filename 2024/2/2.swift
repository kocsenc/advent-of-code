import Foundation

// https://adventofcode.com/2024/day/2

func parseFileIntoIntegerListReports(filePath: String) -> [[Int]] {
  do {
    let lines = try String(contentsOfFile: filePath, encoding: .utf8)
      .split(separator: "\n")

    let integerLists = lines.map { line -> [Int] in
      return line.split(separator: " ")
        .compactMap { Int($0) }
    }

    return integerLists
  } catch {
    print("Error reading file: \(error)")
    return [[]]
  }
}

func isMonotonicWithinRange(_ nums: [Int], minDifference: Int, maxDifference: Int) -> Bool {
  guard nums.count > 1 else { return true }

  let differences = zip(nums, nums.dropFirst()).map { $1 - $0 }

  if (differences.allSatisfy { $0 >= 0 && minDifference...maxDifference ~= $0 }
    || differences.allSatisfy { $0 <= 0 && minDifference...maxDifference ~= abs($0) })
  {
    return true
  }

  // We try to remove at most one element to see if we can get a satisfactory report.
  // This approach is brute force trial and error since algorithmically finding a pattern
  // still cost the same time complexity and keeps the code simple.
  for num in nums.enumerated() {
    var numsCopy = nums
    numsCopy.remove(at: num.offset)

    let tolerantDifferences = zip(numsCopy, numsCopy.dropFirst()).map { $1 - $0 }
    if (tolerantDifferences.allSatisfy { $0 >= 0 && minDifference...maxDifference ~= $0 }
      || tolerantDifferences.allSatisfy { $0 <= 0 && minDifference...maxDifference ~= abs($0) })
    {
      return true
    }
  }

  return false
}

func main(filePath: String) {
  var safeCount = 0

  if FileManager.default.fileExists(atPath: filePath) {
    let reports = parseFileIntoIntegerListReports(filePath: filePath)
    for report in reports {
      if isMonotonicWithinRange(report, minDifference: 1, maxDifference: 3) {
        safeCount += 1
      }
    }
  } else {
    print("File not found at path: \(filePath)")
  }

  print(safeCount)
}
main(filePath: "./input.txt")
