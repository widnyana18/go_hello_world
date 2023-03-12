package main

func (msc *music) updateMusic(name string) string {
	msc.title = name
	return msc.title
}
