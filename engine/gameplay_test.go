package engine

import "testing"

import (
  "gopoker/model"
  "gopoker/model/game"
)

func TestGameplay_rotateGame(t *testing.T) {
  mix := model.NewMix(game.Eight, 8)

  ctx := &Context{
    Mix: mix,
    Table: model.NewTable(9),
    Stake: model.NewStake(10.0),
  }

  g := NewGameplay(ctx)
  if g.Game != mix.Games[0] {
    t.Fatalf("first game should be %s; got %s", mix.Games[0], g.Game)
  }
  for i := 0; i <= 8; i++ {
    g.rotateGame()
  }
  if g.Game != mix.Games[1] {
    t.Fatalf("rotated game should be %s; got %s", mix.Games[1], g.Game)
  }
}

func TestGameplay_dealing(t *testing.T) {
  game := model.NewGame(game.Texas, game.NoLimit, 9)

  table := model.NewTable(9)

  table.AddPlayer(model.Player("A"), 1, 1000.0)
  table.AddPlayer(model.Player("B"), 2, 1000.0)
  table.AddPlayer(model.Player("C"), 3, 1000.0)
  table.AddPlayer(model.Player("D"), 4, 1000.0)

  t.Logf("%s", table)

  ctx := &Context{
    Game: game,
    Table: table,
    Stake: model.NewStake(10.0),
  }

  g := NewGameplay(ctx)

  g.startDeal()
  g.dealHole(2)

}

func TestGameplay_antes(t *testing.T) {
  game := model.NewGame(game.Texas, game.NoLimit, 9)

  table := model.NewTable(9)

  table.AddPlayer(model.Player("A"), 1, 1000.0)
  table.AddPlayer(model.Player("B"), 2, 1000.0)
  table.AddPlayer(model.Player("C"), 3, 1000.0)
  table.AddPlayer(model.Player("D"), 4, 1000.0)

  t.Logf("%s", table)

  ctx := &Context{
    Game: game,
    Table: table,
    Stake: model.NewStake(10.0),
  }

  g := NewGameplay(ctx)

  g.startDeal()
  g.postAntes()

}

func TestGameplay_blinds(t *testing.T) {
  game := model.NewGame(game.Texas, game.NoLimit, 9)

  table := model.NewTable(9)

  table.SetButton(2)

  table.AddPlayer(model.Player("A"), 1, 1000.0)
  table.AddPlayer(model.Player("B"), 2, 1000.0)
  table.AddPlayer(model.Player("C"), 3, 1000.0)
  table.AddPlayer(model.Player("D"), 4, 1000.0)

  t.Logf("%s", table)

  ctx := &Context{
    Game: game,
    Table: table,
    Stake: model.NewStake(10.0),
  }

  g := NewGameplay(ctx)

  g.startDeal()
  g.postBlinds()
  //t.FailNow()

}
