package main

import (
	"strings"
	"testing"
)

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{
		{
			name: "find h1",
			inputURL: `<html>
						  <body>
						    <h1>Welcome</h1>
							  <h2>to Boot.dev</h2>
						      <main>
						      	<p>Learn to code by building real projects.</p>
						      	<p>This is the second paragraph.</p>
						      </main>
						  </body>
						</html>`,
			expected: "Welcome",
		},
		{
			name: "find h2",
			inputURL: `<html>
						  <body>
						    <h2>Welcome to Boot.dev</h2>
						    <main>
						      <p>Learn to code by building real projects.</p>
						      <p>This is the second paragraph.</p>
						    </main>
						  </body>
						</html>`,
			expected: "Welcome to Boot.dev",
		},
		{
			name: "no closing tag",
			inputURL: `<html>
						  <body>
						    <h1>Welcome to Boot.dev</h1
						    <main>
						      <p>Learn to code by building real projects.</p>
						      <p>This is the second paragraph.</p>
						    </main>
						  </body>
						</html>`,
			expected: "Welcome to Boot.dev",
		},
		{
			name: "no opening tag",
			inputURL: `<html>
						  <body>
						    h1>Welcome to Boot.dev</h1>
						    <main>
						      <p>Learn to code by building real projects.</p>
						      <p>This is the second paragraph.</p>
						    </main>
						  </body>
						</html>`,
			expected: "",
		},
		{
			name: "not html",
			inputURL: `if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return`,
			expected: "",
			//errorContains: "couldn't parse URL",
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getHeadingFromHTML(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{
		{
			name: "find first p in main",
			inputURL: `<html>
						  <body>
						  	<p>Whatever, skip this</p>
						    <h1>Welcome to Boot.dev</h1>
						    <main>
						      <p>Learn to code by building real projects.</p>
						      <p>This is the second paragraph.</p>
						    </main>
						  </body>
						</html>`,
			expected: "Learn to code by building real projects.",
		},
		{
			name: "find first p without main",
			inputURL: `<html>
						  <body>
						    <h1>Welcome to Boot.dev</h1>
						    <p>Learn to code by building real projects.</p>
						    <p>This is the second paragraph.</p>
						  </body>
						</html>`,
			expected: "Learn to code by building real projects.",
		},
		{
			name: "no closing tag",
			inputURL: `<html>
						  <body>
						    <h1>Welcome to Boot.dev</h1>
						    <main>
						      <p>Learn to code by building real projects.</p
						      <p>This is the second paragraph.</p>
						    </main>
						  </body>
						</html>`,
			expected: "Learn to code by building real projects.",
		},
		{
			name: "no opening tag",
			inputURL: `<html>
						  <body>
						    <h1>Welcome to Boot.dev</h1>
						    <main>
						      p>Learn to code by building real projects.</p>
						      <p>This is the second paragraph.</p>
						    </main>
						  </body>
						</html>`,
			expected: "",
		},
		{
			name: "not html",
			inputURL: `if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return`,
			expected: "",
			//errorContains: "couldn't parse URL",
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getFirstParagraphFromHTML(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
