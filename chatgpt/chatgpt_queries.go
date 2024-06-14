package chatgpt

const (
	apiEndpoint    = "https://api.openai.com/v1/chat/completions"
	inititalPrompt = "You are summarising reviews people left of a mass transport system"
	directions     = `The first field "overallOpinion" should be general view the reviewer has. It can only have 5 values (Positive, Slightly Positive, Mixed, Slightly Negative, Negative). For each of the 9 following topics seperated by commas, (Parentheses provide more context and are not their own topic) and provide answers to 3 questions for the previous review. Do they mention the topic? do they view the topic positively, negatively, or mixed (The value for those fields should only be 'Positive', 'Negative', or 'Mixed')? Summarize the topic in 10 words or less. Only fill out questions you are confident in. If you do not have a confident answer leave the mentions field false and opinions field and summary field as empty stringsl. Drivers (drivers of either trains, buse or Trams, streetcars), Purchasing (Buying tickets, etc), Homeless (Panhandlers),Accessibility for Disabled People, Personal safey (How safe do people feel), Customer Service (Can people find help), Time (Scheduling, are things on time, Reliability), Signage (How clear are the times), Cleanliness. Provide a JSON response in this format replace 'boolean' with true/false and 'text' with the answer: {\"overallOpinion\": \"text\", \"mentionsDrivers\": \"boolean\", \"opinionOfDriver\": \"text\", \"driversSummary\": \"text\", \"mentionsPurchasing\": \"boolean\", \"opinionOfPurchasing\": \"text\", \"purchasingSummary\": \"text\", \"mentionsHomeless\": \"boolean\", \"opinionOfHomeless\": \"text\", \"homelessSummary\": \"text\", \"mentionsAccessibility\": \"boolean\", \"opinionOfAccessibility\": \"text\", \"accessibilitySummary\": \"text\", \"mentionsSafety\": \"boolean\", \"opinionOfSafety\": \"text\", \"safetySummary\": \"text\", \"mentionsCustomerService\": \"boolean\", \"opinionOfCustomerService\": \"text\", \"customerServiceSummary\": \"text\", \"mentionsTime\": \"boolean\", \"opinionOfTime\": \"text\", \"timeSummary\": \"text\", \"mentionsSignage\": \"boolean\", \"opinionOfSignage\": \"text\", \"signageSummary\": \"text\", \"mentionsCleanliness\": \"boolean\", \"opinionOfCleanliness\": \"text\", \"cleanlinessSummary\": \"text\"}`
)

type ChatGPTResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     any    `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}
