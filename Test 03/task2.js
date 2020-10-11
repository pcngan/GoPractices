function solution (S) {
    let numOfCharacter = 0;
    while (numOfCharacter <= S.length) {
        let mapCountSubString = {};
        for (let i = 0; i < S.length; i++) {
            if (i+numOfCharacter < S.length) {
                let element = S.substring(i, i+numOfCharacter + 1)
                if (mapCountSubString[element] != null) {
                    let value = mapCountSubString[element];
                    mapCountSubString[element] = value + 1;
                } else {
                    mapCountSubString[element] = 1;
                }
            }
        }

        for (var key in mapCountSubString){
            if (mapCountSubString[key] == 1) {
                return key.length
            }
        } 
        numOfCharacter ++;
    }
    return 0;
}

console.log(solution("abaaba"))
console.log(solution("zyzyzyz"))
console.log(solution("aabbbabaaa"))