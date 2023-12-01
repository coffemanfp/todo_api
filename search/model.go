package search

type Search struct {
	ClientID       int
	Title          *string
	IsDone         *bool
	IsAddedToMyDay *bool
	IsImportant    *bool
	HasDueDate     *bool
	ExpireSoon     *bool
}

func New(clientID int, title *string, isDone, isAddedToMyDay, isImportant, hasDueDate, expireSoon *bool) (s Search, err error) {
	s.ClientID = clientID
	s.Title = title
	s.IsDone = isDone
	s.IsAddedToMyDay = isAddedToMyDay
	s.IsImportant = isImportant
	s.HasDueDate = hasDueDate
	s.ExpireSoon = expireSoon
	return
}
