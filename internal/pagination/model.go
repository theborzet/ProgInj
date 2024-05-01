package pagination

type Paginator struct {
	CurrentPage  int
	TotalPages   int
	PageSize     int
	TotalItems   int
	HasPrevious  bool
	HasNext      bool
	PreviousPage int
	NextPage     int
	PageNumbers  []PageNumber
}

type PageNumber struct {
	Number   int
	URL      string
	IsActive bool
}
