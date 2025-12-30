include "base.thrift"

namespace go sparkhire_runtime
namespace dart sparkhire.runtime

struct PingRequest {
	1:   optional string    ping
	255: required base.Base Base
}

struct PingResponse {
	1:   required string        pong
	255: required base.BaseResp BaseResp
}

const list<string> UnAuthPathList = [
    "/api/v1/ice/sparkhire/runtime/user/register",
    "/api/v1/ice/sparkhire/runtime/user/login",
    "/api/v1/ice/sparkhire/runtime/ping",
    "/api/v1/ice/sparkhire/runtime/verify/code/send"
]

// =============================================== user ===============================================

enum UserRole {
	Visitor   = 1
	Candidate = 2
	HR        = 3
	Admin     = 4
}

struct RegisterUserRequest {
	1:   required string    email
	2:   required string    verifyCode
	255: required base.Base Base
}

struct RegisterUserResponse {
	1:            string        accessToken
	255: required base.BaseResp BaseResp
}

struct UserBasicInfo {
	1:  i64      id
	2:  string   username
	3:  UserRole role
	4:  string   userAvatar
	5:  i8       gender
	6:  string   email
}

struct UserMailLoginRequest {
	1:   required string    email
	2:   required string    verifyCode
	255: required base.Base Base
}

struct UserMailLoginResponse {
	2:            string        accessToken
	255: required base.BaseResp BaseResp
}

struct FetchCurrentUserRequest {
	255: required base.Base Base
}

struct FetchCurrentUserResponse {
	1:            UserBasicInfo basicInfo
	255: required base.BaseResp BaseResp
}

struct SwitchUserRoleRequest {
	1:   required UserRole  userRole
	255: required base.Base Base
}

struct SwitchUserRoleResponse {
	255: required base.BaseResp BaseResp
}


// =============================================== Candidate ===============================================

enum EducationStatus {
	Bachelor        = 1    // 本科
	Master          = 2    // 研究生
	Doctor          = 3    // 博士生
	JuniorCollege   = 4    // 大专
	HighSchool      = 5    // 高中
	BelowHighSchool = 6    // 高中以下
}

const list<EducationStatus> EducationStatusList = [
    EducationStatus.Bachelor,
    EducationStatus.Master,
    EducationStatus.Doctor,
    EducationStatus.JuniorCollege,
    EducationStatus.HighSchool,
    EducationStatus.BelowHighSchool
]

enum JobStatus {
	Available       = 1    // 随时到岗
	WithInMonth     = 2    // 月内到岗
	OpenOpportunity = 3    // 考虑机会
	NotInterested   = 4    // 暂不考虑
}

const list<JobStatus> JobStatusList = [JobStatus.Available, JobStatus.WithInMonth, JobStatus.OpenOpportunity, JobStatus.NotInterested]

struct GeoDetailInfo {
	1:  required i64    firstGeoLevelId
	2:  required i64    secondGeoLevelId
	3:  required i64    thirdGeoLevelId
	4:  required i64    forthGeoLevelId
	5:  optional string address
	6:  optional double latitude
	7:  optional double longitude

	8:  optional string firstGeoLevelName
	9:  optional string secondGeoLevelName
	10: optional string thirdGeoLevelName
	11: optional string forthGeoLevelName
}

struct GeoModifyInfo {
	1: required i64    firstGeoLevelId
	2: required i64    secondGeoLevelId
	3: required i64    thirdGeoLevelId
	4: required i64    forthGeoLevelId
	5: required string address
	6: required double latitude
	7: required double longitude
}

struct ContractInfo {
	1: required string        phone
	2: required string        email
	3: required GeoDetailInfo geoInfo
}

struct CandidateInfo {
	1: required i32             age
	3: required JobStatus       jobStatus
	4: required ContractInfo    ContractInfo
	5: required i32             graduationYear
	6: required EducationStatus educationStatus
	7: optional i64             id
	8: optional list<TagInfo>   tagList
}

//struct AddCandidateRequest {
//	1:   required CandidateInfo candidateInfo
//	255: required base.Base     Base
//}
//
//struct AddCandidateResponse {
//	1:            i64           id
//	255: required base.BaseResp BaseResp
//}
//
//struct EditCandidateRequest {
//	1:   required CandidateInfo candidateInfo
//	255: required base.Base     Base
//}
//
//struct EditCandidateResponse {
//	255: required base.BaseResp BaseResp
//}

