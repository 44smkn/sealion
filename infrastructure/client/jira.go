package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"sealion/domain/model"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var Set = wire.NewSet(NewJira)

type JiraClient struct {
	URL        *url.URL
	HTTPClient *http.Client

	Username, Password string
}

type auth struct {
	Session struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"session"`
}

func NewJira() (*JiraClient, error) {
	urlStr := os.Getenv("JIRA_BASE_URL")
	username := os.Getenv("JIRA_USERNAME")
	password := os.Getenv("JIRA_PASSWORD")

	if len(username) == 0 {
		return nil, errors.New("missing  username")
	}

	if len(password) == 0 {
		return nil, errors.New("missing user password")
	}

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", urlStr)
	}

	httpClient := &http.Client{}

	client := &JiraClient{
		URL:        parsedURL,
		Username:   username,
		Password:   password,
		HTTPClient: httpClient,
	}
	return client, nil

}

func (c *JiraClient) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	setCookie(ctx, req, c)

	return req, nil
}

func (c *JiraClient) GetMyIssue(ctx context.Context) ([]model.Issue, error) {
	q := &url.Values{}
	q.Add("fields", "issuetype,project,summary,status,duedate")
	jql := fmt.Sprintf("assignee=%s AND status!=closed AND project!=TROUBLE", c.Username)
	q.Add("jql", jql)

	req, err := c.newRequest(ctx, http.MethodGet, "/rest/api/2/search", nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to request jira issues api")
	}
	req.URL.RawQuery = q.Encode()
	logrus.Debugf("path: %v, query: %v\n", req.URL.Path, req.URL.RawQuery)
	logrus.Debugf("header: %v\n", req.Header)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read respose body")
		}
		return nil, errors.New("StatusCode is not 200. response body is " + string(body))
	}

	response := struct {
		Issues []model.Issue `json:"issues"`
	}{}

	if err := decodeBody(res, &response); err != nil {
		return nil, errors.Wrapf(err, "failed to decode respose from jira. response is %v", res)
	}

	return response.Issues, nil
}

func setCookie(ctx context.Context, req *http.Request, c *JiraClient) error {

	body := fmt.Sprintf("{\"username\": \"%s\",\"password\": \"%s\"}", c.Username, c.Password)
	spath := fmt.Sprintf("%s/rest/auth/1/session", c.URL.String())
	preReq, err := http.NewRequest(http.MethodPost, spath, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return errors.Wrapf(err, "failed to set body stirings for authentification.")
	}
	preReq.Header.Set("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(preReq)
	if err != nil {
		return errors.Wrapf(err, "failed to request for authentification.")
	}

	var auth auth
	if err := decodeBody(res, &auth); err != nil {
		return errors.Wrapf(err, "failed to decode auth from jira. response is %v", res)
	}

	req.Header.Add("cookie", fmt.Sprintf("%s=%s", auth.Session.Name, auth.Session.Value))

	return nil
}
