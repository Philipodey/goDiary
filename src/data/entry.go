package model

		
type Entry struct{
	Id int `json:"id"`
	Date string `json:"date"`
	Title string `json:"title"`
	Body string `json:"body"`
	DiaryId string `json:"diaryId"`
}

func (id *Entry) SetId(identity int){
	id.Id = identity
}
func(id *Entry) GetId() int{
	return id.Id
}
func(date *Entry) SetDateCreated(dateTime string){
	date.Date = dateTime
}

func(date *Entry) GetDateCreated() string{
	return date.Date
}

func(titles *Entry) SetTitle(title string){
	titles.Title = title
}

func(titles *Entry) GetTitle() string{
	return titles.Title
}
func(body *Entry) SetBody(bodies string){
	body.Body = bodies
}

func(body *Entry) GetBody() string{
	return body.Body
}
func(diaryId *Entry) SetDairyId(diaryIdentity string){
	diaryId.DiaryId = diaryIdentity
}

func(diaryId *Entry) GetDiaryId() string{
	return diaryId.DiaryId
}


