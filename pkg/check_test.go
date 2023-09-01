package pkg

import (
	"fmt"
	"riskanly/conf"
	"riskanly/store"
	"testing"
)

func TestA(t *testing.T) {
	tmp :=store.Task{
		DSN: conf.DsnNewBarinSaas,
	}
	table_info :=TableInfo{
		Schema: "ads",
		IfCheck: true,
	}
	tmp.Start()

	var allname = []string{
		"ads_stds_park_location_traffic",
		"ads_stds_park_carrier",
		"ads_stds_industry_strategic_label",
		"ads_lget_company_mktval",
		"ads_lget_finance_reporte_deri",
		"ads_psif_person_info",
		"ads_lget_company_legal_rep_inv_info",
		"ads_lget_company_actual_control",
		"ads_lget_company_beneficiary",
		"ads_lget_bond_info",
		"ads_psif_tech_award",
		"ads_lget_standard_info",
		"ads_lget_ic_layout",
		"ads_lget_company_illegal_info",
		"ads_lget_company_lose_trust",
		"ads_lget_company_simple_cancel",
		"ads_lget_company_event",
	}
	for i := 0; i < len(allname); i++ {
		checktable := &TableInfo{
			Schema: table_info.Schema,
			Name:   allname[i],
			CheckFields: map[string]string{
				"id":          "不存在",
				"update_time": "不存在",
				"is_delete":   "不存在",
				"create_time": "不存在",
			},
			IfCheck: table_info.IfCheck,
		}
		checktable.QueryAllColumns(tmp.Engine)
		fmt.Println(checktable)
	}
}
