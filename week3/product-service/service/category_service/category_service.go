package category_service

import (
	"io/ioutil"
	"net/http"
	"product-service/config"
	"product-service/constants"
	"strings"
)

const PATH_CATEGORY_LIST = "categories"

func FetchCategoryByIdList(idList []string) (response []byte, err error) {
	req, err := http.NewRequest("GET", config.Config.Services.CategoryService+PATH_CATEGORY_LIST, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("ids", strings.Join(idList, ","))

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	client.Timeout = constants.CATEGORY_SERVICE_TIMEOUT
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return body, nil
}
