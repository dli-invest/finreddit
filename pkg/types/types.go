package types

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

type DiscordPayload struct {
	Content string `json:"content,omitempty" xml:"content,omitempty" form:"content,omitempty" query:"content,omitempty"`
	Embeds []DiscordEmbed `json:"embeds,omitempty" xml:"embeds,omitempty" form:"embeds,omitempty" query:"embeds,omitempty"`
}
