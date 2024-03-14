package adapter

import "fmt"

// new Interface - Music play
type MusicPlayer interface {
	play(fileType string, fileName string)
	// ExistPlayerInter // donot need
}

// 比较正规的做法是写一个接口 并写struct和其方法
type ExistPlayerInter interface {
	playMp3(fileName string)
	playWma(fileName string)
}

// old Interface - Exist Player
// 此方法是直接用 struct + method = Interface
// 而省略了写 Type Interface
type ExistPlayer struct{}

func (*ExistPlayer) playMp3(fileName string) {
	fmt.Println("play mp3 :", fileName)
}

func (*ExistPlayer) playWma(fileName string) {
	fmt.Println("play wma :", fileName)
}

// Adapter
// 比较规范的做法是
//
// 在新的适配器调用 old Interface 而不是struct
type PlayerAdapter struct {
	existPlayer ExistPlayer
	// ExistPlayerInter ExistPlayerInter // 在新的struct调用旧的接口 比较规范
}

// Target: new Interface implement
//
// summary: 在新的struct(通常是adapter)中调用旧的接口(结构体和接口都行)
// 然后 Adapter.OldInterface.method() 调用
func (player *PlayerAdapter) play(fileType string, fileName string) {
	switch fileType {
	case "mp3":
		player.existPlayer.playMp3(fileName)
	case "wma":
		player.existPlayer.playWma(fileName)
	default:
		fmt.Println("Dont support this filetype to play.")
	}
}
