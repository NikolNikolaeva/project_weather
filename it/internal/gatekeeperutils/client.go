package gatekeeperutils

//
//import (
//	"context"
//	encodingjson "encoding/json"
//	"fmt"
//	"github.com/NikolNikolaeva/project_weather/config"
//	"github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/http"
//	"io"
//	nethttp "net/http"
//	"net/url"
//)
//
////go:generate mockgen --build_flags=--mod=mod -destination ../../generated/go-mocks/utils/gatekeeperutils/mock_client.go . Client
//type Client interface {
//	GetDevice(ctx context.Context, id string) (*Device, http.APIError)
//}
//
//func NewClient(configuration config.ApplicationConfiguration) Client {
//	return &_Client{
//		client: nethttp.DefaultClient,
//		//logger: logger.WithGroup("gatekeeperutils.Client"),
//		//url: lang.Must(url.Parse(configuration.GatekeeperURL)),
//	}
//}
//
//type _Client struct {
//	url *url.URL
//	//logger logs.Logger
//	client *nethttp.Client
//}
//
//func (self *_Client) GetDevice(ctx context.Context, id string) (*Device, http.APIError) {
//	request, err := self.createRequest(ctx, nethttp.MethodGet, nethttp.NoBody, "devices", id)
//	if err != nil {
//		return nil, err
//	}
//
//	return self.processRequest(request, &Device{})
//}
//
//func (self *_Client) parseBody(body io.ReadCloser, target *Device) (*Device, http.APIError) {
//	bytes, err := io.ReadAll(body)
//	if err != nil {
//		return target, http.NewAPIError(
//			err,
//			nethttp.StatusInternalServerError,
//			nethttp.StatusText(nethttp.StatusInternalServerError),
//		)
//	}
//
//	return target, http.NewAPIError(
//		encodingjson.Unmarshal(bytes, target),
//		nethttp.StatusInternalServerError,
//		nethttp.StatusText(nethttp.StatusInternalServerError),
//	)
//}
//
//func (self *_Client) processRequest(request *nethttp.Request, target *Device) (*Device, http.APIError) {
//	response, err := self.client.Do(request)
//	if err != nil {
//		return nil, http.NewAPIError(
//			err,
//			nethttp.StatusInternalServerError,
//			nethttp.StatusText(nethttp.StatusInternalServerError),
//		)
//	}
//
//	return self.processResponse(response, target)
//}
//
//func (self *_Client) processResponse(response *nethttp.Response, target *Device) (*Device, http.APIError) {
//	defer response.Body.Close()
//
//	if response.StatusCode < 200 || response.StatusCode > 299 {
//		return target, http.NewAPIError(
//			fmt.Errorf("unexpected response code received: %v", response.StatusCode),
//			response.StatusCode,
//			nethttp.StatusText(response.StatusCode),
//		)
//	}
//
//	return self.parseBody(response.Body, target)
//}
//
//func (self *_Client) createRequest(ctx context.Context, method string, body io.Reader, path ...string) (*nethttp.Request, http.APIError) {
//	request, err := nethttp.NewRequestWithContext(
//		ctx,
//		method,
//		http.NewURLBuilder(self.url).
//			AppendPath(path...).
//			Build().
//			String(),
//		body,
//	)
//
//	if err == nil {
//		//request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", ctx.Value(oauth.RAW_JWT)))
//		request.Header.Set("Authorization", fmt.Sprintf("Bearer"))
//	}
//
//	return request, http.NewAPIError(
//		err,
//		nethttp.StatusInternalServerError,
//		nethttp.StatusText(nethttp.StatusInternalServerError),
//	)
//}
