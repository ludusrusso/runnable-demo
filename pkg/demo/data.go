package demo

type DataStatus string

const (
	DataStatusPending DataStatus = "pending"
	DataStatusLoading DataStatus = "loading"
	DataStatusDone    DataStatus = "done"
)

type Data struct {
	value         int
	Status        DataStatus
	ComputedValue int
}

func NewData(value int) Data {
	return Data{
		value:  value,
		Status: DataStatusPending,
	}
}
