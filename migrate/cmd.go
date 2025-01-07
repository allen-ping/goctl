package migrate

import "github.com/allen-ping/goctl/v2/internal/cobrax"

var (
	boolVarVerbose   bool
	stringVarVersion string
	// Cmd describes a migrate command.
	Cmd = cobrax.NewCommand("migrate", cobrax.WithRunE(migrate))
)

func init() {
	migrateCmdFlags := Cmd.Flags()
	migrateCmdFlags.BoolVarP(&boolVarVerbose, "verbose", "v")
	migrateCmdFlags.StringVarWithDefaultValue(&stringVarVersion, "version", defaultMigrateVersion)
}
