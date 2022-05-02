package cli

import (
	"fmt"
	"strings"
)

type Handler func(*Context) error

type command struct {
	app        *App
	subCommand []*command
	name       string
	desc       string
	handler    Handler
	flags      []Flag
}

func NewCommand(app *App, name, desc string, flags []Flag, handler Handler) *command {
	return &command{
		app:     app,
		name:    name,
		desc:    desc,
		handler: handler,
		flags:   flags,
	}
}

func emptyCommand(app *App) *command {
	return &command{
		app: app,
	}
}

func (cmd *command) SubCommand(sub *command) {
	sub.fillHelpFlag()
	cmd.subCommand = append(cmd.subCommand, sub)
}

func (cmd *command) fillHelpFlag() {
	cmd.flags = append(cmd.flags, &BoolFlag{
		Name:         "help",
		Desc:         "Display the help information",
		Aliases:      []string{"h"},
		DefaultValue: false,
	})
}

func (cmd *command) fillVersionFlag() {
	cmd.flags = append(cmd.flags, &BoolFlag{
		Name:         "version",
		Desc:         "Print version information and quit",
		Aliases:      []string{"v"},
		DefaultValue: false,
	})
}

func (cmd *command) isRootCommand() bool {
	return cmd.app.cmd == cmd
}

func (cmd *command) printHelp(cmdCrumbs []string) {
	_, _ = fmt.Fprint(cmd.app.output, cmd.formatHelpHeader(cmdCrumbs))
	_, _ = fmt.Fprint(cmd.app.output, cmd.app.cmd.formatHelpOptions())
	if !cmd.isRootCommand() {
		_, _ = fmt.Fprint(cmd.app.output, cmd.formatHelpOptions())
	}
	_, _ = fmt.Fprint(cmd.app.output, cmd.formatHelpAvailableCmd(cmdCrumbs))
}

func (cmd *command) formatHelpAvailableCmd(cmdCrumbs []string) string {
	cmdNum := len(cmd.subCommand)
	if cmdNum == 0 {
		return ""
	}
	var (
		cmdString  = make([]string, 0, cmdNum)
		descString = make([]string, 0, cmdNum)
	)
	for _, v := range cmd.subCommand {
		cmdString = append(cmdString, v.name)
		descString = append(descString, v.desc)
	}
	outputs := []string{
		"Available Commands:\n",
	}
	cmdMaxLen := cmd.flagMaxLen(cmdString) + 2
	for i := 0; i < cmdNum; i++ {
		format := fmt.Sprintf("%%%ds    %%s\n", cmdMaxLen)
		outputs = append(outputs, fmt.Sprintf(format, cmdString[i], descString[i]))
	}
	outputs = append(outputs, "\n")

	cmdCrumbsSlice := fillCmdCrumbs(cmdCrumbs)
	if len(cmdCrumbsSlice) > 0 {
		cmdCrumbsSlice = append(cmdCrumbsSlice, "<command> --help")
		outputs = append(outputs, fmt.Sprintf(`Use "%s" for more information about a given command.`, strings.Join(cmdCrumbsSlice, " ")), "\n")
	}

	return strings.Join(outputs, "")
}

func (cmd *command) formatHelpOptions() string {
	flagsNum := len(cmd.flags)
	if flagsNum == 0 {
		return ""
	}
	var (
		aliasesString = make([]string, 0, flagsNum)
		flagsString   = make([]string, 0, flagsNum)
		descString    = make([]string, 0, flagsNum)
	)
	for _, v := range cmd.flags {
		aliases := v.GetAliases()
		for ak, alias := range aliases {
			aliases[ak] = cmd.fillOptionFlag(alias)
		}
		var aliasesStr string
		if len(aliases) > 0 {
			aliasesStr = strings.Join(aliases, ", ") + ", "
		} else {
			aliasesStr = ""
		}
		aliasesString = append(aliasesString, aliasesStr)
		flagsString = append(flagsString, cmd.fillOptionFlag(v.GetName()))
		descString = append(descString, v.GetUsage())
	}
	aliasMaxLen := cmd.flagMaxLen(aliasesString) + 2
	flagMaxLen := cmd.flagMaxLen(flagsString) + 4

	var outputs []string
	if cmd.isRootCommand() {
		outputs = append(outputs, "Global Options:\n")
	} else {
		outputs = append(outputs, "Options:\n")
	}
	for i := 0; i < flagsNum; i++ {
		format := fmt.Sprintf("%%%ds%%-%ds%%s\n", aliasMaxLen, flagMaxLen)
		outputs = append(outputs, fmt.Sprintf(format, aliasesString[i], flagsString[i], descString[i]))
	}
	outputs = append(outputs, "\n")
	return strings.Join(outputs, "")
}

func (cmd *command) formatHelpHeader(cmdCrumbs []string) string {
	var (
		cmdPrints []string
		cmdLines  []string
	)
	for k, v := range cmdCrumbs {
		opts := "[--options ...]"
		if k == 0 {
			opts = "[global options]"
		}
		cmdPrints = append(cmdPrints, fmt.Sprintf("%s %s", v, opts))
	}
	cmdLines = append(cmdLines, fmt.Sprintf("\nUsage: %s\n\n", strings.Join(cmdPrints, " ")))
	if cmd.desc != "" {
		cmdLines = append(cmdLines, fmt.Sprintf("%s\n\n", cmd.desc))
	}
	return strings.Join(cmdLines, "")
}

func (cmd *command) fillOptionFlag(name string) string {
	switch {
	case len(name) > 1:
		return fmt.Sprintf("--%s", name)
	case len(name) == 1:
		return fmt.Sprintf("-%s", name)
	default:
		return ""
	}
}

func (cmd *command) flagMaxLen(names []string) int {
	l := 0
	for _, v := range names {
		if ll := len(v); ll > l {
			l = ll
		}
	}
	return l
}
