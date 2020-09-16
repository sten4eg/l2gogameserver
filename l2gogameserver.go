package main

import "l2gogameserver/gameserver"

func main()  {
 x := gameserver.New()
 x.Init()
 x.Start()
}