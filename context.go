package cli

import (
	"flag"
	"time"
)

type Context struct {
	rootFlagSet *flag.FlagSet
	flagSet     *flag.FlagSet
	cmd         *command
	cmdCrumbs   []string
}

func (c *Context) PrintHelp() {
	c.cmd.printHelp(c.cmdCrumbs)
}

func (c *Context) String(name string) (value string) {
	value, _ = c.valueAliases(name).(string)
	return
}

func (c *Context) Bool(name string) (value bool) {
	value, _ = c.valueAliases(name).(bool)
	return
}

func (c *Context) Float64(name string) (value float64) {
	value, _ = c.valueAliases(name).(float64)
	return
}

func (c *Context) Int64(name string) (value int64) {
	value, _ = c.valueAliases(name).(int64)
	return
}

func (c *Context) Int(name string) int {
	value, _ := c.valueAliases(name).(int64)
	return int(value)
}

func (c *Context) Uint64(name string) (value uint64) {
	value, _ = c.valueAliases(name).(uint64)
	return
}

func (c *Context) Uint(name string) uint {
	value, _ := c.valueAliases(name).(uint64)
	return uint(value)
}

func (c *Context) Duration(name string) (value time.Duration) {
	value, _ = c.valueAliases(name).(time.Duration)
	return
}

func (c *Context) valueAliases(name string) interface{} {
	names := []string{name}
	for _, f := range c.cmd.flags {
		if name == f.GetName() {
			names = append(names, f.GetAliases()...)
		}
	}
	if val := c.value(names); val != nil {
		return val
	}
	return ""
}

func (c *Context) value(names []string) interface{} {
	var flagSets = []*flag.FlagSet{c.flagSet, c.rootFlagSet}
	for _, fs := range flagSets {
		if result, find := c.findValueFromActual(fs, names); find {
			return result
		}
	}
	for _, fs := range flagSets {
		if result, find := c.findValueFromFormal(fs, names); find {
			return result
		}
	}
	return nil
}

func (c *Context) findValueFromActual(flagSet *flag.FlagSet, names []string) (result interface{}, hasFind bool) {
	if flagSet == nil {
		return
	}
	flagSet.Visit(func(f *flag.Flag) {
		if stringIndex(f.Name, names) > -1 {
			if getter, ok := f.Value.(flag.Getter); ok {
				result = getter.Get()
				hasFind = true
			}
		}
	})
	return
}

func (c *Context) findValueFromFormal(flagSet *flag.FlagSet, names []string) (result interface{}, hasFind bool) {
	if flagSet == nil {
		return
	}
	flagSet.VisitAll(func(f *flag.Flag) {
		if stringIndex(f.Name, names) > -1 {
			if getter, ok := f.Value.(flag.Getter); ok {
				result = getter.Get()
				hasFind = true
			}
		}
	})
	return
}
