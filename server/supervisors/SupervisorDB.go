package supervisors

import (
	"strconv"

	"github.com/lightfeather_coding-challenge/server/libs"
)

type SupervisorDB struct {
	// DataMap map[int32]*SupervisorStruct
}

func NewDB() *SupervisorDB {
	r := new(SupervisorDB)
	return r
}

// CRUD
func (x *SupervisorDB) GetAll() []*SupervisorStruct {
	r := make([]*SupervisorStruct, 0)
	libs.HttpGet("https://o3m5qixdng.execute-api.us-east-1.amazonaws.com/api/managers", &r)

	// filter numeric jurisdictions
	for i := 0; i < len(r); i++ {
		if _, err := strconv.Atoi(r[i].Jurisdiction); err == nil {
			// remove
			r = append(r[:i], r[i+1:]...)
			i--
		}
	}
	return r
}
