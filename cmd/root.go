package cmd

import (
	"fmt"
	"os"
	"vict-devv/s3-batch-uploader/constants"

	"github.com/spf13/cobra"
)

var (
	// Update it with ldflags in the go build command. E.g: go build ... -ldflags="-X 'root.version=x.y.z'" ...
	Version string = "0.0.1"

	AwsProfile string
	AwsRegion  string

	Folder string
)

var rootCmd = &cobra.Command{
	Use:   "s3-batch-uploader",
	Short: "S3 Batch Uploader is a simple tool that helps you perform upload of folders to a AWS S3 bucket",
	Long: `A Go commmand line tool that performs upload of folders to AWS S3 buckets, built with Go Cobra library
			to make it powerful and easy to extend. For more information access the README
			at my GitHub: https://github.com/vict-devv/s3-batch-uploader`,
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Profile: ", AwsProfile)
		fmt.Println("Region: ", AwsRegion)
		fmt.Println("Folder: ", Folder)

		// TODO: do the upload here
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&AwsProfile, "profile", "p", constants.AwsDefaultProfile, "AWS credential profile to be used (required)")
	rootCmd.MarkPersistentFlagRequired("profile")

	rootCmd.PersistentFlags().StringVarP(&AwsRegion, "region", "r", constants.AwsDefaultRegion, "The bucket region (required)")
	rootCmd.MarkPersistentFlagRequired("region")

	rootCmd.PersistentFlags().StringVarP(&Folder, "folder", "f", constants.DefaultFolder, "The path of the folder to be upload (required)")
	rootCmd.MarkPersistentFlagRequired("folder")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
