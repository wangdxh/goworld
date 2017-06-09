package game

import (
	"flag"

	"github.com/xiaonanln/goworld/common"
)

var (
	gameid      int
	gameService *GameService
)

func init() {
	parseArgs()
}

func parseArgs() {
	flag.IntVar(&gameid, "gid", 0, "set gameid")
	flag.Parse()
}

func Run(delegate IGameDelegate) {
	gameService = newGameService(gameid, delegate)
	gameService.run()
}

func GetServiceProviders(serviceName string) []common.EntityID {
	return gameService.registeredServices[serviceName].ToList()
}
