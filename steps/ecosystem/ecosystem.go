package ecosystem

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/getgauge-contrib/gauge-go/testsuit"
	"github.com/openshift-pipelines/release-tests/pkg/oc"
)

var _ = gauge.Step("Verify that jib-maven image registry variable is exported", func() {
	if os.Getenv("jib_maven_repository") == "" {
		testsuit.T.Errorf("'jib_maven_repository' environment variable is not exported")
	}
})

var _ = gauge.Step("Create secret with image registry credentials for jib-maven", func() {
	quayUser := os.Getenv("QUAY_USER")
	quayPass := os.Getenv("QUAY_PASS")

	if quayUser == "" || quayPass == "" {
		testsuit.T.Errorf("QUAY_USER or QUAY_PASS environment variables are not exported")
	} else {
		auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", quayUser, quayPass)))
		dockerConfig := fmt.Sprintf(`{"auths":{"quay.io":{"auth":"%s"}}}`, auth)
		oc.CreateJibMavenImageRegistrySecret(dockerConfig)
	}
})
