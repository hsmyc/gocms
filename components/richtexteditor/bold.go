package richtexteditor

func Bold(data string, specArray []string) []byte {
	specArray = append(specArray, "bold")
	return RichTextEditor(data, specArray)
}
