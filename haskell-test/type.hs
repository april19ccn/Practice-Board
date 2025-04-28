removeNonUppercase :: [Char] -> [Char]  
removeNonUppercase st = [ c | c <- st, c `elem` ['A'..'Z']]

addThree :: Int -> Int -> Int -> Int  
addThree x y z = x + y + z

factorial :: Integer -> Integer  
factorial n = product [1..n]

circumference :: Float -> Float  
circumference r = 2 * pi * r

circumference' :: Double -> Double  
circumference' r = 2 * pi * r

*Note*: 判断相等的==运算子是函数，``+-*/``之类的运算子也是同样。
在缺省条件下，它们多为中缀函数。
若要检查它的型别，就必须得用括号括起使之作为另一个函数，或者说以首码函数的形式调用它。