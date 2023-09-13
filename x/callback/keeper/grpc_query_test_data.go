package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/evmos/evmos/v9/x/callback/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TestDataAll(c context.Context, req *types.QueryAllTestDataRequest) (*types.QueryAllTestDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var testDatas []types.TestData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	testDataStore := prefix.NewStore(store, types.KeyPrefix(types.TestDataKey))

	pageRes, err := query.Paginate(testDataStore, req.Pagination, func(key []byte, value []byte) error {
		var testData types.TestData
		if err := k.cdc.Unmarshal(value, &testData); err != nil {
			return err
		}

		testDatas = append(testDatas, testData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTestDataResponse{TestData: testDatas, Pagination: pageRes}, nil
}

func (k Keeper) TestData(c context.Context, req *types.QueryGetTestDataRequest) (*types.QueryGetTestDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	testData, found := k.GetTestData(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTestDataResponse{TestData: testData}, nil
}
