// Write your deepDifferences function here! ✨
// You'll need to export it so the tests can run it.

export function deepDifferences(a: string[][], b: string[][]): ((string | undefined)[] | undefined)[] {
    if (a.length !== b.length) {
        return Array(Math.max(a.length, b.length)).fill(undefined)
    }

    const allResults: ((string | undefined)[] | undefined)[] = []
    for (let i = 0; i < a.length; i++) {
        if (a[i].length !== b[i].length) {
            allResults.push(undefined)
            continue
        }

        const temp: (string | undefined)[] = []
        for (let j = 0; j < b[i].length; j++) {
            a[i][j] === b[i][j] ? temp.push(a[i][j]) : temp.push(undefined)
        }

        allResults.push(temp)
    }

    return allResults
}

console.log(deepDifferences([["a","g"],["g","t"]], [["a","c"],["g","t","c"]]))