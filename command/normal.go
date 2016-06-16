package ts3command

import (
	"fmt"
	"github.com/ssttevee/go-ts3/enum"
	"github.com/ssttevee/go-ts3/util"
)

// Authenticates with the TeamSpeak 3 Server
// instance using given ServerQuery login credentials.
func Login(username, password string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("login %s %s", username, password))
}

// Displays the database ID of the virtual server
// running on the UDP port specified by virtualserver_port.
func ServerIdGetByPort(port string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("serveridgetbyport virtualserver_port=%s", port))
}

// Deletes the virtual server specified with sid.
//
// Please note that only virtual servers in stopped state can be deleted.
func ServerDelete(serverId string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("serverdelete sid=%s", serverId))
}

// Starts the virtual server specified with sid.
//
// Depending on your permissions, you're able to start either your
// own virtual server only or all virtual servers in the server instance
func ServerStart(serverId string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("serverstart sid=%s", serverId))
}

// Stops the virtual server specified with sid.
//
// Depending on your permissions, you're able to stop either your
// own virtual server only or all virtual servers in the server instance.
func ServerStop(serverId string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("serverstop sid=%s", serverId))
}

// Sends a text message a specified target.
//
// The type of the target is determined by targetmode while target specifies
// the ID of the recipient, whether it be a virtual server, a channel or a client.
func SendTextMessage(targetMode ts3enum.TextMessageTargetMode, target, message string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("sendtextmessage targetmode=%s target=%s msg=%s", targetMode, target, ts3util.Escape(message)))
}

// Writes a custom entry into the servers log.
//
// Depending on your permissions, you'll be able to add entries
// into the server instance log and/or your virtual servers log.
//
// The loglevel parameter specifies the type of the entry
func LogAdd(level ts3enum.LogLevel, message string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("logadd loglevel=%s logmsg=%s", level, ts3util.Escape(message)))
}

// Sends a text message to all clients on all virtual servers in the TeamSpeak 3 Server instance.
func GlobalMessage(message string) SimpleCommand {
	return SimpleCommand(fmt.Sprintf("gm msg=%s", ts3util.Escape(message)))
}
