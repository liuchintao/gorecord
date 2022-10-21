package main

import (
	"fmt"
	"regexp"
)

func main() {
	var (
		transferred     = regexp.MustCompile(`Transferred: (\d+) \/ (\d+), \d+%`)
		transferredSize = regexp.MustCompile(`Transferred: (\d+\.?\d*\s?\S+) \/ (\d+\.?\d*\s?\S+), .+, (\d+\.*\d*\s?\S+\/s), ETA (.+)`)
		deleted         = regexp.MustCompile(`Deleted: (\d+) \(files\), (\d+) \(dirs\)`)
	)

	cases := []string{
		"Transferred: 2.772 MiB / 2.772 MiB, 100%, 215.290 KiB/s, ETA 0s",
		"Transferred: 0 B / 0 B, -, 0 B/s, ETA -",
		"Transferred: 604 / 604, 100%",
		"Deleted: 4 (files), 1 (dirs)",
		"2022/10/12 03:02:53 INFO : Transferred/config: Copied (new)",
	}

	for _, c := range cases {
		fmt.Println(c)
		fmt.Println("[Transferred]", transferred.FindStringSubmatch(c))
		fmt.Println("[TransferredSize]", transferredSize.FindStringSubmatch(c))
		fmt.Println("[Deleted]", deleted.FindStringSubmatch(c))
	}
}
