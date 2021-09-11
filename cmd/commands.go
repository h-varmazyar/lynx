package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "gob",
	Short: "GoB is a simple golang project builder.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func initNewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "new",
		Short: "Create new golang project",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
			} else {
				createNewProject(args[0], cmd.Flags())
			}
		},
	}
	command.Flags().BoolP("skip-env", "e", false, "Ignore adding env file in the project")
	command.Flags().BoolP("skip-go-module", "u", false, "Ignore adding go module to the project")
	command.Flags().BoolP("skip-docker", "d", false, "Ignore adding docker to the project")
	command.Flags().BoolP("skip-git", "g", false, "Ignore adding git to the project")
	command.Flags().BoolP("force", "f", false, "Forcibly create project. delete all previous files if available")
	return command
}
