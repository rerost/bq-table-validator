package validate

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/rerost/bq-table-validator/domain/validator"
	"github.com/rerost/bq-table-validator/types"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

func NewCmd(ctx context.Context, validate validator.Validator) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "validate",
		Args: cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			validateFilePath := args[0]
			validateFile, err := ioutil.ReadFile(validateFilePath)
			if err != nil {
				return errors.WithStack(err)
			}
			var validates []types.Validate
			if err := yaml.Unmarshal(validateFile, &validates); err != nil {
				return errors.WithStack(err)
			}

			for _, v := range validates {
				out, err := validate.Valid(ctx, v)
				if err != nil {
					return errors.WithStack(err)
				}

				// TODO format & color
				fmt.Printf(out)
			}
			return nil
		},
	}

	return cmd
}
