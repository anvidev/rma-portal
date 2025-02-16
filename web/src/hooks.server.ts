import { validateAuth } from '$lib/server/auth'
import { redirect, type Handle } from '@sveltejs/kit'

export const handle: Handle = async ({ event, resolve }) => {
	event.locals.user = await validateAuth(event)

	if (event.url.pathname.startsWith('/admin')) {
		if (event.locals.user == null) {
			throw redirect(303, '/login')
		}
	}

	const response = resolve(event)
	return response
}
