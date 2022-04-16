package kongclient

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

// // Gives a pointer to a string, returns nil for empty strings
// func handleEmptyInt(i int) *int {
// 	if i == 0 {
// 		return nil
// 	}
// 	return &i
// }
