package adapter

import "testing"

func TestAdapter(t *testing.T) {
	player := PlayerAdapter{}
	player.play("mp3", "新造的人")
	player.play("wma", "Glock shot")
	player.play("mp4", "周处除三害")
}
