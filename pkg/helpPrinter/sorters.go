package helpPrinter

import (
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"sort"
)

func getSortedCommands(commands map[argParserConfig.Command]bool) (res []string) {
	if len(commands) == 0 {
		return nil
	}
	res = make([]string, 0, len(commands))
	for command := range commands {
		res = append(res, string(command))
	}
	sort.Strings(res)

	return res
}

func getSortedFlags(groupFlagNameMap map[argParserConfig.Flag]bool) (res []string) {
	if len(groupFlagNameMap) == 0 {
		return nil
	}
	res = make([]string, 0, len(groupFlagNameMap))
	for flag := range groupFlagNameMap {
		res = append(res, string(flag))
	}
	sort.Strings(res)

	return res
}

func getSortedStrings(strings map[string]bool) (res []string) {
	if len(strings) == 0 {
		return nil
	}
	res = make([]string, 0, len(strings))
	for s := range strings {
		res = append(res, s)
	}
	sort.Strings(res)

	return res
}

func getSortedFlagsForDescription(flagDescriptions map[argParserConfig.Flag]*argParserConfig.FlagDescription) (res []string) {
	res = make([]string, 0, len(flagDescriptions))
	for flag := range flagDescriptions {
		res = append(res, string(flag))
	}
	sort.Strings(res)

	return res
}
