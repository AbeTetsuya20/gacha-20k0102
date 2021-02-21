package main

import (
	"fmt"

	"AbeTetsuya20/gacha-gacha/gacha"
)

var rrea = &gacha.Rea{
	Normal: []string{"スライム1", "スライム2"},
	Rare:   []string{"オーク1", "オーク2"},
	Super:  []string{"ドラゴン1", "ドラゴン2"},
	Xrare:  []string{"イフリート1", "イフリート2"},
}

func main() {

	p := gacha.NewPlayer(10, 100)

	n := inputN(p)
	results, summary := gacha.DrawN(p, n, rrea)

	fmt.Println(results)
	fmt.Println(summary)
}

func inputN(p *gacha.Player) int {

	max := p.DrawableNum()
	fmt.Printf("ガチャを引く回数を入力してください（最大:%d回）\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		if n > 0 && n <= max {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", max)
	}

	return n
}
