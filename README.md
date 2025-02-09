# Wikipedia RecentChanges Discord Bot

This is a Discord bot that sends a message to a Discord channel whenever a
new edit is made to a Wikipedia page. The bot uses the Wikipedia's
[endpoint](https://stream.wikimedia.org/v2/stream/recentchange)
to get the recent changes and sends a message to the Discord channel using
the Discord Bot.

## Installation

### Prerequisites
- Docker, docker compose
---

### Steps

1. **Clone the repository:**

   ```bash
   git clone https://github.com/FIFSAK/WikipediaRecentChangesDiscordBot.git
   cd WikipediaRecentChangesDiscordBot
   ```

2. **Run and build the bot and Redis:**

   Start the bot using:
   ```bash
   docker-compose up --build
   ```

Note: You can use [Makefile](./Makefile) to run the bot.

## Architecture

- [bot](./bot) contains the discord bot setup and its handlers
- [services](./services) contains the services like Wikipedia and Redis
- [cmd](./cmd) contains the main file to run the bot
- [config](./config) contains the configuration for the application

## Additional
- CI/CD [pipeline](https://gitlab.com/anuar200572/WikipediaRecentChangesDiscordBot/-/blob/master/.gitlab-ci.yml?ref_type=heads)
- The bot filtering changes by fixed languages they can be easily changes by adding
  them [here](https://github.com/FIFSAK/WikipediaRecentChangesDiscordBot/blob/master/services/wikipedia/listener.go#L20)
- The bot stores up to the last 10
  changes ([change it here](https://github.com/FIFSAK/WikipediaRecentChangesDiscordBot/blob/master/services/wikipedia/listener.go#L17))
  and sends them all, ensuring the total message length does not exceed Discord's 2000-character limit.

## How it can be scaled

- Kafka is an ideal fit for this system as it is inherently designed to work with a pull model.
- The project can be scaled by integrating Kafka as the system for processing Wikipedia
Recent Changes. 
- All recent changes can be published to a single Kafka topic, and for optimization (optional), data can
be partitioned by language. 
- The bot will subscribe to this topic and, upon user command, use the pull model to fetch
messages from Kafka. 
- A similar approach can be implemented with Redis. Redis could aggregate change statistics and
provide quick responses to requests such as `!stats`. 
- Kafka is specifically built to handle large volumes of data, that's why
it ideal for such tasks. 
- This architecture will create good base for horizontal scalability in the future.

