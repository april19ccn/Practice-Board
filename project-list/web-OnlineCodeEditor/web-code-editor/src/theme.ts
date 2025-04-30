import { createTheme, MantineColorsTuple } from '@mantine/core';

// const standard: MantineColorsTuple = [
//     '#e4f4ff',
//     '#cde3ff',
//     '#9ac4ff',
//     '#64a3ff',
//     '#3888fe',
//     '#1d76fe',
//     '#096dff',
//     '#005de4',
//     '#0052cd',
//     '#0046b5'
// ];

const standard: MantineColorsTuple = [
    "#ffe9f6",
    "#ffd1e6",
    "#faa1c9",
    "#f66eab",
    "#f24391",
    "#f02981",
    "#f01879",
    "#d60867",
    "#c0005c",
    "#a9004f"
];

export const theme = createTheme({
    colors: {
        standard,
    }
});