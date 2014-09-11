package channels

import "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"

type VolumeDevice interface {
	SetVolume(volume float64) error
	VolumeUp() error
	VolumeDown() error
	SetMuted(muted bool) error
	ToggleMuted() error
}

type VolumeState struct {
	Volume *float64 `json:"volume,omitempty"`
	Muted  *bool    `json:"muted,omitempty"`
}

type VolumeChannel struct {
	baseChannel
	device VolumeDevice
}

func NewVolumeChannel(device VolumeDevice) *VolumeChannel {
	return &VolumeChannel{baseChannel{}, device}
}

func (c *VolumeChannel) Set(message mqtt.Message, state *float64, reply *interface{}) error {
	return c.device.SetVolume(*state)
}

func (c *VolumeChannel) VolumeUp(message mqtt.Message, _, reply *interface{}) error {
	return c.device.VolumeUp()
}

func (c *VolumeChannel) VolumeDown(message mqtt.Message, _, reply *interface{}) error {
	return c.device.VolumeDown()
}

func (c *VolumeChannel) Mute(message mqtt.Message, _, reply *interface{}) error {
	return c.device.SetMuted(true)
}

func (c *VolumeChannel) Unmute(message mqtt.Message, _, reply *interface{}) error {
	return c.device.SetMuted(false)
}

func (c *VolumeChannel) ToggleMuted(message mqtt.Message, _, reply *interface{}) error {
	return c.device.ToggleMuted()
}

func (c *VolumeChannel) SendState(volume *float64, muted *bool) error {
	return c.SendEvent("state", &VolumeState{
		Volume: volume,
		Muted:  muted,
	})
}
