package json

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

type Class struct {
	Name  string
	Param []Param
}

type Param struct {
	Name     string
	Property string
	List     bool
}

var class_group []Class

func createFile(body string) (string, error) {
	dir, err := getDir(body)
	if err != nil {
		return "", err
	}
	getClass("root", dir)
	// data := []Class{}
	// data = append(data, Class{Name: "root", Item: get(dir)})
	// for _, v := range class_group {
	// 	log.Printf("%s: \n", v.Name)
	// 	for _, v1 := range v.Param {
	// 		log.Printf("%s : %s \n", v1.Name, v1.Property)
	// 	}
	// }
	//data, _ := ioutil.ReadFile("./pkg/json/c#.tpl")
	render := template.Must(template.New("model").
		Parse(`
namespace project
{ {{range $}}
	public class {{.Name}}
	{ {{range .Param}}
		public {{.Name}} {{if .List}}List<{{.Property}}>{{else}}{{.Property}}{{end}} {get;set;} 	
		{{end}}
	}
	{{end}}
}`))
	err = render.Execute(os.Stdout, class_group)
	return "", nil
}

func getClass(class_name string, item map[string]interface{}) {
	class_group = append(class_group,
		Class{
			Name:  class_name,
			Param: getProperty(class_name, item),
		})

}

func getProperty(class_name string, item map[string]interface{}) []Param {
	param := []Param{}
	for k, v := range item {
		p := Param{
			Name: fmt.Sprintf("%s", k),
		}
		var temp map[string]interface{}
		switch v.(type) {
		case []interface{}:
			list := v.([]interface{})
			if len(list) == 0 {
				return nil
			}
			p.Property = typeChange(list[0])
			p.List = true
			if p.Property == "class" {
				temp = list[0].(map[string]interface{})
			}
		default:
			p.Property = typeChange(v)
			if p.Property == "class" {
				temp = v.(map[string]interface{})
			}
		}

		if temp != nil {
			p.Property = fmt.Sprintf("class_%s", k)
			getClass(p.Property, temp)
		}

		param = append(param, p)
	}

	return param
}

func getDir(body string) (res map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func typeChange(value interface{}) string {
	switch value.(type) {
	case string:
		return "string"
	case float32, float64:
		return "decimal"
	case int, int8, int16, int32:
		return "int"
	case int64:
		return "long"
	case map[string]interface{}:
		return "class"
	default:
		return "unknown"
	}

}
