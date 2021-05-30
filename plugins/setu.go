package plugins

import (
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/cqCode"
	"github.com/3343780376/leafBot/message"
)

func UseSetuHandle() {
	leafBot.OnCommand("/tu").
		AddAllies("来点色图").
		SetWeight(10).
		SetBlock(false).
		SetPluginName("色图").
		AddHandle(
			func(event leafBot.Event, bot *leafBot.Bot, args []string) {
				if len(args) < 1 {
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.ParseMessageFromString(cqCode.Image("https://laosepi.org/pic.php", map[string]interface{}{"cache": 0, "c": 3})))
				} else if args[0] == "mc" {
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.ParseMessageFromString(cqCode.Image("https://laosepi.org/mcpic.php", map[string]interface{}{"cache": 0, "c": 3})))
				} else if args[0] == "sj" {
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.ParseMessageFromString(cqCode.Image("https://laosepi.org/sjpic.php", map[string]interface{}{"cache": 0, "c": 3})))
				} else if args[0] == "gq" {
					bot.SendMsg(event.MessageType, event.UserId, event.GroupId, message.ParseMessageFromString(cqCode.Image("https://laosepi.org/gqpic.php", map[string]interface{}{"cache": 0, "c": 3})))
				}
			})

}
