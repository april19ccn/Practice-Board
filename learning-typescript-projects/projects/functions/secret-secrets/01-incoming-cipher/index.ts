// Write your createCipher function here! ✨
// You'll need to export it so the tests can run it.

type Cipher = (ch: string) => string

export const createCipher  = (cipher: Cipher) => {
    return (str: string) => {
        let res = ""
        for(let ch of str){
            res += cipher(ch)
        }
        return res
    }
}