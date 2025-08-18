// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	"image"
	"sync"
)


type Block struct {
	Border	bool
	BorderStyle Style
	BorderLeft, BorderRight, BorderTop, BorderBottom bool
	PaddingLeft, PaddingRight, PaddingTop, PaddingBottom int
	image.Rectangle
	Inner image.Rectangle
	Title string
	TitleStyle Style
	sync.Mutex
}

func NewBlock() *Block {
	return &Block{
		Border: true,
		BorderStyle: Theme.Block.Border,
		BorderLeft:  true,
		BorderRight: true,
		BorderTop:   true,
		BorderBottom:true,
		TitleStyle: Theme.Block.Title,
	}
}

func (self *Block) drawBorder(buf *Buffer) {
	
}