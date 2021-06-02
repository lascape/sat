package sat

import (
	"testing"
)

var (
	dicter = DefaultDict()
)

func TestConvertPlain(t *testing.T) {
	cases := map[string]string{
		"繁体中文": "繁體中文",
		"还带来了音乐等方面的重大更新": "還帶來了音樂等方麵的重大更新",
		"如果你是 iOS 14.5.1 用户，或者 Apple Music 的忠实粉丝，那么根本找不到理由去拒绝 iOS 14.6。": "如菓妳昰 iOS 14.5.1 用戶，或者 Apple Music 的忠實粉絲，那麼根本找不到理由去拒絕 iOS 14.6。",
	}

	for sc, expected := range cases {
		actually := dicter.ReadReverse(sc)
		if expected != actually {
			t.Errorf("\nexpected: %s\nactually: %s", expected, actually)
		}
	}

}

func TestConvertHTML(t *testing.T) {
	source := `<p id="KN1928">如果你最近看到有人打<strong>王者荣耀</strong>时情绪波动较大，不一定是因为 <em>ta</em> 的队友太坑或对手太强，而或许只因 ta 升级了 iOS 14.5.1。</p>
	<p><img src="https://img.example.com/not/exist/image.jpg" alt="无效的图片"></p>
	<p>发热、掉帧、卡顿，iOS 14.5.1 的性能 bug，成为<a href="https://example.com" target="_blank" rel="nofollow">近期</a> iPhone 用户吐槽最多的话题。无论你是哪款 iPhone，升级 iOS 14.5.1 之后，都有可能遇到性能缩水的问题。</p>`

	expect := `<p id="KN1928">如菓妳最近看到有人打<strong>王者榮燿</strong>時情緒波動較大，不一定昰囙爲 <em>ta</em> 的隊友太阬或對手太強，而或許隻囙 ta 陞級了 iOS 14.5.1。</p>
	<p><img src="https://img.example.com/not/exist/image.jpg" alt="無傚的圖片"></p>
	<p>髮熱、掉幀、卡頓，iOS 14.5.1 的性能 bug，成爲<a href="https://example.com" target="_blank" rel="nofollow">近期</a> iPhone 用戶吐槽最多的話題。無論妳昰哪欵 iPhone，陞級 iOS 14.5.1 之后，都有可能遇到性能縮水的問題。</p>`

	out := dicter.ReadReverse(source)
	if expect != out {
		t.Error(out)
	}
}
