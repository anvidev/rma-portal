// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	type User = {
		id: number
		name: string
		email: string
		isActive: boolean
		inserted: string
	}

	namespace App {
		// interface Error {}
		interface Locals {
			user: User | null
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {}
