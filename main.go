package main

import (
	"music/core"
	"sync"
)

func main() {
	mainNotes := `
	7-  1   2   3   0   5-  5   3   0   0   0   0   0   0   0   0
	7-  1   2   3   0   5-  5   3   2   3   1   2   7-  1   5-  0
	7-  1   2   3   0   5-  5   3   0   0   0   0   0   0   0   0
	7-  1   2   3   0   5-  5   3   2   3   1   2   7-  1   5-  0
	7   1+  2+  3+  0   5   5+  3+  0   0   0   0   0   0   0   0
	7   1+  2+  3+  0   5   5+  3+  2+  3+  1+  2+  7   1+  5   0
	7   1+  2+  3+  0   5   5+  3+  0   0   0   0   0   0   0   0
	2   0   0   0   0   0   0   0   1   0   0   0   0   0   0   0
	2   0   0   1   2   0   0   1   2   0   3   0   5   0   3   0
	2   0   0   1   2   0   0   1   2   3   2   1   6-  0   0   0
	2   0   0   1   2   0   0   1   2   0   3   0   5   0   3   0
	2   0   0   3   2   0   1   2   2   0   0   0   0   0   0   0
	2   0   0   1   2   0   0   1   2   0   3   0   5   0   3   0
	2   0   0   3   2   0   1   0   6-  0   0   0
	3   2   1   2   1   0   0   0
	3   2   1   2   1   0   0
	5-  3   2   1   2   0   0   1   0   0   0   0   0
	1   0   2   0   3   0   1   0   6   0   5   6   0   0   0
	2   7   0   6   7   0   0   0   0
	7   0   6   7   0   0   3   0   1+  2+  1+  7   6   0   0
	5   6   0   5   6   0   5   6   5   6   0   5   1   0   5   0   3   3   0   0   0   0   0   0   0
	1   0   2   0   3   0   1   0   6   0   5   6   0   0   0
	2   7   0   6   7   0   0   0   0
	7   0   6   7   0   0   3   0   1+  2+  1+  7   6   0   0
	5   6   0   3+  3+  0   0   5   0   6   0   3+  3+  0
	5   0   6   6   0   3-  0   3-  0   3-  0   3-  0   0   0
	1+  0   2+  0   3+  0   6+  5+  0   0   6+  5+  0   0   6+  5+  0   2+  0   0
	3+  0   6+  5+  0   0   6+  5+  0   0   6+  5+  0   3+  0   0
	2+  0   1+  6   0   1+  0   1+  2+  0   1+  6   0   0   1+  0   3+  0   0   0   0   0   3+  0   2+  0   0   0
	1+  0   2+  0   3+  0   6+  5+  0   0   6+  5+  0   0   6+  5+  0   0
	2+  0   3+  0   6+  5+  0   0   6+  5+  0   0   6+  5+  0   0
	3+  0   2+  0   1+  6   0   0   3+  0   2+  0   1+
	6   0   1+  0   0   1+  0   0   0   0   0   0   0   0   0   0   0
	6   3+  0   0   2+  0   1+  6   0   3+  0   0   2+  0   1+
	6   0   1+  0   0   1+  0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
	7   1+  2+  3+  0   5   5+  3+  2+  3+  7   1+  6   7   5   0
	7   1+  2+  3+  0   5   5+  3+  0   0   0   0   0   0   0   0
	6+  3+  2+  6   3   6   2+  3+  6+  0   0   0   0   0   0   0
	  2   0   0   1   2   0   0   1   2   0   3   0   5   0   3   0
	2   0   0   1   2   0   0   1   2   3   2   1   5-  0   0   0
	2   0   0   1   2   0   0   1   2   0   3   0   5   0   3   0
	2   0   0   3   2   0   1   0   2   0   0   0   0   0   0   0
	2   0   0   1   2   0   0   1   2   0   3   0   5   0   3   0
	2   0   0   3   2   0   1   0   6-  0   0   0   3   2   1   2
	1   0   0   0   3   2   1   2   1   0   0   5-  3   2   1   2
	1   0   0   0   0   0   0   0   1   0   2   0   3   0   1   0
	6   0   5   6   0   0   0   1   7   0   6   7   0   0   0   0
	7   0   6   7   0   0   3   0   1+  2+  1+  7   6   0   5   0
	6   0   5   6   0   5   6   5   6   0   5   2   0   5   0   0
	3   0   0   0   0   0   0   0   1   0   2   0   3   0   1   0
	6   0   5   6   0   0   0   1   7   0   6   7   0   0   0   0
	7   0   6   7   0   0   3   0   1+  2+  1+  7   6   0   5   0
	6   0   3+  3+  0   0   5   0   6   0   3+  3+  0   5   0   6
	6   0   0   0   0   0   0   0   0   0   0   0   1+  0   2+  0
	3+  0   6+  5+  0   0   6+  5+  0   0   6+  5+  0   0   2+  3+
	0   0   6+  5+  0   0   6+  5+  0   0   6+  5+  0   3+  0   0
	2+  0   1+  6   0   1+  0   0   2+  0   1+  6   0   1+  0   0
	3+  0   0   0   0   4+  3+  0   3+  2+  0   0   1+  0   2+  0
	3+  0   6+  5+  0   0   6+  5+  0   0   6+  5+  0   0   2+  0
	3+  0   6+  5+  0   0   6+  5+  0   0   6+  5+  0   3+  0   0
	2+  0   1+  6   0   3+  0   0   2+  0   1+  6   0   1+  0   0
	1+  0   0   0   0   0   0   0   0   0   0   0   6   3+  0   0
	2+  0   0   0   1+  0   6   0   0   0   3+  0   0   0   0   0
	2+  0   0   0   1+  0   6   0   0   0   1+  0   0   0   0   0
	1+  0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
	0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
	`

	accompaniments := `
	4-- 0   1-  0   3-  0   0   0   5-- 0   7-- 0   2-  0   0   0
	3-- 0   7-- 0   2-  0   0   0   6-- 0   1-  0   3-  0   0   0
	4-- 0   1-  0   3-  0   0   0   5-- 0   7-- 0   2-  0   0   0
	3-- 0   7-- 0   2-  0   0   0   6-- 0   1-  0   3-  0   0   0
	4-  0   1-  0   3-  0   0   0   5-- 0   7-- 0   2-  0   0   0
	3-- 0   7-- 0   2-  0   0   0   6-- 0   1-  0   3-  0   0   0
	4-- 0   1-  0   3-  0   0   0   5-- 0   7-- 0   2-  0   0   0
	3-  0   0   0   0   0   0   0   1-  0   0   0   0   0   0   0
	1-- 0   0   0   3-  0   0   0   3-  0   0   0   3-  0   0   0
	7-- 0   0   0   2-  0   0   0   2-  0   0   0   2-  0   0   0
	7-- 0   0   0   2-  0   0   0   2-  0   0   0   2-  0   0   0   4--
	0   0   0   4-  0   0   0   4-  0   0   0   4-  0   0   0
	2-- 0   0   0   2-  0   0   0   5-- 0   0   0   2-  0   0   0   6-- 0   0   0   3-  0   0   0   6-- 0   0   0
	0   0   0   0   4-- 0   0   0
	0   0   0   0   4-  0   0
	0   0   0   0   0   1-- 0   5-- 0   1-  0   3-  0
	1   0   0   0   1-  0   0   0   4-- 1-  4-  6-  1   0   4-
	0   5-- 2-  5-  7-  2   0   5-  0
	3-- 7-- 3-  5-  7-  0   0   0   6-- 3-  6-  3-  1   0   0
	0   4-- 1-  4-  6-  1   0   4-  0   5-- 2-  5-  7-  2-  0   5-  0   1-- 5-- 1-  3-  5-  0   3-  0
	1   0   0   0   5-  0   0   0   4-- 1-  4-  1-  6-  0   1-
	0   5-- 2-  5-  2-  7-  0   2-  0
	3-- 7-- 3-  5-  7-  0   3-  0   6-- 3-  6-  3-  1   0   3-
	0   4-- 1-  4-  6-  3   0   4-  0   5-- 2-  5-  7-  2   0
	5-  0   6-- 0   6-- 0   6-- 0   6-- 0   6-- 0   0   0
	0   0   0   0   4-- 0   1-  0   4-  0   0   0   5-- 0   2-  0   5-  0   0   0
	3-- 0   7-- 0   3-  0   0   0   6-- 0   3-  0   6-  0   0   0
	4-- 0   1-  0   4-  0   0   0   5-- 0   2-  0   5-  0   0   0
	1-  0   5-  0   1   0   0   0   3-  0   7-  0
	3   0   0   0   4-- 0   1-  0   4-  0   0   0   5-- 0   2-  0   5-  0   0
	0   3-- 0   7-- 0   3-  0   0   0   6-- 0   3-  0   6-  0
	0   0   4-- 0   1-  0   4-  0   0   0   5-- 0   2-  0   5-  0   0
	0   1-  0   5-  0   1   0   5-  0   3-  0   0   0
	0   0   0   0   4-- 0   0   0   0   0   0   0   5-- 0   0   0   0   0   0   0
	4-- 0   1-  0   6-  0   1-  0   5-- 0   2-  0   5-  0   2-  0
	3-- 0   7-- 0   5-  0   7-- 0   6-- 0   3-  0   1   0   3-  0
	4-- 0   1-  0   6-  0   1-  0   5-- 0   2-  0   7-  0   2-  0
	6-- 0   3-  0   6-  0   3-  0   1   0   0   0   3-  0   0   0
	  1-- 0   0   0   3-  0   0   0   3-  0   0   0   3-  0   0   0
	7-- 0   0   0   2-  0   0   0   2-  0   0   0   2-  0   0   0
	7-- 0   0   0   2-  0   0   0   2-  0   0   0   2-  0   0   0
	4-- 0   0   0   4-  0   0   0   4-  0   0   0   4-  0   0   0
	2-- 0   0   0   2-  0   0   0   5-- 0   0   0   2-  0   0   0
	6-- 0   0   0   3-  0   0   0   6-- 0   0   0   0   0   0   0
	4-- 0   0   0   0   0   0   0   5-- 0   0   0   0   0   0   0
	1-  0   0   0   0   0   0   0   1-  0   0   0   0   0   0   0
	4-- 0   0   0   0   0   0   0   5-- 0   0   0   0   0   0   0
	3-- 0   0   0   0   0   0   0   6-- 0   0   0   0   0   0   0
	4-- 0   0   0   0   0   0   0   5-- 0   0   0   0   0   0   0
	1-  0   0   0   0   0   0   0   1-  0   0   0   0   0   0   0
	4-- 0   0   0   0   0   0   0   5-- 0   0   0   0   0   0   0
	3-- 0   0   0   0   0   0   0   6-- 0   0   0   0   0   0   0
	4-- 0   0   0   0   0   0   0   5-- 0   0   0   0   0   0   0
	6-- 0   0   0   3-  0   0   0   6-  0   0   0   0   0   0   0
	4-- 0   0   0   4-  0   0   0   5-- 0   0   0   5-  0   0   0
	3-- 0   0   0   3-  0   0   0   6-- 0   0   0   6-  0   0   0
	4-- 0   0   0   4-  0   0   0   5-- 0   0   0   5-  0   0   0
	1-- 0   0   0   1-  0   0   0   3-- 0   0   0   3-  0   0   0
	4-- 0   0   0   4-  0   0   0   5-- 0   0   0   5-  0   0   0
	3-- 0   0   0   3-  0   0   0   6-- 0   0   0   6-  0   0   0
	4-- 0   0   0   4-  0   0   0   5-- 0   0   0   5-  0   0   0
	6-- 0   0   0   3-  0   0   0   6-- 0   0   0   0   0   0   0
	4-- 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
	5-- 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
	0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
	0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
	`

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		player := core.AudioPlay{Times: 400}
		player.LoadNotes(mainNotes)
		player.Play()
		defer wg.Done()
	}()

	go func() {
		player := core.AudioPlay{Times: 400}
		player.LoadNotes(accompaniments)
		player.Play()
		defer wg.Done()
	}()

	wg.Wait()
}
