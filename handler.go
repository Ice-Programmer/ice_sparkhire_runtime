package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"ice_sparkhire_runtime/handler"
	"ice_sparkhire_runtime/handler/biz"
	"ice_sparkhire_runtime/handler/candidate"
	"ice_sparkhire_runtime/handler/career_exp"
	"ice_sparkhire_runtime/handler/company"
	"ice_sparkhire_runtime/handler/education_exp"
	"ice_sparkhire_runtime/handler/geo"
	"ice_sparkhire_runtime/handler/industry"
	"ice_sparkhire_runtime/handler/major"
	"ice_sparkhire_runtime/handler/ping"
	"ice_sparkhire_runtime/handler/school"
	"ice_sparkhire_runtime/handler/tag"
	"ice_sparkhire_runtime/handler/user"
	"ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
)

// SparkhireRuntimeServiceImpl implements the last service interface defined in the IDL.
type SparkhireRuntimeServiceImpl struct{}

// Ping implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) Ping(ctx context.Context, req *sparkhire_runtime.PingRequest) (resp *sparkhire_runtime.PingResponse, err error) {
	resp, err = ping.Ping(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "ping failed: %v", err)
		resp = &sparkhire_runtime.PingResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// =============================================== user ===============================================

// RegisterUser implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) RegisterUser(ctx context.Context, req *sparkhire_runtime.RegisterUserRequest) (resp *sparkhire_runtime.RegisterUserResponse, err error) {
	resp, err = user.RegisterUser(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "register user failed: %v", err)
		resp = &sparkhire_runtime.RegisterUserResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// UserMailLogin implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) UserMailLogin(ctx context.Context, req *sparkhire_runtime.UserMailLoginRequest) (resp *sparkhire_runtime.UserMailLoginResponse, err error) {
	resp, err = user.UserMailLogin(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "user mail login failed: %v", err)
		resp = &sparkhire_runtime.UserMailLoginResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// FetchCurrentUser implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) FetchCurrentUser(ctx context.Context, req *sparkhire_runtime.FetchCurrentUserRequest) (resp *sparkhire_runtime.FetchCurrentUserResponse, err error) {
	resp, err = user.FetchCurrentUser(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "fetch current user failed: %v", err)
		resp = &sparkhire_runtime.FetchCurrentUserResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// SwitchUserRole implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) SwitchUserRole(ctx context.Context, req *sparkhire_runtime.SwitchUserRoleRequest) (resp *sparkhire_runtime.SwitchUserRoleResponse, err error) {
	resp, err = user.SwitchUserRole(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "switch user role failed: %v", err)
		resp = &sparkhire_runtime.SwitchUserRoleResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// =============================================== candidate ===============================================

// AddCandidate implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) AddCandidate(ctx context.Context, req *sparkhire_runtime.AddCandidateRequest) (resp *sparkhire_runtime.AddCandidateResponse, err error) {
	resp, err = candidate.AddCandidate(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "add candidate failed: %v", err)
		resp = &sparkhire_runtime.AddCandidateResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// EditCandidate implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) EditCandidate(ctx context.Context, req *sparkhire_runtime.EditCandidateRequest) (resp *sparkhire_runtime.EditCandidateResponse, err error) {
	resp, err = candidate.EditCandidate(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "edit candidate failed: %v", err)
		resp = &sparkhire_runtime.EditCandidateResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// GetCurrentCandidate implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) GetCurrentCandidate(ctx context.Context, req *sparkhire_runtime.GetCurrentCandidateRequest) (resp *sparkhire_runtime.GetCurrentCandidateResponse, err error) {
	resp, err = candidate.GetCurrentCandidate(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "get current candidate failed: %v", err)
		resp = &sparkhire_runtime.GetCurrentCandidateResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// =============================================== tag ===============================================

// QueryTag implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) QueryTag(ctx context.Context, req *sparkhire_runtime.QueryTagRequest) (resp *sparkhire_runtime.QueryTagResponse, err error) {
	resp, err = tag.QueryTag(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "query tag failed: %v", err)
		resp = &sparkhire_runtime.QueryTagResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// AddTag implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) AddTag(ctx context.Context, req *sparkhire_runtime.AddTagRequest) (resp *sparkhire_runtime.AddTagResponse, err error) {
	resp, err = tag.AddTag(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "add tag failed: %v", err)
		resp = &sparkhire_runtime.AddTagResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// BindTags implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) BindTags(ctx context.Context, req *sparkhire_runtime.BindTagsRequest) (resp *sparkhire_runtime.BindTagsResponse, err error) {
	resp, err = tag.BindTags(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "bind tags failed: %v", err)
		resp = &sparkhire_runtime.BindTagsResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// UnbindTags implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) UnbindTags(ctx context.Context, req *sparkhire_runtime.UnbindTagsRequest) (resp *sparkhire_runtime.UnbindTagsResponse, err error) {
	resp, err = tag.UnbindTags(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "unbind tags failed: %v", err)
		resp = &sparkhire_runtime.UnbindTagsResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// =============================================== education experience ===============================================

