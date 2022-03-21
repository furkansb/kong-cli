package kongclient

func strPointerToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
