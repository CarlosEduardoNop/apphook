package calls

type Calls struct {
	ID        string `bson:"_id,omiempty" json:"id"`
	Url       string `bson:"url" json:"url max:255"`
	Event     int    `bson:"event" json:"event"`
	Payload   string `bson:"payload" json:"payload"`
	Response  string `bson:"response" json:"response"`
	CreatedAt string `bson:"created_at" json:"created_at"`
	UpdatedAt string `bson:"updated_at" json:"updated_at"`
}
