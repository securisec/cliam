package cmd

import (
	"context"
	"errors"
	"fmt"

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
	firebaseAdminCmd.Flags().Bool("firestore-documents", false, "Enumerate all documents")
	firebaseAdminCmd.Flags().BoolP("firestore-dump-data", "d", false, "Dump data")
	firebaseAdminCmd.Flags().String("rtdb-database", "default", "Name of RTDB database")
}

func firebaseAdminCmdFunc(cmd *cobra.Command, _ []string) {
	rtdbDatabaseName, _ := cmd.Flags().GetString("rtdb-database")
	firebaseServiceAccount = expandPath(firebaseServiceAccount)
	if firebaseServiceAccount == "" {
		logger.LogPanic(errors.New("Service account not found"))
	}

	gcpSA, err := gcpReadServiceAccount(firebaseServiceAccount)
	if err != nil {
		logger.LogError(err)
	}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, &firebase.Config{
		DatabaseURL: fmt.Sprintf("https://%s-%s-rtdb.firebaseio.com/", gcpSA.ProjectID, rtdbDatabaseName),
	}, option.WithCredentialsFile(firebaseServiceAccount))
	if err != nil {
		logger.LogError(err)
	}

	// enumerate firebase firestore
	firebaseAdminFirestore(ctx, cmd, app)
	// enumerate firebase rtdb
	firebaseAdminRTDB(ctx, app)
	// enumerate firebase functions
	// TODO 🔥

}

func firebaseAdminFirestore(ctx context.Context, cmd *cobra.Command, app *firebase.App) {
	dumpData, _ := cmd.Flags().GetBool("firestore-dump-data")
	enumDocs, _ := cmd.Flags().GetBool("firestore-documents")

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
		firebaseSuccessLogging("firestore", map[string]any{
			"collection": coll.ID,
		}, shared.GetMessageColor("success"))
		if enumDocs {
			colRef := client.Collection(coll.ID)
			docs, err := colRef.Documents(ctx).GetAll()
			if err != nil {
				logger.LogError(err)
			}
			for _, doc := range docs {
				firebaseSuccessLogging("firestore", map[string]any{
					"document":   doc.Ref.ID,
					"collection": coll.ID,
				}, shared.GetMessageColor("success"))
				if dumpData {
					snap, err := doc.Ref.Get(ctx)
					data := snap.Data()
					if err != nil {
						logger.LogError(err)
					}
					data["collection"] = coll.ID
					data["document"] = doc.Ref.ID
					firebaseSuccessLogging("firestore", data, shared.GetMessageColor("info"))
				}
			}
		}
	}
}

func firebaseAdminRTDB(ctx context.Context, app *firebase.App) {
	client, err := app.Database(ctx)
	if err != nil {
		logger.LogError(err)
	}
	ref := client.NewRef("/")

	var hold map[string]any
	if err := ref.Get(ctx, &hold); err != nil {
		logger.LogError(err)
	}

	for k, v := range hold {
		firebaseSuccessLogging("rtdb", map[string]any{
			k: v,
		}, shared.GetMessageColor("success"))
	}
}

func firebaseSuccessLogging(resource string, obj map[string]any, color string) {
	l := logger.Logger.Info().Str("_resource", resource)
	for k, v := range obj {
		l.Interface(k, v)
	}
	l.Msg(color)
}
