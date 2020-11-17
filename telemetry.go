package thingsboard

// Telemetry
// https://thingsboard.io/docs/reference/http-api/#telemetry-upload-api
// HTTP basics
// HTTP is a general-purpose network protocol that can be used in IoT applications. You can find more information about HTTP here. HTTP protocol is TCP based and uses request-response model.
// ThingsBoard server nodes act as an HTTP Server that supports both HTTP and HTTPS protocols.
// Client libraries setup
// You can find HTTP client libraries for different programming languages on the web. Examples in this article will be based on curl. In order to setup this tool, you can use instructions in our Hello World guide.
// HTTP Authentication and error codes
// We will use access token device credentials in this article and they will be referred to later as $ACCESS_TOKEN. The application needs to include $ACCESS_TOKEN as a path parameter in each HTTP request. Possible error codes and their reasons:

// 400 Bad Request - Invalid URL, request parameters or body.
// 401 Unauthorized - Invalid $ACCESS_TOKEN.
// 404 Not Found - Resource not found.
// Key-value format
// By default, ThingsBoard supports key-value content in JSON. Key is always a string, while value can be either string, boolean, double, long or JSON. Using custom binary format or some serialization framework is also possible. See protocol customization for more details. For example:

// {
//  "stringKey":"value1",
//  "booleanKey":true,
//  "doubleKey":42.0,
//  "longKey":73,
//  "jsonKey": {
//     "someNumber": 42,
//     "someArray": [1,2,3],
//     "someNestedObject": {"key": "value"}
//  }
// }
// Telemetry upload API
// In order to publish telemetry data to ThingsBoard server node, send POST request to the following URL:

// http(s)://host:port/api/v1/$ACCESS_TOKEN/telemetry
// The simplest supported data formats are:

// {"key1":"value1", "key2":"value2"}
// or

// [{"key1":"value1"}, {"key2":"value2"}]
// Please note that in this case, the server-side timestamp will be assigned to uploaded data!

// In case your device is able to get the client-side timestamp, you can use following format:

// {"ts":1451649600512, "values":{"key1":"value1", "key2":"value2"}}
// In the example above, we assume that “1451649600512” is a unix timestamp with milliseconds precision. For example, the value ‘1451649600512’ corresponds to ‘Fri, 01 Jan 2016 12:00:00.512 GMT’

// TelemetryData holds telemetry data that's going to be uploaded to ThingsBoard
type TelemetryData struct{}

// SaveTelemetry uploads Telemetry to ThingsBoard, by deviceID
func (tb *Thingsboard) SaveTelemetry(deviceToken string, td TelemetryData) error {

	_, err := tb.resty.R().
		SetBody(&td).
		Post(tb.api + "/" + deviceToken + "/telemetry")

	return err
}
