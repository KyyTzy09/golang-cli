package structs

type AIToolsResponse struct {
	ToolName   string            `json:"tool_name"`
	Suggestion string            `json:"suggestion"`
	Params     map[string]string `json:"params"`
}