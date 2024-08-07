package rest

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"io"
	"log"
	"net/http"
	"strings"
	"svc/proxy-service/internal/common"
	"svc/proxy-service/internal/config"
	"time"
)

var Client *http.Client

func init() {
	// Create a http.Transport and configure it
	transport := &http.Transport{
		IdleConnTimeout:       90 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives:     false,
	}
	// Enable HTTP/2 for the transport
	if err := http2.ConfigureTransport(transport); err != nil {
		log.Fatalf("Failed to configure HTTP/2 transport: %v", err)
	}

	// Initialize the http.Client with the configured transport
	Client = &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}
}
func ProxyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ProxyRequest(c)
		c.Abort()
	}
}

func ProxyRequest(c *gin.Context) {
	targetURL := targetUrlResolve(c)
	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		panic(common.ServerError{Code: 500, Message: "Error when create request " + c.Request.URL.Path})
		return
	}
	prepareRequest(c, req)

	resp, err := Client.Do(req)
	if err != nil {
		panic(common.ServerError{Code: 500, Message: "Error during request " + c.Request.URL.Path})
		return
	}

	handleResponse(c, resp)
}

func targetUrlResolve(c *gin.Context) string {
	path := c.Request.URL.Path
	integrationConfig := config.GetConfig().Integration

	// Define a map of URL prefixes to service URLs
	urlMap := map[string]string{
		"/api/monitoring": integrationConfig.MonitorService.URL,
		"/api/job":        integrationConfig.JobService.URL,
	}

	// Find the appropriate service URL based on the path prefix
	for prefix, url := range urlMap {
		if strings.HasPrefix(path, prefix) {
			return url
		}
	}

	// Default to the core service URL
	return integrationConfig.CoreService.URL + c.Request.URL.RequestURI()
}

func prepareRequest(c *gin.Context, req *http.Request) {
	// Copy the headers
	for k, v := range c.Request.Header {
		req.Header[k] = v
	}
}

func handleResponse(c *gin.Context, resp *http.Response) {

	// header
	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}

	// body
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(common.ServerError{Code: 500, Message: "Error handle request " + c.Request.URL.Path})
	}
	_, err = c.Writer.Write(buf)

	// status
	c.Writer.WriteHeader(resp.StatusCode)

	// close body stream
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
}
