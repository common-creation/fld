package commands

import (
	"github.com/common-creation/fld/constants"
	"github.com/common-creation/fld/i18n"
	"github.com/common-creation/fld/utils"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type (
	payload struct {
		target string
		err    error
		code   int
	}
)

func NewDeployCmd() *cobra.Command {
	var (
		all          bool
		adminWeb     bool
		platform     bool
		scenario     bool
		bosai        bool
		bi           bool
		survey       bool
		liff         bool
		distribution bool
	)

	var (
		useContainer bool
		sequential   bool
	)

	cmd := &cobra.Command{
		Use:   "deploy",
		Short: i18n.T("deploy.help.short", nil),
		Run: func(cmd *cobra.Command, args []string) {
			basePath, _ := filepath.Abs("./")
			// basePath, _ := filepath.Abs("../../linefukuoka/line-smart-city-develop/aws_back")
			lscPath, ok := findLsc(basePath)
			if !ok {
				panic("invalid pwd")
			}

			targets := make([]string, 0)
			if adminWeb || all {
				targets = append(targets, "admin-web")
			}
			if platform || all {
				targets = append(targets, "platform")
			}
			if scenario || all {
				targets = append(targets, "scenario")
			}
			if bosai || all {
				targets = append(targets, "bosai")
			}
			if bi || all {
				targets = append(targets, "bi")
			}
			if survey || all {
				targets = append(targets, "survey")
			}
			if liff || all {
				targets = append(targets, "liff")
			}
			if distribution || all {
				targets = append(targets, "distribution")
			}

			br, dr := deploy(lscPath, targets, useContainer, sequential)
			utils.Infoln()
			utils.Infoln(i18n.T("deploy.summary.build", nil))
			printResult(br)
			utils.Infoln()
			utils.Infoln(i18n.T("deploy.summary.deploy", nil))
			printResult(dr)
		},
	}

	cmd.PersistentFlags().BoolVarP(&all, "all", "", false, i18n.T("deploy.all.usage", nil))

	cmd.PersistentFlags().BoolVarP(&adminWeb, "admin-web", "", false, i18n.T("deploy.admin-web.usage", nil))
	cmd.PersistentFlags().BoolVarP(&platform, "platform", "", false, i18n.T("deploy.platform.usage", nil))
	cmd.PersistentFlags().BoolVarP(&scenario, "scenario", "", false, i18n.T("deploy.scenario.usage", nil))
	cmd.PersistentFlags().BoolVarP(&bosai, "bosai", "", false, i18n.T("deploy.bosai.usage", nil))
	cmd.PersistentFlags().BoolVarP(&bi, "bi", "", false, i18n.T("deploy.bi.usage", nil))
	cmd.PersistentFlags().BoolVarP(&survey, "survey", "", false, i18n.T("deploy.survey.usage", nil))
	cmd.PersistentFlags().BoolVarP(&liff, "liff", "", false, i18n.T("deploy.liff.usage", nil))
	cmd.PersistentFlags().BoolVarP(&distribution, "distribution", "", false, i18n.T("deploy.distribution.usage", nil))

	cmd.PersistentFlags().BoolVarP(&useContainer, "useContainer", "", false, i18n.T("deploy.useContainer.usage", nil))
	cmd.PersistentFlags().BoolVarP(&sequential, "sequential", "", false, i18n.T("deploy.sequential.usage", nil))

	return cmd
}

func printResult(results []payload) {
	for _, r := range results {
		switch {
		case r.err != nil:
			utils.Infoln(r.target, ":", i18n.T("deploy.common.execerror", nil))
		case r.code != 0:
			utils.Infoln(r.target, ":", i18n.T("deploy.common.codeerror", nil))
		default:
			utils.Infoln(r.target, ":", i18n.T("deploy.common.ok", nil))
		}
	}
}

func findLsc(basePath string) (string, bool) {
	targetPath := path.Join(basePath, "lsc.sh")
	_, err := os.Stat(targetPath)
	if err != nil {
		list := strings.Split(basePath, string(os.PathSeparator))
		list = list[:len(list)-1]
		if len(list) == 0 {
			return "", false
		}
		parentPath := filepath.Join(list...)
		if !strings.HasPrefix(parentPath, string(os.PathSeparator)) {
			parentPath = string(os.PathSeparator) + parentPath
		}
		return findLsc(parentPath)
	}

	return targetPath, true
}

func deploy(lscPath string, targets []string, useContainer bool, sequential bool) ([]payload, []payload) {
	checkEnv(useContainer)

	built := make(chan payload, len(targets))
	deployed := make(chan payload, len(targets))

	buildResult := make([]payload, len(targets))
	deployResult := make([]payload, len(targets))

	go func() {
		for i := 0; i < len(targets); i++ {
			p := <-built
			buildResult[i] = p

			if p.code != 0 || p.err != nil {
				utils.Errorln(i18n.T("deploy.common.builderror", nil), ":", p.target)
				continue
			}
			go lscDeploy(p.target, lscPath, &deployed)
		}
	}()

	for _, target := range targets {
		if sequential {
			lscBuild(target, lscPath, useContainer, &built)
		} else {
			go lscBuild(target, lscPath, useContainer, &built)
		}
	}

	for i := 0; i < len(targets); i++ {
		p := <-deployed
		deployResult[i] = p

		if p.code != 0 || p.err != nil {
			utils.Errorln(i18n.T("deploy.common.deployerror", nil), ":", p.target)
			continue
		}
	}

	return buildResult, deployResult
}

func checkEnv(useContainer bool) {
	if useContainer {
		ok := utils.HasCommands([]utils.Command{
			{Cmd: "docker", Args: []string{"--version"}},
			{Cmd: "docker-compose", Args: []string{"--version"}},
		})
		if !ok {
			os.Exit(1)
			return
		}
	} else {
		ok := utils.HasCommands([]utils.Command{
			{Cmd: "python", Args: []string{"--version"}},
			{Cmd: "pip", Args: []string{"--version"}},
		})
		if !ok {
			os.Exit(1)
			return
		}
	}
	ok := utils.HasCommands([]utils.Command{
		{Cmd: "node", Args: []string{"--version"}},
		{Cmd: "npm", Args: []string{"--version"}},
	})
	if !ok {
		os.Exit(1)
		return
	}
}

func lscBuild(target string, lscPath string, useContainer bool, ch *chan payload) {
	cmdInfo := constants.LSC_COMMANDS[target]
	cmd := utils.NewCommand(target, lscPath, cmdInfo.Build)
	if useContainer && cmdInfo.UseContainer != nil {
		cmd.ExtendArgs(*cmdInfo.UseContainer)
	}
	code, err := cmd.ExecAsync()
	*ch <- payload{
		target: target,
		code:   code,
		err:    err,
	}
}

func lscDeploy(target string, lscPath string, ch *chan payload) {
	cmdInfo := constants.LSC_COMMANDS[target]
	cmd := utils.NewCommand(target, lscPath, cmdInfo.Deploy)
	code, err := cmd.ExecAsync()
	*ch <- payload{
		target: target,
		code:   code,
		err:    err,
	}
}
