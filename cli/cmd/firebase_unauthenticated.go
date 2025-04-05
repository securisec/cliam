package cmd

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"sync"

	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var firebaseUnauthenticatedeCmd = &cobra.Command{
	Use:               "unauthenticated [--known-value database=<database>,collection=<collection>,document=<document>,url=<url>...]",
	Example:           "cliam firebase unauthenticated",
	Short:             "Enumerate unauthenticated Firebase permissions",
	Run:               firebaseUnauthenticatedeCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	firebaseCmd.AddCommand(firebaseUnauthenticatedeCmd)
}

func firebaseUnauthenticatedeCmdFunc(_ *cobra.Command, _ []string) {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// the is the maximum concurrent goroutines
	max := make(chan struct{}, MaxThreads)

	fs := []func() (int, error){
		firebaseRTDB,
		firebaseFirestore,
		firebaseStorage,
		firebaseHostingInit,
	}

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

func firebaseHostingInit() (int, error) {
	if _, ok := firebaseKnownValues["url"]; !ok {
		return 0, nil
	}
	u, err := url.Parse(firebaseKnownValues["url"])
	if err != nil {
		logger.Logger.Error().Err(errors.New("Hosting: URL not valid")).Send()
		return 0, nil
	}
	u.Path = path.Join(u.Path, "/__/firebase/init.json")
	s := u.String()
	return getRequest(s, "HostingInit")
}

func firebaseRTDB() (int, error) {
	if firebaseProjectId == "" {
		logger.Logger.Warn().Err(errors.New("RTDB: Project ID not provided")).Send()
		return 0, nil
	}
	// check if db is specified
	if _, ok := firebaseKnownValues["database"]; !ok {
		firebaseKnownValues["database"] = fmt.Sprintf("%s-default-rtdb", firebaseProjectId)
	}

	t := "https://{{.database}}.firebaseio.com/.json"
	url, err := templateBuilder(t, firebaseKnownValues)
	if err != nil {
		return 0, err
	}

	return getRequest(url, "RTDB")
}

func firebaseFirestore() (int, error) {
	if firebaseProjectId == "" {
		logger.Logger.Warn().Err(errors.New("Firestore: Project ID not found")).Send()
		return 0, nil
	}
	// check if collection is specified
	if _, ok := firebaseKnownValues["collection"]; !ok {
		firebaseKnownValues["collection"] = "default"
	}
	if _, ok := firebaseKnownValues["projectID"]; !ok {
		firebaseKnownValues["projectID"] = firebaseProjectId
	}
	if _, ok := firebaseKnownValues["document"]; !ok {
		firebaseKnownValues["document"] = ""
	}

	t := "https://firestore.googleapis.com/v1/projects/{{.projectID}}/databases/(default)/documents/{{.collection}}/{{.document}}"
	url, err := templateBuilder(t, firebaseKnownValues)
	if err != nil {
		return 0, err
	}

	return getRequest(url, "Firestore")
}

func firebaseStorage() (int, error) {
	if _, ok := firebaseKnownValues["path"]; !ok {
		firebaseKnownValues["path"] = ""
	}
	t := "https://firebasestorage.googleapis.com/v0/b/{{.projectID}}.appspot.com/o{{.path}}"
	url, err := templateBuilder(t, firebaseKnownValues)
	if err != nil {
		return 0, err
	}
	return getRequest(url, "Storage")
}
