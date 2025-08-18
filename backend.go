// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	tb "github.com/nsf/termbox-go"
)

func Init() error {
	if err := tb.Init(); err != nil {
		return err
	}

	tb.SetInputMode(tb.InputEsc | tb.InputMouse)
	tb.SetOutputMode(tb.Output256)
	return nil
}

func Close() {
	tb.Close()
}

func TerminalDimensions() (int, int) {
	tb.Sync()
	width, height := tb.Size()
	return width, height	
}

func Clear() {
	tb.Clear(tb.ColorDefault, tb.Attribute(Theme.Default.Bg+1))
}