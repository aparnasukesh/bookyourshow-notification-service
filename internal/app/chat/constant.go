package chat

var responses = map[string]string{
	"hi":      "Hi, how can I help you?",
	"hello":   "Hello, how can I assist you?",
	"hey":     "Hey there! How can I help?",
	"hai":     "Hai, how can I assist you today?",
	"support": "Our support team is here for you. What do you need help with?",
}

func ResponseMsg(input string) string {
	if exist, ok := responses[input]; ok {
		return exist
	}
	return "I'm here to help with any questions you have!"
}
