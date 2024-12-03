import Foundation

// https://adventofcode.com/2024/day/1
func main() {
    var leftList = [3, 4, 2, 1, 3, 3]
    var rightList = [4, 3, 5, 3, 9, 3]
    var distances: [Int] = []
    
    assert(leftList.count == rightList.count, "Lists are not the same length, a historian has done an oopsie.")
    for _ in 1...leftList.count {
        if let leftMin = findAndPopMinimum(in: &leftList), 
           let rightMin = findAndPopMinimum(in: &rightList) {
            distances.append(abs(leftMin - rightMin))
        }
    }

    print(distances.reduce(0, +))
}


func findAndPopMinimum(in list: inout [Int]) -> Int? {
    guard let minValue = list.min(), let index = list.firstIndex(where: { $0 == minValue }) else {
        return nil
    }
    return list.remove(at: index)
}

// Execute the main function
main()
