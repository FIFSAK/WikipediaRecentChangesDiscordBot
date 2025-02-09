redis_up:
	docker run --name redis -p 6379:6379 -d redis

start:
	go build WikipediaRecentChangesDiscordBot/cmd

