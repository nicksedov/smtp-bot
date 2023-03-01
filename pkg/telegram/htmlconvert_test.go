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
			<li>Bullet Item <b>1</b></li>
			<li>Bullet Item <b>2</b></li>
			<li>Bullet Item <b>3</b></li>
		</ul>
		<ol>
			<li>Numbered Item 1</li>
			<li>Numbered Item 2</li>
			<li>Numbered Item 3</li>
		</ol>
	`
	// Call the function and check the output
	output := TryMakeHtmlTelegramCompatible(input)
	fmt.Printf("=>%s<=", output)
}
