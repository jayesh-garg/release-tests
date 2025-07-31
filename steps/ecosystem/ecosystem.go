package ecosystem

import (
	"encoding/base64"
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
	dockerConfigJson := oc.GetDockerConfigJson("quay-io-dockerconfig", "openshift-pipelines")
	dockerConfig, err := base64.StdEncoding.DecodeString(dockerConfigJson)
	if err != nil {
		testsuit.T.Errorf("failed to decode docker config json: %v", err)
	}
	oc.CreateJibMavenImageRegistrySecret(string(dockerConfig))
})
