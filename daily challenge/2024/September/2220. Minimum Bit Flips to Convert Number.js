/**
 * @param {number} start
 * @param {number} goal
 * @return {number}
 */
var minBitFlips = function(start, goal) {
    let result = 0
    let bits = start ^ goal
    while(bits > 0) {
        result += bits & 1
        bits = bits >> 1
    }
    return result
};