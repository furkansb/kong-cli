package subcommands

import "github.com/urfave/cli/v2"

func mergeCommands(commands ...[]*cli.Command) []*cli.Command {
	var result []*cli.Command
	for _, f := range commands {
		result = append(result, f...)
	}
	return result
}

func AllCommands() []*cli.Command {
	return mergeCommands(ServiceCommands, RouteCommands, ConsumerCommands, Oauth2Commands)
}
