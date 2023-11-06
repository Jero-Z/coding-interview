package replace

/**
请实现一个函数，将一个字符串中的空格替换成“%20”。 例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。

We Are Happy

We%20Are%20Happy

如果不考虑在原来的字符串上替换的话, 那么我们直接再开一个数组，从前往后依次赋值

遇见空格，就填上%20, 否则就填当前字符。

但是这个肯定不是面试官期待的

那么怎么在原字符串上进行高效的替换呢？
*/

func replaceStr(subject []byte, charLength int) {

	count := 0
	for i := 0; i < charLength; i++ {
		if subject[i] == ' ' {
			count++
		}
	}
	newCount := charLength + count<<1
	if cap(subject) < newCount {
		return
	}
	oldIndex, newIndex := charLength-1, newCount-1
	for oldIndex >= 0 && newIndex >= 0 {
		if subject[oldIndex] == ' ' {
			subject[newIndex] = '0'
			newIndex--
			subject[newIndex] = '2'
			newIndex--
			subject[newIndex] = '%'
			oldIndex--
		} else {
			subject[newIndex] = subject[oldIndex]
			newIndex--
			oldIndex--
		}
	}
}
