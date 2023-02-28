package telegram

import (
	"fmt"
	"testing"
)

func TestTryMakeHtmlTelegramCompatible(t *testing.T) {
	// Test input and expected output
	input := `
		<p><strong>Bold Text</strong></p>
		<p><em>Italic Text</em></p>
		<p><u>Underlined Text</u></p>
		<p><s>Strikethrough Text</s></p>
		<p><code>Code Text</code></p>
		<p><a href="https://example.com">Link Text</a></p>
		<ul>
			<li>List Item 1</li>
			<li>List Item 2</li>
			<li>List Item 3</li>
		</ul>
		<ol>
			<li>List Item 1</li>
			<li>List Item 2</li>
			<li>List Item 3</li>
		</ol>
	`
	// Call the function and check the output
	output := TryMakeHtmlTelegramCompatible(input)
	fmt.Printf("=>%s<=", output)
}
