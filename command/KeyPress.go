package command

type KeyPress struct {
	KeyPressCommand Command
}

func (press *KeyPress) Press() {
	press.KeyPressCommand.execute()
}

type OnCommand struct {
	CommandItem ToggleItem
}

func (on *OnCommand) execute() {
	on.CommandItem.on()
}

type OffCommand struct {
	CommandItem ToggleItem
}

func (off *OffCommand) execute() {
	off.CommandItem.off()
}
