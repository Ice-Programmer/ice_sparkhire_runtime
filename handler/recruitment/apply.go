package recruitment

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"ice_sparkhire_runtime/handler"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func ApplyRecruitment(ctx context.Context, req *sparkruntime.ApplyRecruitmentRequest) (*sparkruntime.ApplyRecruitmentResponse, error) {
	if req.GetRecruitmentId() <= 0 {
		return nil, fmt.Errorf("invalid recruitment id")
	}

	// fetch current user
	currentUserId, err := utils.GetCurrentUserId(ctx)
	if err != nil {
		return nil, fmt.Errorf("get current user id %w", err)
	}

	// find recruitment
	recruitment, err := db.FindRecruitmentById(ctx, db.DB, req.RecruitmentId)
	if err != nil {
		return nil, fmt.Errorf("find recruitment %w", err)
	}

	// exist recruitment application
	exist, err := db.ExistRecruitmentCandidateApplication(ctx, db.DB, currentUserId, req.RecruitmentId)
	if err != nil {
		return nil, fmt.Errorf("find recruitment error: %w", err)
	}
	if exist {
		return nil, fmt.Errorf("recruitment application already exist")
	}

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		// create apply
		application := &db.RecruitmentApplication{
			Id:            utils.GetId(),
			UserId:        currentUserId,
			RecruitmentId: req.RecruitmentId,
			CompanyId:     recruitment.CompanyId,
			Status:        int8(sparkruntime.RecruitmentApplyStatus_APPLIED),
		}
		if err := db.CreateRecruitmentApplication(ctx, tx, application); err != nil {
			return fmt.Errorf("create recruitment application %w", err)
		}

		// create apply log
		if err := db.CreateRecruitmentApplicationLog(ctx, tx, &db.RecruitmentApplicationLog{
			Id:            utils.GetId(),
			ApplicationId: application.Id,
			ToStatus:      int8(sparkruntime.RecruitmentApplyStatus_APPLIED),
			OperatorId:    currentUserId,
		}); err != nil {
			return fmt.Errorf("create recruitment application log %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("create recruitment application error: %w", err)
	}

	return &sparkruntime.ApplyRecruitmentResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
