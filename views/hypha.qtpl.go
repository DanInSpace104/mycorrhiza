// Code generated by qtc from "hypha.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/hypha.qtpl:1
package views

//line views/hypha.qtpl:1
import "path/filepath"

//line views/hypha.qtpl:2
import "strings"

//line views/hypha.qtpl:4
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/hypha.qtpl:5
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/hypha.qtpl:6
import "github.com/bouncepaw/mycorrhiza/user"

//line views/hypha.qtpl:7
import "github.com/bouncepaw/mycorrhiza/util"

//line views/hypha.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/hypha.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/hypha.qtpl:9
func streamnonExistentHyphaNotice(qw422016 *qt422016.Writer, h *hyphae.Hypha, u *user.User) {
//line views/hypha.qtpl:9
	qw422016.N().S(`
<section class="non-existent-hypha">
	<h2 class="non-existent-hypha__title">This hypha does not exist</h2>
	`)
//line views/hypha.qtpl:12
	if user.AuthUsed && u.Group == "anon" {
//line views/hypha.qtpl:12
		qw422016.N().S(`
	<p>You are not authorized to create new hyphae. Here is what you can do:</p>
	<ul>
		<li><a href="/login">Log in to your account, if you have one</a></li>
		`)
//line views/hypha.qtpl:16
		if cfg.UseRegistration {
//line views/hypha.qtpl:16
			qw422016.N().S(`<li><a href="/register">Register a new account</a></li>`)
//line views/hypha.qtpl:16
		}
//line views/hypha.qtpl:16
		qw422016.N().S(`
	</ul>
	`)
//line views/hypha.qtpl:18
	} else {
//line views/hypha.qtpl:18
		qw422016.N().S(`

	<div class="non-existent-hypha__ways">
	<section class="non-existent-hypha__way">
		<h3 class="non-existent-hypha__subtitle">📝 Write a text</h3>
		<p>Write a note, a diary, an article, a story or anything textual using <a href="https://mycorrhiza.lesarbr.es/hypha/mycomarkup" class="shy-link">Mycomarkup</a>. Full history of edits to the document will be saved.</p>
		<p>Make sure to follow this wiki&apos;s writing conventions if there are any.</p>
		<a class="btn btn_accent stick-to-bottom" href="/edit/`)
//line views/hypha.qtpl:25
		qw422016.E().S(h.Name)
//line views/hypha.qtpl:25
		qw422016.N().S(`">Create</a>
	</section>

	<section class="non-existent-hypha__way">
		<h3 class="non-existent-hypha__subtitle">🖼 Upload a media</h3>
		<p>Upload a picture, a video or an audio. Most common formats can be accessed from the browser, others can be only downloaded afterwards. You can write a description for the media later.</p>
		<form action="/upload-binary/`)
//line views/hypha.qtpl:31
		qw422016.E().S(h.Name)
//line views/hypha.qtpl:31
		qw422016.N().S(`"
        		method="post" enctype="multipart/form-data"
        		class="upload-binary">
        	<label for="upload-binary__input"></label>
        	<input type="file" id="upload-binary__input" name="binary">

        	<input type="submit" class="btn stick-to-bottom" value="Upload">
        </form>
	</section>
	</div>
	`)
//line views/hypha.qtpl:41
	}
//line views/hypha.qtpl:41
	qw422016.N().S(`
</section>
`)
//line views/hypha.qtpl:43
}

//line views/hypha.qtpl:43
func writenonExistentHyphaNotice(qq422016 qtio422016.Writer, h *hyphae.Hypha, u *user.User) {
//line views/hypha.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/hypha.qtpl:43
	streamnonExistentHyphaNotice(qw422016, h, u)
//line views/hypha.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line views/hypha.qtpl:43
}

//line views/hypha.qtpl:43
func nonExistentHyphaNotice(h *hyphae.Hypha, u *user.User) string {
//line views/hypha.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/hypha.qtpl:43
	writenonExistentHyphaNotice(qb422016, h, u)
//line views/hypha.qtpl:43
	qs422016 := string(qb422016.B)
//line views/hypha.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/hypha.qtpl:43
	return qs422016
//line views/hypha.qtpl:43
}

//line views/hypha.qtpl:45
func StreamNaviTitleHTML(qw422016 *qt422016.Writer, h *hyphae.Hypha) {
//line views/hypha.qtpl:45
	qw422016.N().S(`
`)
//line views/hypha.qtpl:47
	var (
		prevAcc = "/hypha/"
		parts   = strings.Split(h.Name, "/")
	)

//line views/hypha.qtpl:51
	qw422016.N().S(`
<h1 class="navi-title">
`)
//line views/hypha.qtpl:53
	qw422016.N().S(`<a href="/hypha/`)
//line views/hypha.qtpl:54
	qw422016.E().S(cfg.HomeHypha)
//line views/hypha.qtpl:54
	qw422016.N().S(`">`)
//line views/hypha.qtpl:55
	qw422016.N().S(cfg.NaviTitleIcon)
//line views/hypha.qtpl:55
	qw422016.N().S(`<span aria-hidden="true" class="navi-title__colon">:</span></a>`)
//line views/hypha.qtpl:59
	for i, part := range parts {
//line views/hypha.qtpl:60
		if i > 0 {
//line views/hypha.qtpl:60
			qw422016.N().S(`<span aria-hidden="true" class="navi-title__separator">/</span>`)
//line views/hypha.qtpl:62
		}
//line views/hypha.qtpl:62
		qw422016.N().S(`<a href="`)
//line views/hypha.qtpl:64
		qw422016.E().S(prevAcc + part)
//line views/hypha.qtpl:64
		qw422016.N().S(`" rel="`)
//line views/hypha.qtpl:64
		if i == len(parts)-1 {
//line views/hypha.qtpl:64
			qw422016.N().S(`bookmark`)
//line views/hypha.qtpl:64
		} else {
//line views/hypha.qtpl:64
			qw422016.N().S(`up`)
//line views/hypha.qtpl:64
		}
//line views/hypha.qtpl:64
		qw422016.N().S(`">`)
//line views/hypha.qtpl:65
		qw422016.N().S(util.BeautifulName(part))
//line views/hypha.qtpl:65
		qw422016.N().S(`</a>`)
//line views/hypha.qtpl:67
		prevAcc += part + "/"

//line views/hypha.qtpl:68
	}
//line views/hypha.qtpl:69
	qw422016.N().S(`
</h1>
`)
//line views/hypha.qtpl:71
}

