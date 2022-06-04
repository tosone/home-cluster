package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var helmreleaseCmd = &cobra.Command{
	Use:     "helmrelease",
	Aliases: []string{"hr"},
	Short:   "Install helmrelease",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			for _, arg := range args {
				err := installHelmRelease(arg)
				if err != nil {
					return err
				}
			}
			return nil
		}
		return installAllHelmReleases()
	},
}

func installAllHelmReleases() error {
	for _, release := range deploy.HelmReleases {
		err := installHelmRelease(release.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func installHelmRelease(name string) error {
	exist := false
	for _, release := range deploy.HelmReleases {
		if release.Name == name {
			exist = true
			return installHelmReleaseHelper(kubeconfig, release.Namespace, release.Name, release.Chart, release.Source, release.Values)
		}
	}
	if !exist {
		return fmt.Errorf("helmrelease not found: %s", name)
	}
	return nil
}

func installHelmReleaseHelper(kubeconfig, namespace, name, chart, source, values string) error {
	cmd := exec.Command("flux", "create", "helmrelease", name, "--source", source, "--chart", chart)
	if values != "" {
		file, err := ioutil.TempFile("", fmt.Sprintf("%s.*.yaml", name))
		if err != nil {
			return err
		}
		defer file.Close()           // nolint: errcheck
		defer os.Remove(file.Name()) // nolint: errcheck
		_, err = file.WriteString(values)
		if err != nil {
			return err
		}
		cmd.Args = append(cmd.Args, "--values", file.Name())
	}
	addArgs(cmd, kubeconfig, namespace)
	return runCommand(cmd)
}
