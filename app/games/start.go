package games

import "ddz/app/scene"

// 开始游戏
func Start() {
	game := NewGame()
	scene.CreateSceneFlow(game)

}
