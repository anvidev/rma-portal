import { writable } from 'svelte/store'
import type { Lightbox } from './lightbox-state.svelte'

export const activeLightbox = writable<Lightbox | undefined>(undefined)
