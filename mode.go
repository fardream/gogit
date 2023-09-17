package gogit

// Mode
type Mode int32

const (
	Mode_RegularFile Mode = 100644 // 100644
	Mode_Tree        Mode = 40000  // 40000
)
