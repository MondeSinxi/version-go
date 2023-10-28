package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

type version struct {
	major int
	minor int
	patch int
}

func (v *version) bumpVersion(part string) {
	switch part {
	case "major":
		v.major += 1
		v.minor = 0
	case "minor":
		v.minor += 1
		v.patch = 0
	case "patch":
		v.patch += 1
	}
}

func (v *version) strVersion() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	app := &cli.App{
		Usage: "Update a semantic version",
		Commands: []*cli.Command{
			{
				Name:   "bump-version",
				Action: execute,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "write",
						Aliases: []string{"w"},
						Usage:   "File to write version after bump.",
					},
					&cli.StringFlag{
						Name:    "read",
						Aliases: []string{"r"},
						Usage:   "File to get current version.",
					},
				},
			},
			{
				Name: "get-version",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "read",
						Aliases: []string{"r"},
						Usage:   "File to get current version.",
					},
				},
				Action: func(cCtx *cli.Context) error {
					fileName := cCtx.String("read")
					if len(fileName) > 0 {
						version, err := getVersion(fileName)
						check(err)
						fmt.Printf("Current version: %s", version)
					} else {
						log.Fatal("No file to get current version.")
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func execute(cCtx *cli.Context) error {
	versionType := cCtx.Args().First()
	fmt.Println(versionType)
	fmt.Println(cCtx.String("write"))
	v, err := BumpVersion(versionType, cCtx.String("read"))
	check(err)
	if len(cCtx.String("write")) > 0 {
		f, err := os.Create(cCtx.String("write"))
		check(err)
		fmt.Println(v)
		w := bufio.NewWriter(f)
		w.WriteString(v)
		fmt.Printf("Write version file %s", cCtx.String("write"))
		w.Flush()
	}

	return nil
}

func BumpVersion(versionType string, versionFile string) (string, error) {
	var parts [3]int
	version_str, err := getVersion(versionFile)
	check(err)
	versionPartsStr := strings.Split(version_str, ".")
	for idx, a := range versionPartsStr {
		parts[idx] = convertStrtoInt(a)
	}
	check(err)

	v := version{major: parts[0], minor: parts[1], patch: parts[2]}
	v.bumpVersion(versionType)
	strVersion := v.strVersion()
	fmt.Println(strVersion)
	return strVersion, nil

}

func getVersion(versionFile string) (string, error) {
	version, err := os.ReadFile(versionFile)
	check(err)
	return strings.TrimSpace(string(version)), nil
}

func convertStrtoInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}
