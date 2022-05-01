package worker

type Worker interface {
	Run()
	CheckInput(interface{}) bool
}
