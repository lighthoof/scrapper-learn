package main

import (
	"net/url"
	"reflect"
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

// URL getter tests
func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLRelative(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body><a href="boot"><span>Boot.dev</span></a></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"http://crawler-test.com/boot"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLAll(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html>
				    <body>
					  <a href="https://crawler-test.com">
					    <span>Boot.dev</span>
					  </a>
					  <main>
					    <a href="news/latest.html>Latest news</a>
					  </main>
					</body>
				  </html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com", "https://crawler-test.com/news/latest.html"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

// Image getter tests
func TestGetImagesFromHTMLRelative(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body><img src="/logo.png" alt="Logo"></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLMultiple(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html>	
				    <body>
					  <img src="/logo.png" alt="Logo">
					  <main>
					    <img src="older/cat.gif alt="Cat">
					</body>
				  </html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
