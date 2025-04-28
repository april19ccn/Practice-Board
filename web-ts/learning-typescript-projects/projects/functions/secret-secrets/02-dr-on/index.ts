// Write your createAdvancedCipher function here! ✨
// You'll need to export it so the tests can run it.

type passwordFunc = (text: string) => string

export function createAdvancedCipher(onVowel: passwordFunc, onConsonant: passwordFunc, onPunctuation: passwordFunc) {
    return (text: string) => {
        let res = ""
        for (let ch of text) {
            if (ch.match(/[aeiou]/i)) {
                res += onVowel(ch)
            }
            else if (ch.match(/[bcdfghjklmnpqrstvwxyz]/i)) {
                res += onConsonant(ch)
            }
            else {
                res += onPunctuation(ch)
            }
        }
        return res
    }
}