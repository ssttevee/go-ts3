#GO TS3
TeamSpeak 3 Server Query interface for Go

##Quick Start
###Connecting to a server
It's as simple as
```
client, err := ts3.Dial("ts.example.com:10011")
defer client.Close()
```

###Sending a command - part 1
```
client.Command("clientlist")
```

###Sending a command - part 2
```
client.Commandf("login %s %s", "serveradmin", "12345")
```

###Sending a command - part 3
```
client.DoCommand(ts3command.GlobalMessage("TS3 is da bom"))
```

##Reading Responses
For example, if I wanted to log all the names of connected clients
```
r, err := client.Command("clientlist")
arr := r.Array()

for _, c := range arr {
    log.Println(c["client_nickname"])
}
```

Or if I wanted to get my current client id
```
r, err := client.Command("whoami")
data := r.Map()

clid := data["client_id"]
```