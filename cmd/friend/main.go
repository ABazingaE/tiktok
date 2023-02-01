package main

import (
	"log"
	friend "tiktok/kitex_gen/friend/friendservice"
)

func main() {
	svr := friend.NewServer(new(FriendServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
