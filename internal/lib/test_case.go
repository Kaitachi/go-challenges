package lib

type TestCase struct {
	Name		string
	Algorithm	string
	Input		string
	Output		string

	Actual		string
}


// Create Test Case with scenario data
func NewTestCase(c *Challenge, scenario string) *TestCase {

	var input, output string

	switch scenario {
	case "": // Get real data
		input, output = c.getFileData("")
		break
	
	default: // Get scenario data
		input, output = c.getFileData(scenario)
		break
	}

	return &TestCase{
		Name: scenario,
		Input: input,
		Output: output,
		Algorithm: c.Algorithm,
	}
}

