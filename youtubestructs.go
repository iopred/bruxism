package main

import (
	"errors"
	"fmt"
	"html"
)

type Error struct {
	Errors []struct {
		Domain  string `json:"domain"`
		Reason  string `json:"reason"`
		Message string `json:"message"`
	}
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) NewError(message string) error {
	return errors.New(fmt.Sprintf("Error %v: %v (%v)", message, e.Message, e.Code))
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type LiveBroadcastSnippetThumbnails struct {
	Url    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type LiveBroadcastSnippet struct {
	PublishedAt        string                                     `json:"publishedAt,omitempty"`
	ChannelId          string                                     `json:"channelId,omitempty"`
	Title              string                                     `json:"title,omitempty"`
	Description        string                                     `json:"description,omitempty"`
	Thumbnails         map[string]*LiveBroadcastSnippetThumbnails `json:"thumbnails,omitempty"`
	ScheduledStartTime string                                     `json:"scheduledStartTime,omitempty"`
	ScheduledEndTime   string                                     `json:"scheduledEndTime,omitempty"`
	ActualStartTime    string                                     `json:"actualStartTime,omitempty"`
	ActualEndTime      string                                     `json:"actualEndTime,omitempty"`
	IsDefaultBroadcast bool                                       `json:"isDefaultBroadcast,omitempty"`
	LiveChatId         string                                     `json:"liveChatId,omitempty"`
}

type LiveBroadcastStatus struct {
	LifeCycleStatus string `json:"lifeCycleStatus,omitempty"`
	PrivacyStatus   string `json:"privacyStatus,omitempty"`
	RecordingStatus string `json:"recordingStatus,omitempty"`
}

type LiveBroadcastContentDetailsMonitorStream struct {
	EnableMonitorStream    bool   `json:"enableMonitorStream,omitempty"`
	BroadcastStreamDelayMs uint   `json:"broadcastStreamDelayMs,omitempty"`
	EmbedHtml              string `json:"embedHtml,omitempty"`
}

type LiveBroadcastContentDetails struct {
	BoundStreamId           string                                    `json:"boundStreamId,omitempty"`
	MonitorStream           *LiveBroadcastContentDetailsMonitorStream `json:"monitorStream,omitempty"`
	EnableEmbed             bool                                      `json:"enableEmbed,omitempty"`
	EnableDvr               bool                                      `json:"enableDvr,omitempty"`
	EnableContentEncryption bool                                      `json:"enableContentEncryption,omitempty"`
	StartWithSlate          bool                                      `json:"startWithSlate,omitempty"`
	RecordFromStart         bool                                      `json:"recordFromStart,omitempty"`
	EnableClosedCaptions    bool                                      `json:"enableClosedCaptions,omitempty"`
}

type LiveBroadcast struct {
	Error          *Error                       `json:"error,omitempty"`
	Kind           string                       `json:"kind,omitempty"`
	Etag           string                       `json:"etag,omitempty"`
	Id             string                       `json:"id,omitempty"`
	Snippet        *LiveBroadcastSnippet        `json:"snippet,omitempty"`
	Status         *LiveBroadcastStatus         `json:"status,omitempty"`
	ContentDetails *LiveBroadcastContentDetails `json:"contentDetails,omitempty"`
}

type LiveBroadcasts struct {
	Error         *Error           `json:"error"`
	Kind          string           `json:"kind"`
	Etag          string           `json:"etag"`
	NextPageToken string           `json:"nextPageToken"`
	PageInfo      *PageInfo        `json:"pageInfo"`
	Items         []*LiveBroadcast `json:"items"`
}

type LiveChatMessageSnippetType string

const (
	LiveChatMessageSnippetTypeText       LiveChatMessageSnippetType = "textMessageEvent"
	LiveChatMessageSnippetTypeFanFunding LiveChatMessageSnippetType = "fanFundingEvent"
)

type LiveChatMessageSnippetFanFundingEventDetails struct {
	AmountMicros        int    `json:"amountMicros,string,omitempty"`
	Currency            string `json:"currency,omitempty"`
	AmountDisplayString string `json:"amountDisplayString,omitempty"`
	UserComment         string `json:"userComment,omitempty"`
}

type LiveChatMessageSnippetTextMessageDetails struct {
	MessageText string `json:"messageText,omitempty"`
}

type LiveChatMessageSnippet struct {
	Type                   LiveChatMessageSnippetType                    `json:"type,omitempty"`
	LiveChatId             string                                        `json:"liveChatId,omitempty"`
	AuthorChannelId        string                                        `json:"authorChannelId,omitempty"`
	PublishedAt            string                                        `json:"publishedAt,omitempty"`
	HasDisplayContent      bool                                          `json:"hasDisplayContent,omitempty"`
	DisplayMessage         string                                        `json:"displayMessage,omitempty"`
	FanFundingEventDetails *LiveChatMessageSnippetFanFundingEventDetails `json:"fanFundingEventDetails,omitempty"`
	TextMessageDetails     *LiveChatMessageSnippetTextMessageDetails     `json:"textMessageDetails,omitempty"`
}

type LiveChatMessageAuthorDetails struct {
	ChannelId       string `json:"channelId,omitempty"`
	ChannelUrl      string `json:"channelUrl,omitempty"`
	DisplayName     string `json:"displayName,omitempty"`
	ProfileImageUrl string `json:"profileImageUrl,omitempty"`
	IsVerified      bool   `json:"isVerified,omitempty"`
	IsChatOwner     bool   `json:"isChatOwner,omitempty"`
	IsChatSponsor   bool   `json:"isChatSponsor,omitempty"`
	IsChatModerator bool   `json:"isChatModerator,omitempty"`
}

const LiveChatMessageKind string = "youtube#liveChatMessage"

type LiveChatMessage struct {
	Error         *Error                        `json:"error,omitempty"`
	Kind          string                        `json:"kind,omitempty"`
	Etag          string                        `json:"etag,omitempty"`
	Id            string                        `json:"id,omitempty"`
	Snippet       *LiveChatMessageSnippet       `json:"snippet,omitempty"`
	AuthorDetails *LiveChatMessageAuthorDetails `json:"authorDetails,omitempty"`
}

func (m *LiveChatMessage) Channel() string {
	return m.Snippet.LiveChatId
}

func (m *LiveChatMessage) User() string {
	return m.AuthorDetails.DisplayName
}

func (m *LiveChatMessage) Message() string {
	switch m.Snippet.Type {
	case LiveChatMessageSnippetTypeText:
		return html.UnescapeString(m.Snippet.TextMessageDetails.MessageText)
	}
	return html.UnescapeString(m.Snippet.DisplayMessage)
}

type LiveChatMessageListResponse struct {
	Error                 *Error `json:"error"`
	Kind                  string `json:"kind"`
	Etag                  string `json:"etag"`
	NextPageToken         string `json:"nextPageToken"`
	PollingIntervalMillis int    `json:"pollingIntervalMillis"`
	PageInfo              *PageInfo
	Items                 []*LiveChatMessage `json:"items"`
}
