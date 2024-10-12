package cmd

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
	"v2ray.com/core/app/router"
)

var (
	ruleSitesDomains []*router.Domain
)

var geositeCmd = &cobra.Command{
	Use:               "geosite",
	Short:             "generate geosite resources from text files",
	PersistentPreRunE: PersistentPreFunc,
	Run: func(cmd *cobra.Command, args []string) {
		MakeGeoSiteDatFile(srcDir, outDir)
	},
}

func MakeGeoSiteDatFile(srcDir string, outDir string) {
	siteFile := path.Join(srcDir, "sites")
	datFile := path.Join(outDir, "geosite.dat")

	v2SiteList := new(router.GeoSiteList)
	ruleFiles, err := os.ReadDir(siteFile)

	if err != nil {
		panic(err)
	}

	for _, rf := range ruleFiles {
		filename := rf.Name()
		fmt.Println(filename)
		v2SiteList.Entry = append(v2SiteList.Entry, &router.GeoSite{
			CountryCode: strings.ToUpper(filename),
			Domain:      GetSitesList(siteFile + "/" + filename),
		})
	}

	v2SiteListBytes, err := proto.Marshal(v2SiteList)

	if err != nil {
		fmt.Println("failed to marshal v2sites:", err)
		return
	}

	if err := os.WriteFile(datFile, v2SiteListBytes, FilePerm); err != nil {
		fmt.Println("failed to write %s.", datFile, err)
	}

	fmt.Println(siteFile, "-->", datFile)
}

func processOutDir() error {
	if outDir == "" {
		outDir = srcDir
	}
	return nil
}

func init() {
	geositeCmd.PersistentFlags().StringVar(&srcDir, "src", "", "source directory")
	geositeCmd.PersistentFlags().StringVar(&outDir, "out", "", "output directory")
	rootCmd.AddCommand(geositeCmd)
}

func GetSitesList(fileName string) []*router.Domain {
	d, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	domains := strings.Split(string(d), "\n")

	ruleSitesDomains = make([]*router.Domain, len(domains))
	for idx, pattern := range domains {
		ruleSitesDomains[idx] = &router.Domain{
			Type:  router.Domain_Domain,
			Value: pattern,
		}
	}
	return ruleSitesDomains
}
