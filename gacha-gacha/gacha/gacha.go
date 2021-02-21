package gacha

import (
	"math/rand"
	"time"
)

type Rea struct {
	Normal []string
	Rare   []string
	Super  []string
	Xrare  []string
}

func init() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())
}

func DrawN(p *Player, n int, r *Rea) ([]*Card, map[Rarity]int) {
	p.draw(n)

	results := make([]*Card, n)
	summary := make(map[Rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw(r)
		summary[results[i].Rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw(r *Rea) *Card {

	num := rand.Intn(100)

	//numでレア度が決まる→choiceでインデックスが決まる→そのインデックスに対応したキャラ名がreturnされる
	switch {
	case num < 80:
		choice := rand.Intn(len(r.Normal))
		return &Card{Rarity: RarityN, Name: r.Normal[choice]}
	case num < 95:
		choice := rand.Intn(len(r.Rare))
		return &Card{Rarity: RarityR, Name: r.Rare[choice]}
	case num < 99:
		choice := rand.Intn(len(r.Super))
		return &Card{Rarity: RaritySR, Name: r.Super[choice]}
	default:
		choice := rand.Intn(len(r.Xrare))
		return &Card{Rarity: RarityXR, Name: r.Xrare[choice]}
	}
}
