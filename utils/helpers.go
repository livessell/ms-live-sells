package utils

// ExtractProductCode extracts the product code from a comment
// If a product code is found (indicated by a hashtag), it returns it
func ExtractProductCode(comment string) string {
	// Simple logic to check for a product code in the comment (e.g., starting with "#")
	if len(comment) > 0 && comment[0] == '#' {
		return comment[1:] // Return the product code after the "#" symbol
	}
	return ""
}
