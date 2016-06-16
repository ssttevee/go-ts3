// contains commands which have a constant number of arguments
package ts3command

import (
	"fmt"
	"github.com/ssttevee/go-ts3/enum"
	"github.com/ssttevee/go-ts3/util"
)

func Login(username, password string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("login %s %s", username, password))
}

func ServerIdGetByPort(port string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("serveridgetbyport virtualserver_port=%s", port))
}

func ServerDelete(serverId string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("serverdelete sid=%d", serverId))
}

func ServerStart(serverId string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("serverstart sid=%d", serverId))
}

func ServerStop(serverId string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("serverstop sid=%d", serverId))
}

func SendTextMessage(targetMode ts3enum.TextMessageTargetMode, target, message string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("sendtextmessage targetmode=%s target=%s msg=%s", targetMode, target, ts3util.Escape(message)))
}

func LogAdd(level ts3enum.LogLevel, message string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("logadd loglevel=%s logmsg=%s", level, ts3util.Escape(message)))
}

func GlobalMessage(message string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("gm msg=%s", ts3util.Escape(message)))
}
