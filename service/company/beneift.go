package company

import (
	"context"
	"fmt"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime"
	"ice_sparkhire_runtime/model/db"
	"ice_sparkhire_runtime/utils"
)

func BuildBenefitInfoList(ctx context.Context, companyId int64) ([]*sparkruntime.BenefitInfo, error) {
	// 2. find benefit category
	benefitCategories, err := db.FindCompanyBenefitCategoryByCompanyId(ctx, db.DB, companyId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch company benefit categories: %w", err)
	}

	if len(benefitCategories) == 0 {
		return nil, nil
	}

	categoryIds := utils.MapStructListDistinct(benefitCategories, func(benefitCategory *db.CompanyBenefitCategory) int64 {
		return benefitCategory.Id
	})

	benefitItemList, err := db.FindCompanyBenefitItemByCategoryIds(ctx, db.DB, categoryIds)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch company benefit items: %w", err)
	}

	benefitDetailListMap := make(map[int64][]*sparkruntime.BenefitDetail)
	for _, item := range benefitItemList {
		detail := &sparkruntime.BenefitDetail{
			Id:      item.Id,
			Title:   item.Title,
			Content: item.Content,
		}

		benefitDetailListMap[item.CategoryId] = append(
			benefitDetailListMap[item.CategoryId],
			detail,
		)
	}

	// 组装最后返回结果
	benefitInfoList := make([]*sparkruntime.BenefitInfo, 0, len(benefitItemList))
	for _, category := range benefitCategories {
		benefitInfoList = append(benefitInfoList, &sparkruntime.BenefitInfo{
			Id:       category.Id,
			Title:    category.Title,
			SubTitle: category.Subtitle,
			ItemList: benefitDetailListMap[category.Id],
		})
	}

	return benefitInfoList, nil
}
