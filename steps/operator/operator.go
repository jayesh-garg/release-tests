package olm

import (
	"fmt"
	"log"

	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/getgauge-contrib/gauge-go/models"
	"github.com/openshift-pipelines/release-tests/pkg/cmd"
	"github.com/openshift-pipelines/release-tests/pkg/operator"
	"github.com/openshift-pipelines/release-tests/pkg/store"
)

var _ = gauge.Step("Update TektonConfig CR to use param with name createRbacResource and value <value> to <action> auto creation of RBAC resources", func(value, action string) {
	patchData := fmt.Sprintf("{\"spec\":{\"params\":[{\"name\":\"createRbacResource\",\"value\":\"%s\"}]}}", value)
	log.Println(action, "auto creation of RBAC resources")
	log.Printf("output: %s\n", cmd.MustSucceed("oc", "patch", "TektonConfig", "config", "--type=merge", "-p", patchData).Stdout())
})

var _ = gauge.Step("Verify RBAC resources disabled successfully", func() {
	operator.ValidateRBACAfterDisable(store.Clients(), store.GetCRNames())
})

var _ = gauge.Step("Verify RBAC resources are auto created successfully", func() {
	operator.ValidateRBAC(store.Clients(), store.GetCRNames())
})

var _ = gauge.Step("Verify the following roles are present in namespace <namespace>: <table>", func(namespace string, rolesTable *models.Table) {
	for _, row := range rolesTable.Rows {
		role := row.Cells[0]
		operator.VerifyRole(store.Clients(), role, namespace)
	}
})

var _ = gauge.Step("Verify total number of roles in namespace <namespace> is <expectedCount>", func(namespace, expectedCount string) {
	operator.VerifyNumberOfRoles(store.Clients(), namespace, expectedCount)
})

