/**
 * @param {number[]} arr
 * @param {number[][]} queries
 * @return {number[]}
 */
var xorQueries = function(arr, queries) {
    const size = arr.length
    const xorArr = Array(size+1).fill(0)
    for (let i = 0 ; i < size; i++) {
        xorArr[i+1] = xorArr[i]^arr[i]
    }
    const result = Array(queries.length).fill(0)
    for (let i = 0 ; i < queries.length; i++) {
        result[i] = xorArr[queries[i][0]]^xorArr[queries[i][1]+1]
    }

    return result
};