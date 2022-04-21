package kongclient

import "os"

func strPointerToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Gives a pointer to a string, returns nil for empty strings
func StrToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func PsliceFromStrSlice(sl []string) []*string {
	result := make([]*string, len(sl))
	for i, str := range sl {
		strLocal := str
		result[i] = &strLocal
	}
	return result
}

func StrSliceFromPSlice(sl []*string) []string {
	result := make([]string, len(sl))
	for i, str := range sl {
		strLocal := str
		result[i] = *strLocal
	}
	return result
}

func BoolHandler(b bool) *bool {
	if !b {
		return nil
	}
	return &b
}

func GetBaseUrl() string {
	baseUrl := os.Getenv("KONG_ADMIN_URL")
	if baseUrl == "" {
		return "http://localhost:8001"
	}
	return baseUrl
}
