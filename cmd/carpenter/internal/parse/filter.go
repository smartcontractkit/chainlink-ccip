package parse

/*
// Filter returns true if any of the filters match.
func Filter(data *Data, filters []string) bool {

}
*/

// Filter decides if the data should be displayed based on the provided filters.
func Filter(data *Data, filters []string) (bool, error) {
	return true, nil
}
