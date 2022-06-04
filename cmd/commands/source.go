package commands

import (
	"errors"
	"os/exec"

	"github.com/spf13/cobra"
)

var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Install source",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			for _, arg := range args {
				err := installSource(arg)
				if err != nil {
					return err
				}
			}
			return nil
		}
		return installAllSources()
	},
}

func installAllSources() error {
	for _, deploy := range deploy.Sources {
		for _, source := range deploy {
			err := installSource(source.Name)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func installSource(name string) error {
	exist := false
	for typ, deploy := range deploy.Sources {
		for _, source := range deploy {
			if source.Name == name {
				exist = true
				switch typ {
				case "git":
					return installGitSource(kubeconfig, source.Namespace, source.Name, source.URL, source.Branch, source.Tag)
				case "helm":
					return installHelmSource(kubeconfig, source.Namespace, source.Name, source.URL)
				default:
					return errors.New("unknown source type")
				}
			}
		}
	}
	if !exist {
		return errors.New("source not found")
	}
	return nil
}

func installHelmSource(kubeconfig, namespace, name, url string) error {
	cmd := exec.Command("flux", "create", "source", "helm", name, "--url", url)
	addArgs(cmd, kubeconfig, namespace)
	return runCommand(cmd)
}

func installGitSource(kubeconfig, namespace, name, url, branch, tag string) error {
	cmd := exec.Command("flux", "create", "source", "git", name, "--url", url)
	if branch != "" {
		cmd.Args = append(cmd.Args, "--branch", branch)
	}
	if tag != "" {
		cmd.Args = append(cmd.Args, "--tag", tag)
	}
	addArgs(cmd, kubeconfig, namespace)
	return runCommand(cmd)
}
