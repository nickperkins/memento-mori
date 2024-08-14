package flags

import (
	"flag"
	"fmt"
	"os"
	"syscall"
)

type Options struct {
	DryRun  bool
	RunOnce bool
	Help    bool
	Version bool
}

var Flags Options

var Version string

func init() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage of %s: a Mastodon bot that posts random quotes about death.\n", os.Args[0])

		flag.PrintDefaults()

		//fmt.Fprintf(w, "...custom postamble ... \n")

	}
}

func SetupFlags(version string) {
	Version = version
	const (
		defaultDryRun  = false
		dryRunUsage    = "Run the bot in dry-run mode"
		defaultRunOnce = false
		runOnceUsage   = "Run the bot once and exit"
	)
	flag.BoolVar(&Flags.DryRun, "dry-run", defaultDryRun, dryRunUsage)
	flag.BoolVar(&Flags.DryRun, "d", defaultDryRun, dryRunUsage+" (shorthand)")
	flag.BoolVar(&Flags.RunOnce, "run-once", defaultRunOnce, runOnceUsage)
	flag.BoolVar(&Flags.RunOnce, "r", defaultRunOnce, runOnceUsage+" (shorthand)")
	flag.BoolVar(&Flags.Help, "help", false, "Show help")
	flag.BoolVar(&Flags.Version, "version", false, "Show version")
}

func ParseFlags() {
	flag.Parse()
	if Flags.Help {
		flag.Usage()
		syscall.Exit(0)
	}

	if Flags.Version {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Version: %s\n", Version)
		syscall.Exit(0)
	}
}
