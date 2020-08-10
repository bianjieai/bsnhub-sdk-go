package oracle

import (
	time "time"

	service "github.com/bianjieai/irita-sdk-go/modules/service"
	sdk "github.com/bianjieai/irita-sdk-go/types"
)

type OracleI interface {
	sdk.Module
	OracleTx
	OracleQuery
}

type OracleTx interface {
	CreateFeed(request FeedCreateRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	StartFeed(feedName string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	CreateAndStartFeed(request FeedCreateRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	PauseFeed(feedName string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	EditFeed(request FeedEditRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	SubscribeFeedValue(feedName string, callback func(value FeedValue)) (sdk.Subscription, sdk.Error)
}

type OracleQuery interface {
	QueryFeed(feedName string) (QueryFeedResponse, sdk.Error)
	QueryFeeds(state string) ([]QueryFeedResponse, sdk.Error)
	QueryFeedValue(feedName string) ([]QueryFeedValueResponse, sdk.Error)
}

// FeedCreateRequest - struct for create a feed
type FeedCreateRequest struct {
	FeedName          string       `json:"feed_name"`
	LatestHistory     uint64       `json:"latest_history"`
	Description       string       `json:"description"`
	ServiceName       string       `json:"service_name"`
	Providers         []string     `json:"providers"`
	Input             string       `json:"input"`
	Timeout           int64        `json:"timeout"`
	ServiceFeeCap     sdk.DecCoins `json:"service_fee_cap"`
	RepeatedFrequency uint64       `json:"repeated_frequency"`
	AggregateFunc     string       `json:"aggregate_func"`
	ValueJsonPath     string       `json:"value_json_path"`
	ResponseThreshold uint32       `json:"response_threshold"`
}

// FeedEditRequest - struct for edit a existed feed
type FeedEditRequest struct {
	FeedName          string       `json:"feed_name"`
	Description       string       `json:"description"`
	LatestHistory     uint64       `json:"latest_history"`
	Providers         []string     `json:"providers"`
	Timeout           int64        `json:"timeout"`
	ServiceFeeCap     sdk.DecCoins `json:"service_fee_cap"`
	RepeatedFrequency uint64       `json:"repeated_frequency"`
	ResponseThreshold uint32       `json:"response_threshold"`
}

type QueryFeedResponse struct {
	Feed              *Feed                       `json:"feed"`
	ServiceName       string                      `json:"service_name"`
	Providers         []sdk.AccAddress            `json:"providers"`
	Input             string                      `json:"input"`
	Timeout           int64                       `json:"timeout"`
	ServiceFeeCap     sdk.Coins                   `json:"service_fee_cap"`
	RepeatedFrequency uint64                      `json:"repeated_frequency"`
	ResponseThreshold uint32                      `json:"response_threshold"`
	State             service.RequestContextState `json:"state"`
}

type QueryFeedValueResponse struct {
	Data      string    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}
