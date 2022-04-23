package todos

var Toos = []Todo{
	{
		ID:       1,
		UserID:   1,
		Activity: "Some Activity",
		Time:     "Some Time",
		Date:     "Some Date",
		Status:   "Completed",
	},
	{
		Activity: "Some Activity",
		UserID:   2,
		ID:       3,
		Time:     "Some Time",
		Date:     "Some Date",
		Status:   "Completed",
	},
}

func GetData(total int) []Todo {
	return Toos
}