// CreateEducationExp implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) CreateEducationExp(ctx context.Context, req *sparkhire_runtime.CreateEducationExpRequest) (resp *sparkhire_runtime.CreateEducationExpResponse, err error) {
	resp, err = education_exp.CreateEducationExp(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "create education exp failed: %v", err)
		resp = &sparkhire_runtime.CreateEducationExpResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// EditEducationExp implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) EditEducationExp(ctx context.Context, req *sparkhire_runtime.EditEducationExpRequest) (resp *sparkhire_runtime.EditEducationExpResponse, err error) {
	resp, err = education_exp.EditEducationExp(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "education exp failed: %v", err)
		resp = &sparkhire_runtime.EditEducationExpResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// DeleteEducationExp implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) DeleteEducationExp(ctx context.Context, req *sparkhire_runtime.DeleteEducationExpRequest) (resp *sparkhire_runtime.DeleteEducationExpResponse, err error) {
	resp, err = education_exp.DeleteEducationExp(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "education exp failed: %v", err)
		resp = &sparkhire_runtime.DeleteEducationExpResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// GetCurrentEducationExp implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) GetCurrentEducationExp(ctx context.Context, req *sparkhire_runtime.GetCurrentUserEducationExpRequest) (resp *sparkhire_runtime.GetCurrentUserEducationExpResponse, err error) {
	resp, err = education_exp.GetCurrentEducationExp(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "get current education exp failed: %v", err)
		resp = &sparkhire_runtime.GetCurrentUserEducationExpResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// =============================================== career experience ===============================================

// CreateCareerExperience implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) CreateCareerExperience(ctx context.Context, req *sparkhire_runtime.CreateCareerExperienceRequest) (resp *sparkhire_runtime.CreateCareerExperienceResponse, err error) {
	resp, err = career_exp.CreateCareerExperience(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "create career experience failed: %v", err)
		resp = &sparkhire_runtime.CreateCareerExperienceResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// EditCareerExperience implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) EditCareerExperience(ctx context.Context, req *sparkhire_runtime.EditCareerExperienceRequest) (resp *sparkhire_runtime.EditCareerExperienceResponse, err error) {
	resp, err = career_exp.EditCareerExperience(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "edit career experience failed: %v", err)
		resp = &sparkhire_runtime.EditCareerExperienceResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// GetCurrentUserCareerExperience implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) GetCurrentUserCareerExperience(ctx context.Context, req *sparkhire_runtime.GetCurrentUserCareerExperienceRequest) (resp *sparkhire_runtime.GetCurrentUserCareerExperienceResponse, err error) {
	resp, err = career_exp.GetCurrentUserCareerExperience(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "get current user career experience failed: %v", err)
		resp = &sparkhire_runtime.GetCurrentUserCareerExperienceResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// DeleteCareerExperience implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) DeleteCareerExperience(ctx context.Context, req *sparkhire_runtime.DeleteCareerExperienceRequest) (resp *sparkhire_runtime.DeleteCareerExperienceResponse, err error) {
	resp, err = career_exp.DeleteCareerExperience(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "delete career experience failed: %v", err)
		resp = &sparkhire_runtime.DeleteCareerExperienceResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// =============================================== information ===============================================

// ListMajor implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) ListMajor(ctx context.Context, req *sparkhire_runtime.ListMajorRequest) (resp *sparkhire_runtime.ListMajorResponse, err error) {
	resp, err = major.ListMajor(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "list major failed: %v", err)
		resp = &sparkhire_runtime.ListMajorResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// ListIndustry implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) ListIndustry(ctx context.Context, req *sparkhire_runtime.ListIndustryRequest) (resp *sparkhire_runtime.ListIndustryResponse, err error) {
	resp, err = industry.ListIndustry(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "list industry failed: %v", err)
		resp = &sparkhire_runtime.ListIndustryResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// ListSchool implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) ListSchool(ctx context.Context, req *sparkhire_runtime.ListSchoolRequest) (resp *sparkhire_runtime.ListSchoolResponse, err error) {
	resp, err = school.ListSchool(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "list school failed: %v", err)
		resp = &sparkhire_runtime.ListSchoolResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// ListGeo implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) ListGeo(ctx context.Context, req *sparkhire_runtime.ListGeoRequest) (resp *sparkhire_runtime.ListGeoResponse, err error) {
	resp, err = geo.ListGeo(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "list geo failed: %v", err)
		resp = &sparkhire_runtime.ListGeoResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// =============================================== company ===============================================

// CreateCompany implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) CreateCompany(ctx context.Context, req *sparkhire_runtime.CreateCompanyRequest) (resp *sparkhire_runtime.CreateCompanyResponse, err error) {
	resp, err = company.CreateCompany(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "create company failed: %v", err)
		resp = &sparkhire_runtime.CreateCompanyResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// EditCompany implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) EditCompany(ctx context.Context, req *sparkhire_runtime.EditCompanyRequest) (resp *sparkhire_runtime.EditCompanyResponse, err error) {
	resp, err = company.EditCompany(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "edit company failed: %v", err)
		resp = &sparkhire_runtime.EditCompanyResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// =============================================== biz ===============================================

// SendVerifyCode implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) SendVerifyCode(ctx context.Context, req *sparkhire_runtime.SendVerifyCodeRequest) (resp *sparkhire_runtime.SendVerifyCodeResponse, err error) {
	resp, err = biz.SendVerifyCode(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "send verify code failed: %v", err)
		resp = &sparkhire_runtime.SendVerifyCodeResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}

// UploadFile implements the SparkhireRuntimeServiceImpl interface.
func (s *SparkhireRuntimeServiceImpl) UploadFile(ctx context.Context, req *sparkhire_runtime.UploadFileRequest) (resp *sparkhire_runtime.UploadFileResponse, err error) {
	resp, err = biz.UploadFile(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "upload file failed: %v", err)
		resp = &sparkhire_runtime.UploadFileResponse{
			BaseResp: handler.GenErrorBaseResp(err.Error()),
		}
	}
	return resp, nil
}
