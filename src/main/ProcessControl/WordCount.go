package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
词频计数器
*/
func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage : %s <file1> [<file2>[...<fileN>]]", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	// 创建一个空映射用于以词为键,记录出现的次数
	frequencyWord := make(map[string]int)

	// 读取从命令行中获取到的文件,这里的文件是用绝对路径标注的
	// 遍历每一个文件,记录他们的词频
	for _, filename := range commandLineFiles(os.Args[1:]) {
		updateFrequencies(filename, frequencyWord)
	}
	reportByWords(frequencyWord)
	wordsForFrequency := invertStringIntMap(frequencyWord)
	reportByFrequency(wordsForFrequency)
}

// 用于简单处理文件通配符的情况
func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name)
			} else if matches != nil {
				args = append(args, matches...) //将匹配的文件名都放进去
			}
		}
		return args
	}
	return files
}

// 这个函数是用于处理文件的
// 他会打开指定的文件,并在函数返回的时候关闭对应的文件
func updateFrequencies(filename string, frequencyWord map[string]int) {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	//defer用于释放资源,在函数任何地方声明的defer都会在函数结束时进行调用
	defer file.Close()
	readAndUpdateFre(bufio.NewReader(file), frequencyWord)
}

// 使用换行符作为分割进行文件读取
// 当文件读到结尾的时候,就会退出
func readAndUpdateFre(reader *bufio.Reader, fre map[string]int) {
	for {
		line, err := reader.ReadString('\n')
		// SplitOnNonLetters会忽略非单词的字符
		// TrimSpace过滤掉开头和结尾的空白
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				fre[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				fmt.Println("无法完成读文件操作:", err)
			}
			//文件读取完成后退出
			break
		}
	}
}

// 识别非单词的字符主要原理是调用FieldsFunc来判断每一个rune是否为字母
// 如果是字符,就返回false,如果不是则返回true
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}

func invertStringIntMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string)
	for k, v := range intForString {
		stringsForInt[v] = append(stringsForInt[v], k)
	}
	return stringsForInt
}

func reportByWords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0
	for k, v := range frequencyForWord {
		words = append(words, k)
		if width := utf8.RuneCountInString(k); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(v)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	//用*s可以打印固定长度的空白
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth, frequencyForWord[word])
	}
}

func reportByFrequency(wordForFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordForFrequency))
	for frequency := range wordForFrequency {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency -> Words")
	for _, frequency := range frequencies {
		words := wordForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ","))
	}
}
