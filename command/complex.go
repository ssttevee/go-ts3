package ts3command

import (
	"github.com/ssttevee/go-ts3/enum"
)

// Changes the server instance configuration using given properties.
type InstanceEdit map[string]string

// Command returns a string to send
func (this InstanceEdit) Command() string {
	ret := "instanceedit"

	for property, value := range this {
		ret += " " + property + "=" + value
	}

	return ret
}

// Selects the virtual server specified with sid or port to allow further interaction.
//
// The ServerQuery client will appear on the virtual server and acts like a real
// TeamSpeak 3 Client, except it's unable to send or receive voice data.
//
// If your database contains multiple virtual servers using the same UDP port,
// use will select a random virtual server using the specified port.
type Use struct {
	// Select a server by it's id
	ServerId string
	// Select a server by it's port, a random server will
	// be selected if the port is used by more than one server
	Port     string
	// Start the server in virtual mode, which
	// allows changes but do not allow clients to connect
	Virtual  bool
}

// Command returns a string to send
func (this Use) Command() string {
	ret := "use "

	if this.ServerId != "" {
		ret += this.ServerId
	} else if this.Port != "" {
		ret += "port=" + this.Port
	}

	if this.Virtual {
		ret += " -virtual"
	}

	return ret
}

// Displays a list of virtual servers including
// their ID, status, number of clients online, etc.
//
// If you're using the -all option, the server will list all virtual servers stored in the database.
// This can be useful when multiple server instances with different machine IDs are using the same database.
// The machine ID is used to identify the server instance a virtual server is associated with.
//
// The status of a virtual server can be either online, offline,
// deploy running, booting up, shutting down and virtual online.
// While most of them are self-explanatory, virtual online is a bit more complicated.
//
// Please note that whenever you select a virtual server which is currently stopped,
// it will be started in virtual mode which means you are able to change its configuration,
// create channels or change permissions, but no regular TeamSpeak 3 Client can connect.
// As soon as the last ServerQuery client deselects the virtual server, its status will be changed back to offline.
type ServerList struct {
	// whether or not to return the unique id
	UniqueId    bool
	// Only return server id, port, and status (Overrides UniqueId)
	Short       bool
	// List all servers in the database, including those
	// that are not associated which the current machine
	All         bool
	// Only include servers whose status is offline
	OnlyOffline bool
}

// Command returns a string to send
func (this ServerList) Command() string {
	ret := "serverlist"

	if this.Short {
		ret += " -short"
	} else if this.UniqueId {
		ret += " -uid"
	}

	if this.All {
		ret += " -all"
	}

	if this.OnlyOffline {
		ret += " -onlyoffline"
	}

	return ret
}

// Creates a new virtual server using the given properties and displays
// its ID, port and initial administrator privilege key.
//
// If virtualserver_port is not specified, the server will test for the first unused UDP port.
//
// The first virtual server will be running on UDP port 9987 by default.
// Subsequently started virtual servers will be running on increasing UDP port numbers.
type ServerCreate struct {
	Name       string            // Name of the new server
	Properties map[string]string // Properties for the new server
}

// Command returns a string to send
func (this ServerCreate) Command() string {
	ret := "servercreate virtualserver_name=" + this.Name

	if this.Properties != nil && len(this.Properties) > 0 {
		for property, value := range this.Properties {
			ret += " " + property + "=" + value
		}
	}

	return ret
}

// Changes the selected virtual servers configuration using given properties.
//
// Note that this command accepts multiple properties which means that you're
// able to change all settings of the selected virtual server at once.
type ServerEdit map[string]string

// Command returns a string to send
func (this ServerEdit) Command() string {
	ret := "serveredit"

	if len(this) > 0 {
		for property, value := range this {
			ret += " " + property + "=" + value
		}
	}

	return ret
}

// Creates a new server group using the name specified with name and displays its ID.
//
// The optional type parameter can be used to create ServerQuery groups and template groups.
type ServerGroupAdd struct {
	// Name of the server group
	Name string
	// Either a template group, regular group, or query group
	// Default is regular group
	Type ts3enum.PermissionGroupDatabaseType
}

// Command returns a string to send
func (this ServerGroupAdd) Command() string {
	ret := "servergroupadd name=" + this.Name

	if this.Type != "" && this.Type != ts3enum.PermissionGroupDatabaseTypeRegular {
		ret += " type=" + string(this.Type)
	}

	return ret
}

// Deletes the server group specified with sgid.
//
// If force is enabled, the server group will be deleted even if there are clients within.
type ServerGroupDelete struct {
	Id    string // The id of the server group to delete
	Force bool   // Force the server group to be delete even while there are clients within
}

// Command returns a string to send
func (this ServerGroupDelete) Command() string {
	ret := "servergroupdel sgid=" + this.Id

	if this.Force {
		ret += " force=1"
	}

	return ret
}

// Registers for a specified category of events on a virtual server to receive notification messages.
//
// Depending on the notifications you've registered for, the server will
// send you a message on every event in the view of your ServerQuery client
// (e.g. clients joining your channel, incoming text messages, server configuration changes, etc).
//
// The event source is declared by the event parameter while
// id can be used to limit the notifications to a specific channel.
type ServerNotifyRegister struct {
	// Events to register for
	Event   ts3enum.NotifyEvent
	// Channel to listen to if the event is a channel.
	// Channel 0 means all channels
	Channel string
}

// Command returns a string to send
func (this ServerNotifyRegister) Command() string {
	ret := "servernotifyregister"

	if this.Event != "" {
		ret += " event=" + string(this.Event)
	}

	if this.Event == ts3enum.NotifyEventChannel || this.Event == ts3enum.NotifyEventTextChannel {
		if this.Channel != "" {
			ret += " id=" + this.Channel
		}
	}

	return ret
}

// Change your ServerQuery clients settings using given properties.
type ClientUpdate map[string]string

// Command returns a string to send
func (this ClientUpdate) Command() string {
	ret := "clientupdate"

	for property, value := range this {
		ret += " " + property + "=" + value
	}

	return ret
}
