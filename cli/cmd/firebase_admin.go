package cmd

import (
	"context"
	"errors"

	firebase "firebase.google.com/go/v4"
	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var firebaseAdminCmd = &cobra.Command{
	Use:               "admin",
	Example:           "cliam firebase admin",
	Short:             "Enumerate Firebase resources via admin credentials",
	Run:               firebaseAdminCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	firebaseCmd.AddCommand(firebaseAdminCmd)
	firebaseAdminCmd.Flags().BoolP("dump-data", "d", false, "Dump data")
}

func firebaseAdminCmdFunc(cmd *cobra.Command, _ []string) {
	firebaseServiceAccount = expandPath(firebaseServiceAccount)
	if firebaseServiceAccount == "" {
		logger.LogPanic(errors.New("Service account not found"))
	}

	dumpData, _ := cmd.Flags().GetBool("dump-data")

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(firebaseServiceAccount))
	if err != nil {
		logger.LogError(err)
	}

	client, err := app.Firestore(ctx)
	defer client.Close()
	if err != nil {
		logger.LogError(err)
	}

	collIter := client.Collections(ctx)
	collections, err := collIter.GetAll()
	if err != nil {
		logger.LogError(err)
	}
	for _, coll := range collections {
		logger.LogSuccess("collection", coll.ID)
		colRef := client.Collection(coll.ID)
		docs, err := colRef.Documents(ctx).GetAll()
		if err != nil {
			logger.LogError(err)
		}
		for _, doc := range docs {
			logger.Logger.Info().Str("document", doc.Ref.ID).Str("collection", coll.ID).Msg(shared.GetMessageColor("success"))
			if dumpData {
				snap, err := doc.Ref.Get(ctx)
				data := snap.Data()
				if err != nil {
					logger.LogError(err)
				}
				data["collection"] = coll.ID
				data["object_id"] = doc.Ref.ID
				logger.Logger.Info().Interface("dump", data).Msg(shared.GetMessageColor("info"))
			}
		}
	}
}