struct EditCandidateContractInfoRequest {
	1:   required GeoModifyInfo geoInfo
	2:   optional string        phoneNumber
	255: required base.Base     Base
}

struct EditCandidateContractInfoResponse {
	255: required base.BaseResp BaseResp
}

struct GetCurrentCandidateRequest {
	255: required base.Base Base
}

struct GetCurrentCandidateResponse {
	1:            CandidateInfo candidateInfo
	255: required base.BaseResp BaseResp
}

// =============================================== tag ===============================================

struct TagInfo {
	1:  i64    id
	2:  string tagName
}

struct QueryTagRequest {
	1:   optional string    searchText
	2:            i32       pageNum
	3:            i32       pageSize
	255: required base.Base Base
}

struct QueryTagResponse {
	1:            i64           total
	2:            list<TagInfo> tagList
	255: required base.BaseResp BaseResp
}

struct AddTagRequest {
	1:   required string    tagName
	255: required base.Base Base
}

struct AddTagResponse {
	1:            i64           id
	255: required base.BaseResp BaseResp
}

enum TagObjType {
	Candidate   = 1
	Recruitment = 2
}

struct BindTagsRequest {
	1:   required TagObjType objType
	2:   required list<i64>  tagIdList
	3:   required i64        objId
	255: required base.Base  Base
}

struct BindTagsResponse {
	1:            i64           num
	255: required base.BaseResp BaseResp
}

struct UnbindTagsRequest {
	1:   required TagObjType objType
	2:   required list<i64>  tagIdList
	3:   required i64        objId
	255: required base.Base  Base
}

struct UnbindTagsResponse {
	1:            i64           num
	255: required base.BaseResp BaseResp
}

// =============================================== information ===============================================

struct MajorInfo {
	1:  i64    id
	2:  string majorName
}

struct ListMajorRequest {
	255: required base.Base Base
}

struct ListMajorResponse {
	1:            list<MajorInfo> majorList
	255: required base.BaseResp   BaseResp
}

struct IndustryDetail {
	1:  i64    id
	2:  string industryName
}

struct IndustryInfo {
	1:  string               industryTypeName
	2:  list<IndustryDetail> industryList
}

struct ListIndustryRequest {
	255: required base.Base Base
}

struct ListIndustryResponse {
	1:            list<IndustryInfo> industryInfoList
	255: required base.BaseResp      BaseResp
}

struct SchoolInfo {
	1:  i64    id
	2:  string schoolName
}

struct ListSchoolRequest {
	255: required base.Base Base
}

struct ListSchoolResponse {
	1:            list<SchoolInfo> schoolList
	255: required base.BaseResp    BaseResp
}

enum GeoLevel {
	FirstGeo  = 1
	SecondGeo = 2
	ThirdGeo  = 3
	ForthGeo  = 4
}

struct GeoInfo {
	1: required i64           id
	2: required string        name
	3: required string        code
	4: required GeoLevel      level
}

struct ListGeoRequest {
	1:   required GeoLevel  level
	2:   optional i64       parentId
	255: required base.Base Base
}

struct ListGeoResponse {
	1:            list<GeoInfo> geoList
	255: required base.BaseResp BaseResp
}

// =============================================== education experience ===============================================

struct CreateEducationExpRequest {
	1:   required i64             schoolId
	2:   required EducationStatus status
	3:   required i32             beginYear
	4:   required i32             endYear
	5:   required i64             majorId
	6:   required string          activity
	255: required base.Base       Base
}

struct CreateEducationExpResponse {
	255: required base.BaseResp BaseResp
}

struct EditEducationExpRequest {
	1:   required i64       id
	2:   required i64       schoolId
	3:   required i32       beginYear
	4:   required i32       endYear
	5:   required i64       majorId
	6:   required string    activity
	255: required base.Base Base
}

struct EditEducationExpResponse {
	255: required base.BaseResp BaseResp
}

struct DeleteEducationExpRequest {
	1:   required i64       id
	255: required base.Base Base
}

struct DeleteEducationExpResponse {
	255: required base.BaseResp BaseResp
}

struct EducationExpInfo {
	1:  i64        id
	2:  SchoolInfo schoolInfo
	3:  i32        beginYear
	4:  i32        endYear
	5:  MajorInfo  majorInfo
	6:  string     activity
}

struct GetCurrentUserEducationExpRequest {
	255: required base.Base Base
}

