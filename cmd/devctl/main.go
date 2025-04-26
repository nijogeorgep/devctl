package main

import (
	"devctl/internal/awshelper"
	"devctl/internal/githelper"
	"devctl/internal/kubehelper"
	"devctl/internal/netcheck"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	version   = "dev"
	gitSha    = "none"
	buildDate = "unknown"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("Version: %s\nGit SHA: %s\nBuilt at: %s\n", version, gitSha, buildDate)
		return
	}
	
	var rootCmd = &cobra.Command{Use: "devctl"}

	rootCmd.AddCommand(netcheck.NewNetCheckCmd())
	rootCmd.AddCommand(kubehelper.NewKubeHelperCmd())
	rootCmd.AddCommand(awshelper.NewAwsHelperCmd())
	rootCmd.AddCommand(githelper.NewGitHelperCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