//line views/hypha.qtpl:71
func WriteNaviTitleHTML(qq422016 qtio422016.Writer, h *hyphae.Hypha) {
//line views/hypha.qtpl:71
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/hypha.qtpl:71
	StreamNaviTitleHTML(qw422016, h)
//line views/hypha.qtpl:71
	qt422016.ReleaseWriter(qw422016)
//line views/hypha.qtpl:71
}

//line views/hypha.qtpl:71
func NaviTitleHTML(h *hyphae.Hypha) string {
//line views/hypha.qtpl:71
	qb422016 := qt422016.AcquireByteBuffer()
//line views/hypha.qtpl:71
	WriteNaviTitleHTML(qb422016, h)
//line views/hypha.qtpl:71
	qs422016 := string(qb422016.B)
//line views/hypha.qtpl:71
	qt422016.ReleaseByteBuffer(qb422016)
//line views/hypha.qtpl:71
	return qs422016
//line views/hypha.qtpl:71
}

//line views/hypha.qtpl:73
func StreamAttachmentHTML(qw422016 *qt422016.Writer, h *hyphae.Hypha) {
//line views/hypha.qtpl:73
	qw422016.N().S(`
	`)
//line views/hypha.qtpl:74
	switch filepath.Ext(h.BinaryPath) {
//line views/hypha.qtpl:76
	case ".jpg", ".gif", ".png", ".webp", ".svg", ".ico":
//line views/hypha.qtpl:76
		qw422016.N().S(`
	<div class="binary-container binary-container_with-img">
		<a href="/binary/`)
//line views/hypha.qtpl:78
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:78
		qw422016.N().S(`"><img src="/binary/`)
//line views/hypha.qtpl:78
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:78
		qw422016.N().S(`"/></a>
	</div>

	`)
//line views/hypha.qtpl:81
	case ".ogg", ".webm", ".mp4":
//line views/hypha.qtpl:81
		qw422016.N().S(`
	<div class="binary-container binary-container_with-video">
		<video controls>
			<source src="/binary/`)
//line views/hypha.qtpl:84
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:84
		qw422016.N().S(`"/>
			<p>Your browser does not support video. <a href="/binary/`)
//line views/hypha.qtpl:85
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:85
		qw422016.N().S(`">Download video</a></p>
		</video>
	</div>

	`)
//line views/hypha.qtpl:89
	case ".mp3":
//line views/hypha.qtpl:89
		qw422016.N().S(`
	<div class="binary-container binary-container_with-audio">
		<audio controls>
			<source src="/binary/`)
//line views/hypha.qtpl:92
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:92
		qw422016.N().S(`"/>
			<p>Your browser does not support audio. <a href="/binary/`)
//line views/hypha.qtpl:93
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:93
		qw422016.N().S(`">Download audio</a></p>
		</audio>
	</div>

	`)
//line views/hypha.qtpl:97
	default:
//line views/hypha.qtpl:97
		qw422016.N().S(`
	<div class="binary-container binary-container_with-nothing">
		<p><a href="/binary/`)
//line views/hypha.qtpl:99
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:99
		qw422016.N().S(`">Download media</a></p>
	</div>
`)
//line views/hypha.qtpl:101
	}
//line views/hypha.qtpl:101
	qw422016.N().S(`
`)
//line views/hypha.qtpl:102
}

//line views/hypha.qtpl:102
func WriteAttachmentHTML(qq422016 qtio422016.Writer, h *hyphae.Hypha) {
//line views/hypha.qtpl:102
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/hypha.qtpl:102
	StreamAttachmentHTML(qw422016, h)
//line views/hypha.qtpl:102
	qt422016.ReleaseWriter(qw422016)
//line views/hypha.qtpl:102
}

//line views/hypha.qtpl:102
func AttachmentHTML(h *hyphae.Hypha) string {
//line views/hypha.qtpl:102
	qb422016 := qt422016.AcquireByteBuffer()
//line views/hypha.qtpl:102
	WriteAttachmentHTML(qb422016, h)
//line views/hypha.qtpl:102
	qs422016 := string(qb422016.B)
//line views/hypha.qtpl:102
	qt422016.ReleaseByteBuffer(qb422016)
//line views/hypha.qtpl:102
	return qs422016
//line views/hypha.qtpl:102
}
