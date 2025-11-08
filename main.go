package main

import "fmt"

func main() {
	cases := []string{
		"SRSSRRR",
		"RSSRR",
		"SSSRRRRS",
		"SRRSSR",
		"SSRSRR",
		"C",
		"R",
		"S",
		"RSSRR",
	}
	for _, c := range cases {
		fmt.Println("Input: ", c, "-> Output:", BossBabayRevenge(c))
	}
}
