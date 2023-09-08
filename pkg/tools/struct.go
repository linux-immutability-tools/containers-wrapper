package tools

/*	License: GPLv3
	Authors:
		Mirko Brombin <mirko@fabricators.ltd>
		Lit Contributors <https://github.com/linux-immutability-tools/>
	Copyright: 2023
	Description:
		Containers Wrapper is a Go library that provides a convenient
		and unified interface for interacting with container engines
		such as Docker, Podman, and Containerd.
*/

import (
	"fmt"
	"reflect"
)

// Struct2Args converts a struct to a list of command line arguments
// where the argument name is the json tag of the struct field and the
// argument value is the value of the struct field.
// I'm ot happy with this, but it's the only way I found to pass the options
// to the command without having to parse each option and build the command
// manually so, let's say it's a temporary solution
func Struct2Args(options interface{}, original interface{}) (args []string) {
	optVal := reflect.ValueOf(options)
	origVal := reflect.ValueOf(original)

	if optVal.Type() != origVal.Type() {
		panic("options and original must have the same type")
	}

	for i := 0; i < optVal.NumField(); i++ {
		optField := optVal.Field(i)
		origField := origVal.Field(i)

		if optField.CanInterface() {
			fieldName := optVal.Type().Field(i).Tag.Get("json")
			optValue := optField.Interface()
			origValue := origField.Interface()

			if fieldName != "" && !reflect.DeepEqual(optValue, origValue) {
				if (reflect.TypeOf(optValue).Kind() == reflect.Slice && reflect.ValueOf(optValue).Len() == 0) ||
					(reflect.TypeOf(optValue).Kind() == reflect.Map && reflect.ValueOf(optValue).Len() == 0) {
					continue
				}

				args = append(args, fmt.Sprintf("--%s", fieldName))
				args = append(args, fmt.Sprintf("%v", optValue))
			}
		}
	}

	return args
}
