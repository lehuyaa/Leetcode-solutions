/**
 * @param {string} allowed
 * @param {string[]} words
 * @return {number}
 */
var checkConsistent = (allowed, words) => {
    for(let i = 0 ; i<words.length;i++){
        if(!allowed.includes(words[i])){
            return false;
        }
    }
    return true;
}

var countConsistentStrings = function(allowed, words) {
    let count =0;
    for(const item of words) {
            if(checkConsistent(allowed, item))count++;
    }
    return count;
};