package models

/*
這裡是定義訊息資料格式的地方
*/
type Message struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Time   string `json:"time"`
	Exit   bool   `json:"exit"`
	Join   bool   `json:"join"`
}
