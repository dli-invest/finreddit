package types

// TODO DELETE BEFORE RELEASE

// type conf struct {
//     Hits int64 `yaml:"hits"`
//     Time int64 `yaml:"time"`
// }


// type Config struct {
//     Server struct {
//         // Host is the local machine IP Address to bind the HTTP Server to
//         Host string `yaml:"host"`

//         // Port is the local machine TCP Port to bind the HTTP Server to
//         Port    string `yaml:"port"`
//         Timeout struct {
//             // Server is the general server timeout to use
//             // for graceful shutdowns
//             Server time.Duration `yaml:"server"`

//             // Write is the amount of time to wait until an HTTP server
//             // write opperation is cancelled
//             Write time.Duration `yaml:"write"`

//             // Read is the amount of time to wait until an HTTP server
//             // read operation is cancelled
//             Read time.Duration `yaml:"read"`

//             // Read is the amount of time to wait
//             // until an IDLE HTTP session is closed
//             Idle time.Duration `yaml:"idle"`
//         } `yaml:"timeout"`
//     } `yaml:"server"`
// }


// search settings for subreddits
type SRConfig struct {
    Name string `yaml:"name,omitempty" url:"name,omitempty"`
    Time string `yaml:"t,omitempty" url:"t,omitempty"`
    Limit int `yaml:"limit,omitempty" url:"t,omitempty"`
    After   string `yaml:"after,omitempty" url:"after,omitempty"`
	Before  string `yaml:"before,omitempty" url:"before,omitempty"`
	Count   int    `yaml:"count,omitempty" url:"count,omitempty"`
    Show    string `yaml:"show,omitempty" url:"show,omitempty"`
    Article string `yaml:"article,omitempty" url:"article,omitempty"`
    MinScore int `yaml:"minScore,omitempty"`
    MinComments int `yaml:"minComments,omitempty"`
}

type SearchConfig struct {
	Data struct {
        SubReddits []SRConfig `yaml:"subreddits" json:"subreddits"`
	} `yaml:"data" json:"data"`
}

type DiscordEmbed struct {
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

type DiscordWebhook struct {
	Content string `json:"content" xml:"content" form:"content" query:"content"`
	Embeds []DiscordEmbed `json:"embeds" xml:"embeds" form:"embeds" query:"embeds"`
}