struct GetCurrentUserEducationExpResponse {
	1:            list<EducationExpInfo> educationExpList
	255: required base.BaseResp          BaseResp
}

// =============================================== career experience ===============================================

struct CreateCareerExperienceRequest {
	1:   required string    experienceName
	2:   required string    jobRole
	3:   required string    description
	4:   required i64       startTS
	5:   required i64       endTS
	255: required base.Base Base
}

struct CreateCareerExperienceResponse {
	255: required base.BaseResp BaseResp
}

struct EditCareerExperienceRequest {
	1:   required string    experienceName
	2:   required string    jobRole
	3:   required string    description
	4:   required i64       startTS
	5:   required i64       endTS
	6:   required i64       id
	255: required base.Base Base
}

struct EditCareerExperienceResponse {
	255: required base.BaseResp BaseResp
}

struct CareerExperienceInfo {
	1: required string experienceName
	2: required string jobRole
	3: required string description
	4: required i64    startTS
	5: required i64    endTS
	6: required i64    id
}

struct GetCurrentUserCareerExperienceRequest {
	255: required base.Base Base
}

struct GetCurrentUserCareerExperienceResponse {
	1:            list<CareerExperienceInfo> careerExperienceInfoList
	255: required base.BaseResp              BaseResp
}

struct DeleteCareerExperienceRequest {
	1:   required i64       id
	255: required base.Base Base
}

struct DeleteCareerExperienceResponse {
	255: required base.BaseResp BaseResp
}

// =============================================== biz ===============================================

struct SendVerifyCodeRequest {
	1:   required string    email
	255: required base.Base Base
}

struct SendVerifyCodeResponse {
	255: required base.BaseResp BaseResp
}

enum FileType {
	CompanyAvatar = 1
}

const map<FileType, string> FileTypeMap = {
    FileType.CompanyAvatar: "company/avatar",
}

struct UploadFileRequest {
	1:   required binary    file
	255: required base.Base Base
}

struct UploadFileResponse {
	1:            string        fileKey
	255: required base.BaseResp BaseResp
}

// =============================================== company ===============================================

struct CreateCompanyRequest {
	1:   required string        companyName
	2:   required string        description
	3:   optional string        logo
	4:   required i64           industryId
	5:   optional GeoModifyInfo geoInfo
	255: required base.Base     Base
}

struct CreateCompanyResponse {
	255: required base.BaseResp BaseResp
}

struct EditCompanyRequest {
	1:   required i64           id
	2:   required string        companyName
	3:   required string        description
	4:   optional string        logo
	5:   required i64           industryId
	6:   optional GeoModifyInfo geoInfo
	255: required base.Base     Base
}

struct EditCompanyResponse {
	255: required base.BaseResp BaseResp
}

struct DeleteCompanyRequest {
	1:   required i64       id
	255: required base.Base Base
}

struct DeleteCompanyResponse {
	255: required base.BaseResp BaseResp
}

