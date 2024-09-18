package dependancies

type Dependancies struct {
	controllers *controllerDependancies
	services    *serviceDependancies
	apis        *apiDependancies
}

func Initilize() *Dependancies {
	return &Dependancies{
		controllers: &controllerDependancies{},
		services:    &serviceDependancies{},
		apis:        &apiDependancies{},
	}
}
