/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/utils"
	"github.com/spf13/cobra"
)

var (
	eth1Client      string
	consensusClient string
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Quick start 1Click",
	Long: `Run the setup tool on-premise in a quick way. Provide only the command line
options and the tool will do all the work.

First it will check if dependencies like docker and docker-compose are installed on your machine
and provide instructions for installing them if they are not installed.

Second, it will generate docker-compose scripts to run the full setup according to your selection.

Finally, it will run the generated docker-compose script`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if dependencies are installed
		dependencies := configs.GetDependencies()
		log.Infof(configs.CheckingDependencies, strings.Join(dependencies, ", "))
		pending := utils.CheckDependencies(dependencies)

		if len(pending) > 0 {
			log.Infof(configs.DependenciesPending, strings.Join(pending, ", "))
			// Let the user decide to see the instructions for installing dependencies and exit or let the tool install them and continue
			err := utils.InstallDependencies(pending)
			if err != nil {
				log.Fatalf(configs.InstallingDependenciesError, err)
			}
		}

		log.Info(configs.DependenciesOK)
		log.Info(configs.GeneratingDockerComposeScript)
		//TODO: Implement logic for generating docker-compose scripts
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	// Local flags
	cliCmd.Flags().StringVar(&eth1Client, "execution", "", "Execution engine client, e.g. Geth, Nethermind, Besu, Erigon")
	cliCmd.MarkFlagRequired("execution")

	cliCmd.Flags().StringVar(&consensusClient, "consensus", "", "Consensus engine client, e.g. Teku, Lodestar, Prysm, Lighthouse, Nimbus")
	cliCmd.MarkFlagRequired("consensus")
}