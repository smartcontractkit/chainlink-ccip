package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func main() {
	var pattern, src, dest string

	var rootCmd = &cobra.Command{
		Use:   "copy-to-pods",
		Short: "Copy a file to pods filtered by a regex pattern, ensuring target directory exist",
		Run: func(cmd *cobra.Command, args []string) {
			reg, err := regexp.Compile(pattern)
			if err != nil {
				panic(errors.Wrap(err, "invalid regex pattern"))
			}

			fmt.Printf("Copying %s to pods matching %s at %s\n", src, pattern, dest)

			out, err := exec.Command("kubectl", "get", "pods", "--no-headers", "-o", "custom-columns=:metadata.name").Output()
			if err != nil {
				panic(errors.Wrap(err, "failed to get pods"))
			}

			lines := strings.Split(string(out), "\n")
			var pods []string
			for _, line := range lines {
				podName := strings.TrimSpace(line)
				if podName != "" && reg.MatchString(podName) {
					pods = append(pods, podName)
				}
			}

			if len(pods) == 0 {
				panic(fmt.Sprintf("no pods found matching regex: %s", pattern))
			}

			destDir := path.Dir(dest)

			for _, pod := range pods {
				mkdirCmd := exec.Command("kubectl", "exec", pod, "--", "mkdir", "-p", destDir)
				var mkdirStderr bytes.Buffer
				mkdirCmd.Stderr = &mkdirStderr
				if err := mkdirCmd.Run(); err != nil {
					panic(fmt.Sprintf("failed to create directory %s in pod %s: %v - %s", destDir, pod, err, mkdirStderr.String()))
				}

				copyCmd := exec.Command("kubectl", "cp", src, fmt.Sprintf("%s:%s", pod, dest))
				var copyStderr bytes.Buffer
				copyCmd.Stderr = &copyStderr
				if err := copyCmd.Run(); err != nil {
					panic(fmt.Sprintf("failed to copy to pod %s: %v - %s", pod, err, copyStderr.String()))
				}
				fmt.Printf("Successfully copied %s to pod %s at %s\n", src, pod, dest)
			}
		},
	}

	rootCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "Regex pattern to filter pods by name")
	rootCmd.Flags().StringVarP(&src, "source", "s", "", "Path of the source file to copy")
	rootCmd.Flags().StringVarP(&dest, "destination", "d", "", "Destination file path in the pod")
	rootCmd.MarkFlagRequired("pattern")
	rootCmd.MarkFlagRequired("source")
	rootCmd.MarkFlagRequired("destination")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
