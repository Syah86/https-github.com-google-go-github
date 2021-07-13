package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type EndpointPattern = *regexp.Regexp

// Users
var UsersGetEndpoint EndpointPattern = regexp.MustCompile(`^\/users\/[a-zA-Z]+`)

// Orgs
var OrgsListEndpoint = regexp.MustCompile(`^\/users\/([a-z]+\/orgs|orgs)$`)
var OrgsGetEndpoint = regexp.MustCompile(`^\/orgs\/[a-z]+`)

type RequestMatch struct {
	EndpointPattern EndpointPattern
	Method          string // GET or POST
}

func (rm *RequestMatch) Match(r *http.Request) bool {
	if (r.Method == rm.Method) &&
		r.URL.Path == rm.EndpointPattern.FindString(r.URL.Path) {
		return true
	}

	return false
}

var RequestMatchUsersGet = RequestMatch{
	EndpointPattern: UsersGetEndpoint,
	Method:          http.MethodGet,
}

var RequestMatchOrganizationsList = RequestMatch{
	EndpointPattern: OrgsListEndpoint,
	Method:          http.MethodGet,
}

type MockRoundTripper struct {
	RequestMocks map[RequestMatch][][]byte
}

// RoundTrip implements http.RoundTripper interface
func (mrt *MockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	for requestMatch, respBodies := range mrt.RequestMocks {
		if requestMatch.Match(r) {
			if len(respBodies) == 0 {
				fmt.Printf(
					"no more available mocked responses for endpoit %s\n",
					r.URL.Path,
				)

				fmt.Println("please add the required RequestMatch to the MockHttpClient. Eg.")
				fmt.Println(`
				mockedHttpClient := NewMockHttpClient(
					WithRequestMatch(
						RequestMatchUsersGet,
						MustMarshall(github.User{
							Name: github.String("foobar"),
						}),
					),
					WithRequestMatch(
						RequestMatchOrganizationsList,
						MustMarshall([]github.Organization{
							{
								Name: github.String("foobar123"),
							},
						}),
					),
				)
				`)

				panic(nil)
			}

			resp := respBodies[0]

			defer func(mrt *MockRoundTripper, rm RequestMatch) {
				mrt.RequestMocks[rm] = mrt.RequestMocks[rm][1:]
			}(mrt, requestMatch)

			re := bytes.NewReader(resp)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(re),
			}, nil
		}
	}

	return nil, fmt.Errorf(
		"couldn find a mock request that matches the request sent to: %s",
		r.URL.Path,
	)

}

var _ http.RoundTripper = &MockRoundTripper{}

type MockHttpClientOption func(*MockRoundTripper)

func WithRequestMatch(
	rm RequestMatch,
	marshalled []byte,
) MockHttpClientOption {
	return func(mrt *MockRoundTripper) {
		if _, found := mrt.RequestMocks[rm]; !found {
			mrt.RequestMocks[rm] = make([][]byte, 0)
		}

		mrt.RequestMocks[rm] = append(
			mrt.RequestMocks[rm],
			marshalled,
		)
	}
}

func NewMockHttpClient(options ...MockHttpClientOption) *http.Client {
	rt := &MockRoundTripper{
		RequestMocks: make(map[RequestMatch][][]byte),
	}

	for _, o := range options {
		o(rt)
	}

	return &http.Client{
		Transport: rt,
	}
}

func MustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)

	if err == nil {
		return b
	}

	panic(err)
}