package foo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	app "github.com/ignite/cli/ignite/pkg/cosmosanalysis/app/testdata/modules/registration_not_in_app_go"
)

type Foo struct {
	FooKeeper foo.keeper
}

func (f Foo) RegisterAPIRoutes()         {}
func (f Foo) RegisterTxService()         {}
func (f Foo) RegisterTendermintService() {}
func (f Foo) Name() string               { return app.BaseApp.Name() }
func (f Foo) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (f Foo) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}
