package types

// Exercise is struct for a exercise
type Exercise struct {
	ID          interface{} `bson:"_id,omitempty"`
	Title       string      `bson:"title"`
	Description string      `bson:"description"`
	Testcase    string      `bson:"testcase"`
}
