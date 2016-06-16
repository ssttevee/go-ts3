package ts3

import "strings"

// A ServerResponse is a wrapper for raw responses received from the server query
type ServerResponse string

// String returns the raw response
func (r ServerResponse) String() string {
	return string(r)
}

// Header returns the first word of the raw response.
// Includes but is not limited to notify callbacks and error responses.
// Does not include words that are part of the response data.
func (r ServerResponse) Header() string {
	pos := strings.IndexRune(string(r), ' ')
	header := string(r[:pos])

	if strings.IndexRune(header, '=') != -1 {
		return header
	}

	return ""
}

// Map parses and returns the response data into a map
// when there is only one set of data in the response.
func (r ServerResponse) Map() map[string]string {
	return parseResponse(string(r))
}

// Array parses and returns the response data into an array
// when there is more than one set of data in the response.
func (r ServerResponse) Array() []map[string]string {
	items := strings.Split(string(r), "|")

	ret := make([]map[string]string, 0, len(items))
	for _, item := range items {
		ret = append(ret, parseResponse(item))
	}

	return ret
}

func parseResponse(response string) map[string]string {
	parts := strings.Split(response, " ")

	ret := make(map[string]string)
	for _, part := range parts {
		if pos := strings.Index(part, "="); pos != -1 {
			ret[part[:pos]] = strings.Replace(part[pos+1:], "\\s", " ", -1)
		}
	}

	return ret
}
