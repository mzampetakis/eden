package cmd

import (
	"fmt"
	"github.com/docker/distribution/context"
	"github.com/docker/docker/client"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/daemon"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/lf-edge/eden/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path"
	"runtime"
	"sort"
)

const (
	defaultEveTag         = "latest"
	defaultEvePrefixInTar = "bits"
)

var (
	eveArch   string
	eveTag    string
	outputDir string
	saveLocal bool
)

var downloaderCmd = &cobra.Command{
	Use:   "download",
	Short: "download eve from docker",
	Long:  `Download eve from docker.`,
	Run: func(cmd *cobra.Command, args []string) {
		if eveTag == "master" {
			eveTag = "latest"
		}
		image = fmt.Sprintf("lfedge/eve:%s-%s", eveTag, eveArch)
		ref, err := name.ParseReference(image)
		if err != nil {
			log.Fatalf("parsing reference %q: %v", image, err)
		}
		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			log.Fatalf("client.NewClientWithOpts: %s", err)
		}
		cli.NegotiateAPIVersion(ctx)
		options := daemon.WithClient(cli)
		img, err := daemon.Image(ref, options)
		if err != nil {
			desc, err := remote.Get(ref)
			if err != nil {
				log.Fatalf("remote.Get: %s", err)
			}
			img, err = desc.Image()
			if err != nil {
				log.Fatalf("desc.Image: %s", err)
			}
			if saveLocal {
				tag, err := name.NewTag(image)
				if err != nil {
					log.Fatalf("name.NewTag: %s", err)
				}
				_, err = daemon.Write(tag, img)
				if err != nil {
					log.Fatalf("daemon.Write: %s", err)
				}
				img, err = daemon.Image(ref, options)
				if err != nil {
					log.Fatalf("daemon.Image on saved image: %s", err)
				}
			}
		}
		layers, err := img.Layers()
		if err != nil {
			log.Fatal(err)
		}
		if len(layers) == 0 {
			log.Fatalf("no layers in image")
		}
		sort.SliceStable(layers, func(i, j int) bool {
			layerISize, err := layers[i].Size()
			if err != nil {
				log.Fatal(err)
			}
			layerJSize, err := layers[j].Size()
			if err != nil {
				log.Fatal(err)
			}
			return layerISize > layerJSize
		})
		neededLayer := layers[0]
		u, err := neededLayer.Uncompressed()
		if err != nil {
			log.Fatal(err)
		}
		defer u.Close()
		err = utils.ExtractFilesFromDocker(u, outputDir, defaultEvePrefixInTar)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func downloaderInit() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	downloaderCmd.Flags().StringVarP(&eveTag, "tag", "", defaultEveTag, "tag to download")
	downloaderCmd.Flags().StringVarP(&eveArch, "arch", "", runtime.GOARCH, "arch of EVE")
	downloaderCmd.Flags().StringVarP(&outputDir, "output", "o", path.Join(currentPath, "dist", "eve", "dist", runtime.GOARCH), "output directory")
	downloaderCmd.Flags().BoolVarP(&saveLocal, "save", "", true, "save image to local docker registry")
}
