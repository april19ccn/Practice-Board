package inset

import (
	"bytes"
	"fmt"
)

// IntSet 是一组小的非负整数。
// 其零值表示空集。
type IntSet struct {
	words []uint64
}

const WORDBITS = 32 << (^uint(0) >> 63)

// const WORDBITS = 32

// Has 报告集是否包含非负值 x。
func (s *IntSet) Has(x int) bool {
	word, bit := x/WORDBITS, uint(x%WORDBITS)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add 将非负值 x 添加到集合中。
func (s *IntSet) Add(x int) {
	word, bit := x/WORDBITS, uint(x%WORDBITS)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith 将 s 设置为 s 和 t 的并集。
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String 以 “{1， 2， 3}” 形式的字符串形式返回集合。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WORDBITS; j++ {
			if word&(1<<uint(j)) != 0 {
				// 在首次添加元素之前，缓冲区里只有左花括号 {，此时 buf.Len() 等于 1，所以不会添加空格。
				// 当添加后续元素时，buf.Len() 会大于 1，这时就会添加空格，从而确保元素之间用空格分隔。
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}

				// fmt.Fprintf 会把格式化之后的字符串输出到 buf 这个缓冲区中。这里的格式化字符串 "%d" 意味着要以十进制整数的格式来输出。
				// 之所以使用 Fprintf 而不是直接用 buf.WriteString(strconv.Itoa(WORDBITS*i+j))，是因为 Fprintf 能够更便捷地处理整数到字符串的转换工作。
				fmt.Fprintf(&buf, "%d", WORDBITS*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

/*
练习6.1: 为bit数组实现下面这些方法
func (*IntSet) Len() int      // return the number of elements
func (*IntSet) Remove(x int)  // remove x from the set
func (*IntSet) Clear()        // remove all elements from the set
func (*IntSet) Copy() *IntSet // return a copy of the set
*/
// return the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}

		for i := 0; i < WORDBITS; i++ {
			if word&(1<<uint(i)) != 0 {
				count++
			}
		}
	}
	return count
}
func (s *IntSet) Remove(x int) {
	// x:   0101 1100
	// y:   0011 0100
	// ^y:  1100 1011    // y 按位取反
	// x&^y: 0100 1000   // x 中对应 y 为 1 的位（第2、3、5位）被置零

	word, bit := x/WORDBITS, uint(x%WORDBITS)
	if word < len(s.words) {
		// 如果没有越界, 则将word的第bit位置为0
		s.words[word] &= ^(1 << bit)
	}
}
func (s *IntSet) Clear() {
	s.words = nil
}
func (s *IntSet) Copy() *IntSet {
	t := &IntSet{}
	t.words = make([]uint64, len(s.words))
	copy(t.words, s.words)
	return t
}

/*
练习 6.2： 定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)。
*/
func (s *IntSet) AddAll(x ...int) {
	for _, i := range x {
		s.Add(i)
	}
}

/*
练习 6.3： (*IntSet).UnionWith会用|操作符计算两个集合的并集，我们再为IntSet实现另外的几个函数IntersectWith（交集：元素在A集合B集合均出现），
DifferenceWith（差集：元素出现在A集合，未出现在B集合），SymmetricDifference（并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A）。
*/
// 交集：元素在A集合B集合均出现
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words { // 在 Go 语言里，当使用 for...range 循环遍历切片时，迭代变量 x 确实是元素的一个副本，并非元素本身。所以不用循环值。
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// 差集：元素出现在A集合，未出现在B集合
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

// 并查集：元素出现在A但没有出现在B，或者出现在B没有出现在A
func (s *IntSet) SymmetricDifference(t *IntSet) {
	temp := s.Copy()
	s.UnionWith(t)
	temp.IntersectWith(t)
	s.DifferenceWith(temp)
}

/*
练习 6.4： 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作。
*/
func (s *IntSet) Elems() []int {
	result := []int{}

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WORDBITS; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, WORDBITS*i+j)
			}
		}
	}
	return result
}

/*
练习 6.5： 我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。
修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。
当然了，这里我们可以不用简单粗暴地除64，可以定义一个常量来决定是用32还是64，这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)
*/

// 1. ^uint(0) 的作用
// 	在 Go 语言里，^ 是按位取反运算符。^uint(0) 会生成一个全为 1 的 uint 类型数值。uint 类型的位数和具体架构相关，32 位系统上是 32 位，64 位系统上则是 64 位。
// 		32 位系统中：^uint(0) 得到的是 0xFFFFFFFF，即 2^32 - 1。
// 		64 位系统中：^uint(0) 得到的是 0xFFFFFFFFFFFFFFFF，也就是 2^64 - 1。

// 2. (^uint(0) >> 63) 的计算逻辑
// 	右移操作符 >> 会把二进制位向右移动指定的位数。
// 		32 位系统中：^uint(0) 是 32 位的 0xFFFFFFFF，右移 63 位后结果为 0。
// 		64 位系统中：^uint(0) 是 64 位的 0xFFFFFFFFFFFFFFFF，右移 63 位后结果为 1。

// 3. 32 << (^uint(0) >> 63) 的计算结果
// 	左移操作符 << 会让二进制位向左移动指定的位数。
// 		32 位系统中：32 << 0 的结果是 32。
// 		64 位系统中：32 << 1 的结果是 64。
