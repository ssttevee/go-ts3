package ts3enum

type HostMessageMode string

const (
	HostMessageModeLog       HostMessageMode = "1" // 1: display message in chatlog
	HostMessageModeModal     HostMessageMode = "2" // 2: display message in modal dialog
	HostMessageModeModalQuit HostMessageMode = "3" // 3: display message in modal dialog and close connection
)

type HostBannerMode string

const (
	HostBannerModeNoAdjust     HostBannerMode = "0" // 0: do not adjust
	HostBannerModeIgnoreAspect HostBannerMode = "1" // 1: adjust but ignore aspect ratio (like TeamSpeak 2)
	HostBannerModeKeepAspect   HostBannerMode = "2" // 2: adjust and keep aspect ratio
)

type Codec string

const (
	CodecSpeexNarrowband    Codec = "0" // 0: speex narrowband (mono, 16bit, 8kHz)
	CodecSpeexWideband      Codec = "1" // 1: speex wideband (mono, 16bit, 16kHz)
	CodecSpeexUltraWideband Codec = "2" // 2: speex ultra-wideband (mono, 16bit, 32kHz)
	CodecCeltMono           Codec = "3" // 3: celt mono (mono, 16bit, 48kHz)
)

type CodecEncryptionMode string

const (
	CodecEncryptionModeIndividual CodecEncryptionMode = "0" // 0: configure per channel
	CodecEncryptionModeDisabled   CodecEncryptionMode = "1" // 1: globally disabled
	CodecEncryptionModeEnabled    CodecEncryptionMode = "2" // 2: globally enabled
)

type TextMessageTargetMode string

const (
	TextMessageTargetModeClient  TextMessageTargetMode = "1" // 1: target is a client
	TextMessageTargetModeChannel TextMessageTargetMode = "2" // 2: target is a channel
	TextMessageTargetModeServer  TextMessageTargetMode = "3" // 3: target is a virtual server
)

type LogLevel string

const (
	LogLevelError   LogLevel = "1" // 1: everything that is really bad
	LogLevelWarning LogLevel = "2" // 2: everything that might be bad
	LogLevelDebug   LogLevel = "3" // 3: output that might help find a problem
	LogLevelInfo    LogLevel = "4" // 4: informational output
)

type ReasonIdentifier string

const (
	ReasonKickChannel ReasonIdentifier = "4" // 4: kick client from channel
	ReasonKickServer  ReasonIdentifier = "5" // 5: kick client from server
)

type PermissionGroupDatabaseType string

const (
	PermissionGroupDatabaseTypeTemplate PermissionGroupDatabaseType = "0" // 0: template group (used for new virtual servers)
	PermissionGroupDatabaseTypeRegular  PermissionGroupDatabaseType = "1" // 1: regular group (used for regular clients)
	PermissionGroupDatabaseTypeQuery    PermissionGroupDatabaseType = "2" // 2: global query group (used for ServerQuery clients)
)

type PermissionGroupType string

const (
	PermissionGroupTypesServerGroup   PermissionGroupType = "0" // 0: server group permission
	PermissionGroupTypesGlobalClient  PermissionGroupType = "1" // 1: client specific permission
	PermissionGroupTypesChannel       PermissionGroupType = "2" // 2: channel specific permission
	PermissionGroupTypesChannelGroup  PermissionGroupType = "3" // 3: channel group permission
	PermissionGroupTypesChannelClient PermissionGroupType = "4" // 4: channel-client specific permission
)

type TokenType string

const (
	TokenTypeServerGroup  TokenType = "0" // 0: server group token (id1={groupID} id2=0)
	TokenTypeChannelGroup TokenType = "1" // 1: channel group token (id1={groupID} id2={channelID})
)

type NotifyEvent string

const (
	NotifyEventServer      NotifyEvent = "server"
	NotifyEventChannel     NotifyEvent = "channel"
	NotifyEventTextServer  NotifyEvent = "textserver"
	NotifyEventTextChannel NotifyEvent = "textchannel"
	NotifyEventTextPrivate NotifyEvent = "textprivate"
)

type NoticeType string

const (
	// server events
	NoticeTypeServerEdit               NoticeType = "notifyserveredited"              // On Server Edited
	NoticeTypeClientConnect            NoticeType = "notifycliententerview"           // On Client Join Server
	NoticeTypeClientDisconnect         NoticeType = "notifyclientleftview"            // On Client Leave Server

	// channel events
	NoticeTypeClientMove               NoticeType = "notifyclientmoved"               // On Client Change Channel
	NoticeTypeChannelEdit              NoticeType = "notifychanneledited"             // On Channel Edited
	NoticeTypeChannelDescriptionChange NoticeType = "notifychanneldescriptionchanged" // On Channel Description Changed
	NoticeTypeChannelPasswordChange    NoticeType = "notifychannelpasswordchanged"    // On Channel Password Changed
	NoticeTypeChannelMove              NoticeType = "notifychannelmoved"              // On Channel Position Changed
	NoticeTypeChannelCreate            NoticeType = "notifychannelcreated"            // On Channel Creation
	NoticeTypeChannelDelete            NoticeType = "notifychanneldeleted"            // On Channel Deletion

	// text event
	NoticeTypeTextMessage              NoticeType = "notifytextmessage"               // On Receive Text Message
	NoticeTypePrivateTextMessage       NoticeType = "notifyprivatetextmessage"        // On Receive Text Message
	NoticeTypeChannelTextMessage       NoticeType = "notifychanneltextmessage"        // On Receive Text Message
	NoticeTypeServerTextMessage        NoticeType = "notifyservertextmessage"         // On Receive Text Message
)

func (t NoticeType) NotifyEvent() NotifyEvent {
	switch t {
	case NoticeTypeServerEdit, NoticeTypeClientConnect, NoticeTypeClientDisconnect:
		return NotifyEventServer
	case NoticeTypeClientMove, NoticeTypeChannelEdit,
		NoticeTypeChannelDescriptionChange, NoticeTypeChannelPasswordChange,
		NoticeTypeChannelMove, NoticeTypeChannelCreate, NoticeTypeChannelDelete:
		return NotifyEventChannel
	case NoticeTypePrivateTextMessage:
		return NotifyEventTextPrivate
	case NoticeTypeChannelTextMessage:
		return NotifyEventTextChannel
	case NoticeTypeServerTextMessage:
		return NotifyEventTextServer
	default:
		panic(string(t) + " cannot be registered")
	}
}