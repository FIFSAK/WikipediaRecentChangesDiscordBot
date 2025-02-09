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



