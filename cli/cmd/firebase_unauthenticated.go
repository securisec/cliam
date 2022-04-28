package cmd

import (
	"fmt"
	"sync"

	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var firebaseUnauthenticatedeCmd = &cobra.Command{
	Use:               "unauthenticated",
	Short:             "Enumerate unauthenticated Firebase permissions",
	Run:               firebaseUnauthenticatedeCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	firebaseCmd.AddCommand(firebaseUnauthenticatedeCmd)
}

func firebaseUnauthenticatedeCmdFunc(cmd *cobra.Command, _ []string) {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// the is the maximum concurrent goroutines
	max := make(chan struct{}, MaxThreads)

	fs := []func() (int, error){firebaseRTDB, firebaseFirestore}

	go func() {
		defer wg.Done()
		for _, f := range fs {

			wg.Add(1)

			go func(wg *sync.WaitGroup, function func() (int, error)) {

				// block the channel until the maximum number of goroutines is reached
				max <- struct{}{}
				defer func() {
					<-max
				}()

				status, err := function()
				if err != nil {
					logger.LogDebug("Bad status", status)
				}

				// done with this goroutine
				wg.Done()
			}(wg, f)
		}

	}()

	wg.Wait()

}

func firebaseRTDB() (int, error) {
	// check if db is specified
	if _, ok := firebaseExtras["database"]; !ok {
		firebaseExtras["database"] = fmt.Sprintf("%s-default-rtdb", firebaseProjectId)
	}

	t := "https://{{.database}}.firebaseio.com/.json"
	url, err := templateBuilder(t, firebaseExtras)
	if err != nil {
		return 0, err
	}

	return getRequest(url, "RTDB")
}

func firebaseFirestore() (int, error) {
	// check if collection is specified
	if _, ok := firebaseExtras["collection"]; !ok {
		firebaseExtras["collection"] = "default"
	}
	if _, ok := firebaseExtras["projectID"]; !ok {
		firebaseExtras["projectID"] = firebaseProjectId
	}

	t := "https://firestore.googleapis.com/v1/projects/{{.projectID}}/databases/(default)/documents/{{.collection}}/{{.key}}"
	url, err := templateBuilder(t, firebaseExtras)
	if err != nil {
		return 0, err
	}

	return getRequest(url, "Firestore")
}
