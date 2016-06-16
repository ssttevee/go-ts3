// contains commands with optional or variable numbers of arguments
package ts3command

import (
	"github.com/ssttevee/go-ts3/enum"
)

type ServerList struct {
	UniqueId bool // whether or not to return the unique id
	Short bool // Only return server id, port, and status (Overrides UniqueId)
	All bool // List all servers in the database, including those of which that are not associated which the current machine
	OnlyOffline bool // Only include servers whose status is offline
}

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

type InstanceEdit map[string]string

func (this InstanceEdit) Command() string {
	ret := "instanceedit"

	for property, value := range this {
		ret += " " + property + "=" + value
	}

	return ret
}

type Use struct {
	ServerId string // Select a server by it's id
	Port     string // Select a server by it's port, a random server will be selected if the port is used by more than one server
	Virtual  bool   // Start the server in virtual mode, which allows changes but do not allow clients to connect
}

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

type ServerCreate struct {
	Name string                  // Name of the new server
	Properties map[string]string // Properties for the new server
}

func (this ServerCreate) Command() string {
	ret := "servercreate virtualserver_name=" + this.Name

	if this.Properties != nil && len(this.Properties) > 0 {
		for property, value := range this.Properties {
			ret += " " + property + "=" + value
		}
	}

	return ret
}

type ServerEdit map[string]string

func (this ServerEdit) Command() string {
	ret := "serveredit"

	if len(this) > 0 {
		for property, value := range this {
			ret += " " + property + "=" + value
		}
	}

	return ret
}

type ServerGroupAdd struct {
	Name string                              // Name of the server group
	Type ts3enum.PermissionGroupDatabaseType // Type of
}

func (this ServerGroupAdd) Command() string {
	ret := "servergroupadd name=" + this.Name

	if this.Type != "" && this.Type != ts3enum.PermissionGroupDatabaseTypeRegular {
		ret += " type=" + string(this.Type)
	}

	return ret
}

type ServerGroupDelete struct {
	Id string  // The id of the server group to delete
	Force bool // Force the server group to be delete even while there are clients within
}

func (this ServerGroupDelete) Command() string {
	ret := "servergroupdel sgid=" + this.Id

	if this.Force {
		ret += " force=1"
	}

	return ret
}

type ServerNotifyRegister struct {
	Event   ts3enum.NotifyEvent // Event to register for
	Channel string              // Channel to listen to if the event is a channel
}

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

type ClientUpdate map[string]string

func (this ClientUpdate) Command() string {
	ret := "clientupdate"

	for property, value := range this {
		ret += " " + property + "=" + value
	}

	return ret
}
