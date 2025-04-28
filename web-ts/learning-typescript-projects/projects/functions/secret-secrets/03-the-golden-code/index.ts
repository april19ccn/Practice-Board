// Write your createCodeCracker function here! ✨
// You'll need to export it so the tests can run it.

type AttackKey = {
    attempts: number,
    makeGuess: (text: string, attempt: number) => string,
    validateGuess: (guess: string) => boolean,
}

export function createCodeCracker(key: AttackKey) {
    return (text: string) => {
        for (let i = 0; i < key.attempts; i++) {
            const getKey = key.makeGuess(text, i)
            if ( key.validateGuess(getKey)) {
                return getKey
            }
        }
        return undefined
    }
}