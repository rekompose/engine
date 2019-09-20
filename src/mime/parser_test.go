package mime

import "testing"

func TestParseSubjectFromRawMessage(t *testing.T)  {
	raw := "000"

	m := Parse([]byte(raw))

	t.Error(m.Subject)
}
