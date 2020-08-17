package weread

type ShelfBookResp struct {
	Synckey        int              `json:"synckey"`
	BookProgress   []WeBookProgress `json:"bookProgress"`
	Removed        []interface{}    `json:"removed"`
	RemovedArchive []interface{}    `json:"removedArchive"`
	LectureRemoved []interface{}    `json:"lectureRemoved"`
	Archive        []interface{}    `json:"archive"`
	Books          []WeBook         `json:"books"`
	LectureBooks   []WeBook         `json:"lectureBooks"`
	LectureSynckey int              `json:"lectureSynckey"`
	LectureUpdate  []interface{}    `json:"lectureUpdate"`
}

type WeBookProgress struct {
	BookID        string `json:"bookId"`
	Progress      int    `json:"progress"`
	ChapterUID    int    `json:"chapterUid"`
	ChapterOffset int    `json:"chapterOffset"`
	AppID         string `json:"appId"`
	UpdateTime    int    `json:"updateTime"`
	Synckey       int    `json:"synckey,omitempty"`
}

type WeBook struct {
	BookID                string  `json:"bookId"`
	Title                 string  `json:"title"`
	Author                string  `json:"author"`
	Cover                 string  `json:"cover"`
	Version               int     `json:"version"`
	Format                string  `json:"format"`
	Type                  int     `json:"type"`
	Price                 float64 `json:"price"`
	OriginalPrice         int     `json:"originalPrice"`
	Soldout               int     `json:"soldout"`
	BookStatus            int     `json:"bookStatus"`
	PayType               int     `json:"payType"`
	Finished              int     `json:"finished"`
	MaxFreeChapter        int     `json:"maxFreeChapter"`
	Free                  int     `json:"free"`
	McardDiscount         int     `json:"mcardDiscount"`
	Ispub                 int     `json:"ispub"`
	FinishReading         int     `json:"finishReading"`
	Category              string  `json:"category"`
	Paid                  int     `json:"paid"`
	HasLecture            int     `json:"hasLecture"`
	UpdateTime            int     `json:"updateTime"`
	LastChapterIdx        int     `json:"lastChapterIdx"`
	Secret                int     `json:"secret"`
	ReadUpdateTime        int     `json:"readUpdateTime"`
	AuthorVids            string  `json:"authorVids,omitempty"`
	ShouldHideTTS         int     `json:"shouldHideTTS,omitempty"`
	CpName                string  `json:"cpName,omitempty"`
	LPushName             string  `json:"lPushName,omitempty"`
	LectureReadUpdateTime int     `json:"lectureReadUpdateTime"`
	LectureOffline        int     `json:"lectureOffline"`
	LecturePaid           int     `json:"lecturePaid"`
}
