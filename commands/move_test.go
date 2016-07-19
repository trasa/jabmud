package commands

import (
	"github.com/stretchr/testify/suite"
	"github.com/trasa/jabmud/world"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MoveFullyConnectedSuite struct {
	suite.Suite
	center, north, south, west, east, up, down *world.Room
	player                                     *world.Player
}

func TestMoveFullyConnectedSuite(t *testing.T) {
	suite.Run(t, new(MoveFullyConnectedSuite))
}

func (suite *MoveFullyConnectedSuite) SetupTest() {
	zone := &world.Zone{}
	suite.center = world.NewRoom(zone, "center", "", "")
	suite.north = world.NewRoom(zone, "north", "", "")
	suite.south = world.NewRoom(zone, "south", "", "")
	suite.west = world.NewRoom(zone, "west", "", "")
	suite.east = world.NewRoom(zone, "east", "", "")
	suite.up = world.NewRoom(zone, "up", "", "")
	suite.down = world.NewRoom(zone, "down", "", "")

	suite.center.North = suite.north
	suite.center.South = suite.south
	suite.center.East = suite.east
	suite.center.West = suite.west
	suite.center.Up = suite.up
	suite.center.Down = suite.down

	suite.player = &world.Player{
		Id: "foo",
	}
	suite.center.AddPlayer(suite.player)
}


func (suite *MoveFullyConnectedSuite) TestMoveNorthSuccess() {
	result := MoveDirection(suite.player, []string{"n"}).(MoveResult)
	assert.True(suite.T(), result.Success)
	assert.False(suite.T(), inRoom(suite.center, suite.player.Id))
	assert.True(suite.T(), inRoom(suite.north, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.south, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.east, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.west, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.up, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.down, suite.player.Id))
}


func (suite *MoveFullyConnectedSuite) TestMoveSouthSuccess() {
	result := MoveDirection(suite.player, []string{"s"}).(MoveResult)
	assert.True(suite.T(), result.Success)
	assert.False(suite.T(), inRoom(suite.center, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.north, suite.player.Id))
	assert.True(suite.T(), inRoom(suite.south, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.east, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.west, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.up, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.down, suite.player.Id))
}


func (suite *MoveFullyConnectedSuite) TestMoveEastSuccess() {
	result := MoveDirection(suite.player, []string{"e"}).(MoveResult)
	assert.True(suite.T(), result.Success)
	assert.False(suite.T(), inRoom(suite.center, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.north, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.south, suite.player.Id))
	assert.True(suite.T(), inRoom(suite.east, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.west, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.up, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.down, suite.player.Id))
}

func (suite *MoveFullyConnectedSuite) TestMoveWestSuccess() {
	result := MoveDirection(suite.player, []string{"w"}).(MoveResult)
	assert.True(suite.T(), result.Success)
	assert.False(suite.T(), inRoom(suite.center, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.north, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.south, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.east, suite.player.Id))
	assert.True(suite.T(), inRoom(suite.west, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.up, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.down, suite.player.Id))
}


func (suite *MoveFullyConnectedSuite) TestMoveUpSuccess() {
	result := MoveDirection(suite.player, []string{"u"}).(MoveResult)
	assert.True(suite.T(), result.Success)
	assert.False(suite.T(), inRoom(suite.center, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.north, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.south, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.east, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.west, suite.player.Id))
	assert.True(suite.T(), inRoom(suite.up, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.down, suite.player.Id))
}


func (suite *MoveFullyConnectedSuite) TestMoveDownSuccess() {
	result := MoveDirection(suite.player, []string{"d"}).(MoveResult)
	assert.True(suite.T(), result.Success)
	assert.False(suite.T(), inRoom(suite.center, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.north, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.south, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.east, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.west, suite.player.Id))
	assert.False(suite.T(), inRoom(suite.up, suite.player.Id))
	assert.True(suite.T(), inRoom(suite.down, suite.player.Id))
}

type MoveFailSuite struct {
	suite.Suite
	center *world.Room
	player *world.Player
}

func TestMoveFailSuite(t *testing.T) {
	suite.Run(t, new(MoveFailSuite))
}

func (suite *MoveFailSuite) SetupTest() {
	zone := &world.Zone{}
	suite.center = world.NewRoom(zone, "center", "", "")

	suite.player = &world.Player{
		Id: "foo",
	}
	suite.center.AddPlayer(suite.player)
}

func (suite *MoveFailSuite) TestBadDirection() {
	result := MoveDirection(suite.player, []string{"x"}).(MoveResult)
	assert.False(suite.T(), result.Success)
}

func inRoom(room *world.Room, id string) (exists bool) {
	_, exists = room.Players[id]
	return
}