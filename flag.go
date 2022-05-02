package cli

import "time"

type Flag interface {
	GetName() string
	GetDefaultValue() interface{}
	GetUsage() string
	GetAliases() []string
}

type StringFlag struct {
	Name         string
	Desc         string
	Aliases      []string
	DefaultValue string
}

func (s *StringFlag) GetName() string {
	return s.Name
}

func (s *StringFlag) GetDefaultValue() interface{} {
	return s.DefaultValue
}

func (s *StringFlag) GetUsage() string {
	return s.Desc
}

func (s *StringFlag) GetAliases() []string {
	return s.Aliases
}

type BoolFlag struct {
	Name         string
	Desc         string
	Aliases      []string
	DefaultValue bool
}

func (s *BoolFlag) GetName() string {
	return s.Name
}

func (s *BoolFlag) GetDefaultValue() interface{} {
	return s.DefaultValue
}

func (s *BoolFlag) GetUsage() string {
	return s.Desc
}

func (s *BoolFlag) GetAliases() []string {
	return s.Aliases
}

type Float64Flag struct {
	Name         string
	Desc         string
	Aliases      []string
	DefaultValue float64
}

func (s *Float64Flag) GetName() string {
	return s.Name
}

func (s *Float64Flag) GetDefaultValue() interface{} {
	return s.DefaultValue
}

func (s *Float64Flag) GetUsage() string {
	return s.Desc
}

func (s *Float64Flag) GetAliases() []string {
	return s.Aliases
}

type Int64ValueFlag struct {
	Name         string
	Desc         string
	Aliases      []string
	DefaultValue int64
}

func (s *Int64ValueFlag) GetName() string {
	return s.Name
}

func (s *Int64ValueFlag) GetDefaultValue() interface{} {
	return s.DefaultValue
}

func (s *Int64ValueFlag) GetUsage() string {
	return s.Desc
}

func (s *Int64ValueFlag) GetAliases() []string {
	return s.Aliases
}

type Uint64ValueFlag struct {
	Name         string
	Desc         string
	Aliases      []string
	DefaultValue uint64
}

func (s *Uint64ValueFlag) GetName() string {
	return s.Name
}

func (s *Uint64ValueFlag) GetDefaultValue() interface{} {
	return s.DefaultValue
}

func (s *Uint64ValueFlag) GetUsage() string {
	return s.Desc
}

func (s *Uint64ValueFlag) GetAliases() []string {
	return s.Aliases
}

type DurationValueFlag struct {
	Name         string
	Desc         string
	Aliases      []string
	DefaultValue time.Duration
}

func (s *DurationValueFlag) GetName() string {
	return s.Name
}

func (s *DurationValueFlag) GetDefaultValue() interface{} {
	return s.DefaultValue
}

func (s *DurationValueFlag) GetUsage() string {
	return s.Desc
}

func (s *DurationValueFlag) GetAliases() []string {
	return s.Aliases
}
