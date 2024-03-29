package main

import (
	"flag"
	"reflect"
	"github.com/robfig/revel"
	controllers "emberdata/app/controllers"
	tests "emberdata/tests"
	controllers0 "github.com/robfig/revel/modules/testrunner/app/controllers"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	rev.Init(*runMode, *importPath, *srcPath)
	rev.INFO.Println("Running revel server")
	
	rev.RegisterController((*controllers.Application)(nil),
		[]*rev.MethodType{
			&rev.MethodType{
				Name: "Index",
				Args: []*rev.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					10: []string{ 
					},
				},
			},
			&rev.MethodType{
				Name: "SelectedPerson",
				Args: []*rev.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&rev.MethodType{
				Name: "People",
				Args: []*rev.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&rev.MethodType{
				Name: "Person",
				Args: []*rev.MethodArg{ 
					&rev.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	rev.RegisterController((*controllers0.TestRunner)(nil),
		[]*rev.MethodType{
			&rev.MethodType{
				Name: "Index",
				Args: []*rev.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					46: []string{ 
						"testSuites",
					},
				},
			},
			&rev.MethodType{
				Name: "Run",
				Args: []*rev.MethodArg{ 
					&rev.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&rev.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					69: []string{ 
						"error",
					},
				},
			},
			&rev.MethodType{
				Name: "List",
				Args: []*rev.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	rev.DefaultValidationKeys = map[string]map[int]string{ 
	}
	rev.TestSuites = []interface{}{ 
		(*tests.ApplicationTest)(nil),
	}

	rev.Run(*port)
}
