package commands

import (
	"errors"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"home-cluster/cmd/types"
)

var kubeconfig string
var deployFile string
var envFile string

var deploy *types.Deploy

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "home-cluster",
	Short: "home cluster",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if kubeconfig == "" {
			kubeconfig = os.Getenv("KUBECONFIG")
			if kubeconfig == "" {
				home, err := os.UserHomeDir()
				if err != nil {
					return err
				}
				fileinfo, err := os.Stat(path.Join(home, ".kube", "config"))
				if errors.Is(err, os.ErrNotExist) {
					return errors.New("kubeconfig not found")
				} else if err != nil {
					return err
				}
				kubeconfig = fileinfo.Name()
			}
		}
		if deployFile == "" {
			fileinfo, err := os.Stat("deploy.yaml")
			if errors.Is(err, os.ErrNotExist) {
				return errors.New("deploy file not found")
			} else if err != nil {
				return err
			}
			deployFile = fileinfo.Name()
		}
		if envFile == "" {
			fileinfo, err := os.Stat("env.yaml")
			if errors.Is(err, os.ErrNotExist) {
				return errors.New("env file not found")
			} else if err != nil {
				return err
			}
			envFile = fileinfo.Name()
		}
		var err error
		deploy, err = compileDeploy()
		if err != nil {
			logrus.Fatalf("compile deploy with error: %v", err)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		deploy, err := compileDeploy()
		if err != nil {
			logrus.Fatalf("compile deploy with error: %v", err)
		}
		logrus.Info(deploy)
		logrus.Info(viper.GetString("KUBECONFIG"))
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if outside of cluster")
	rootCmd.PersistentFlags().StringVar(&deployFile, "deploy", "", "Path to a deploy file")
	rootCmd.PersistentFlags().StringVar(&envFile, "env", "", "Path to a env file")
	rootCmd.AddCommand(
		sourceCmd,
		helmreleaseCmd,
	)
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
