package systemTests

import (
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Docker Image", func() {

	commands := map[string][]string{
		"aws":     []string{"--version"},
		"bosh":    []string{"-v"},
		"bundle":  []string{"-v"},
		"cf":      []string{"version"},
		"credhub": []string{"--version"},
		"dep":     []string{"version"},
		"gcloud":  []string{"--version"},
		"ginkgo":  []string{"help"},
		"go":      []string{"doc", "cmd/vet"},
		"govc":    []string{"version"},
		"haproxy": []string{"-v"},
		"ipcalc":  []string{},
		"jq":      []string{"--version"},
		"kubectl": []string{"help"},
		"make":    []string{"-v"},
		"om":      []string{"--version"},
		"rake":    []string{"--version"},
		"ruby":    []string{"-v"},
		"semver":  []string{"--help"},
		"spruce":  []string{"--version"},
		"vegeta":  []string{"-version"},
		"which":   []string{"sshuttle"},
		"zip":     []string{"--version"},
	}

	for executable, args := range commands {
		executable := executable
		args := args

		It(fmt.Sprintf("has %v installed", executable), func() {
			command := exec.Command(executable, args...)
			session, err := Start(command, GinkgoWriter, GinkgoWriter)

			Expect(err).ToNot(HaveOccurred())
			Eventually(session, "5s").Should(Exit(0))
		})
	}

	It("has an OpenBSD netcat", func() {
		command := exec.Command("nc", "-h")
		session, err := Start(command, GinkgoWriter, GinkgoWriter)

		Expect(err).ToNot(HaveOccurred())
		Eventually(session.Err).Should(gbytes.Say("OpenBSD"))
		Eventually(session, "5s").Should(Exit(0))
	})
})
