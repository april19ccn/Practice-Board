import { Either, IO, compose, map, toUpperCase } from "../../utils/support.js";
import Task from "data.task";

import fs from "fs";

// readFile :: String -> IO String
const readFile = filename => new IO(() => {
    console.log("**********")
    fs.readFileSync(filename, 'utf-8')
});

// print :: String -> IO String
const print = x => new IO(() => {
  console.log(x);
  return x;
});

// cat :: String -> IO (IO String)
const cat = compose(map(print), readFile);

console.log(cat('').unsafePerformIO());

console.log(cat('d:/Star_Code/A_Study/unocss-practice/fp-test/examples/monad/test.txt').unsafePerformIO().unsafePerformIO());
// IO(IO('[core]\nrepositoryformatversion = 0\n'))