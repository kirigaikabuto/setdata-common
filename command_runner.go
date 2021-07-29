package setdata_common

type (
	Command interface {
		Exec(pub interface{}) (interface{}, error)
	}

	CommandHandler interface {
		ExecCommand(Command) (interface{}, error)
	}
)

type commandHandler struct {
	pub interface{}
}

func (c *commandHandler) ExecCommand(cmd Command) (interface{}, error) {
	return cmd.Exec(c.pub)
}

func NewCommandHandler(pub interface{}) CommandHandler {
	return &commandHandler{pub}
}
