package subcommands

import "github.com/urfave/cli/v2"

func mergeCommands(commands ...[]*cli.Command) []*cli.Command {
	var result []*cli.Command
	for _, f := range commands {
		result = append(result, f...)
	}
	return result
}

func getSliceElementsPointer(sl []string) []*string {
	result := make([]*string, len(sl))
	for i, str := range sl {
		strLocal := str
		result[i] = &strLocal
	}
	return result
}

func AllCommands() []*cli.Command {
	return mergeCommands(ServiceCommands, RouteCommands, ConsumerCommands, Oauth2Commands)
}
