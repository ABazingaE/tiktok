package main

import (
	"log"
	follow "tiktok/kitex_gen/follow/followservice"
)

func main() {
	svr := follow.NewServer(new(FollowServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
