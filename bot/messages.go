package utils

const (
	invalidSyntaxMsg = `❌ Invalid expression\. Please use only digits, \+ \- \* / and parentheses\.
Use \/help for more info\.`
	defaultMsg = "Sorry, I don't know that command"
	startMsg   = "👋 Welcome! I am a simple calculator bot.\n\nJust send me a math expression like:\n`(2 + 3) * 4 - 1`\n\nType /help to learn more."
	helpMsg    = `*Help Menu*

Send me any valid arithmetic expression containing:
\- Digits: 0\-9
\- Operators: \+ \- \* /
\- Parentheses: \( \)

Examples:
2 \+ 2
10 / \(2 \+ 3\)
\(4 \+ 5\) \* \(6 \- 3\)

I’ll calculate and reply with the result\.
`
)
