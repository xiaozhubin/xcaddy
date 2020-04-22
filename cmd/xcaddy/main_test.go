package main

import "testing"

func TestSplitWith(t *testing.T) {
	for i, tc := range []struct {
		input         string
		expectModule  string
		expectVersion string
		expectReplace string
		expectErr     bool
	}{
		{
			input:        "module",
			expectModule: "module",
		},
		{
			input:         "module@version",
			expectModule:  "module",
			expectVersion: "version",
		},
		{
			input:         "module@version=replace",
			expectModule:  "module",
			expectVersion: "version",
			expectReplace: "replace",
		},
		{
			input:         "module=replace",
			expectModule:  "module",
			expectReplace: "replace",
		},
		{
			input:     "=replace",
			expectErr: true,
		},
		{
			input:     "@version",
			expectErr: true,
		},
		{
			input:     "@version=replace",
			expectErr: true,
		},
		{
			input:     "",
			expectErr: true,
		},
	} {
		actualModule, actualVersion, actualReplace, actualErr := splitWith(tc.input)
		if actualModule != tc.expectModule {
			t.Errorf("Test %d: Expected module '%s' but got '%s' (input=%s)",
				i, tc.expectModule, actualModule, tc.input)
		}
		if tc.expectErr {
			if actualErr == nil {
				t.Errorf("Test %d: Expected error but did not get one (input='%s')", i, tc.input)
			}
			continue
		}
		if !tc.expectErr && actualErr != nil {
			t.Errorf("Test %d: Expected no error but got: %s (input='%s')", i, actualErr, tc.input)
		}
		if actualVersion != tc.expectVersion {
			t.Errorf("Test %d: Expected version '%s' but got '%s' (input='%s')",
				i, tc.expectVersion, actualVersion, tc.input)
		}
		if actualReplace != tc.expectReplace {
			t.Errorf("Test %d: Expected module '%s' but got '%s' (input='%s')",
				i, tc.expectReplace, actualReplace, tc.input)
		}
	}
}
