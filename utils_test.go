package gorm_test

import (
	"testing"

	"github.com/Cappta/gorm"
)

func TestDefaultNameHandlerGenerateFriendlyName(t *testing.T) {
	var maps = map[string]string{
		"":                          "",
		"X":                         "x",
		"ThisIsATest":               "this_is_a_test",
		"PFAndESI":                  "pf_and_esi",
		"AbcAndJkl":                 "abc_and_jkl",
		"EmployeeID":                "employee_id",
		"SKU_ID":                    "sku_id",
		"FieldX":                    "field_x",
		"HTTPAndSMTP":               "http_and_smtp",
		"HTTPServerHandlerForURLID": "http_server_handler_for_url_id",
		"UUID":     "uuid",
		"HTTPURL":  "http_url",
		"HTTP_URL": "http_url",
		"ThisIsActuallyATestSoWeMayBeAbleToUseThisCodeInGormPackageAlsoIdCanBeUsedAtTheEndAsID": "this_is_actually_a_test_so_we_may_be_able_to_use_this_code_in_gorm_package_also_id_can_be_used_at_the_end_as_id",
	}

	for key, value := range maps {
		if gorm.DefaultNameHandler(key) != value {
			t.Errorf("%v DefaultNameHandler should equal %v, but got %v", key, value, gorm.DefaultNameHandler(key))
		}
	}
}
