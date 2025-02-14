package dao

import (
	"github.com/laizy/log"
	"github.com/laizy/web3"
	"github.com/laizy/web3/contract"
	"github.com/laizy/web3/utils"
	"github.com/ontology-layer-2/rollup-contracts/binding"
	"github.com/ontology-layer-2/rollup-contracts/cmd/rollupcli/common"
	"github.com/ontology-layer-2/rollup-contracts/cmd/rollupcli/flags"
	"github.com/urfave/cli/v2"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:        "dao",
		Subcommands: SubCommand(),
		Flags: []cli.Flag{
			flags.ConfigFlag,
		},
	}
}

func SubCommand() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "sequencerWhitelist",
			Usage: "set sequencer whitelist",
			Flags: []cli.Flag{
				flags.AccountFlag,
				flags.EnabledFlag,
				flags.SubmitFlag,
			},
			Action: sequencerWhitelist,
		},
		{
			Name:  "proposerWhitelist",
			Usage: "set proposer whitelist",
			Flags: []cli.Flag{
				flags.AccountFlag,
				flags.EnabledFlag,
				flags.SubmitFlag,
			},
			Action: proposerWhitelist,
		},
	}
}

func sequencerWhitelist(ctx *cli.Context) error {
	params, err := parseCtx(ctx)
	if err != nil {
		return err
	}
	setSequencerWhitelist(params)
	return nil
}

func proposerWhitelist(ctx *cli.Context) error {
	params, err := parseCtx(ctx)
	if err != nil {
		return err
	}
	setProposerWhitelist(params)
	return nil
}

type param struct {
	signer  *contract.Signer
	dao     web3.Address
	target  web3.Address
	enabled bool
}

func parseCtx(ctx *cli.Context) (*param, error) {
	path := ctx.String(flags.ConfigFlag.Name)
	signer, conf, err := common.SetUpL1(path)
	if err != nil {
		return nil, err
	}
	acc := ctx.String(flags.AccountFlag.Name)
	enabled := ctx.Bool(flags.EnabledFlag.Name)
	submit := ctx.Bool(flags.SubmitFlag.Name)
	signer.Submit = submit
	return &param{signer, conf.L1Addresses.DAO, web3.HexToAddress(acc), enabled}, nil
}

func setSequencerWhitelist(params *param) {
	dao := params.dao
	signer := params.signer
	sequencer := params.target
	enabled := params.enabled
	c := binding.NewDAO(dao, signer.Client)
	c.Contract().SetFrom(signer.Address())
	receipt := c.SetSequencerWhitelist(sequencer, enabled).Sign(signer).SendTransaction(signer).EnsureNoRevert()
	log.Info("set sequencer whitelist", "dao", dao, "sequencer", sequencer, "enabled", enabled, "receipt", utils.JsonStr(receipt))
}

func setProposerWhitelist(params *param) {
	dao := params.dao
	signer := params.signer
	proposer := params.target
	enabled := params.enabled
	c := binding.NewDAO(dao, signer.Client)
	c.Contract().SetFrom(signer.Address())
	receipt := c.SetProposerWhitelist(proposer, enabled).Sign(signer).SendTransaction(signer).EnsureNoRevert()
	log.Info("set proposer whitelist", "dao", dao, "proposer", proposer, "enabled", enabled, "receipt", utils.JsonStr(receipt))
}
