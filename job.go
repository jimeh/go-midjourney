package midjourney

import "fmt"

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

type Event struct {
	Height       int      `json:"height,omitempty"`
	TextPrompt   []string `json:"textPrompt,omitempty"`
	ImagePrompts []string `json:"imagePrompts,omitempty"`
	Width        int      `json:"width,omitempty"`
	BatchSize    int      `json:"batchSize,omitempty"`
	SeedImageURL string   `json:"seedImageURL,omitempty"`
}

type Job struct {
	CurrentStatus     JobStatus `json:"current_status,omitempty"`
	EnqueueTime       Time      `json:"enqueue_time,omitempty"`
	Event             *Event    `json:"event,omitempty"`
	Flagged           bool      `json:"flagged,omitempty"`
	FollowedByUser    bool      `json:"followed_by_user,omitempty"`
	GridID            string    `json:"grid_id,omitempty"`
	GridNum           string    `json:"grid_num,omitempty"`
	GuildID           string    `json:"guild_id,omitempty"`
	Hidden            bool      `json:"hidden,omitempty"`
	ID                string    `json:"id,omitempty"`
	ImagePaths        []string  `json:"image_paths,omitempty"`
	IsPublished       bool      `json:"is_published,omitempty"`
	LikedByUser       bool      `json:"liked_by_user,omitempty"`
	LowPriority       bool      `json:"low_priority,omitempty"`
	Metered           bool      `json:"metered,omitempty"`
	ModHidden         bool      `json:"mod_hidden,omitempty"`
	Platform          string    `json:"platform,omitempty"`
	PlatformChannel   string    `json:"platform_channel,omitempty"`
	PlatformChannelID string    `json:"platform_channel_id,omitempty"`
	PlatformMessageID string    `json:"platform_message_id,omitempty"`
	PlatformThreadID  string    `json:"platform_thread_id,omitempty"`
	Prompt            string    `json:"prompt,omitempty"`
	RankedByUser      bool      `json:"ranked_by_user,omitempty"`
	RankingByUser     int       `json:"ranking_by_user,omitempty"`
	Type              JobType   `json:"type,omitempty"`
	UserID            string    `json:"user_id,omitempty"`
	Username          string    `json:"username,omitempty"`
	FullCommand       string    `json:"full_command,omitempty"`
	ReferenceJobID    string    `json:"reference_job_id,omitempty"`
	ReferenceImageNum string    `json:"reference_image_num,omitempty"`
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
