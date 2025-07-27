package options

import "github.com/jsas4coding/utify/pkg/messages"

type Options struct {
	Bold      bool
	Italic    bool
	NoColor   bool
	NoIcon    bool
	NoStyle   bool
	Exit      bool
	ShowIcons bool
	Callback  func(messages.Type, string)
}

func Default() *Options {
	return &Options{}
}

func (o *Options) WithBold() *Options {
	o.Bold = true
	return o
}

func (o *Options) WithItalic() *Options {
	o.Italic = true
	return o
}

func (o *Options) WithoutColor() *Options {
	o.NoColor = true
	return o
}

func (o *Options) WithoutIcon() *Options {
	o.NoIcon = true
	o.ShowIcons = false
	return o
}

func (o *Options) WithIcon() *Options {
	o.ShowIcons = true
	o.NoIcon = false
	return o
}

func (o *Options) WithoutStyle() *Options {
	o.NoStyle = true
	return o
}

func (o *Options) WithExit() *Options {
	o.Exit = true
	o.Callback = nil
	return o
}

func (o *Options) WithCallback(cb func(messages.Type, string)) *Options {
	o.Callback = cb
	o.Exit = false
	return o
}
