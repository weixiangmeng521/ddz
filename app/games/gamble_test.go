package games

import (
	"ddz/app/cards"
	"ddz/app/constant"
	"ddz/app/players"
	"testing"
)

func PrepareGame(t *testing.T) constant.GameInterface {
	game := NewGame("test1")
	game.Debug()

	people := []constant.PlayerInterface{}
	peopleName := []string{
		"Kenny", "Kyle", "Stan",
	}
	// 创建三个玩家
	for i := 0; i < 3; i++ {
		p := players.NewPlayer(peopleName[i])
		people = append(people, p)
	}
	// 加入游戏
	for _, p := range people {
		game.JoinPlayer(p)
	}
	// 三个玩家准备好游戏
	for _, p := range people {
		p.SetState(constant.Already)
	}
	// 游戏是否能开始
	if !game.CanStart() {
		t.Error("Cannot start this game.")
	}

	// *** 叫地主环节 ***

	// 三个玩家都叫地主
	for _, p := range people {
		if p.GetName() == "Kenny" {
			p.CallLord()
			game.Turn()
			continue
		}
		p.NotCall()
		game.Turn()
	}
	return game
}

func TestReadyGame(t *testing.T) {
	game := PrepareGame(t)
	// *** 出牌环节 ***
	// cur := game.GetCurPlayer()
	list := []*cards.Card{
		cards.NewCard("3", cards.Spade),
		cards.NewCard("3", cards.Club),
	}
	// Kenny 对三
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Kyle 不出
	list = []*cards.Card{}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Stan 对五
	list = []*cards.Card{
		cards.NewCard("5", cards.Spade),
		cards.NewCard("5", cards.Club),
	}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Kenny 不出
	list = []*cards.Card{}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Kyle 不出
	list = []*cards.Card{}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Stan 4
	list = []*cards.Card{
		cards.NewCard("4", cards.Heart),
	}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Kenny 5
	list = []*cards.Card{
		cards.NewCard("5", cards.Heart),
	}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Kyle 6
	list = []*cards.Card{
		cards.NewCard("6", cards.Heart),
	}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Stan
	list = []*cards.Card{
		cards.NewCard("2", cards.Diamond),
	}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Kenny boom
	list = []*cards.Card{
		cards.NewCard("Jack", cards.Real),
		cards.NewCard("Jack", cards.Freak),
	}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Kyle 不出
	list = []*cards.Card{}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Stan 不出
	list = []*cards.Card{}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	// Kenny 出对6
	list = []*cards.Card{
		cards.NewCard("6", cards.Spade),
		cards.NewCard("6", cards.Club),
	}
	if err := game.DealCards(list); err != nil {
		t.Error(err)
		return
	}

	game.Display()

}
