package internal

type Exported struct {
	A int
	B int
}

type ExportedNested struct {
	Exported
	C int
}
