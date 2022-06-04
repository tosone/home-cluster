package commands

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"home-cluster/cmd/types"
)

func runCommand(cmd *exec.Cmd) error {
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	logrus.Info("run command: ", cmd.String())
	return cmd.Run()
}

// addArgs adds the args to the command.
func addArgs(cmd *exec.Cmd, kubeconfig, namespace string) {
	if namespace != "" {
		cmd.Args = append(cmd.Args, "--namespace", namespace)
	}
	if kubeconfig != "" {
		cmd.Args = append(cmd.Args, "--kubeconfig", kubeconfig)
	}
}

func compileDeploy() (*types.Deploy, error) {
	deploy, err := getDeploy()
	if err != nil {
		logrus.Fatalf("get deploy with error: %v", err)
	}
	t := template.Must(template.New("deploy").Funcs(sprig.TxtFuncMap()).Parse(deploy))
	config, err := getConfig()
	if err != nil {
		logrus.Fatalf("get config with error: %v", err)
	}
	var deployed bytes.Buffer
	err = t.Execute(&deployed, config)
	if err != nil {
		logrus.Fatalf("execute template with error: %v", err)
	}
	d := types.Deploy{}
	err = yaml.NewDecoder(strings.NewReader(deployed.String())).Decode(&d)
	if err != nil {
		logrus.Fatal("Error decoding file:", err)
	}
	return &d, err
}

// getConfig ...
func getConfig() (*types.Config, error) {
	file, err := os.Open(envFile)
	if err != nil {
		return nil, err
	}
	defer file.Close() // nolint: errcheck
	config := types.Config{}
	err = yaml.NewDecoder(file).Decode(&config)
	return &config, err
}

// getDeploy ...
func getDeploy() (string, error) {
	file, err := os.Open(deployFile)
	if err != nil {
		return "", err
	}
	defer file.Close() // nolint: errcheck
	content, err := ioutil.ReadAll(file)
	return string(content), err
}
