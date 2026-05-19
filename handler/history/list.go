package history

import (
	"context"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func ListCurrentUserSearchHistoryL20(ctx context.Context, req *sparkruntime.ListCurrentUserSearchHistoryL20Request) (*sparkruntime.ListCurrentUserSearchHistoryL20Response, error) {
	userId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, err
	}

	historyList, err := db.ListUserSearchHistoryByUserId(ctx, db.DB, req.GetType(), userId, 20)
	if err != nil {
		return nil, err
	}

	historyInfoList := utils.MapStructList(historyList, func(history *db.UserSearchHistory) *sparkruntime.UserHistoryInfo {
		return &sparkruntime.UserHistoryInfo{
			Id:        history.Id,
			Content:   history.SearchContent,
			CreatedAt: history.CreatedAt.Unix(),
		}
	})

	return &sparkruntime.ListCurrentUserSearchHistoryL20Response{
		HistoryList: historyInfoList,
		BaseResp:    handler.ConstructSuccessResp(),
	}, nil
}
