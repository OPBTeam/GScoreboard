# GScoreboard

Simple scoreboard for [dragonfly](https://github.com/df-mc/dragonfly) server

## Get the package
```
go get github.com/opbteam/gscoreboard
```
# Import the package
```golang
import "github.com/opbteam/gscoreboard/gscoreboard"
```
# Create scoreboard
```go
score := gscoreboard.Scoreboard{
  Title: "Minecraft server",
  Line: []string{"Hello world", "Have fun:)"},
  Delay: 5,
}
score.Send(*player.Player)
```

`Delay`: scoreboard will be reloaded after 5 seconds
