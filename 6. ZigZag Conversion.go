// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string text, int nRows);
// convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".


/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    var dir int
    index := 0
    var result string = ""
    strArr := make([]string, numRows)
    for i := 0; i < len(s); i++ {
        strArr[index] = strArr[index] + string(s[i])
        if index == numRows - 1 {
            dir = 0
        } else if index == 0 {
            dir = 1
        }
        if dir == 0 {
            index--
        } else {
            index++
        }
    }
    for _,v := range strArr {
        result = result + v
    }
    return result
}