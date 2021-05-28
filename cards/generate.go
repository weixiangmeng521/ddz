package cards

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 一共54张牌
type CardsBoot struct {
	list [54]*Card
}

func NewCardsBoot() *CardsBoot {
	return &CardsBoot{
		list: [54]*Card{},
	}
}

// 创建牌
func (t *CardsBoot) Init() {
	cur := 0
	// 装入一般牌
	for _, v := range cardValue {
		types := [4]CardType{Spade, Heart, Diamond, Club}
		for _, tp := range types {
			t.list[cur] = NewCard(v, tp)
			cur++
		}
	}
	// 装入特殊牌
	t.list[cur] = NewCard("Jack", Freak)
	t.list[cur+1] = NewCard("Jack", Real)
}

// 展示卡牌
func (t *CardsBoot) Display() {
	for _, v := range t.list {
		fmt.Println(v.ToString())
	}
}

// 洗牌
func (t *CardsBoot) Shuffle() error {
	// 如果没有初始化牌，就洗牌
	if t.list[0] == nil {
		fmt.Print()
		return errors.New("Pls call Init method first. ")
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(t.list), func(i, j int) {
		t.list[i], t.list[j] = t.list[j], t.list[i]
	})
	return nil
}

// 获取指定卡牌
func (t *CardsBoot) getCard(index int) *Card {
	if index < 0 || index >= 54 {
		return nil
	}
	return t.list[index]
}

func (t *CardsBoot) Interator() *CardsBootIterator {
	return NewCardsBootIterator(t)
}

type CardsBootIterator struct {
	cur  int
	boot *CardsBoot
}

// 创建迭代器
func NewCardsBootIterator(boot *CardsBoot) *CardsBootIterator {
	return &CardsBootIterator{
		cur:  0,
		boot: boot,
	}
}

func (t *CardsBootIterator) Next() (card *Card) {
	card = t.boot.getCard(t.cur)
	t.cur++
	return
}

func (t *CardsBootIterator) HasNext() bool {
	return t.boot.getCard(t.cur) != nil
}
