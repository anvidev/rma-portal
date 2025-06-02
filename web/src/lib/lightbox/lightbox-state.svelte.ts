import { activeLightbox } from './lightbox-store'

export type LightboxImage = {
	id: string | number
	url: string
}

export type LightboxOptions = { images: LightboxImage[]; loop?: boolean }

export class Lightbox {
	private imgs: LightboxImage[]
	private cur: LightboxImage | undefined = $state(undefined)
	private loop: boolean
	private curIdx: number = $state(-1)

	constructor(opts: LightboxOptions) {
		this.imgs = opts.images
		this.loop = opts.loop ?? false
	}

	#toggleBody() {
		document.getElementsByTagName('body')[0].classList.toggle('overflow-y-hidden')
	}

	open(id: string | number) {
		this.#toggleBody()
		const foundIdx = this.imgs.findIndex(img => img.id === id)
		if (foundIdx !== -1) {
			this.curIdx = foundIdx
			this.cur = this.imgs[foundIdx]
			activeLightbox.set(this)
		}
		activeLightbox.set(this)
	}

	close() {
		this.#toggleBody()
		this.cur = undefined
		this.curIdx = -1
		activeLightbox.set(undefined)
	}

	next() {
		if (this.curIdx === -1) return
		if (this.curIdx < this.imgs.length - 1) {
			this.curIdx++
		} else if (this.loop) {
			this.curIdx = 0
		} else {
			return
		}
		this.cur = this.imgs[this.curIdx]
	}

	previous() {
		if (this.curIdx === -1) return
		if (this.curIdx > 0) {
			this.curIdx--
		} else if (this.loop) {
			this.curIdx = this.imgs.length - 1
		} else {
			return
		}
		this.cur = this.imgs[this.curIdx]
	}

	current(): LightboxImage | undefined {
		return this.cur
	}

	// only works if url is a download url bla
	download() {
		if (!this.cur) return
		const a = document.createElement('a')
		a.href = this.cur.url
		a.click()
	}

	images(): LightboxImage[] {
		return this.imgs
	}
}
