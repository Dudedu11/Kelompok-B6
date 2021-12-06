package datastruct

type   Menyukai struct {
	User_id  	int64  `json:"user_id"`
	Feed_id  	int64  `json:"feed_id"`
	Status_like bool `json:"status_like"`
	Tgl_like 	string `json:"tgl_like"`
}

type Member struct {
	Username       string `json:"username"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Phonenumber    string `json:"phonenumber"`
	Password       string `json:"password"`
	Email_verified string `json:"email_verified"`
	Image_file     string `json:"image_file"`
	Identity_type  string `json:"identity_type"`
	Identity_no    string `json:"identity_no"`
	Emergency_call string `json:"emergency_call"`
	Address_ktp    string `json:"address_ktp"`
	Domisili       string `json:"domisili"`
	Create_date    string `json:"create_date"`
	Update_date    string `json:"update_date"`
	Email          string `json:"email"`
	Isprivate      bool   `json:"isPrivate"`
	User_id        string `json:"user_id"`
}

type Feed struct {
    Image_feed   string `json:"image_feed"`
	Caption_feed string `json:"caption_feed"`
	Post_date	 string	`json:"post_date"`
	Feed_id      string `json:"feed_id"`
	
}

type Response1 struct {
	ID_user 	int64  `json:"id_user,omitempty"`
	Name_user 	string `json:"name_user,omitempty"`
	ID_feed 	int64  `json:"id_feed,omitempty"`
	Is_like  	bool   `json:"is_like,omitempty"`
	Message     string `json:"message,omitempty"`
}

type Response2 struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Menyukai `json:"data"`
}

type Response3 struct {
	Status     int      `json:"status"`
	Message    string   `json:"message"`
	JumlahLike int64      `json:"jml_like"`
}

type Response4 struct {
	Status     int      `json:"status"`
	Message    string   `json:"message"`
}
