// contains commands which do not require any arguments
package ts3command

type SimpleCommand string

func (cmd SimpleCommand) Command() string {
	return string(cmd)
}

var (
	Quit                        SimpleCommand = "quit"
	Logout                      SimpleCommand = "logout"
	Version                     SimpleCommand = "version"
	HostInfo                    SimpleCommand = "hostinfo"
	InstanceInfo                SimpleCommand = "instanceinfo"
	BindingList                 SimpleCommand = "bindinglist"
	ServerProcessStop           SimpleCommand = "serverprocessstop"
	ServerInfo                  SimpleCommand = "serverinfo"
	ServerRequestConnectionInfo SimpleCommand = "serverrequestconnectioninfo"
	ServerGroupList             SimpleCommand = "servergrouplist"
	ServerNotifyUnregister      SimpleCommand = "servernotifyunregister"
	WhoAmI                      SimpleCommand = "whoami"
)
