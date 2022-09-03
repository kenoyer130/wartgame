package models

type StatusMessage struct {
	Phase    string
	Messsage string
	Keys     string
}

func (re *StatusMessage) Clear() {
	re.Phase = ""
	re.Messsage = ""
	re.Keys = ""
}