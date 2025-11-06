TheRedDevilsData-API
-
TheRedDevilsData-API is a backend service dedicated to collecting, storing, and serving football data for Manchester United.
It fetches historical data from the API-Football provider, processes and stores it in a PostgreSQL database, and exposes endpoints for easy consumption by any client (web, CLI, or future analytics tools).

Features
-
* Data fetching
  * Fetch detailed data for Manchester United (fixtures, results, statistics, injuries, squad, etc.)
  * Fetch a bit of a general data about available countries for football data and all venues in England
  * Built-in CLI tool using Cobra for triggering data fetches from API-Football
  * Modular client design that can easily be extended to support other teams
* Data storage
  * Store fetched football data in PostgreSQL
  * Structured models
* API service
  * Serve structured JSON responses for football data
  * Endpoints ready to be consumed by any external client (web, CLI, or other apps)

Technologies used
-
* Go
* GORM
* PostgreSQL
* Docker
* Cobra CLI
* API-Football

Running guide
-
* Prerequisites
  * Go installed
  * Docker installed
  * PostgreSQL installed
* Locally
  * Clone the repo
  * Run go mod download
  * Set up PostgreSQL and ensure a database named thereddevilsdata exists
  * Create and adjust the .env file with your database credentials and API-Football key
  * Run
    * go run cli/main.go {command) -> to fetch data via CLI
    * go run api/main.go -> to start REST API
* Docker
  * Run docker compose up
  * Execute in terminal docker compose exec cli ./theRedDevilsData-cli {command} -> to fetch data via CLI

Workflow
-
* Fetch data -> Use the CLI to fetch Manchester United data from API-Football
* Store data -> Service saves new and updated information into PostgreSQL
* Serve data –> The API exposes endpoints returning football data in JSON format
* Extend –> Can easily expand to include other teams from the same API provider

Commands CLI
-
* fetch-countries -> Fetch and save  available countries for football data from API-Football
* fetch-all-leagues-for-team -> Fetch and save all leagues in which Manchester United has played at least one match
* fetch-team -> Fetch and save Manchester United from api-football
* fetch-team-stats -> Fetch and save  Manchester United stats for seasons: 2021, 2022, 2023
* fetch-venues -> Fetch and save all venues from England
* fetch-standings -> Fetch and save Manchester United standings for seasons:  2021, 2022, 2023
* fetch-fixtures -> Fetch and save all Manchester United fixtures in the premier league for seasons: 2021, 2022, 2023
* fetch-injuries -> Fetch and save all Manchester United injuries for seasons: 2021, 2022, 2023
* fetch-squad -> Fetch and save Manchester United squad for season 2025/2026

API Endpoints
-
| Method  | Endpoint                                  | Description                                                                        |
| ------- | ----------------------------------------- | ---------------------------------------------------------------------------------- |
| **GET** | `{host}/country`                          | Retrieve all available countries for football data                                 |
| **GET** | `{host}/country/{name}`                   | Retrieve details for a specific country by name                                    |
| **GET** | `{host}/league`                           | Retrieve all football leagues where Manchester United participated atleast one time|
| **GET** | `{host}/team`                             | Retrieve information about the team (Manchester United)                            |
| **GET** | `{host}/teamStats/games/{season}`         | Retrieve information about all premier league games for a given season             |
| **GET** | `{host}/teamStats/goals/{season}`         | Retrieve goal statistics for the team by season                                    |
| **GET** | `{host}/teamStats/streak/{season}`        | Retrieve win/loss/draw streak data by season                                       |
| **GET** | `{host}/teamStats/biggest/{season}`       | Retrieve biggest wins, losses and goals scored by season                           |
| **GET** | `{host}/teamStats/cleansheet/{season}`    | Retrieve clean sheet statistics by season                                          |
| **GET** | `{host}/teamStats/failedtoscore/{season}` | Retrieve data for matches where the team failed to score                           |
| **GET** | `{host}/teamStats/penalty/{season}`       | Retrieve penalty statistics for the team by season                                 |
| **GET** | `{host}/teamStats/cards/{season}`         | Retrieve yellow/red card statistics by season                                      |
| **GET** | `{host}/teamStats/lineup/{season}`        | Retrieve information about lineups and formations for a given season               |
| **GET** | `{host}/venue`                            | Retrieve all available venues in England                                           |
| **GET** | `{host}/venue/{city}`                     | Retrieve venue information by city name                                            |
| **GET** | `{host}/venue/biggest&smallest`           | Retrieve the biggest and smallest venues in England                                |
| **GET** | `{host}/standings/{season}`               | Retrieve league standings for the given season                                     |
| **GET** | `{host}/fixtures/{season}`                | Retrieve all fixtures for the given season                                         |
| **GET** | `{host}/injuries/{season}`                | Retrieve players injury data for the given season                                  |
| **GET** | `{host}/squad`                            | Retrieve the current (2025/2026) Manchester United squad information               |
