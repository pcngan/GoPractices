function solution(N, K) {
    let numOfGlasses = N,
        liters = K;

    let glasses = [];

    for (let num = 0; num < numOfGlasses; num++) {
        glasses.push(num + 1)
    }

    let sum = glasses.reduce((a, b) => a + b, 0);

    if (sum/liters == 1) {
        return glasses.length;
    } else if (sum/liters < 1) {
        return -1;
    }

    let glassCount = 0;
    let deltaLiters = liters;
    for (let glassInx = glasses.length - 1; glassInx >= 0; glassInx --) {
        let delta = deltaLiters - glasses[glassInx]
        if (delta >= 0) {
            deltaLiters = delta;
            glassCount ++;
            if (deltaLiters == 0) {
                return glassCount;
            }
        }
    }
    if (deltaLiters > 0 || glassCount == 0) return -1;
    return glassCount;
}

console.log(solution(5, 8)); // expect 2
console.log(solution(4, 10)); // expect 4
console.log(solution(1, 2)); // expect -1
console.log(solution(10, 5)); // expect 1