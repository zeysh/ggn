package work

import "github.com/n0rad/go-erlog/logs"

func (e Env) Generate() {
	logs.WithFields(e.fields).Debug("Generating units")
	services := e.ListServices()

	for _, service := range services {
		service := e.LoadService(service)
		if err := service.Generate(); err != nil {
			logs.WithE(err).Error("Generate failed")
		}
	}
}
