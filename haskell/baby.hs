doubleMe :: Num a => a -> a
doubleMe x = x + x

doubleUs :: Num a => a -> a -> a
doubleUs x y = doubleMe x + doubleMe y

doubleSmallNumber :: (Ord a, Num a) => a -> a
doubleSmallNumber x = if x > 100
                      then x
                      else  x*2

doubleSmallNumber' :: (Ord a, Num a) => a -> a
doubleSmallNumber' x = (if x > 100 then x else x*2) + 1

conanO'Brien :: String
conanO'Brien = "It's a-me, Conan O'Brien!"

-- *Note*: 在 ghci 下，我们可以使用 ``let`` 关键字来定义一个常量。在 ghci 下执行 ``let a=1`` 与在脚本中编写 ``a=1`` 是等价的。
a :: Integer
a = 1

-- *Note*: ``[],[[]],[[],[],[]]`` 是不同的。第一个是一个空的 List，第二个是含有一个空 List 的 List，第三个是含有三个空 List 的 List。

head :: [a] -> a
head y = y !! 2