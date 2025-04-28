export const lottery = () => {
    return Math.random() > 0.5 ? "WIN" : "LOSE";
}