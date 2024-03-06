package services

import "model"


type DiaryService interface{
	Save(model.Diary) model.Diary
	
	findAll() [] model.Diary
}


