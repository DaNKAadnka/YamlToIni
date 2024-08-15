package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func test(in *bufio.Reader, out *bufio.Writer) {

	n := readInt(in)
	s := make([]ent, n)
	for i := 0; i < n; i++ {
		sa := readEnt(in)
		s[i] = sa
		fmt.Println(sa)
	}
	currentLevel := 0
	var ans bytes.Buffer
	var level []string
	var queue []ent
	writeAns := func() {
		if len(level) > 0 {
			ans.WriteString(getLevel(level))
			ans.WriteRune('\n')
		}
		for _, e := range queue {
			ans.WriteString(kv(e))
		}
		ans.WriteRune('\n')
		queue = nil
	}
	for i := 0; i < n; i++ {
		if s[i].Level != currentLevel {
			// Drop buffer
			// First - level, than - queue values

			// Change level
			if currentLevel > s[i].Level {
				if len(queue) > 0 {
					writeAns()
				}
				for currentLevel != s[i].Level {
					level = level[:len(level)-1]
					currentLevel--
				}
			}
			currentLevel = s[i].Level
		}
		if s[i].Value == "" {
			if len(queue) > 0 {
				writeAns()
			}

			level = append(level, s[i].Key)
		} else {
			queue = append(queue, s[i])
		}
	}
	if len(queue) > 0 {
		writeAns()
	}
	fmt.Fprint(out, ans.String())
}

func getLevel(s []string) string {
	var res bytes.Buffer
	res.WriteRune('[')
	for i, c := range s {
		res.WriteString(c)
		if i != len(s)-1 {
			res.WriteRune('.')
		}
	}
	res.WriteRune(']')
	return res.String()
}

func kv(e ent) string {
	return e.Key + " = " + e.Value + "\n"
}

type ent struct {
	Key, Value string
	Level      int
}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func readEnt(in *bufio.Reader) ent {
	line, _ := in.ReadString('\n')
	space := 0
	for _, c := range line {
		if c != ' ' {
			break
		}
		space++
	}
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, " ", "")
	line = strings.ReplaceAll(line, "\t", "")
	line = strings.ReplaceAll(line, "\n", "")
	s := strings.Split(line, ":")

	return ent{Key: s[0], Value: s[1], Level: space / 4}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	t := readInt(in)
	for ; t > 0; t-- {
		test(in, out)
	}
}
