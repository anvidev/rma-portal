// See https://svelte.dev/docs/kit/types#app.d.ts

import type { api } from '$lib/server/api'

// for information about these interfaces
declare global {
	namespace App {
		interface Error {
			requestId?: string
		}
		interface Locals {
			user: User
			api: typeof api
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {}
