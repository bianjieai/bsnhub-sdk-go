package oracle

import (
	"errors"
	"strings"

	"github.com/bianjieai/irita-sdk-go/codec"
	"github.com/bianjieai/irita-sdk-go/codec/types"
	sdk "github.com/bianjieai/irita-sdk-go/types"
)

const (
	// ModuleName define module name
	ModuleName = "oracle"

	RUNNING   = "running"
	PAUSED    = "paused"
	COMPLETED = "completed"

	eventTypeSetFeed = "set_feed"

	attributeKeyFeedName  = "feed_name"
	attributeKeyFeedValue = "feed_value"
)

var (
	_ sdk.Msg = MsgCreateFeed{}
	_ sdk.Msg = MsgStartFeed{}
	_ sdk.Msg = MsgPauseFeed{}
	_ sdk.Msg = MsgEditFeed{}

	amino = codec.New()

	ModuleCdc = codec.NewHybridCodec(amino, types.NewInterfaceRegistry())
)

func init() {
	registerCodec(amino)
}

//______________________________________________________________________

func (msg MsgCreateFeed) Route() string {
	return ModuleName
}

// Type implements Msg.
func (msg MsgCreateFeed) Type() string {
	return "create_feed"
}

// ValidateBasic implements Msg.
func (msg MsgCreateFeed) ValidateBasic() error {
	feedName := strings.TrimSpace(msg.FeedName)
	if len(feedName) == 0 {
		return errors.New("feedName missed")
	}

	serviceName := strings.TrimSpace(msg.ServiceName)
	if len(serviceName) == 0 {
		return errors.New("serviceName missed")
	}

	if len(msg.Providers) == 0 {
		return errors.New("providers missed")
	}

	aggregateFunc := strings.TrimSpace(msg.AggregateFunc)
	if len(aggregateFunc) == 0 {
		return errors.New("aggregateFunc missed")
	}

	valueJsonPath := strings.TrimSpace(msg.ValueJsonPath)
	if len(valueJsonPath) == 0 {
		return errors.New("valueJsonPath missed")
	}

	if len(msg.Creator) == 0 {
		return errors.New("creator missed")
	}
	return nil
}

// GetSignBytes implements Msg.
func (msg MsgCreateFeed) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners implements Msg.
func (msg MsgCreateFeed) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

//______________________________________________________________________

func (msg MsgStartFeed) Route() string {
	return ModuleName
}

// Type implements Msg.
func (msg MsgStartFeed) Type() string {
	return "start_feed"
}

// ValidateBasic implements Msg.
func (msg MsgStartFeed) ValidateBasic() error {
	feedName := strings.TrimSpace(msg.FeedName)
	if len(feedName) == 0 {
		return errors.New("feedName missed")
	}
	if len(msg.Creator) == 0 {
		return errors.New("creator missed")
	}
	return nil
}

// GetSignBytes implements Msg.
func (msg MsgStartFeed) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners implements Msg.
func (msg MsgStartFeed) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

//______________________________________________________________________

func (msg MsgPauseFeed) Route() string {
	return ModuleName
}

// Type implements Msg.
func (msg MsgPauseFeed) Type() string {
	return "pause_feed"
}

// ValidateBasic implements Msg.
func (msg MsgPauseFeed) ValidateBasic() error {
	feedName := strings.TrimSpace(msg.FeedName)
	if len(feedName) == 0 {
		return errors.New("feedName missed")
	}
	if len(msg.Creator) == 0 {
		return errors.New("creator missed")
	}
	return nil
}

// GetSignBytes implements Msg.
func (msg MsgPauseFeed) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners implements Msg.
func (msg MsgPauseFeed) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

//______________________________________________________________________

func (msg MsgEditFeed) Route() string {
	return ModuleName
}

// Type implements Msg.
func (msg MsgEditFeed) Type() string {
	return "edit_feed"
}

// ValidateBasic implements Msg.
func (msg MsgEditFeed) ValidateBasic() error {
	feedName := strings.TrimSpace(msg.FeedName)
	if len(feedName) == 0 {
		return errors.New("feedName missed")
	}

	if len(msg.Creator) == 0 {
		return errors.New("creator missed")
	}
	return nil
}

// GetSignBytes implements Msg.
func (msg MsgEditFeed) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners implements Msg.
func (msg MsgEditFeed) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

//______________________________________________________________________

func (this FeedContext) Convert() interface{} {
	return QueryFeedResponse{
		Feed:              this.Feed,
		ServiceName:       this.ServiceName,
		Providers:         this.Providers,
		Input:             this.Input,
		Timeout:           this.Timeout,
		ServiceFeeCap:     this.ServiceFeeCap,
		RepeatedFrequency: this.RepeatedFrequency,
		ResponseThreshold: this.ResponseThreshold,
		State:             this.State,
	}
}

type feedsContext []FeedContext

func (this feedsContext) Convert() interface{} {
	responses := make([]QueryFeedResponse, len(this))
	for i, response := range this {
		responses[i] = response.Convert().(QueryFeedResponse)
	}
	return responses
}

func (this FeedValue) Convert() interface{} {
	return QueryFeedValueResponse{
		Data:      this.Data,
		Timestamp: this.Timestamp,
	}
}

type feedValues []FeedValue

func (this feedValues) Convert() interface{} {
	feedValues := make([]QueryFeedValueResponse, len(this))
	for i, feedValue := range this {
		feedValues[i] = feedValue.Convert().(QueryFeedValueResponse)
	}
	return feedValues
}

//______________________________________________________________________

func registerCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateFeed{}, "irita/modules/MsgCreateFeed", nil)
	cdc.RegisterConcrete(MsgStartFeed{}, "irita/modules/MsgStartFeed", nil)
	cdc.RegisterConcrete(MsgPauseFeed{}, "irita/modules/MsgPauseFeed", nil)
	cdc.RegisterConcrete(MsgEditFeed{}, "irita/modules/MsgEditFeed", nil)

	cdc.RegisterConcrete(Feed{}, "irita/modules/Feed", nil)
	cdc.RegisterConcrete(FeedContext{}, "irita/modules/FeedContext", nil)
}
