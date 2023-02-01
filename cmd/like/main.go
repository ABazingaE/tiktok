package main

import (
	"log"
	like "tiktok/kitex_gen/like/likeservice"
)

func main() {
	svr := like.NewServer(new(LikeServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
