package validator

import (
	"fmt"
	"net/url"
	"os"

	"github.com/logrusorgru/aurora"
)

//URLValidator validates ursl from command arguments and print error and exits in case of invalid url
func URLValidator(u string) {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		error := fmt.Errorf("%v", aurora.Index(197, err))
		fmt.Println(error.Error())
		os.Exit(1)
	}
}
