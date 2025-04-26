package githelper

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func NewGitHelperCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "git",
		Short: "Perform common Git actions quickly",
	}

	cmd.AddCommand(cloneCmd())
	cmd.AddCommand(checkoutCmd())
	cmd.AddCommand(commitCmd())
	cmd.AddCommand(pushCmd())

	return cmd
}

func cloneCmd() *cobra.Command {
	var repo string
	var dir string

	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone a Git repository",
		Run: func(cmd *cobra.Command, args []string) {
			if repo == "" {
				log.Fatal("❌ --repo is required")
			}

			args = []string{"clone", repo}
			if dir != "" {
				args = append(args, dir)
			}

			runGitCommand(args...)
		},
	}

	cmd.Flags().StringVarP(&repo, "repo", "r", "", "Repository URL")
	cmd.Flags().StringVarP(&dir, "dir", "d", "", "Target directory (optional)")
	return cmd
}

func checkoutCmd() *cobra.Command {
	var branch string

	cmd := &cobra.Command{
		Use:   "checkout",
		Short: "Checkout a Git branch",
		Run: func(cmd *cobra.Command, args []string) {
			if branch == "" {
				log.Fatal("❌ --branch is required")
			}
			runGitCommand("checkout", branch)
		},
	}

	cmd.Flags().StringVarP(&branch, "branch", "b", "", "Branch name")
	return cmd
}

func commitCmd() *cobra.Command {
	var message string

	cmd := &cobra.Command{
		Use:   "commit",
		Short: "Create a Git commit",
		Run: func(cmd *cobra.Command, args []string) {
			if message == "" {
				log.Fatal("❌ --message is required")
			}

			runGitCommand("add", ".")
			runGitCommand("commit", "-m", message)
		},
	}

	cmd.Flags().StringVarP(&message, "message", "m", "", "Commit message")
	return cmd
}

func pushCmd() *cobra.Command {
	var remote string
	var branch string

	cmd := &cobra.Command{
		Use:   "push",
		Short: "Push changes to a Git remote",
		Run: func(cmd *cobra.Command, args []string) {
			if remote == "" {
				remote = "origin"
			}
			if branch == "" {
				out, err := exec.Command("git", "branch", "--show-current").Output()
				if err != nil {
					log.Fatalf("❌ Failed to determine current branch: %v", err)
				}
				branch = strings.TrimSpace(string(out))
			}
			runGitCommand("push", remote, branch)
		},
	}

	cmd.Flags().StringVarP(&remote, "remote", "r", "origin", "Git remote")
	cmd.Flags().StringVarP(&branch, "branch", "b", "", "Git branch")
	return cmd
}

func runGitCommand(args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Printf("▶️ git %s\n", strings.Join(args, " "))
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ git command failed: %v", err)
	}
}