service SparkhireRuntimeService {
    PingResponse Ping(1: PingRequest req) (api.post="/api/v1/ice/sparkhire/runtime/ping", api.serializer="json")

    // =============================================== user ===============================================
    RegisterUserResponse RegisterUser(1: RegisterUserRequest req) (api.post="/api/v1/ice/sparkhire/runtime/user/register", api.serializer="json")
    UserMailLoginResponse UserMailLogin(1: UserMailLoginRequest req) (api.post="/api/v1/ice/sparkhire/runtime/user/login", api.serializer="json")
    FetchCurrentUserResponse FetchCurrentUser(1: FetchCurrentUserRequest req) (api.post="/api/v1/ice/sparkhire/runtime/user/current/fetch", api.serializer="json")
    SwitchUserRoleResponse SwitchUserRole(1: SwitchUserRoleRequest req) (api.post="/api/v1/ice/sparkhire/runtime/user/role/switch", api.serializer="json")

    // =============================================== candidate ===============================================
//    AddCandidateResponse AddCandidate(1: AddCandidateRequest req) (api.post="/api/v1/ice/sparkhire/runtime/user/candidate/add", api.serializer="json")
//    EditCandidateResponse EditCandidate(1: EditCandidateRequest req) (api.post="/api/v1/ice/sparkhire/runtime/user/candidate/edit", api.serializer="json")
    GetCurrentCandidateResponse GetCurrentCandidate(1: GetCurrentCandidateRequest req) (api.post="/api/v1/ice/sparkhire/runtime/user/candidate/current/get", api.serializer="json")
    EditCandidateContractInfoResponse EditCandidateContractInfo(1: EditCandidateContractInfoRequest req) (api.post="/api/v1/ice/sparkhire/runtime/user/candidate/contract/edit", api.serializer="json")

    // =============================================== tag ===============================================
    QueryTagResponse QueryTag(1: QueryTagRequest req) (api.post="/api/v1/ice/sparkhire/runtime/tag/query", api.serializer="json")
    AddTagResponse AddTag(1: AddTagRequest req) (api.post="/api/v1/ice/sparkhire/runtime/tag/add", api.serializer="json")
    BindTagsResponse BindTags(1: BindTagsRequest req) (api.post="/api/v1/ice/sparkhire/runtime/tag/bind", api.serializer="json")
    UnbindTagsResponse UnbindTags(1: UnbindTagsRequest req) (api.post="/api/v1/ice/sparkhire/runtime/tag/unbind", api.serializer="json")

    // =============================================== education experience ===============================================
    CreateEducationExpResponse CreateEducationExp(1: CreateEducationExpRequest req) (api.post="/api/v1/ice/sparkhire/runtime/education/exp/create", api.serializer="json")
    EditEducationExpResponse EditEducationExp(1: EditEducationExpRequest req) (api.post="/api/v1/ice/sparkhire/runtime/education/exp/edit", api.serializer="json")
    DeleteEducationExpResponse DeleteEducationExp(1: DeleteEducationExpRequest req) (api.post="/api/v1/ice/sparkhire/runtime/education/exp/delete", api.serializer="json")
    GetCurrentUserEducationExpResponse GetCurrentEducationExp(1: GetCurrentUserEducationExpRequest req) (api.post="/api/v1/ice/sparkhire/runtime/education/exp/current", api.serializer="json")

    // =============================================== career experience ===============================================
    CreateCareerExperienceResponse CreateCareerExperience(1: CreateCareerExperienceRequest req) (api.post="/api/v1/ice/sparkhire/runtime/career/exp/create", api.serializer="json")
    EditCareerExperienceResponse EditCareerExperience(1: EditCareerExperienceRequest req) (api.post="/api/v1/ice/sparkhire/runtime/career/exp/edit", api.serializer="json")
    GetCurrentUserCareerExperienceResponse GetCurrentUserCareerExperience(1: GetCurrentUserCareerExperienceRequest req) (api.post="/api/v1/ice/sparkhire/runtime/career/exp/current", api.serializer="json")
    DeleteCareerExperienceResponse DeleteCareerExperience(1: DeleteCareerExperienceRequest req) (api.post="/api/v1/ice/sparkhire/runtime/career/exp/delete", api.serializer="json")

    // =============================================== information ===============================================
    ListMajorResponse ListMajor(1: ListMajorRequest req) (api.post="/api/v1/ice/sparkhire/runtime/major/list", api.serializer="json")
    ListIndustryResponse ListIndustry(1: ListIndustryRequest req) (api.post="/api/v1/ice/sparkhire/runtime/industry/list", api.serializer="json")
    ListSchoolResponse ListSchool(1: ListSchoolRequest req) (api.post="/api/v1/ice/sparkhire/runtime/school/list", api.serializer="json")
    ListGeoResponse ListGeo(1: ListGeoRequest req) (api.post="/api/v1/ice/sparkhire/runtime/geo/list", api.serializer="json")

    // =============================================== company ===============================================
    CreateCompanyResponse CreateCompany(1: CreateCompanyRequest req) (api.post="/api/v1/ice/sparkhire/runtime/user/hr/company/create", api.serializer="json")
    EditCompanyResponse EditCompany(1: EditCompanyRequest req) (api.post="/api/v1/ice/sparkhire/runtime/user/hr/company/edit", api.serializer="json")
    DeleteCompanyResponse DeleteCompany(1: DeleteCompanyRequest req) (api.post="/api/v1/ice/sparkhire/runtime/user/hr/company/delete", api.serializer="json")

    // =============================================== biz ===============================================
    SendVerifyCodeResponse SendVerifyCode(1: SendVerifyCodeRequest req) (api.post="/api/v1/ice/sparkhire/runtime/verify/code/send", api.serializer="json")
    UploadFileResponse UploadFile(1: UploadFileRequest req) (api.post="/api/v1/ice/sparkhire/runtime/verify/file/upload", api.serializer="json")

} (agw.preserve_base="true", agw.js_conv="str")