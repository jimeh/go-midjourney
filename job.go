package midjourney

import (
	"fmt"
	"regexp"
)

type JobType string

const (
	JobTypeNull    JobType = "null"
	JobTypeGrid    JobType = "grid"
	JobTypeUpscale JobType = "upscale"
)

type JobStatus string

const (
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
)

type Job struct {
	JobType           string           `json:"_job_type,omitempty"`
	Service           string           `json:"_service,omitempty"`
	ParsedParams      *ParsedJobParams `json:"_parsed_params,omitempty"`
	CurrentStatus     JobStatus        `json:"current_status,omitempty"`
	EnqueueTime       Time             `json:"enqueue_time,omitempty"`
	Event             *Event           `json:"event,omitempty"`
	Flagged           bool             `json:"flagged,omitempty"`
	FollowedByUser    bool             `json:"followed_by_user,omitempty"`
	GridID            string           `json:"grid_id,omitempty"`
	GridNum           string           `json:"grid_num,omitempty"`
	GuildID           string           `json:"guild_id,omitempty"`
	Hidden            bool             `json:"hidden,omitempty"`
	ID                string           `json:"id,omitempty"`
	ImagePaths        []string         `json:"image_paths,omitempty"`
	IsPublished       bool             `json:"is_published,omitempty"`
	LikedByUser       bool             `json:"liked_by_user,omitempty"`
	LowPriority       bool             `json:"low_priority,omitempty"`
	Metered           bool             `json:"metered,omitempty"`
	ModHidden         bool             `json:"mod_hidden,omitempty"`
	Platform          string           `json:"platform,omitempty"`
	PlatformChannel   string           `json:"platform_channel,omitempty"`
	PlatformChannelID string           `json:"platform_channel_id,omitempty"`
	PlatformMessageID string           `json:"platform_message_id,omitempty"`
	PlatformThreadID  string           `json:"platform_thread_id,omitempty"`
	Prompt            string           `json:"prompt,omitempty"`
	RankedByUser      bool             `json:"ranked_by_user,omitempty"`
	RankingByUser     int              `json:"ranking_by_user,omitempty"`
	Type              JobType          `json:"type,omitempty"`
	UserID            string           `json:"user_id,omitempty"`
	Username          string           `json:"username,omitempty"`
	FullCommand       string           `json:"full_command,omitempty"`
	ReferenceJobID    string           `json:"reference_job_id,omitempty"`
	ReferenceImageNum string           `json:"reference_image_num,omitempty"`
}

func (j *Job) DiscordURL() string {
	if j.Platform != "discord" || j.GuildID == "" ||
		j.PlatformChannelID == "" || j.PlatformMessageID == "" {
		return ""
	}

	return fmt.Sprintf("https://discord.com/channels/%s/%s/%s",
		j.GuildID,
		j.PlatformChannelID,
		j.PlatformMessageID,
	)
}

func (j *Job) MainImageURL() string {
	return fmt.Sprintf("https://mj-gallery.com/%s/grid_0.png", j.ID)
}

var imageFilenameRegexp = regexp.MustCompile(`[^a-zA-Z0-9\._]+`)

func (j *Job) ImageFilename() string {
	s := fmt.Sprintf("%s_%s", j.Username, j.Prompt)
	s = imageFilenameRegexp.ReplaceAllString(s, "_")
	if len(s) > 63 {
		s = s[:63]
	}

	return fmt.Sprintf("%s_%s.png", s, j.ID)
}

func (j *Job) VideoURL() string {
	if j.Type != JobTypeGrid {
		return ""
	}

	return fmt.Sprintf("https://i.mj.run/%s/video.mp4", j.ID)
}

type Event struct {
	Height       int      `json:"height,omitempty"`
	TextPrompt   []string `json:"textPrompt,omitempty"`
	ImagePrompts []string `json:"imagePrompts,omitempty"`
	Width        int      `json:"width,omitempty"`
	BatchSize    int      `json:"batchSize,omitempty"`
	SeedImageURL string   `json:"seedImageURL,omitempty"`
}

type ParsedJobParams struct {
	Anime    bool             `json:"anime,omitempty"`
	Aspect   string           `json:"aspect,omitempty"`
	Creative bool             `json:"creative,omitempty"`
	Fast     bool             `json:"fast,omitempty"`
	HD       bool             `json:"hd,omitempty"`
	No       []string         `json:"no,omitempty"`
	Style    string           `json:"style,omitempty"`
	Stylize  int              `json:"stylize,omitempty"`
	Test     bool             `json:"test,omitempty"`
	Testp    bool             `json:"testp,omitempty"`
	Tile     bool             `json:"tile,omitempty"`
	Upanime  bool             `json:"upanime,omitempty"`
	Upbeta   bool             `json:"upbeta,omitempty"`
	Uplight  bool             `json:"uplight,omitempty"`
	Version  AlgorithmVersion `json:"version,omitempty"`
	Vibe     bool             `json:"vibe,omitempty"`
	Video    bool             `json:"video,omitempty"`
}
