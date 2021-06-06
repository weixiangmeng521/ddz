package constant

// 游戏钩子
type GameHookType int32

const (
	GAME_JOINED_PLYAER        = "GAME_JOINED_PLYAER"        // 玩家加入游戏时
	GAME_LEAVED_PLYAER        = "GAME_LEAVED_PLYAER"        // 玩家离开游戏时
	GAME_PLAYER_STATE_CHANGED = "GAME_PLAYER_STATE_CHANGED" // 玩家信息更新时
	GAME_CARDS_CHANGED        = "GAME_CARDS_CHANGED"        // 洗牌成功后
	GAME_TURN_CHANGED         = "GAME_TURN_CHANGED"         // 出牌权移交后
	GAME_STATE_CHANGED        = "GAME_STATE_CHANGED"        // 游戏状态发生变化
	GAME_PLAYER_LEAVED        = "GAME_PLAYER_LEAVED"        // 当游戏里面某个玩家离开房间
	GAME_PLAYER_PLAYED_CARDS  = "GAME_PLAYER_PLAYED_CARDS"  // 玩家出完牌后
)

// 游戏状态
type GameState int8

const (
	GameReady   GameState = iota // 准备
	GameStarted                  // 游戏开始后
	GameCalled                   // 叫完地主后
	GameEnd                      // 游戏结束
)

func (t GameState) ToString() string {
	m := map[GameState]string{
		GameReady:   "GameReady",
		GameStarted: "GameStarted",
		GameCalled:  "GameCalled",
		GameEnd:     "GameEnd",
	}
	return m[t]
}
