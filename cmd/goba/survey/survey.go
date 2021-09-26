package survey

import (
	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
	"strings"
)

// RequireMixedFlagsFunc returns a cli.BeforeFunc that
// combines RequireGlobalFlagsFunc and RequireFlagsFunc.
func RequireMixedFlagsFunc(global []string, local ...string) cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		if err := RequireGlobalFlagsFunc(global...)(ctx); err != nil {
			return err
		}
		return RequireFlagsFunc(local...)(ctx)
	}
}

// RequireFlagsFunc returns a cli.BeforeFunc that asks
// the users for flag values for the given flags that
// have not been set.
func RequireFlagsFunc(flags ...string) cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		if len(ctx.Args().Tail()) == 0 {
			return requireFlags(ctx, flags...)
		}
		return nil
	}
}

// RequireGlobalFlagsFunc is like RequireFlagsFunc, but manages global flags.
func RequireGlobalFlagsFunc(flags ...string) cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		if len(ctx.Args().Tail()) == 0 {
			return requireGlobalFlags(ctx, flags...)
		}
		return nil
	}
}

// requireFlags asks the user for flag values
// for the given flags that have not been set.
func requireFlags(ctx *cli.Context, flags ...string) error {
	for _, flag := range flags {
		if !ctx.IsSet(flag) {
			value, err := ask(flag)
			if err != nil {
				return err
			}

			if err = ctx.Set(flag, value); err != nil {
				return err
			}
		}
	}
	return nil
}

// requireGlobalFlags is like requireFlags, but manages global flags.
func requireGlobalFlags(ctx *cli.Context, flags ...string) error {
	for _, flag := range flags {
		if !ctx.GlobalIsSet(flag) {
			value, err := ask(flag)
			if err != nil {
				return err
			}

			if err = ctx.GlobalSet(flag, value); err != nil {
				return err
			}
		}
	}
	return nil
}

// ask prompts the client with the given message
// and asks them to provide a corresponding value.
func ask(msg string) (string, error) {
	var resp string
	return resp, survey.AskOne(prompt(msg), &resp, nil)
}

// prompt prepares a survey.Prompt for the given message.
func prompt(msg string) survey.Prompt {
	msg = formatPromptMsg(msg)
	if strings.Contains(msg, "password") {
		return &survey.Password{Message: msg}
	}
	return &survey.Input{Message: msg}
}

// formatPromptMsg formats msg.
func formatPromptMsg(msg string) string {
	return strings.Title(msg) + ": "
}
