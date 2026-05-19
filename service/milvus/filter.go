package milvus

import (
	"fmt"
	"ice_sparkhire_runtime/utils"
	"strconv"
	"strings"
)

type FilterBuilder struct {
	filterParts []string
}

func NewFilterBuilder() *FilterBuilder {
	return &FilterBuilder{}
}

func (f *FilterBuilder) ExactMatchStr(colName, colValue string) *FilterBuilder {
	f.filterParts = append(f.filterParts, fmt.Sprintf("%s == '%s'", colName, colValue))
	return f
}

func (f *FilterBuilder) ExactMatchInt64(colName string, colValue int64) *FilterBuilder {
	f.filterParts = append(f.filterParts, fmt.Sprintf("%s == %d", colName, colValue))
	return f
}

func (f *FilterBuilder) NotExactMatchInt64(colName string, colValue int64) *FilterBuilder {
	f.filterParts = append(f.filterParts, fmt.Sprintf("%s != %d", colName, colValue))
	return f
}

func (f *FilterBuilder) InInt64Array(colName string, colValues []int64) *FilterBuilder {
	colValuesStr := utils.MapStructList(colValues, func(colValue int64) string {
		return strconv.FormatInt(colValue, 10)
	})
	f.filterParts = append(f.filterParts, fmt.Sprintf("%s in [%s]", colName, strings.Join(colValuesStr, ",")))
	return f
}

func (f *FilterBuilder) LikeMatch(colName, colValue string) *FilterBuilder {
	f.filterParts = append(f.filterParts, fmt.Sprintf("%s LIKE '%%%s%%'", colName, colValue))
	return f
}

func (f *FilterBuilder) AndMatch(conditions ...*FilterBuilder) *FilterBuilder {
	if conditions == nil || len(conditions) == 0 {
		return f
	}
	var conditionStrings []string
	conditionList := f.filterParts[len(f.filterParts)-len(conditions):]
	f.filterParts = f.filterParts[:len(f.filterParts)-len(conditions)]
	for _, condition := range conditionList {
		conditionStrings = append(conditionStrings, "("+condition+")")
	}
	f.filterParts = append(f.filterParts, strings.Join(conditionStrings, " AND "))
	return f
}

func (f *FilterBuilder) OrMatch(conditions ...*FilterBuilder) *FilterBuilder {
	if conditions == nil || len(conditions) == 0 {
		return f
	}
	var conditionStrings []string
	conditionList := f.filterParts[len(f.filterParts)-len(conditions):]
	f.filterParts = f.filterParts[:len(f.filterParts)-len(conditions)]
	for _, condition := range conditionList {
		conditionStrings = append(conditionStrings, "("+condition+")")
	}
	f.filterParts = append(f.filterParts, strings.Join(conditionStrings, " OR "))
	return f
}

func (f *FilterBuilder) Build() string {
	if len(f.filterParts) == 0 {
		return ""
	}

	return strings.Join(f.filterParts, " AND ")
}
