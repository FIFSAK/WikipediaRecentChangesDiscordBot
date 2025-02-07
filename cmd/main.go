package main

import "WikipediaRecentChangesDiscordBot/api"

func main() {
	//err := config.ReadConfig()
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//bot.Start()
	//
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//<-c
	//
	//return
	api.ListenToWikipediaChanges()

}
