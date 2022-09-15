package midjourney

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJob_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		json string
		want *Job
	}{
		{
			name: "full",
			//nolint: lll
			json: `{
		"current_status": "completed",
		"enqueue_time": "2022-09-07 06:58:02.200753",
		"event": {
			"height": 512,
			"textPrompt": [
				"earth, landscape, picturesque, photo, photorealistic"
			],
			"imagePrompts": [],
			"width": 768,
			"batchSize": 1,
			"seedImageURL": null
		},
		"flagged": false,
		"followed_by_user": false,
		"grid_id": null,
		"grid_num": null,
		"guild_id": null,
		"hidden": false,
		"id": "a3052616-372b-42a1-a72b-eb86fa0be633",
		"image_paths": [
			"https://storage.googleapis.com/dream-machines-output/a3052616-372b-42a1-a72b-eb86fa0be633/0_0.png"
		],
		"is_published": true,
		"liked_by_user": false,
		"low_priority": true,
		"metered": false,
		"mod_hidden": false,
		"platform": "discord",
		"platform_channel": "DM",
		"platform_channel_id": "991150132894638170",
		"platform_message_id": "1016966680586506291",
		"platform_thread_id": null,
		"prompt": "earth, landscape, picturesque, photo, photorealistic",
		"ranked_by_user": false,
		"ranking_by_user": null,
		"type": "grid",
		"user_id": "146914681683050496",
		"username": "jimeh",
		"full_command": "earth, landscape, picturesque, photo, photorealistic --testp --ar 16:10  --video",
		"reference_job_id": null,
		"reference_image_num": null
	}`,
			want: &Job{
				CurrentStatus: "completed",
				EnqueueTime: Time{
					time.Date(2022, 9, 7, 6, 58, 2, 200753000, time.UTC),
				},
				Event: &Event{
					Height: 512,
					TextPrompt: []string{
						"earth, landscape, picturesque, photo, photorealistic",
					},
					ImagePrompts: []string{},
					Width:        768,
					BatchSize:    1,
					SeedImageURL: "",
				},
				Flagged:        false,
				FollowedByUser: false,
				GridID:         "",
				GridNum:        "",
				GuildID:        "",
				Hidden:         false,
				ID:             "a3052616-372b-42a1-a72b-eb86fa0be633",
				ImagePaths: []string{
					"https://storage.googleapis.com/dream-machines-output/" +
						"a3052616-372b-42a1-a72b-eb86fa0be633/0_0.png",
				},
				IsPublished:       true,
				LikedByUser:       false,
				LowPriority:       true,
				Metered:           false,
				ModHidden:         false,
				Platform:          "discord",
				PlatformChannel:   "DM",
				PlatformChannelID: "991150132894638170",
				PlatformMessageID: "1016966680586506291",
				PlatformThreadID:  "",
				Prompt: "earth, landscape, picturesque, photo, " +
					"photorealistic",
				RankedByUser:  false,
				RankingByUser: 0,
				Type:          "grid",
				UserID:        "146914681683050496",
				Username:      "jimeh",
				FullCommand: "earth, landscape, picturesque, photo, " +
					"photorealistic --testp --ar 16:10  --video",
				ReferenceJobID:    "",
				ReferenceImageNum: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &Job{}

			err := json.Unmarshal([]byte(tt.json), got)
			require.NoError(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}
