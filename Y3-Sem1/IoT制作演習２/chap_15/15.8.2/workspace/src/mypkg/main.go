package main

import (
	"image/color"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
)

// 四角形のサイズ
const width = 250
const height = 250

// 四角形の作成 … ①
func newRectImage(imgs *glutil.Images) *glutil.Image {
	// 指定したサイズの四角形を作成する
	img := imgs.NewImage(width, height)
	// 四角形を黄色に塗りつぶす
	c := color.RGBA{0x00, 0xFF, 0x00, 0x00}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.RGBA.Set(x, y, c)
		}
	}
	img.Upload()
	return img
}

// 四角形の描画処理 … ②
func paintRect(img *glutil.Image, sz size.Event, te touch.Event) {
	// タッチイベントの場所に四角形を描画する
	// サイズの単位をpixelからpt(ポイント)に変更する
	left := geom.Pt(te.X / sz.PixelsPerPt)
	top := geom.Pt(te.Y / sz.PixelsPerPt)
	right := left + geom.Pt(width/sz.PixelsPerPt)
	bottom := top + geom.Pt(height/sz.PixelsPerPt)
	// 左上、右上、左下の位置を設定する
	leftTop := geom.Point{left, top}
	rightTop := geom.Point{right, top}
	leftBottom := geom.Point{left, bottom}
	// フレームバッファに四角形を描画する
	img.Draw(sz, leftTop, rightTop, leftBottom, img.RGBA.Bounds())
}

// Androidアプリのメイン処理
func mainLoop(a app.App) {
	var ctx gl.Context      // コンテキスト（画面の情報などを格納）
	var imgs *glutil.Images // 画像の集合
	var img *glutil.Image   // 画像（四角形）
	var sz size.Event       // 画面サイズ
	var te touch.Event      // タッチイベント（位置情報などを格納）

	// アプリのイベントを処理するメインループ … ③
	eventCh := a.Events()
	for e := range eventCh {
		switch e := a.Filter(e).(type) {
		case lifecycle.Event:
			// 画面の表示イベント … ④
			switch e.Crosses(lifecycle.StageVisible) {
			case lifecycle.CrossOn:
				// 画面表示（androidのonStart）
				// コンテキストと画像の集合を取得する
				ctx, _ = e.DrawContext.(gl.Context)
				imgs = glutil.NewImages(ctx)
				// 四角形を作成する
				img = newRectImage(imgs)
				// 描画イベントの通知
				a.Send(paint.Event{})
			case lifecycle.CrossOff:
				// 画面非表示（androidのonStop）
				// 画像、画像の集合、コンテキストを破棄する
				img.Release()
				imgs.Release()
				img = nil
				imgs = nil
				ctx = nil
			}
		case size.Event:
			// 画面サイズイベント … ⑤
			sz = e
		case touch.Event:
			// タッチイベント … ⑥
			te = e
		case paint.Event:
			// 描画イベント … ⑦
			// 変数に値がない場合や外部イベントの場合は処理しない
			if ctx == nil || imgs == nil || img == nil || e.External {
				continue
			}

			// 画面の色を白に初期化する
			ctx.ClearColor(1, 1, 1, 1)
			ctx.Clear(gl.COLOR_BUFFER_BIT)

			// 四角形を描画する
			paintRect(img, sz, te)

			// 保留中の描画処理をフラッシュする
			a.Publish()

			// 次回の描画イベントを通知する
			a.Send(paint.Event{})
		}
	}
}

func main() {
	// Androidアプリを起動する
	app.Main(mainLoop)
}
