package maps

import "testing"

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s expected %s", got, want)
	}
}

func assertError(t *testing.T, got error) {
	if got != ErrNotFound {
		t.Fatal("expected to get an error")
	}
}
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"key1": "value1"}
	t.Run("Search for a string that is present", func(t *testing.T) {
		got, _ := dictionary.Search("key1")
		want := "value1"
		assertString(t, got, want)
	})

	t.Run("Search for a string that is not present", func(t *testing.T) {
		_, err := dictionary.Search("key2")
		assertError(t, err)
	})

}
