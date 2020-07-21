package jpushclient

// Notice https://docs.jiguang.cn//jpush/server/push/rest_api_v3_push/#notification
type Notice struct {
	Alert    string          `json:"alert,omitempty"`
	Android  *AndroidNotice  `json:"android,omitempty"`
	IOS      *IOSNotice      `json:"ios,omitempty"`
	WINPhone *WinPhoneNotice `json:"winphone,omitempty"`
}

type AndroidNotice struct {
	Alert             string                 `json:"alert"`
	Title             string                 `json:"title,omitempty"`
	BuilderID         int                    `json:"builder_id,omitempty"`
	ChannelID         string                 `json:"channel_id,omitempty"`
	Priority          int                    `json:"priority,omitempty"`
	Category          string                 `json:"category,omitempty"`
	Style             int                    `json:"style,omitempty"`
	AlertType         int                    `json:"alert_type,omitempty"`
	BigText           string                 `json:"big_text,omitempty"`
	Inbox             interface{}            `json:"inbox,omitempty"`
	BigPicPath        string                 `json:"big_pic_path,omitempty"`
	Extras            map[string]interface{} `json:"extras,omitempty"`
	LargeIcon         string                 `json:"large_icon,omitempty"`
	Intent            interface{}            `json:"intent,omitempty"`
	URIActivity       string                 `json:"uri_activity,omitempty"`
	URIAction         string                 `json:"uri_action,omitempty"`
	BadgeAddNum       int                    `json:"badge_add_num,omitempty"`
	BadgeClass        string                 `json:"badge_class,omitempty"`
	Sound             string                 `json:"sound,omitempty"`
	ShowBeginTime     string                 `json:"show_begin_time,omitempty"`
	ShowEndTime       string                 `json:"show_end_time,omitempty"`
	DisplayForeground string                 `json:"display_foreground,omitempty"`
}

type IOSNotice struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            interface{}            `json:"badge,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
	ThreadID         string                 `json:"thread-id,omitempty"`
}

type WinPhoneNotice struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

func (this *Notice) SetAlert(alert string) {
	this.Alert = alert
}

func (this *Notice) SetAndroidNotice(n *AndroidNotice) {
	this.Android = n
}

func (this *Notice) SetIOSNotice(n *IOSNotice) {
	if n != nil && n.Badge == nil {
		n.Badge = "+1"
	}
	this.IOS = n
}

func (this *Notice) SetWinPhoneNotice(n *WinPhoneNotice) {
	this.WINPhone = n
}
