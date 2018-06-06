package main

import "fmt"
import "strings"
import "unicode"

func main() {
	fmt.Println("// Compare")
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
	fmt.Println("------------")

	fmt.Println("// Contains")
	fmt.Println(strings.Contains("free", "ee"))
	fmt.Println(strings.Contains("learn", "n"))
	fmt.Println(strings.Contains("learn", "not"))
	fmt.Println(strings.Contains("learn", ""))
	fmt.Println("------------")

	fmt.Println("// ContainsAny")
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
	fmt.Println("------------")

	fmt.Println("// Count")
	fmt.Println(strings.Count("free", "e"))
	fmt.Println(strings.Count("free", ""))
	fmt.Println("------------")

	fmt.Println("// EqualFold, 部分大小写")
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println(strings.EqualFold("Java", "JAVA"))
	fmt.Println("------------")

	fmt.Println("// Fields")
	fmt.Println(" bar    baz           foo  |字符串分解后:")
	for _, v := range strings.Fields(" bar baz foo  ") {
		fmt.Printf("v = %+v\n", v)
	}
	fmt.Println("------------")

	fmt.Println("// FieldsFunc 可以传入一个函数自定义")
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	for _, v := range strings.FieldsFunc(" foo1;bar2,baz3...", f) {
		fmt.Printf("v = %+v\n", v)
	}
	fmt.Println("------------")

	fmt.Println("// HasPrefix")
	fmt.Println(strings.HasPrefix("Gopher", "g"))
	fmt.Println(strings.HasPrefix("Gopher", "c"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	fmt.Println("------------")

	fmt.Println("// HasSuffix")
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "igo"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
	fmt.Println("------------")

	fmt.Println("// Index")
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "drm"))
	fmt.Println("------------")

	fmt.Println("// Join")
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ","))
	fmt.Println(strings.Join(s, "-"))
	fmt.Println("------------")

	fmt.Println("// Spilit")
	fmt.Println(strings.Split("abc", ""))
	fmt.Println(strings.Split("a b c", " "))
	fmt.Println(strings.Split("a-b-c", "-"))

	fmt.Println("// Map")
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))

	// 首字母大写
	fmt.Println("// Title")
	fmt.Println(strings.Title("her royal highness"))

}
