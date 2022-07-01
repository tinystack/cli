package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

func NewApp() *App {
	app := &App{
		output: os.Stderr,
	}
	app.cmd = emptyCommand(app)
	return app
}

type App struct {
	cmd     *command
	mu      sync.Mutex
	running bool
	output  io.Writer
	ver     string
}

func (app *App) Command(cmd *command) {
	cmd.fillHelpFlag()
	cmd.fillVersionFlag()
	app.cmd = cmd
}

func (app *App) SetOutput(w io.Writer) {
	app.output = w
}

func (app *App) SetVersion(ver string) {
	app.ver = ver
}

func (app *App) Run() error {
	app.mu.Lock()
	if app.running {
		app.mu.Unlock()
		return errors.New("cli: command is already running")
	}
	app.running = true
	app.mu.Unlock()

	return app.run(os.Args[1:])
}

func (app *App) run(args []string) (err error) {
	var (
		cmd         = app.cmd
		cmdArgs     = args
		flagSet     *flag.FlagSet
		rootFlagSet *flag.FlagSet
		cmdCrumbs   []string
	)
	if app.cmd == nil {
		err = errors.New("cli: command is nil")
		return
	}
	for {
		flagSet, err = app.parseFlag(cmd, cmdArgs)
		if err != nil {
			return
		}
		if rootFlagSet == nil {
			rootFlagSet = flagSet
			cmdCrumbs = append(cmdCrumbs, cmd.name)
		}

		if flagSet.NArg() == 0 {
			break
		}

		findSubCommand := false
		subCommandName := flagSet.Arg(0)
		cmdCrumbs = append(cmdCrumbs, subCommandName)
		for _, v := range cmd.subCommand {
			if v.name == subCommandName {
				cmd = v
				cmdArgs = flagSet.Args()[1:]
				findSubCommand = true
			}
		}
		if !findSubCommand {
			cmdCrumbsSlice := fillCmdCrumbs(cmdCrumbs)
			err = fmt.Errorf(`cli: command not defined: %s`, strings.Join(cmdCrumbsSlice, " "))
			return
		}
	}
	ctx := &Context{
		rootFlagSet: rootFlagSet,
		flagSet:     flagSet,
		cmd:         cmd,
		cmdCrumbs:   cmdCrumbs,
	}
	err = app.runHandler(ctx)
	return
}

func (app *App) runHandler(ctx *Context) error {
	if ctx.Bool("help") {
		app.printHelp(ctx.cmd, ctx.cmdCrumbs)
		return nil
	}
	if ctx.Bool("version") {
		app.printVersion()
		return nil
	}
	if ctx.cmd.handler == nil {
		app.printHelp(ctx.cmd, ctx.cmdCrumbs)
		return nil
	}
	return ctx.cmd.handler(ctx)
}

func (app *App) printHelp(cmd *command, cmdCrumbs []string) {
	cmd.printHelp(cmdCrumbs)
}

func (app *App) printVersion() {
	_, _ = fmt.Fprint(app.output, fmt.Sprintf("version: %s\n", app.ver))
}

func (app *App) parseFlag(cmd *command, args []string) (*flag.FlagSet, error) {
	flagSet := flag.NewFlagSet(cmd.name, flag.ContinueOnError)
	flagSet.SetOutput(io.Discard)
	flagSet.Usage = nil
	for _, v := range cmd.flags {
		app.flagSetVar(flagSet, v, v.GetName(), v.GetDefaultValue(), v.GetUsage())
		for _, name := range v.GetAliases() {
			app.flagSetVar(flagSet, v, name, v.GetDefaultValue(), v.GetUsage())
		}
	}
	if err := flagSet.Parse(args); err != nil {
		return nil, fmt.Errorf("cli: flagSet.Parse err: %s", err.Error())
	}
	return flagSet, nil
}

func (app *App) flagSetVar(flagSet *flag.FlagSet, v Flag, name string, defaultValue interface{}, usage string) {
	switch v.(type) {
	case *StringFlag:
		flagSet.String(name, defaultValue.(string), usage)
	case *BoolFlag:
		flagSet.Bool(name, defaultValue.(bool), usage)
	case *Float64Flag:
		flagSet.Float64(name, defaultValue.(float64), usage)
	case *Int64ValueFlag:
		flagSet.Int64(name, defaultValue.(int64), usage)
	case *Uint64ValueFlag:
		flagSet.Uint64(name, defaultValue.(uint64), usage)
	case *DurationValueFlag:
		flagSet.Duration(name, defaultValue.(time.Duration), usage)
	}
}
