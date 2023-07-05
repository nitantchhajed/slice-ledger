package flags

import (
	"fmt"
	"strings"

	"github.com/sliceledger-blockchain/slice-ledger/op-node/chaincfg"
	"github.com/sliceledger-blockchain/slice-ledger/op-node/flags"
	opservice "github.com/sliceledger-blockchain/slice-ledger/op-service"
	oplog "github.com/sliceledger-blockchain/slice-ledger/op-service/log"
	"github.com/urfave/cli/v2"
)

const envVarPrefix = "OP_BOOTNODE"

func prefixEnvVars(name string) []string {
	return opservice.PrefixEnvVar(envVarPrefix, name)
}

var (
	RollupConfig = &cli.StringFlag{
		Name:    flags.RollupConfig.Name,
		Usage:   "Rollup chain parameters",
		EnvVars: prefixEnvVars("ROLLUP_CONFIG"),
	}
	Network = &cli.StringFlag{
		Name:    flags.Network.Name,
		Usage:   fmt.Sprintf("Predefined network selection. Available networks: %s", strings.Join(chaincfg.AvailableNetworks(), ", ")),
		EnvVars: prefixEnvVars("NETWORK"),
	}
)

var Flags = []cli.Flag{
	RollupConfig,
	Network,
}

func init() {
	Flags = append(Flags, oplog.CLIFlags(envVarPrefix)...)
}
