// Package ts3command binds to various TeamSpeak 3 Server Query commands
package ts3command

// A SimpleCommand holds a command which do not require any arguments or are straight forward
type SimpleCommand string

// Command returns a string to send
func (cmd SimpleCommand) Command() string {
	return string(cmd)
}

var (
	// Closes the ServerQuery connection to the TeamSpeak 3 Server instance.
	Quit                        SimpleCommand = "quit"
	// Deselects the active virtual server and logs out from the server instance.
	Logout                      SimpleCommand = "logout"
	// Displays the servers version information including platform and build number.
	Version                     SimpleCommand = "version"
	// Displays detailed connection information about the server instance
	// including uptime, number of virtual servers online, traffic information, etc.
	HostInfo                    SimpleCommand = "hostinfo"
	// Displays the server instance configuration including database
	// revision number, the file transfer port, default group IDs, etc.
	InstanceInfo                SimpleCommand = "instanceinfo"
	// Displays a list of IP addresses used by the server instance on multi-homed machines.
	BindingList                 SimpleCommand = "bindinglist"
	// Stops the entire TeamSpeak 3 Server instance by shutting down the process.
	ServerProcessStop           SimpleCommand = "serverprocessstop"
	// Displays detailed configuration information about the selected virtual
	// server including unique ID, number of clients online, configuration, etc.
	ServerInfo                  SimpleCommand = "serverinfo"
	// Displays detailed connection information about the selected
	// virtual server including uptime, traffic information, etc.
	ServerRequestConnectionInfo SimpleCommand = "serverrequestconnectioninfo"
	// Displays a list of server groups available. Depending on your permissions,
	// the output may also contain global ServerQuery groups and template groups.
	ServerGroupList             SimpleCommand = "servergrouplist"
	// Unregisters all events previously registered with
	// servernotifyregister so you will no longer receive notification messages.
	ServerNotifyUnregister      SimpleCommand = "servernotifyunregister"
	// Displays information about your current ServerQuery connection including your loginname, etc.
	WhoAmI                      SimpleCommand = "whoami"
)
