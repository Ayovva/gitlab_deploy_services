package main

import (
	"github.com/docopt/docopt-go"
	"fmt"
	"github.com/vovariabov/gitlab_deploy_services/importer"
	"github.com/vovariabov/gitlab_deploy_services"
	"github.com/vovariabov/gitlab_deploy_services/ms_object"
)

const (
	import_  = "import"
	all      = "--all"
	service = "<service>"
	deploy_to_staging = "deploy_to_staging"
	deploy_to_production = "deploy_to_production"
)
func main() {
	usage := `GitLab Deploy Services
	Usage: 	deploy_util import (--all | <service>...)
			deploy_util deploy_to_staging (--all | <service>...)
			deploy_util deploy_to_production (--all | <service>...)
	`
	tgmsDeploy, err := importer.Import(importer.DOMAIN, importer.GROUP, importer.TGMSDEPLOY)
	if err != nil {
		panic(err)
	}
	s, err := gitlab_deploy_services.FetchServices(tgmsDeploy)

	msCollection := ms_object.InitMsObj(s)

	parser := &docopt.Parser{OptionsFirst: false}
	args, _ := parser.ParseArgs(usage, nil, "1.0")
//	fmt.Printf("%+v %T", args, args)
	if args[import_].(bool) {
		if args[all].(bool) {
			for _, ms := range msCollection.Mss {
				err := ms.CloneRepo()
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			for _, s := range args[service].([]string) {
				err := msCollection.Mss[s].CloneRepo()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	if args[deploy_to_staging].(bool) {
		if args[all].(bool) {
			for _, ms := range msCollection.Mss {
				err := ms.DeployServiceToStaging()
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			for _, s := range args[service].([]string) {
				err := msCollection.Mss[s].DeployServiceToStaging()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	if args[deploy_to_production].(bool) {
		if args[all].(bool) {
			for _, ms := range msCollection.Mss {
				err := ms.DeployServiceToProduction()
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			for _, s := range args[service].([]string) {
				err := msCollection.Mss[s].DeployServiceToProduction()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}