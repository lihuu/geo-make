package cmd

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
	"v2ray.com/core/app/router"
	"v2ray.com/core/common"
	"v2ray.com/core/infra/conf"
)

var geoipCmd = &cobra.Command{
	Use:               "geoip",
	Short:             "generate geoip resources from text files",
	PersistentPreRunE: PersistentPreFunc,
	Run: func(cmd *cobra.Command, args []string) {
		MakeGeoIpDatFile(srcDir, outDir)
	},
}

func getIPsList(fileName string) []*router.CIDR {

	d, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	ips := strings.Split(string(d), "\n")

	cidr := make([]*router.CIDR, 0, len(ips))

	for _, ip := range ips {
		c, err := conf.ParseIP(ip)
		common.Must(err)
		cidr = append(cidr, c)
	}
	return cidr
}

func MakeGeoIpDatFile(srcDir string, outDir string) {
	ipFile := path.Join(srcDir, "ips")
	datFile := path.Join(outDir, "geoip.dat")

	geoIPList := new(router.GeoIPList)

	ruleFiles, err := os.ReadDir(ipFile)
	if err != nil {
		panic(err)
	}

	for _, rf := range ruleFiles {
		filename := rf.Name()
		fmt.Println(filename)
		geoIPList.Entry = append(geoIPList.Entry, &router.GeoIP{
			CountryCode: strings.ToUpper(filename),
			Cidr:        getIPsList(ipFile + "/" + filename),
		})
	}

	geoIPBytes, err := proto.Marshal(geoIPList)
	if err != nil {
		fmt.Println("Error marshalling geoip list:", err)
	}

	if err := os.WriteFile("geoip.dat", geoIPBytes, 0777); err != nil {
		fmt.Println("Error writing geoip to file:", err)
	}

	fmt.Println(ipFile, "-->", datFile)
}

func init() {
	geoipCmd.PersistentFlags().StringVar(&srcDir, "src", "", "source directory")
	geoipCmd.PersistentFlags().StringVar(&outDir, "out", "", "output directory")
	rootCmd.AddCommand(geoipCmd)
}
