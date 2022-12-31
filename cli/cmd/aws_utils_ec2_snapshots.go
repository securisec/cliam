package cmd

// TODO 🔥 🚧
// import (
// 	"context"
// 	"encoding/xml"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
// 	"github.com/securisec/cliam/aws/signer"
// 	"github.com/securisec/cliam/logger"
// 	"github.com/securisec/cliam/shared"
// 	"github.com/spf13/cobra"
// )

// var awsUtilsEc2SnapshotCmd = &cobra.Command{
// 	Use:    "ec2-snapshots",
// 	Short:  "Enumerate AWS EC2 snapshots across specified regions",
// 	Run:    awsUtilsEc2SnapshotCmdFunc,
// 	PreRun: awsLoadEnvVarsFirst,
// }

// func init() {
// 	awsUtilsCmd.AddCommand(awsUtilsEc2SnapshotCmd)
// 	awsUtilsEc2SnapshotCmd.Flags().StringSlice("regions", []string{}, "AWS Regions")
// 	awsUtilsEc2SnapshotCmd.Flags().Bool("all", false, "Enumerate all regions")
// 	awsUtilsEc2SnapshotCmd.Flags().Int("max-results", 5, "Max results to get")
// 	awsUtilsEc2SnapshotCmd.Flags().String("description", "", "Description to search for. Can use the format *search query*")
// 	awsUtilsEc2SnapshotCmd.Flags().String("owner-id", "", "Owner ID to search for")

// 	awsUtilsEc2SnapshotCmd.RegisterFlagCompletionFunc("regions", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
// 		return aws_Regions, cobra.ShellCompDirectiveNoFileComp
// 	})
// }

// func awsUtilsEc2SnapshotCmdFunc(cmd *cobra.Command, args []string) {
// 	// parse flags
// 	regions, _ := cmd.Flags().GetStringSlice("regions")
// 	all, _ := cmd.Flags().GetBool("all")
// 	maxResults, _ := cmd.Flags().GetInt("max-results")
// 	description, _ := cmd.Flags().GetString("description")
// 	ownerID, _ := cmd.Flags().GetString("owner-id")

// 	// enumerate all regions
// 	if all {
// 		regions = aws_Regions
// 	}

// 	// sanity check flags
// 	if description == "" && ownerID == "" {
// 		logger.LoggerStdErr.Fatal().Msg("You must specify either a description or an owner ID")
// 	}
// 	if description != "" && ownerID != "" {
// 		logger.LoggerStdErr.Fatal().Msg("You must specify either a description or an owner ID")
// 	}

// 	if len(regions) == 0 {
// 		logger.LoggerStdErr.Fatal().Msg("You must specify at least one region")
// 	}

// 	// get credentials
// 	key, secret, token, _ := getCredsAndRegion()
// 	creds := signer.SetCredentials(key, secret, token, awsProfile)

// 	// create waitgroup
// 	wg := &sync.WaitGroup{}
// 	wg.Add(1)

// 	ch := make(chan _snapshotType)
// 	max := make(chan struct{}, MaxThreads)

// 	// start goroutine to process results
// 	go func() {
// 		defer wg.Done()

// 		for s := range ch {
// 			wg.Add(1)

// 			go func(wg *sync.WaitGroup, s _snapshotType) {
// 				max <- struct{}{}
// 				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(RequestTimeout)*time.Second)

// 				defer func() {
// 					cancel()
// 					wg.Done()
// 					<-max
// 				}()

// 				req, err := _awsUtilsEc2RequestGen(ctx, s)
// 				if err != nil {
// 					logger.LogError(err, "Error generating request")
// 					return
// 				}

// 				res, err := http.DefaultClient.Do(req)
// 				if err != nil {
// 					logger.LogError(err, "Error performing request")
// 					return
// 				}

// 				body, err := io.ReadAll(res.Body)
// 				if err != nil {
// 					logger.LogError(err, "Error reading response body")
// 					return
// 				}
// 				if err := res.Body.Close(); err != nil {
// 					logger.LogError(err, "Error closing response body")
// 					return
// 				}

// 				var snapshots _snapshotRes
// 				if err := xml.Unmarshal(body, &snapshots); err != nil {
// 					logger.LogError(err, "Error unmarshalling response body")
// 				}

// 				if len(snapshots.Snapshots.Items) > 0 {
// 					for _, item := range snapshots.Snapshots.Items {
// 						logger.Logger.Info().
// 							Str("region", s.Region).
// 							Str("snapshot-id", item.SnapshotID).
// 							Str("owner-id", item.OwnerID).
// 							Str("description", item.Description).
// 							Int64("volume-size", item.VolumeSize).
// 							Msg(shared.GetMessageColor("success"))
// 					}
// 				}

// 				// wg.Done()

// 			}(wg, s)

// 		}
// 	}()

// 	// send requests to goroutine
// 	for _, region := range regions {
// 		ch <- _snapshotType{
// 			Region:     region,
// 			Filter:     description,
// 			OwnerId:    ownerID,
// 			Creds:      creds,
// 			MaxResults: maxResults,
// 		}
// 	}

// 	close(ch)
// 	wg.Wait()

// }

// type _snapshotRes struct {
// 	Snapshots struct {
// 		Items []struct {
// 			Description string `xml:"description"`
// 			Encrypted   bool   `xml:"encrypted"`
// 			OwnerID     string `xml:"ownerId"`
// 			Progress    string `xml:"progress"`
// 			SnapshotID  string `xml:"snapshotId"`
// 			StartTime   string `xml:"startTime"`
// 			State       string `xml:"state"`
// 			VolumeID    string `xml:"volumeId"`
// 			VolumeSize  int64  `xml:"volumeSize"`
// 			StorageTier string `xml:"storageTier"`
// 		} `xml:"item"`
// 	} `xml:"snapshotSet"`
// }

// type _snapshotType struct {
// 	Region     string                   `json:"region"`
// 	Filter     string                   `json:"filter"`
// 	OwnerId    string                   `json:"ownerId"`
// 	MaxResults int                      `json:"maxResults"`
// 	Creds      *credentials.Credentials `json:"creds"`
// }

// func _awsUtilsEc2RequestGen(ctx context.Context, t _snapshotType) (*http.Request, error) {
// 	reqUrl := fmt.Sprintf("https://ec2.%s.amazonaws.com/", t.Region)

// 	body := url.Values{}
// 	body.Set("Action", "DescribeSnapshots")
// 	body.Set("Version", "2016-11-15")
// 	body.Set("MaxResults", fmt.Sprintf("%d", t.MaxResults))
// 	if t.Filter != "" {
// 		body.Add("Filter.1.Name", "description")
// 		body.Add("Filter.1.Value.1", t.Filter)
// 	}
// 	// TODO handle ownerId
// 	data := strings.NewReader(body.Encode())

// 	req, err := http.NewRequestWithContext(ctx, "POST", reqUrl, data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Add("Accept", "application/json")

// 	signer := v4.NewSigner(t.Creds)
// 	logger.LogDebugVerbose("url", req.URL.String())

// 	if _, err := signer.Sign(req, data, "ec2", t.Region, time.Now()); err != nil {
// 		return nil, err
// 	}

// 	return req, nil
// }
