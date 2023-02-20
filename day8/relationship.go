package main

import "fmt"

func main() {
	test()
}

/*
	1.一个类型可以实现多个接口
	2.多个类型可以实现同一个接口（多态）
*/

type Player interface {
	playMusic()
}

type Video interface {
	playVideo()
}

type Phone struct {
	name string
}

/*
一个类型实现多个接口
	播放音乐
	播放视频
*/
func (phone Phone) playMusic() {
	fmt.Printf("播放音乐中: %v\n", phone.name)
}

func (phone Phone) playVideo() {
	fmt.Printf("播放视频中: %v\n", phone.name)
}

func test() {
	phone := Phone{
		name: "Iphone13 mini",
	}
	phone.playMusic()
	phone.playVideo()
}

/*
	多个类型实现一个接口
	手机播放音乐
	电脑播放音乐
*/
