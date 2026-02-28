package wish_career

//func CreateWishCareer(ctx context.Context, req *sparkruntime.CreateWishCareerRequest) (*sparkruntime.CreateWishCareerResponse, error) {
//	if err := validateCreateWishCareer(ctx, req); err != nil {
//		return nil, err
//	}
//
//	currentUser, err := user.GetCurrentUser(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	if sparkruntime.UserRole(currentUser.UserRole) != sparkruntime.UserRole_Candidate {
//		return nil, fmt.Errorf("user is not candidate")
//	}
//
//	if err := wish_career.EnsureWishCareerNotExists(ctx, currentUser.Id, req.CareerId); err != nil {
//		return nil, err
//	}
//
//	if err := db.CreateWishCareer(ctx, db.DB, buildWishCareer(req, currentUser.Id)); err != nil {
//		return nil, err
//	}
//
//	return &sparkruntime.CreateWishCareerResponse{
//		BaseResp: handler.ConstructSuccessResp(),
//	}, nil
//}
//
//func validateCreateWishCareer(ctx context.Context, req *sparkruntime.CreateWishCareerRequest) error {
//	if req.IsSetSalaryLower() && req.IsSetSalaryUpper() && req.GetSalaryLower() > req.GetSalaryUpper() {
//		return fmt.Errorf("lowest wish salary can not greater than highest salary")
//	}
//
//	if req.GetCareerId() <= 0 {
//		return fmt.Errorf("career id is invalid")
//	}
//
//	if _, err := db.FindCareerById(ctx, db.DB, req.GetCareerId()); err != nil {
//		return fmt.Errorf("find career error: %v", err)
//	}
//
//	if utils.NotContains(sparkruntime.SalaryCurrencyTypeList, req.GetCurrencyType()) {
//		return fmt.Errorf("currency type is invalid")
//	}
//
//	if utils.NotContains(sparkruntime.SalaryFrequencyTypeList, req.GetFrequencyType()) {
//		return fmt.Errorf("frequency type is invalid")
//	}
//
//	return nil
//}
