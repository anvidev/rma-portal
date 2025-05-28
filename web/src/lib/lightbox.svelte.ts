export type LightboxImage = {
	id: string | number
	url: string
}

export type LightboxOptions = { images: LightboxImage[]; loop?: boolean }

export class Lightbox {
	private imgs: LightboxImage[]
	private cur: LightboxImage | undefined = $state(undefined)
	private loop: boolean

	constructor(opts: LightboxOptions) {
		this.imgs = opts.images
		this.loop = opts.loop ?? false
	}

	#toggleBody() {
		document.getElementsByTagName('body')[0].classList.toggle('overflow-y-hidden')
	}

	open(id: string | number) {
		this.#toggleBody()
		this.cur = this.imgs.find(img => img.id == id)
	}

	close() {
		this.#toggleBody()
		this.cur = undefined
	}

	next() {
		if (!this.cur) return
		const curIdx = this.imgs.findIndex(img => img.id == this.cur?.id)
		if (curIdx === -1) return
		if (curIdx < this.imgs.length - 1) {
			const nextImage = this.imgs[curIdx + 1]
			this.cur = nextImage
		} else {
			if (this.loop) {
				const firstImage = this.imgs[0]
				this.cur = firstImage
			}
		}
	}

	previous() {
		if (!this.cur) return
		const curIdx = this.imgs.findIndex(img => img.id == this.cur?.id)
		if (curIdx == -1) return
		if (curIdx > 0) {
			const previousImage = this.imgs[curIdx - 1]
			this.cur = previousImage
		} else {
			if (this.loop) {
				const lastImage = this.imgs[this.imgs.length - 1]
				this.cur = lastImage
			}
		}
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
}
