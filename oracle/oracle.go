package oracle

import (
	"github.com/bianjieai/irita-sdk-go/codec"
	"github.com/bianjieai/irita-sdk-go/codec/types"
	sdk "github.com/bianjieai/irita-sdk-go/types"
)

type oracleClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) OracleI {
	return oracleClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (a oracleClient) Name() string {
	return ModuleName
}

func (oc oracleClient) RegisterCodec(cdc *codec.Codec) {
	registerCodec(cdc)
}

func (oc oracleClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateFeed{},
		&MsgStartFeed{},
		&MsgPauseFeed{},
		&MsgEditFeed{},
	)
}

func (oc oracleClient) CreateFeed(
	request FeedCreateRequest,
	baseTx sdk.BaseTx,
) (
	result sdk.ResultTx,
	err sdk.Error,
) {
	creator, err := oc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	var providers []sdk.AccAddress
	for _, provider := range request.Providers {
		p, err := sdk.AccAddressFromBech32(provider)
		if err != nil {
			return sdk.ResultTx{}, sdk.Wrapf("%s invalid address", p)
		}
		providers = append(providers, p)
	}

	amt, err := oc.ToMinCoin(request.ServiceFeeCap...)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := MsgCreateFeed{
		FeedName:          request.FeedName,
		LatestHistory:     request.LatestHistory,
		Description:       request.Description,
		Creator:           creator,
		ServiceName:       request.ServiceName,
		Providers:         providers,
		Input:             request.Input,
		Timeout:           request.Timeout,
		ServiceFeeCap:     amt,
		RepeatedFrequency: request.RepeatedFrequency,
		AggregateFunc:     request.AggregateFunc,
		ValueJsonPath:     request.ValueJsonPath,
		ResponseThreshold: request.ResponseThreshold,
	}

	return oc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (oc oracleClient) StartFeed(
	feedName string,
	baseTx sdk.BaseTx,
) (
	result sdk.ResultTx,
	err sdk.Error,
) {
	creator, err := oc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := MsgStartFeed{
		FeedName: feedName,
		Creator:  creator,
	}

	return oc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (oc oracleClient) CreateAndStartFeed(
	request FeedCreateRequest,
	baseTx sdk.BaseTx,
) (
	result sdk.ResultTx,
	err sdk.Error,
) {
	creator, err := oc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	var providers []sdk.AccAddress
	for _, provider := range request.Providers {
		p, err := sdk.AccAddressFromBech32(provider)
		if err != nil {
			return sdk.ResultTx{}, sdk.Wrapf("%s invalid address", p)
		}
		providers = append(providers, p)
	}

	amt, err := oc.ToMinCoin(request.ServiceFeeCap...)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msgCreateFeed := MsgCreateFeed{
		FeedName:          request.FeedName,
		LatestHistory:     request.LatestHistory,
		Description:       request.Description,
		Creator:           creator,
		ServiceName:       request.ServiceName,
		Providers:         providers,
		Input:             request.Input,
		Timeout:           request.Timeout,
		ServiceFeeCap:     amt,
		RepeatedFrequency: request.RepeatedFrequency,
		AggregateFunc:     request.AggregateFunc,
		ValueJsonPath:     request.ValueJsonPath,
		ResponseThreshold: request.ResponseThreshold,
	}

	msgStartFeed := MsgStartFeed{
		FeedName: request.FeedName,
		Creator:  creator,
	}

	return oc.BuildAndSend([]sdk.Msg{msgCreateFeed, msgStartFeed}, baseTx)
}

func (oc oracleClient) PauseFeed(
	feedName string,
	baseTx sdk.BaseTx,
) (
	result sdk.ResultTx,
	err sdk.Error,
) {
	creator, err := oc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := MsgPauseFeed{
		FeedName: feedName,
		Creator:  creator,
	}

	return oc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (oc oracleClient) EditFeed(
	request FeedEditRequest,
	baseTx sdk.BaseTx,
) (
	result sdk.ResultTx,
	err sdk.Error,
) {
	creator, err := oc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	var providers []sdk.AccAddress
	for _, provider := range request.Providers {
		p, err := sdk.AccAddressFromBech32(provider)
		if err != nil {
			return sdk.ResultTx{}, sdk.Wrapf("%s invalid address", p)
		}
		providers = append(providers, p)
	}

	amt, err := oc.ToMinCoin(request.ServiceFeeCap...)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := MsgEditFeed{
		FeedName:          request.FeedName,
		LatestHistory:     request.LatestHistory,
		Description:       request.Description,
		Creator:           creator,
		Providers:         providers,
		Timeout:           request.Timeout,
		ServiceFeeCap:     amt,
		RepeatedFrequency: request.RepeatedFrequency,
		ResponseThreshold: request.ResponseThreshold,
	}

	return oc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (oc oracleClient) SubscribeFeedValue(
	feedName string,
	callback func(value FeedValue),
) (
	subscription sdk.Subscription,
	err sdk.Error,
) {
	feed, err := oc.QueryFeed(feedName)
	if err != nil {
		return subscription, err
	}

	isInValidState := func(state string) bool {
		return state == COMPLETED || state == PAUSED || state == ""
	}

	if isInValidState(feed.State.String()) {
		return subscription, sdk.Wrapf("feed:%s state is invalid:%s", feedName, feed.State)
	}

	handleResult := func(value string, subTx sdk.Subscription) {
		oc.Logger().Debug().
			Str(attributeKeyFeedValue, value).
			Msg("received feed value")

		var fv FeedValue
		if err := oc.UnmarshalJSON([]byte(value), &fv); err == nil {
			callback(fv)
			if f, err := oc.QueryFeed(feedName); err != nil || isInValidState(f.State.String()) {
				_ = oc.Unsubscribe(subTx)
			}
		}
	}

	txBuilder := sdk.NewEventQueryBuilder().AddCondition(
		sdk.NewCond(
			eventTypeSetFeed,
			attributeKeyFeedName,
		).EQ(
			sdk.EventValue(feedName),
		),
	).AddCondition(
		sdk.NewCond(
			sdk.EventTypeMessage,
			sdk.AttributeKeyAction,
		).EQ(
			"respond_service",
		),
	)

	subscription, err = oc.SubscribeTx(
		txBuilder,
		func(tx sdk.EventDataTx) {
			result, err := tx.Result.Events.GetValue(eventTypeSetFeed, attributeKeyFeedValue)
			if err != nil {
				println(err.Error())
			}
			handleResult(result, subscription)
		},
	)

	return subscription, err
}

func (oc oracleClient) QueryFeed(feedName string) (QueryFeedResponse, sdk.Error) {
	param := struct {
		FeedName string
	}{
		FeedName: feedName,
	}

	var feedCtx FeedContext
	if err := oc.QueryWithResponse("custom/oracle/feed", param, &feedCtx); err != nil {
		return QueryFeedResponse{}, sdk.Wrap(err)
	}

	return feedCtx.Convert().(QueryFeedResponse), nil
}

func (oc oracleClient) QueryFeeds(state string) ([]QueryFeedResponse, sdk.Error) {
	param := struct {
		State string
	}{
		State: state,
	}

	var feedsCtx feedsContext
	if err := oc.QueryWithResponse("custom/oracle/feeds", param, &feedsCtx); err != nil {
		return nil, sdk.Wrap(err)
	}

	return feedsCtx.Convert().([]QueryFeedResponse), nil
}

func (oc oracleClient) QueryFeedValue(feedName string) ([]QueryFeedValueResponse, sdk.Error) {
	param := struct {
		FeedName string
	}{
		FeedName: feedName,
	}

	var fvs feedValues
	if err := oc.QueryWithResponse("custom/oracle/feedValue", param, &fvs); err != nil {
		return nil, sdk.Wrap(err)
	}

	return fvs.Convert().([]QueryFeedValueResponse), nil
}
