// Write your alignTexts function here! ✨
// You'll need to export it so the tests can run it.

export type AlignmentOptions = {
	align?: "left" | "middle" | "right";
	width: number;
};

function handleSplitLine(section: string, width: number): string[] {
    const strArray = section.split(" ");

    const result: string[] = []
    while (strArray.length > 0) {
        let str = strArray.shift() as string

        while (strArray.length > 0 && str.length < width) {
            const tempStr = str + " " + strArray[0]
            if (tempStr.length > width) {
                break
            } else {
                strArray.shift()
                str = tempStr
            }
        }

        result.push(str)
    }

    return result
}

function handleAlignLine(split: string[], width: number, align: "left" | "middle" | "right" = "left"): string[] {
    const result = split.map(function (line: string) {
        const count = width - line.length < 0 ? 0 : width - line.length
        if (align === "left") {
            return line + " ".repeat(count)
        }
        if (align === "right") {
            return " ".repeat(count) + line
        }
        if (align === "middle") {
            return " ".repeat(Math.floor((count) / 2)) + line + " ".repeat(Math.ceil((count) / 2))
        }
    })

    return result as string[]
}

export function alignTexts(text: string[], options: AlignmentOptions): string[][] {
    const formatText = text.map(function (section: string) {
        return handleAlignLine(handleSplitLine(section, options.width), options.width, options.align)
    })
    return formatText
}



console.log(alignTexts(["ab cd", "abc def", "a bc def"], { align: "right", width: 4 }))

// Output:
// [
// 	["  ab", "  cd"],
// 	[" abc", " def"],
// 	["a bc", " def"]
// ]