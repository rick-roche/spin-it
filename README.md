# Spin It

![Spin-It CI](https://github.com/rick-roche/spin-it/actions/workflows/spin-it.yml/badge.svg)

Spin-It is a project to connect [Discogs](https://www.discogs.com/ "Music Database and Marketplace") and [Last.fm](https://www.last.fm/ "Play music, find songs, and discover artists") to [Spotify](https://www.spotify.com/ "Music for everyone"), allowing you to create playlists from your listening history or real-world collections.

> :warning: Note: this project is in very early stages of development

## Getting started

### Prerequisites
The easiest way to run locally is to use Docker Compose
- [Docker](https://docs.docker.com/get-docker/ "Docker")

Otherwise, you will need
- [Golang](https://golang.org/doc/install "The Go Programming Language")
- [Air](https://github.com/cosmtrek/air "Live reload for Go apps")
- [go-swagger](https://github.com/go-swagger/go-swagger)

### Developing locally

Both options use [Air](https://github.com/cosmtrek/air "Live reload for Go apps") to allow for a live-reload development environment.

#### Docker Compose
You will need to create a `.env` file in the root of the repo and add the following

```bash
# Spotify App Details
SPOTIFY_ID=your-spotify-app-id
SPOTIFY_SECRET=your-spotify-app-secret

# Last.FM App Details
LASTFM_API_KEY=your-lastfm-app-key
LASTFM_API_SECRET=your-lastfm-app-secret
```

then run

```bash
docker compose -f build/docker-compose.yml --env-file ../.env up --build
```

#### Local

Ensure that you export these variables into your bash context, the simplest way is to have a file called `env_vars.sh` for example

```bash
#!/bin/bash

# Spotify App Details
export SPOTIFY_ID=replace-with-your-spotify-app-id
export SPOTIFY_SECRET=replace-with-your-spotify-app-secret

# Last.FM App Details
export LASTFM_API_KEY=replace-with-your-lastfm-app-key
export LASTFM_API_SECRET=replace-with-your-lastfm-app-secret
```

then run

```bash
source env_vars.sh
air -c .air.toml
```

## Project info

### Project structure
This repo follows the guidelines from https://github.com/golang-standards/project-layout

### TODO list
- [x] Add [Swagger](https://github.com/go-swagger/go-swagger)
- [x] Add [Docker compose](https://docs.docker.com/compose/ "Compose is a tool for defining and running multi-container Docker applications.")
- [x] Add Github workflow to publish images
- [ ] Add [Postman](https://www.postman.com/ "The Collaboration Platform for API Development") collection
- [x] [Last.fm](https://www.last.fm/ "Play music, find songs, and discover artists") integration
- [x] [Spotify](https://www.spotify.com/ "Music for everyone") integration
- [ ] [Discogs](https://www.discogs.com/ "Music Database and Marketplace") integration
